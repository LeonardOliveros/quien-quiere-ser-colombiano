// Package dynamodb is the DynamoDB adapter of the domain.Store port, meant
// for serverless deployments (DB_DRIVER=dynamodb).
//
// Everything lives in one on-demand table (env DYNAMODB_TABLE) with generic
// PK/SK string keys and no GSIs: every access path the port needs is served
// by strongly consistent GetItem/Query calls, so freshly written data (login
// tokens, new sessions, pause state) is immediately readable — GSIs are only
// eventually consistent and would race the read-your-own-write flows of the
// HTTP layer. See keys.go for the full key layout.
//
// The question bank is immutable at runtime and cached in memory (cache.go);
// stats are aggregated in Go over the user's own partitions, which stay small.
//
// Local development/testing: run DynamoDB Local (amazon/dynamodb-local) and
// set DYNAMODB_ENDPOINT=http://localhost:8000 — the adapter then also creates
// the table automatically if it does not exist.
package dynamodb

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

	"quiz-app/internal/domain"
)

// Store implements domain.Store on top of a single DynamoDB table.
type Store struct {
	client *dynamodb.Client
	table  string
	cache  bankCache
}

var _ domain.Store = (*Store)(nil)

// Open connects to DynamoDB. When DYNAMODB_ENDPOINT is set (DynamoDB Local),
// it also applies dummy defaults for region/credentials and auto-creates the
// table, so local runs need no AWS account configuration.
func Open(table string) (*Store, error) {
	if table == "" {
		return nil, errors.New("DYNAMODB_TABLE must be set when DB_DRIVER=dynamodb")
	}
	ctx := context.Background()
	endpoint := os.Getenv("DYNAMODB_ENDPOINT")

	var optFns []func(*config.LoadOptions) error
	if endpoint != "" {
		if os.Getenv("AWS_REGION") == "" && os.Getenv("AWS_DEFAULT_REGION") == "" {
			optFns = append(optFns, config.WithRegion("us-east-1"))
		}
		if os.Getenv("AWS_ACCESS_KEY_ID") == "" {
			optFns = append(optFns, config.WithCredentialsProvider(
				credentials.NewStaticCredentialsProvider("local", "local", "")))
		}
	}
	cfg, err := config.LoadDefaultConfig(ctx, optFns...)
	if err != nil {
		return nil, fmt.Errorf("load AWS config: %w", err)
	}

	client := dynamodb.NewFromConfig(cfg, func(o *dynamodb.Options) {
		if endpoint != "" {
			o.BaseEndpoint = aws.String(endpoint)
		}
	})

	s := &Store{client: client, table: table}
	if endpoint != "" {
		if err := s.ensureTable(ctx); err != nil {
			return nil, err
		}
	}
	return s, nil
}

// ensureTable creates the table when it does not exist (local endpoint only —
// in AWS the table is provisioned by the infrastructure, never by the app).
func (s *Store) ensureTable(ctx context.Context) error {
	_, err := s.client.DescribeTable(ctx, &dynamodb.DescribeTableInput{TableName: aws.String(s.table)})
	if err == nil {
		return nil
	}
	var notFound *types.ResourceNotFoundException
	if !errors.As(err, &notFound) {
		return fmt.Errorf("describe table %s: %w", s.table, err)
	}

	_, err = s.client.CreateTable(ctx, &dynamodb.CreateTableInput{
		TableName:   aws.String(s.table),
		BillingMode: types.BillingModePayPerRequest,
		AttributeDefinitions: []types.AttributeDefinition{
			{AttributeName: aws.String("PK"), AttributeType: types.ScalarAttributeTypeS},
			{AttributeName: aws.String("SK"), AttributeType: types.ScalarAttributeTypeS},
		},
		KeySchema: []types.KeySchemaElement{
			{AttributeName: aws.String("PK"), KeyType: types.KeyTypeHash},
			{AttributeName: aws.String("SK"), KeyType: types.KeyTypeRange},
		},
	})
	if err != nil {
		return fmt.Errorf("create table %s: %w", s.table, err)
	}
	waiter := dynamodb.NewTableExistsWaiter(s.client)
	if err := waiter.Wait(ctx, &dynamodb.DescribeTableInput{TableName: aws.String(s.table)}, 30*time.Second); err != nil {
		return fmt.Errorf("wait for table %s: %w", s.table, err)
	}
	return nil
}

func (s *Store) Users() domain.UserRepository         { return &userRepo{s} }
func (s *Store) Questions() domain.QuestionRepository { return &questionRepo{s} }
func (s *Store) Games() domain.GameRepository         { return &gameRepo{s} }
func (s *Store) Stats() domain.StatsRepository        { return &statsRepo{s} }

func (s *Store) Close() error { return nil }

// ResetUserData deletes the user's sessions (with their answers and question
// history) and study recommendations. The account, its uniqueness markers and
// the auth token are kept, matching the SQLite adapter.
func (s *Store) ResetUserData(userID uint) error {
	ctx := context.Background()

	var deletes []map[string]types.AttributeValue
	addDelete := func(pk, sk string) {
		deletes = append(deletes, map[string]types.AttributeValue{
			"PK": &types.AttributeValueMemberS{Value: pk},
			"SK": &types.AttributeValueMemberS{Value: sk},
		})
	}

	sessions, err := s.queryPrefix(ctx, pkUser(userID), prefixSession)
	if err != nil {
		return err
	}
	for _, item := range sessions {
		var session sessionItem
		if err := attributevalue.UnmarshalMap(item, &session); err != nil {
			return fmt.Errorf("unmarshal session: %w", err)
		}
		// The whole SESSION#<id> partition: pointer, answers, history.
		partition, err := s.queryPartition(ctx, pkSession(session.ID))
		if err != nil {
			return err
		}
		for _, sub := range partition {
			addDelete(pkSession(session.ID), sub["SK"].(*types.AttributeValueMemberS).Value)
		}
		addDelete(pkUser(userID), session.SK)
	}

	recs, err := s.queryPrefix(ctx, pkUser(userID), prefixRec)
	if err != nil {
		return err
	}
	for _, item := range recs {
		addDelete(pkUser(userID), item["SK"].(*types.AttributeValueMemberS).Value)
	}

	return s.batchDelete(ctx, deletes)
}

// batchDelete removes keys in chunks of 25, retrying unprocessed items.
func (s *Store) batchDelete(ctx context.Context, keys []map[string]types.AttributeValue) error {
	for start := 0; start < len(keys); start += 25 {
		chunk := keys[start:min(start+25, len(keys))]
		requests := make([]types.WriteRequest, 0, len(chunk))
		for _, key := range chunk {
			requests = append(requests, types.WriteRequest{
				DeleteRequest: &types.DeleteRequest{Key: key},
			})
		}
		if err := s.batchWrite(ctx, requests); err != nil {
			return err
		}
	}
	return nil
}

// batchWrite issues one BatchWriteItem call (max 25 requests), retrying
// unprocessed items with backoff.
func (s *Store) batchWrite(ctx context.Context, requests []types.WriteRequest) error {
	pending := map[string][]types.WriteRequest{s.table: requests}
	for attempt := 0; len(pending[s.table]) > 0; attempt++ {
		if attempt > 8 {
			return fmt.Errorf("batch write: %d items still unprocessed after retries", len(pending[s.table]))
		}
		if attempt > 0 {
			time.Sleep(time.Duration(50*(1<<attempt)) * time.Millisecond)
		}
		out, err := s.client.BatchWriteItem(ctx, &dynamodb.BatchWriteItemInput{RequestItems: pending})
		if err != nil {
			return fmt.Errorf("batch write: %w", err)
		}
		pending = out.UnprocessedItems
	}
	return nil
}

// getItem fetches one item (strongly consistent) into out.
// Returns domain.ErrNotFound when the item does not exist.
func (s *Store) getItem(ctx context.Context, pk, sk string, out any) error {
	res, err := s.client.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(s.table),
		Key: map[string]types.AttributeValue{
			"PK": &types.AttributeValueMemberS{Value: pk},
			"SK": &types.AttributeValueMemberS{Value: sk},
		},
		ConsistentRead: aws.Bool(true),
	})
	if err != nil {
		return fmt.Errorf("get %s/%s: %w", pk, sk, err)
	}
	if len(res.Item) == 0 {
		return domain.ErrNotFound
	}
	if err := attributevalue.UnmarshalMap(res.Item, out); err != nil {
		return fmt.Errorf("unmarshal %s/%s: %w", pk, sk, err)
	}
	return nil
}

// putItem marshals and writes one item.
func (s *Store) putItem(ctx context.Context, item any) error {
	av, err := attributevalue.MarshalMap(item)
	if err != nil {
		return fmt.Errorf("marshal item: %w", err)
	}
	if _, err := s.client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(s.table),
		Item:      av,
	}); err != nil {
		return fmt.Errorf("put item: %w", err)
	}
	return nil
}
