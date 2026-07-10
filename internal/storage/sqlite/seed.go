package sqlite

import (
	"log"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"quiz-app/internal/domain"
)

// SyncQuestionBank idempotently upserts the taxonomy and question bank inside
// a transaction: categories/subcategories by code, questions by key, choices
// by (question, order). Questions present in the DB but missing from the
// files are logged and kept.
func (s *Store) SyncQuestionBank(taxonomy []domain.SeedCategory, seeds []domain.SeedQuestion) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		return syncQuestionBank(tx, taxonomy, seeds)
	})
}

func syncQuestionBank(tx *gorm.DB, taxonomy []domain.SeedCategory, seeds []domain.SeedQuestion) error {
	subIDs := make(map[string]uint) // "CATEGORY/SUBCATEGORY" -> subcategory ID
	catIDs := make(map[string]uint) // "CATEGORY" -> category ID
	for _, cat := range taxonomy {
		var category domain.Category
		if err := tx.Where(domain.Category{Code: cat.Code}).Assign(domain.Category{Name: cat.Name}).FirstOrCreate(&category, domain.Category{Code: cat.Code}).Error; err != nil {
			return err
		}
		catIDs[cat.Code] = category.ID
		for _, sub := range cat.SubCategories {
			var subCategory domain.SubCategory
			cond := domain.SubCategory{CategoryID: category.ID, Code: sub.Code}
			if err := tx.Where(cond).Assign(domain.SubCategory{Name: sub.Name}).FirstOrCreate(&subCategory, cond).Error; err != nil {
				return err
			}
			subIDs[cat.Code+"/"+sub.Code] = subCategory.ID
		}
	}

	var existing []domain.Question
	if err := tx.Preload("Choices").Find(&existing).Error; err != nil {
		return err
	}
	existingByKey := make(map[string]domain.Question, len(existing))
	for _, q := range existing {
		existingByKey[q.Key] = q
	}

	created, updated := 0, 0
	var toCreate []domain.Question
	for _, s := range seeds {
		want := domain.Question{
			Key:           s.Key,
			CategoryID:    catIDs[s.Category],
			SubCategoryID: subIDs[s.Category+"/"+s.SubCategory],
			Text:          s.Text,
			Difficulty:    s.Difficulty,
			Points:        s.Points,
			Hint:          s.Hint,
			Explanation:   s.Explanation,
		}
		current, exists := existingByKey[s.Key]
		if !exists {
			for _, c := range s.Choices {
				want.Choices = append(want.Choices, domain.Choice{Text: c.Text, IsCorrect: c.IsCorrect, Order: c.Order})
			}
			toCreate = append(toCreate, want)
			created++
			continue
		}
		if questionChanged(current, want) {
			want.ID = current.ID
			want.CreatedAt = current.CreatedAt
			if err := tx.Omit(clause.Associations).Save(&want).Error; err != nil {
				return err
			}
			updated++
		}
		if err := syncChoices(tx, current, s.Choices); err != nil {
			return err
		}
	}
	if len(toCreate) > 0 {
		if err := tx.CreateInBatches(toCreate, 100).Error; err != nil {
			return err
		}
	}

	seedKeys := make(map[string]bool, len(seeds))
	for _, s := range seeds {
		seedKeys[s.Key] = true
	}
	orphans := 0
	for key := range existingByKey {
		if !seedKeys[key] {
			orphans++
			log.Printf("Warning: question %q exists in DB but not in seed files (kept)", key)
		}
	}

	log.Printf("Seed sync: %d questions in files, %d created, %d updated, %d orphaned in DB", len(seeds), created, updated, orphans)
	return nil
}

func questionChanged(current, want domain.Question) bool {
	return current.CategoryID != want.CategoryID ||
		current.SubCategoryID != want.SubCategoryID ||
		current.Text != want.Text ||
		current.Difficulty != want.Difficulty ||
		current.Points != want.Points ||
		current.Hint != want.Hint ||
		current.Explanation != want.Explanation
}

// syncChoices upserts choices by (question_id, order) without deleting rows,
// so existing GameAnswer.ChoiceID references stay valid.
func syncChoices(tx *gorm.DB, question domain.Question, seeds []domain.SeedChoice) error {
	byOrder := make(map[int]domain.Choice, len(question.Choices))
	for _, c := range question.Choices {
		byOrder[c.Order] = c
	}
	for _, s := range seeds {
		current, exists := byOrder[s.Order]
		if !exists {
			c := domain.Choice{QuestionID: question.ID, Text: s.Text, IsCorrect: s.IsCorrect, Order: s.Order}
			if err := tx.Create(&c).Error; err != nil {
				return err
			}
			continue
		}
		if current.Text != s.Text || current.IsCorrect != s.IsCorrect {
			current.Text = s.Text
			current.IsCorrect = s.IsCorrect
			if err := tx.Save(&current).Error; err != nil {
				return err
			}
		}
	}
	return nil
}
