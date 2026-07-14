package dynamodb

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

	"quiz-app/internal/domain"
)

// metricsRepo keeps aggregate usage counters. Lifetime totals piggyback on
// the existing ID allocators (COUNTER#user / COUNTER#session are "entities
// ever created") plus a dedicated COUNTER#guest, so registered users =
// n(user) - n(guest). Per-day activity lives in DAY#<date> rollup items.
type metricsRepo struct{ s *Store }

// dayRollupTTL keeps a bit over a year of daily history; dauMarkerTTL only
// needs to outlive its own day (kept 3 days for timezone/debug slack).
const (
	dayRollupTTL = 400 * 24 * time.Hour
	dauMarkerTTL = 3 * 24 * time.Hour
)

func (r *metricsRepo) dayKey(day string) map[string]types.AttributeValue {
	return map[string]types.AttributeValue{
		"PK": &types.AttributeValueMemberS{Value: pkDay(day)},
		"SK": &types.AttributeValueMemberS{Value: skMeta},
	}
}

// bumpDay increments one counter attribute of the day's rollup item, creating
// it (with its retention TTL) on first write of the day.
func (r *metricsRepo) bumpDay(ctx context.Context, day string, attrs ...string) error {
	add := ""
	values := map[string]types.AttributeValue{
		":one": &types.AttributeValueMemberN{Value: "1"},
		":ttl": &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", time.Now().Add(dayRollupTTL).Unix())},
	}
	for i, attr := range attrs {
		if i > 0 {
			add += ", "
		}
		add += attr + " :one"
	}
	_, err := r.s.client.UpdateItem(ctx, &dynamodb.UpdateItemInput{
		TableName:        aws.String(r.s.table),
		Key:              r.dayKey(day),
		UpdateExpression: aws.String("ADD " + add + " SET #ttl = if_not_exists(#ttl, :ttl)"),
		ExpressionAttributeNames: map[string]string{
			"#ttl": "ttl",
		},
		ExpressionAttributeValues: values,
	})
	if err != nil {
		return fmt.Errorf("bump day %s (%v): %w", day, attrs, err)
	}
	return nil
}

func (r *metricsRepo) RecordUserCreated(isGuest bool, day string) error {
	ctx := context.Background()
	attr := "new_users"
	if isGuest {
		attr = "new_guests"
		if _, err := r.s.nextID(ctx, "guest"); err != nil {
			return err
		}
	}
	return r.bumpDay(ctx, day, attr)
}

func (r *metricsRepo) RecordGameStarted(userID uint, day string) error {
	ctx := context.Background()

	// One marker per (day, user): the conditional put succeeds only on the
	// user's first game of the day, which is when active_users is bumped.
	marker, err := attributevalue.MarshalMap(struct {
		PK  string `dynamodbav:"PK"`
		SK  string `dynamodbav:"SK"`
		TTL int64  `dynamodbav:"ttl"`
	}{pkDay(day), skDayUser(userID), time.Now().Add(dauMarkerTTL).Unix()})
	if err != nil {
		return fmt.Errorf("marshal dau marker: %w", err)
	}
	_, err = r.s.client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName:           aws.String(r.s.table),
		Item:                marker,
		ConditionExpression: aws.String("attribute_not_exists(PK)"),
	})
	firstGameOfDay := true
	var condErr *types.ConditionalCheckFailedException
	if errors.As(err, &condErr) {
		firstGameOfDay = false
	} else if err != nil {
		return fmt.Errorf("dau marker %s/%d: %w", day, userID, err)
	}

	attrs := []string{"games_started"}
	if firstGameOfDay {
		attrs = append(attrs, "active_users")
	}
	return r.bumpDay(ctx, day, attrs...)
}

// counterValue reads one COUNTER item's running total (0 when absent).
func (r *metricsRepo) counterValue(ctx context.Context, entity string) (int64, error) {
	var item struct {
		N int64 `dynamodbav:"n"`
	}
	err := r.s.getItem(ctx, pkCounter(entity), skMeta, &item)
	if errors.Is(err, domain.ErrNotFound) {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}
	return item.N, nil
}

func (r *metricsRepo) Totals() (domain.MetricsTotals, error) {
	ctx := context.Background()
	var t domain.MetricsTotals

	usersEver, err := r.counterValue(ctx, "user")
	if err != nil {
		return t, err
	}
	guestsEver, err := r.counterValue(ctx, "guest")
	if err != nil {
		return t, err
	}
	gamesEver, err := r.counterValue(ctx, "session")
	if err != nil {
		return t, err
	}

	t.RegisteredUsers = usersEver - guestsEver
	t.GuestUsers = guestsEver
	t.TotalGames = gamesEver
	return t, nil
}

func (r *metricsRepo) Daily(days int) ([]domain.DailyMetrics, error) {
	ctx := context.Background()
	now := time.Now().In(domain.MetricsTimezone)

	result := make([]domain.DailyMetrics, days)
	keys := make([]map[string]types.AttributeValue, days)
	index := make(map[string]int, days)
	for i := 0; i < days; i++ {
		day := domain.MetricsDay(now.AddDate(0, 0, -i))
		result[i] = domain.DailyMetrics{Date: day}
		keys[i] = r.dayKey(day)
		index[day] = i
	}

	// BatchGetItem allows 100 keys per call; days is capped at 31 by the
	// handler, so one call suffices (unprocessed keys are still retried).
	pending := keys
	for len(pending) > 0 {
		out, err := r.s.client.BatchGetItem(ctx, &dynamodb.BatchGetItemInput{
			RequestItems: map[string]types.KeysAndAttributes{
				r.s.table: {Keys: pending, ConsistentRead: aws.Bool(true)},
			},
		})
		if err != nil {
			return nil, fmt.Errorf("batch get daily metrics: %w", err)
		}
		for _, raw := range out.Responses[r.s.table] {
			var item dayMetricsItem
			if err := attributevalue.UnmarshalMap(raw, &item); err != nil {
				return nil, fmt.Errorf("unmarshal day metrics: %w", err)
			}
			day := item.PK[len("DAY#"):]
			if i, ok := index[day]; ok {
				result[i].GamesStarted = item.GamesStarted
				result[i].ActiveUsers = item.ActiveUsers
				result[i].NewGuests = item.NewGuests
				result[i].NewUsers = item.NewUsers
			}
		}
		pending = out.UnprocessedKeys[r.s.table].Keys
	}

	return result, nil
}
