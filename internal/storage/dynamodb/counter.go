package dynamodb

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// nextID atomically allocates one ID for the entity kind.
func (s *Store) nextID(ctx context.Context, entity string) (uint, error) {
	return s.nextIDs(ctx, entity, 1)
}

// nextIDs atomically allocates a contiguous range of count IDs and returns the
// first one. Used by the seeder to avoid one counter round-trip per question.
func (s *Store) nextIDs(ctx context.Context, entity string, count int) (uint, error) {
	if count < 1 {
		return 0, fmt.Errorf("nextIDs(%s): count must be >= 1, got %d", entity, count)
	}
	out, err := s.client.UpdateItem(ctx, &dynamodb.UpdateItemInput{
		TableName: aws.String(s.table),
		Key: map[string]types.AttributeValue{
			"PK": &types.AttributeValueMemberS{Value: pkCounter(entity)},
			"SK": &types.AttributeValueMemberS{Value: skMeta},
		},
		UpdateExpression: aws.String("ADD #n :c"),
		ExpressionAttributeNames: map[string]string{
			"#n": "n",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":c": &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", count)},
		},
		ReturnValues: types.ReturnValueUpdatedNew,
	})
	if err != nil {
		return 0, fmt.Errorf("allocate %s id: %w", entity, err)
	}
	nAttr, ok := out.Attributes["n"].(*types.AttributeValueMemberN)
	if !ok {
		return 0, fmt.Errorf("allocate %s id: unexpected counter attribute %T", entity, out.Attributes["n"])
	}
	var last uint64
	if _, err := fmt.Sscanf(nAttr.Value, "%d", &last); err != nil {
		return 0, fmt.Errorf("allocate %s id: parse counter %q: %w", entity, nAttr.Value, err)
	}
	return uint(last) - uint(count) + 1, nil
}
