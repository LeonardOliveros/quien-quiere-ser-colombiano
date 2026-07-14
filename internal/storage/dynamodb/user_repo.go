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
// the UNIQUE constraint of the SQLite schema. Guests skip the marker (they
// never log in by name and their random usernames make collisions negligible)
// and their profile carries a TTL so the account self-destructs after
// domain.GuestDataTTL of inactivity.
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

	item := userItem{
		PK: pkUser(id), SK: skProfile,
		ID: id, Username: user.Username, Password: user.Password,
		Token: user.Token, TokenExpiresAt: user.TokenExpiresAt,
		IsGuest: user.IsGuest, CreatedAt: now, UpdatedAt: now,
	}
	if user.IsGuest {
		item.TTL = guestTTL(now)
	}
	profile, err := attributevalue.MarshalMap(item)
	if err != nil {
		return fmt.Errorf("marshal user: %w", err)
	}

	notExists := aws.String("attribute_not_exists(PK)")
	if user.IsGuest {
		_, err = r.s.client.PutItem(ctx, &dynamodb.PutItemInput{
			TableName: aws.String(r.s.table), Item: profile, ConditionExpression: notExists,
		})
		if err != nil {
			return fmt.Errorf("create guest %q: %w", user.Username, err)
		}
		return nil
	}

	usernameMarker, err := attributevalue.MarshalMap(uniqItem{PK: pkUniqUsername(user.Username), SK: skUniq, UserID: id})
	if err != nil {
		return err
	}

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
	// eventually garbage-collects them. Guest tokens expire without grace —
	// the whole guest account dies with them — and the profile's TTL is
	// extended to match the new token's lifetime.
	tokenTTL := expiresAt.Add(7 * 24 * time.Hour).Unix()
	profileUpdate := "SET #token = :token, token_expires_at = :exp, updated_at = :now"
	updateValues := map[string]types.AttributeValue{
		":token": &types.AttributeValueMemberS{Value: token},
		":exp":   mustMarshalTime(expiresAt),
		":now":   mustMarshalTime(time.Now()),
	}
	if profile.IsGuest {
		tokenTTL = expiresAt.Unix()
		profileUpdate += ", #ttl = :ttl"
		updateValues[":ttl"] = &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", tokenTTL)}
	}

	newToken, err := attributevalue.MarshalMap(tokenItem{
		PK: pkToken(token), SK: skMeta,
		UserID: userID, ExpiresAt: expiresAt,
		TTL: tokenTTL,
	})
	if err != nil {
		return fmt.Errorf("marshal token: %w", err)
	}

	names := map[string]string{"#token": "token"}
	if profile.IsGuest {
		names["#ttl"] = "ttl"
	}
	items := []types.TransactWriteItem{
		{Update: &types.Update{
			TableName: aws.String(r.s.table),
			Key: map[string]types.AttributeValue{
				"PK": &types.AttributeValueMemberS{Value: pkUser(userID)},
				"SK": &types.AttributeValueMemberS{Value: skProfile},
			},
			UpdateExpression:          aws.String(profileUpdate),
			ExpressionAttributeNames:  names,
			ExpressionAttributeValues: updateValues,
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

// TouchGuest extends the guest's profile and token expiry after activity.
// The two updates are deliberately not transactional: a partial failure only
// shortens one item's grace period, and the next touch heals it.
func (r *userRepo) TouchGuest(userID uint, token string, expiresAt time.Time) error {
	ctx := context.Background()
	now := time.Now()
	ttl := &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", expiresAt.Unix())}

	_, err := r.s.client.UpdateItem(ctx, &dynamodb.UpdateItemInput{
		TableName: aws.String(r.s.table),
		Key: map[string]types.AttributeValue{
			"PK": &types.AttributeValueMemberS{Value: pkUser(userID)},
			"SK": &types.AttributeValueMemberS{Value: skProfile},
		},
		// Condition keeps the touch from resurrecting a profile that TTL
		// already deleted (the update would otherwise create a stub item).
		ConditionExpression: aws.String("attribute_exists(PK)"),
		UpdateExpression:    aws.String("SET #ttl = :ttl, token_expires_at = :exp, last_activity_at = :now, updated_at = :now"),
		ExpressionAttributeNames: map[string]string{
			"#ttl": "ttl",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":ttl": ttl,
			":exp": mustMarshalTime(expiresAt),
			":now": mustMarshalTime(now),
		},
	})
	if err != nil {
		return fmt.Errorf("touch guest %d: %w", userID, err)
	}

	_, err = r.s.client.UpdateItem(ctx, &dynamodb.UpdateItemInput{
		TableName: aws.String(r.s.table),
		Key: map[string]types.AttributeValue{
			"PK": &types.AttributeValueMemberS{Value: pkToken(token)},
			"SK": &types.AttributeValueMemberS{Value: skMeta},
		},
		ConditionExpression: aws.String("attribute_exists(PK)"),
		UpdateExpression:    aws.String("SET #ttl = :ttl, expires_at = :exp"),
		ExpressionAttributeNames: map[string]string{
			"#ttl": "ttl",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":ttl": ttl,
			":exp": mustMarshalTime(expiresAt),
		},
	})
	if err != nil {
		return fmt.Errorf("touch guest token: %w", err)
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
