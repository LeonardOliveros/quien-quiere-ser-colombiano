package dynamodb

import (
	"context"
	"math/rand/v2"
	"sort"

	"quiz-app/internal/domain"
)

// questionRepo serves every read from the in-memory bank snapshot (cache.go).
// All results are deep copies — callers mutate them.
type questionRepo struct{ s *Store }

func (r *questionRepo) ByID(id uint) (domain.Question, error) {
	b, err := r.s.bank(context.Background())
	if err != nil {
		return domain.Question{}, err
	}
	q, ok := b.byID[id]
	if !ok {
		return domain.Question{}, domain.ErrNotFound
	}
	return copyQuestion(q), nil
}

func (r *questionRepo) List() ([]domain.Question, error) {
	b, err := r.s.bank(context.Background())
	if err != nil {
		return nil, err
	}
	out := make([]domain.Question, 0, len(b.questions))
	for i := range b.questions {
		out = append(out, copyQuestion(&b.questions[i]))
	}
	return out, nil
}

func (r *questionRepo) ListByCategory(categoryCode string) ([]domain.Question, error) {
	b, err := r.s.bank(context.Background())
	if err != nil {
		return nil, err
	}
	var out []domain.Question
	for i := range b.questions {
		if b.questions[i].Category.Code == categoryCode {
			out = append(out, copyQuestion(&b.questions[i]))
		}
	}
	return out, nil
}

func (r *questionRepo) RandomID(categoryCodes []string, excludeIDs []uint) (uint, error) {
	b, err := r.s.bank(context.Background())
	if err != nil {
		return 0, err
	}
	allowed := make(map[string]bool, len(categoryCodes))
	for _, code := range categoryCodes {
		allowed[code] = true
	}
	excluded := make(map[uint]bool, len(excludeIDs))
	for _, id := range excludeIDs {
		excluded[id] = true
	}
	var candidates []uint
	for i := range b.questions {
		q := &b.questions[i]
		if len(allowed) > 0 && !allowed[q.Category.Code] {
			continue
		}
		if excluded[q.ID] {
			continue
		}
		candidates = append(candidates, q.ID)
	}
	if len(candidates) == 0 {
		return 0, domain.ErrNotFound
	}
	return candidates[rand.IntN(len(candidates))], nil
}

func (r *questionRepo) Count(categoryCodes []string) (int64, error) {
	b, err := r.s.bank(context.Background())
	if err != nil {
		return 0, err
	}
	if len(categoryCodes) == 0 {
		return int64(len(b.questions)), nil
	}
	allowed := make(map[string]bool, len(categoryCodes))
	for _, code := range categoryCodes {
		allowed[code] = true
	}
	var count int64
	for i := range b.questions {
		if allowed[b.questions[i].Category.Code] {
			count++
		}
	}
	return count, nil
}

func (r *questionRepo) CountsByCategory() (map[string]int64, error) {
	b, err := r.s.bank(context.Background())
	if err != nil {
		return nil, err
	}
	counts := make(map[string]int64)
	for i := range b.questions {
		counts[b.questions[i].Category.Code]++
	}
	return counts, nil
}

func (r *questionRepo) CountsBySubcategory() ([]domain.SubcategoryCount, error) {
	b, err := r.s.bank(context.Background())
	if err != nil {
		return nil, err
	}
	type key struct{ cat, sub string }
	counts := make(map[key]int64)
	for i := range b.questions {
		q := &b.questions[i]
		// Category code + subcategory display name, like the SQLite adapter.
		counts[key{q.Category.Code, q.SubCategory.Name}]++
	}
	rows := make([]domain.SubcategoryCount, 0, len(counts))
	for k, n := range counts {
		rows = append(rows, domain.SubcategoryCount{Category: k.cat, SubCategory: k.sub, Count: n})
	}
	sort.Slice(rows, func(i, j int) bool {
		if rows[i].Category != rows[j].Category {
			return rows[i].Category < rows[j].Category
		}
		return rows[i].SubCategory < rows[j].SubCategory
	})
	return rows, nil
}

func (r *questionRepo) ChoiceByID(id uint) (domain.Choice, error) {
	b, err := r.s.bank(context.Background())
	if err != nil {
		return domain.Choice{}, err
	}
	questionID, _ := splitChoiceID(id)
	q, ok := b.byID[questionID]
	if !ok {
		return domain.Choice{}, domain.ErrNotFound
	}
	for _, c := range q.Choices {
		if c.ID == id {
			return c, nil
		}
	}
	return domain.Choice{}, domain.ErrNotFound
}

func (r *questionRepo) CorrectChoices(questionIDs []uint) (map[uint]domain.Choice, error) {
	b, err := r.s.bank(context.Background())
	if err != nil {
		return nil, err
	}
	correct := make(map[uint]domain.Choice, len(questionIDs))
	for _, id := range questionIDs {
		q, ok := b.byID[id]
		if !ok {
			continue
		}
		for _, c := range q.Choices {
			if c.IsCorrect {
				correct[id] = c
				break
			}
		}
	}
	return correct, nil
}

func (r *questionRepo) Categories() ([]domain.Category, error) {
	b, err := r.s.bank(context.Background())
	if err != nil {
		return nil, err
	}
	out := make([]domain.Category, 0, len(b.categories))
	for _, c := range b.categories {
		out = append(out, copyCategory(c))
	}
	return out, nil
}
