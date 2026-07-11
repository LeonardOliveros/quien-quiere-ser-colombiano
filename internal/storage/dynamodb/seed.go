package dynamodb

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

	"quiz-app/internal/domain"
)

// SyncQuestionBank idempotently upserts the taxonomy and question bank:
// categories/subcategories keyed by code, questions keyed by Key. IDs and
// created_at of existing records are preserved; choice IDs are derived
// (question*100+order) and therefore stable across syncs. Nothing is ever
// deleted — questions present in the table but missing from the seed files
// are logged and kept, like the SQLite adapter.
func (s *Store) SyncQuestionBank(taxonomy []domain.SeedCategory, seeds []domain.SeedQuestion) error {
	ctx := context.Background()

	existingCats, existingSubs, err := s.loadTaxonomy(ctx)
	if err != nil {
		return err
	}
	existingQuestions, err := s.loadQuestionItems(ctx)
	if err != nil {
		return err
	}

	catByCode := make(map[string]categoryItem, len(existingCats))
	for _, c := range existingCats {
		catByCode[c.Code] = c
	}
	subBySK := make(map[string]subCategoryItem, len(existingSubs))
	for _, sub := range existingSubs {
		subBySK[sub.SK] = sub
	}
	questionByKey := make(map[string]questionItem, len(existingQuestions))
	for _, q := range existingQuestions {
		questionByKey[q.Key] = q
	}

	// Bulk-allocate IDs for the new records (one counter round-trip per kind).
	newCats, newSubs, newQuestions := 0, 0, 0
	for _, cat := range taxonomy {
		if _, ok := catByCode[cat.Code]; !ok {
			newCats++
		}
		for _, sub := range cat.SubCategories {
			if _, ok := subBySK[skSubCategory(cat.Code, sub.Code)]; !ok {
				newSubs++
			}
		}
	}
	for _, seed := range seeds {
		if _, ok := questionByKey[seed.Key]; !ok {
			newQuestions++
		}
	}
	nextCatID, err := s.allocate(ctx, "category", newCats)
	if err != nil {
		return err
	}
	nextSubID, err := s.allocate(ctx, "subcategory", newSubs)
	if err != nil {
		return err
	}
	nextQuestionID, err := s.allocate(ctx, "question", newQuestions)
	if err != nil {
		return err
	}

	var writes []types.WriteRequest
	putIfChanged := func(existing any, exists bool, want any) error {
		if exists && reflect.DeepEqual(existing, want) {
			return nil
		}
		av, err := attributevalue.MarshalMap(want)
		if err != nil {
			return fmt.Errorf("marshal seed item: %w", err)
		}
		writes = append(writes, types.WriteRequest{PutRequest: &types.PutRequest{Item: av}})
		return nil
	}

	catIDs := make(map[string]uint, len(taxonomy))
	for _, cat := range taxonomy {
		current, exists := catByCode[cat.Code]
		id := current.ID
		if !exists {
			id = nextCatID
			nextCatID++
		}
		catIDs[cat.Code] = id
		want := categoryItem{PK: pkTaxonomy, SK: skCategory(cat.Code), ID: id, Code: cat.Code, Name: cat.Name}
		if err := putIfChanged(current, exists, want); err != nil {
			return err
		}

		for _, sub := range cat.SubCategories {
			sk := skSubCategory(cat.Code, sub.Code)
			currentSub, subExists := subBySK[sk]
			subID := currentSub.ID
			if !subExists {
				subID = nextSubID
				nextSubID++
			}
			wantSub := subCategoryItem{PK: pkTaxonomy, SK: sk, ID: subID, CategoryID: id, Code: sub.Code, Name: sub.Name}
			if err := putIfChanged(currentSub, subExists, wantSub); err != nil {
				return err
			}
		}
	}

	created, updated := 0, 0
	for _, seed := range seeds {
		current, exists := questionByKey[seed.Key]
		id := current.ID
		createdAt := current.CreatedAt
		if !exists {
			id = nextQuestionID
			nextQuestionID++
			createdAt = time.Now()
		}

		choices := make([]choiceItem, 0, len(seed.Choices))
		for _, c := range seed.Choices {
			if c.Order <= 0 || c.Order >= choiceIDFactor {
				return fmt.Errorf("question %s: choice order %d out of range [1,%d)", seed.Key, c.Order, choiceIDFactor)
			}
			choices = append(choices, choiceItem{
				ID: choiceID(id, c.Order), Text: c.Text, IsCorrect: c.IsCorrect, Order: c.Order,
			})
		}

		want := questionItem{
			PK: pkQBank, SK: skQuestion(seed.Key),
			ID: id, Key: seed.Key,
			CategoryCode: seed.Category, SubCategoryCode: seed.SubCategory,
			Text: seed.Text, Difficulty: seed.Difficulty, Points: seed.Points,
			Hint: seed.Hint, Explanation: seed.Explanation,
			CreatedAt: createdAt, Choices: choices,
		}
		if !exists {
			created++
		} else if !reflect.DeepEqual(current, want) {
			updated++
		}
		if err := putIfChanged(current, exists, want); err != nil {
			return err
		}
	}

	for start := 0; start < len(writes); start += 25 {
		if err := s.batchWrite(ctx, writes[start:min(start+25, len(writes))]); err != nil {
			return err
		}
	}

	seedKeys := make(map[string]bool, len(seeds))
	for _, seed := range seeds {
		seedKeys[seed.Key] = true
	}
	orphans := 0
	for key := range questionByKey {
		if !seedKeys[key] {
			orphans++
			log.Printf("Warning: question %q exists in DB but not in seed files (kept)", key)
		}
	}

	s.invalidateBank()
	log.Printf("Seed sync: %d questions in files, %d created, %d updated, %d orphaned in DB", len(seeds), created, updated, orphans)
	return nil
}

// allocate wraps nextIDs, tolerating count == 0.
func (s *Store) allocate(ctx context.Context, entity string, count int) (uint, error) {
	if count == 0 {
		return 0, nil
	}
	return s.nextIDs(ctx, entity, count)
}
