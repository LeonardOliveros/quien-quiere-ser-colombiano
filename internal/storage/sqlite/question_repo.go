package sqlite

import (
	"gorm.io/gorm"

	"quiz-app/internal/domain"
)

type questionRepo struct {
	db *gorm.DB
}

// loaded returns a query with the relations every Question result must carry.
func (r *questionRepo) loaded() *gorm.DB {
	return r.db.Preload("Choices").Preload("Category").Preload("SubCategory")
}

// byCategoryCodes joins categories to filter questions by category code.
func byCategoryCodes(q *gorm.DB, codes []string) *gorm.DB {
	if len(codes) == 0 {
		return q
	}
	return q.Joins("JOIN categories ON categories.id = questions.category_id").
		Where("categories.code IN ?", codes)
}

func (r *questionRepo) ByID(id uint) (domain.Question, error) {
	var question domain.Question
	err := r.loaded().First(&question, id).Error
	return question, translate(err)
}

func (r *questionRepo) List() ([]domain.Question, error) {
	var questions []domain.Question
	err := r.loaded().Find(&questions).Error
	return questions, err
}

func (r *questionRepo) ListByCategory(categoryCode string) ([]domain.Question, error) {
	var questions []domain.Question
	err := byCategoryCodes(r.loaded(), []string{categoryCode}).Find(&questions).Error
	return questions, err
}

func (r *questionRepo) RandomID(categoryCodes []string, excludeIDs []uint) (uint, error) {
	query := byCategoryCodes(r.db.Model(&domain.Question{}), categoryCodes)
	if len(excludeIDs) > 0 {
		query = query.Where("questions.id NOT IN ?", excludeIDs)
	}
	var id uint
	if err := query.Order("RANDOM()").Limit(1).Pluck("questions.id", &id).Error; err != nil {
		return 0, err
	}
	if id == 0 {
		return 0, domain.ErrNotFound
	}
	return id, nil
}

func (r *questionRepo) Count(categoryCodes []string) (int64, error) {
	var count int64
	err := byCategoryCodes(r.db.Model(&domain.Question{}), categoryCodes).Count(&count).Error
	return count, err
}

func (r *questionRepo) CountsByCategory() (map[string]int64, error) {
	var rows []struct {
		Category string
		Count    int64
	}
	err := r.db.Model(&domain.Question{}).
		Select("categories.code as category, COUNT(*) as count").
		Joins("JOIN categories ON categories.id = questions.category_id").
		Group("categories.code").
		Scan(&rows).Error
	if err != nil {
		return nil, err
	}
	counts := make(map[string]int64, len(rows))
	for _, row := range rows {
		counts[row.Category] = row.Count
	}
	return counts, nil
}

func (r *questionRepo) CountsBySubcategory() ([]domain.SubcategoryCount, error) {
	var rows []domain.SubcategoryCount
	err := r.db.Model(&domain.Question{}).
		Select("categories.code as category, sub_categories.name as sub_category, COUNT(*) as count").
		Joins("JOIN categories ON categories.id = questions.category_id").
		Joins("JOIN sub_categories ON sub_categories.id = questions.sub_category_id").
		Group("categories.code, sub_categories.name").
		Scan(&rows).Error
	return rows, err
}

func (r *questionRepo) ChoiceByID(id uint) (domain.Choice, error) {
	var choice domain.Choice
	err := r.db.First(&choice, id).Error
	return choice, translate(err)
}

func (r *questionRepo) CorrectChoices(questionIDs []uint) (map[uint]domain.Choice, error) {
	correct := make(map[uint]domain.Choice, len(questionIDs))
	if len(questionIDs) == 0 {
		return correct, nil
	}
	var choices []domain.Choice
	if err := r.db.Where("question_id IN ? AND is_correct = true", questionIDs).Find(&choices).Error; err != nil {
		return nil, err
	}
	for _, choice := range choices {
		correct[choice.QuestionID] = choice
	}
	return correct, nil
}

func (r *questionRepo) Categories() ([]domain.Category, error) {
	var categories []domain.Category
	err := r.db.Preload("SubCategories").Order("id").Find(&categories).Error
	return categories, err
}
