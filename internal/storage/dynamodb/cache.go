package dynamodb

import (
	"context"
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

	"quiz-app/internal/domain"
)

// bank is an in-memory snapshot of the taxonomy and question bank. The bank
// is small (~750 questions, ~1 MB) and immutable at runtime — it only changes
// through SyncQuestionBank, which invalidates the cache — so every
// QuestionRepository read is served from here. The snapshot persists across
// warm Lambda invocations.
type bank struct {
	questions  []domain.Question // sorted by ID, fully hydrated
	byID       map[uint]*domain.Question
	byKey      map[string]*domain.Question
	categories []domain.Category // sorted by ID, subcategories loaded
	catCodes   map[string]bool
}

type bankCache struct {
	mu sync.Mutex
	b  *bank
}

// bank returns the cached snapshot, loading it on first use.
func (s *Store) bank(ctx context.Context) (*bank, error) {
	s.cache.mu.Lock()
	defer s.cache.mu.Unlock()
	if s.cache.b != nil {
		return s.cache.b, nil
	}
	b, err := s.loadBank(ctx)
	if err != nil {
		return nil, err
	}
	s.cache.b = b
	return b, nil
}

func (s *Store) invalidateBank() {
	s.cache.mu.Lock()
	s.cache.b = nil
	s.cache.mu.Unlock()
}

// queryPartition returns every item of one partition (paginated, consistent).
func (s *Store) queryPartition(ctx context.Context, pk string) ([]map[string]types.AttributeValue, error) {
	var items []map[string]types.AttributeValue
	var startKey map[string]types.AttributeValue
	for {
		out, err := s.client.Query(ctx, &dynamodb.QueryInput{
			TableName:              aws.String(s.table),
			KeyConditionExpression: aws.String("PK = :pk"),
			ExpressionAttributeValues: map[string]types.AttributeValue{
				":pk": &types.AttributeValueMemberS{Value: pk},
			},
			ConsistentRead:    aws.Bool(true),
			ExclusiveStartKey: startKey,
		})
		if err != nil {
			return nil, fmt.Errorf("query %s: %w", pk, err)
		}
		items = append(items, out.Items...)
		if out.LastEvaluatedKey == nil {
			break
		}
		startKey = out.LastEvaluatedKey
	}
	return items, nil
}

// queryPrefix returns every item of a partition whose SK begins with prefix.
func (s *Store) queryPrefix(ctx context.Context, pk, prefix string) ([]map[string]types.AttributeValue, error) {
	var items []map[string]types.AttributeValue
	var startKey map[string]types.AttributeValue
	for {
		out, err := s.client.Query(ctx, &dynamodb.QueryInput{
			TableName:              aws.String(s.table),
			KeyConditionExpression: aws.String("PK = :pk AND begins_with(SK, :prefix)"),
			ExpressionAttributeValues: map[string]types.AttributeValue{
				":pk":     &types.AttributeValueMemberS{Value: pk},
				":prefix": &types.AttributeValueMemberS{Value: prefix},
			},
			ConsistentRead:    aws.Bool(true),
			ExclusiveStartKey: startKey,
		})
		if err != nil {
			return nil, fmt.Errorf("query %s %s*: %w", pk, prefix, err)
		}
		items = append(items, out.Items...)
		if out.LastEvaluatedKey == nil {
			break
		}
		startKey = out.LastEvaluatedKey
	}
	return items, nil
}

func (s *Store) loadTaxonomy(ctx context.Context) ([]categoryItem, []subCategoryItem, error) {
	items, err := s.queryPartition(ctx, pkTaxonomy)
	if err != nil {
		return nil, nil, err
	}
	var cats []categoryItem
	var subs []subCategoryItem
	for _, item := range items {
		sk := item["SK"].(*types.AttributeValueMemberS).Value
		if strings.Contains(sk, "#SUB#") {
			var sub subCategoryItem
			if err := attributevalue.UnmarshalMap(item, &sub); err != nil {
				return nil, nil, fmt.Errorf("unmarshal subcategory %s: %w", sk, err)
			}
			subs = append(subs, sub)
		} else {
			var cat categoryItem
			if err := attributevalue.UnmarshalMap(item, &cat); err != nil {
				return nil, nil, fmt.Errorf("unmarshal category %s: %w", sk, err)
			}
			cats = append(cats, cat)
		}
	}
	return cats, subs, nil
}

func (s *Store) loadQuestionItems(ctx context.Context) ([]questionItem, error) {
	items, err := s.queryPartition(ctx, pkQBank)
	if err != nil {
		return nil, err
	}
	questions := make([]questionItem, 0, len(items))
	for _, item := range items {
		var q questionItem
		if err := attributevalue.UnmarshalMap(item, &q); err != nil {
			return nil, fmt.Errorf("unmarshal question: %w", err)
		}
		questions = append(questions, q)
	}
	return questions, nil
}

func (s *Store) loadBank(ctx context.Context) (*bank, error) {
	catItems, subItems, err := s.loadTaxonomy(ctx)
	if err != nil {
		return nil, err
	}
	qItems, err := s.loadQuestionItems(ctx)
	if err != nil {
		return nil, err
	}

	catsByCode := make(map[string]domain.Category, len(catItems))
	for _, c := range catItems {
		catsByCode[c.Code] = domain.Category{ID: c.ID, Code: c.Code, Name: c.Name}
	}
	subsByCatAndCode := make(map[string]domain.SubCategory, len(subItems))
	subsByCat := make(map[string][]domain.SubCategory)
	for _, sub := range subItems {
		catCode := strings.TrimPrefix(strings.SplitN(sub.SK, "#SUB#", 2)[0], prefixCat)
		sc := domain.SubCategory{ID: sub.ID, CategoryID: sub.CategoryID, Code: sub.Code, Name: sub.Name}
		subsByCatAndCode[catCode+"/"+sub.Code] = sc
		subsByCat[catCode] = append(subsByCat[catCode], sc)
	}

	b := &bank{
		byID:     make(map[uint]*domain.Question, len(qItems)),
		byKey:    make(map[string]*domain.Question, len(qItems)),
		catCodes: make(map[string]bool, len(catsByCode)),
	}

	for code, cat := range catsByCode {
		subs := append([]domain.SubCategory(nil), subsByCat[code]...)
		sort.Slice(subs, func(i, j int) bool { return subs[i].ID < subs[j].ID })
		cat.SubCategories = subs
		b.categories = append(b.categories, cat)
		b.catCodes[code] = true
	}
	sort.Slice(b.categories, func(i, j int) bool { return b.categories[i].ID < b.categories[j].ID })

	b.questions = make([]domain.Question, 0, len(qItems))
	for _, item := range qItems {
		cat := catsByCode[item.CategoryCode]
		sub := subsByCatAndCode[item.CategoryCode+"/"+item.SubCategoryCode]
		q := domain.Question{
			ID: item.ID, Key: item.Key,
			CategoryID: cat.ID, Category: domain.Category{ID: cat.ID, Code: cat.Code, Name: cat.Name},
			SubCategoryID: sub.ID, SubCategory: sub,
			Text: item.Text, Difficulty: item.Difficulty, Points: item.Points,
			Hint: item.Hint, Explanation: item.Explanation, CreatedAt: item.CreatedAt,
		}
		choices := append([]choiceItem(nil), item.Choices...)
		sort.Slice(choices, func(i, j int) bool { return choices[i].Order < choices[j].Order })
		for _, c := range choices {
			q.Choices = append(q.Choices, domain.Choice{
				ID: c.ID, QuestionID: q.ID, Text: c.Text, IsCorrect: c.IsCorrect, Order: c.Order,
			})
		}
		b.questions = append(b.questions, q)
	}
	sort.Slice(b.questions, func(i, j int) bool { return b.questions[i].ID < b.questions[j].ID })
	for i := range b.questions {
		q := &b.questions[i]
		b.byID[q.ID] = q
		b.byKey[q.Key] = q
	}
	return b, nil
}

// copyQuestion returns a deep copy. Callers of the port mutate results (the
// HTTP layer shuffles choices and strips IsCorrect), so the cache must never
// hand out its internal structs.
func copyQuestion(q *domain.Question) domain.Question {
	out := *q
	out.Choices = append([]domain.Choice(nil), q.Choices...)
	out.Category.SubCategories = append([]domain.SubCategory(nil), q.Category.SubCategories...)
	return out
}

func copyCategory(c domain.Category) domain.Category {
	out := c
	out.SubCategories = append([]domain.SubCategory(nil), c.SubCategories...)
	return out
}
