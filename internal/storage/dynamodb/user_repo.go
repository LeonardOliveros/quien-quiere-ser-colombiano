package dynamodb

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

	"quiz-app/internal/domain"
)

type userRepo struct{ s *Store }

// Create inserts the profile plus a uniqueness marker for the username in one
// transaction; a marker collision fails the whole transaction, which mirrors
// the UNIQUE constraint of the SQLite schema.
func (r *userRepo) Create(user *domain.User) error {
	ctx := context.Background()
	id, err := r.s.nextID(ctx, "user")
	if err != nil {
		return err
	}
	now := time.Now()
	user.ID = id
	user.CreatedAt = now
	user.UpdatedAt = now

	profile, err := attributevalue.MarshalMap(userItem{
		PK: pkUser(id), SK: skProfile,
		ID: id, Username: user.Username, Password: user.Password,
		Token: user.Token, TokenExpiresAt: user.TokenExpiresAt,
		CreatedAt: now, UpdatedAt: now,
	})
	if err != nil {
		return fmt.Errorf("marshal user: %w", err)
	}

	usernameMarker, err := attributevalue.MarshalMap(uniqItem{PK: pkUniqUsername(user.Username), SK: skUniq, UserID: id})
	if err != nil {
		return err
	}

	notExists := aws.String("attribute_not_exists(PK)")
	_, err = r.s.client.TransactWriteItems(ctx, &dynamodb.TransactWriteItemsInput{
		TransactItems: []types.TransactWriteItem{
			{Put: &types.Put{TableName: aws.String(r.s.table), Item: profile, ConditionExpression: notExists}},
			{Put: &types.Put{TableName: aws.String(r.s.table), Item: usernameMarker, ConditionExpression: notExists}},
		},
	})
	if err != nil {
		return fmt.Errorf("create user %q: %w", user.Username, err)
	}
	return nil
}

func (r *userRepo) ByUsername(username string) (domain.User, error) {
	ctx := context.Background()
	var marker uniqItem
	if err := r.s.getItem(ctx, pkUniqUsername(username), skUniq, &marker); err != nil {
		return domain.User{}, err
	}
	return r.byID(ctx, marker.UserID)
}

func (r *userRepo) ByToken(token string) (domain.User, error) {
	if token == "" {
		return domain.User{}, domain.ErrNotFound
	}
	ctx := context.Background()
	var tok tokenItem
	if err := r.s.getItem(ctx, pkToken(token), skMeta, &tok); err != nil {
		return domain.User{}, err
	}
	return r.byID(ctx, tok.UserID)
}

func (r *userRepo) byID(ctx context.Context, id uint) (domain.User, error) {
	var profile userItem
	if err := r.s.getItem(ctx, pkUser(id), skProfile, &profile); err != nil {
		return domain.User{}, err
	}
	return profile.toDomain(), nil
}

// SaveSessionToken rotates the user's token: the profile is updated, a new
// TOKEN item is written and the previous TOKEN item is deleted in the same
// transaction, so the old token stops resolving immediately (the SQLite
// schema has a single token column, giving the same one-token semantics).
func (r *userRepo) SaveSessionToken(userID uint, token string, expiresAt time.Time) error {
	ctx := context.Background()
	var profile userItem
	if err := r.s.getItem(ctx, pkUser(userID), skProfile, &profile); err != nil {
		return err
	}

	// TTL is padded so the HTTP layer keeps seeing recently expired tokens
	// (and can report "expired" rather than "invalid") before DynamoDB
	// eventually garbage-collects them.
	newToken, err := attributevalue.MarshalMap(tokenItem{
		PK: pkToken(token), SK: skMeta,
		UserID: userID, ExpiresAt: expiresAt,
		TTL: expiresAt.Add(7 * 24 * time.Hour).Unix(),
	})
	if err != nil {
		return fmt.Errorf("marshal token: %w", err)
	}

	items := []types.TransactWriteItem{
		{Update: &types.Update{
			TableName: aws.String(r.s.table),
			Key: map[string]types.AttributeValue{
				"PK": &types.AttributeValueMemberS{Value: pkUser(userID)},
				"SK": &types.AttributeValueMemberS{Value: skProfile},
			},
			UpdateExpression: aws.String("SET #token = :token, token_expires_at = :exp, updated_at = :now"),
			ExpressionAttributeNames: map[string]string{
				"#token": "token",
			},
			ExpressionAttributeValues: map[string]types.AttributeValue{
				":token": &types.AttributeValueMemberS{Value: token},
				":exp":   mustMarshalTime(expiresAt),
				":now":   mustMarshalTime(time.Now()),
			},
		}},
		{Put: &types.Put{TableName: aws.String(r.s.table), Item: newToken}},
	}
	if old := profile.Token; old != "" && old != token {
		items = append(items, types.TransactWriteItem{
			Delete: &types.Delete{
				TableName: aws.String(r.s.table),
				Key: map[string]types.AttributeValue{
					"PK": &types.AttributeValueMemberS{Value: pkToken(old)},
					"SK": &types.AttributeValueMemberS{Value: skMeta},
				},
			},
		})
	}

	if _, err := r.s.client.TransactWriteItems(ctx, &dynamodb.TransactWriteItemsInput{TransactItems: items}); err != nil {
		return fmt.Errorf("save session token: %w", err)
	}
	return nil
}

func mustMarshalTime(t time.Time) types.AttributeValue {
	av, err := attributevalue.Marshal(t)
	if err != nil {
		panic(fmt.Sprintf("marshal time: %v", err))
	}
	return av
}
