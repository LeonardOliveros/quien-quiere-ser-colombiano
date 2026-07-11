package dynamodb

import (
	"context"
	"fmt"
	"sort"
	"time"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"

	"quiz-app/internal/domain"
)

// statsRepo aggregates in Go over the user's own partitions. A user's
// sessions and answers stay small, so the fan-out (one Query per session)
// is cheap; if volumes ever grow, maintain running totals on a
// USER#<id>/STATS item updated from SaveAnswer.
type statsRepo struct{ s *Store }

func (r *statsRepo) games() *gameRepo { return &gameRepo{r.s} }

func (r *statsRepo) TotalSessions(userID uint) (int64, error) {
	items, err := r.games().userSessions(context.Background(), userID)
	if err != nil {
		return 0, err
	}
	return int64(len(items)), nil
}

func (r *statsRepo) BestScore(userID uint) (int, error) {
	items, err := r.games().userSessions(context.Background(), userID)
	if err != nil {
		return 0, err
	}
	best := 0
	for _, item := range items {
		if item.Status == "COMPLETED" && item.Score > best {
			best = item.Score
		}
	}
	return best, nil
}

// answerTally sums real answers (placeholders excluded) of one session,
// optionally restricted to one category code ("" = all).
func (r *statsRepo) answerTally(ctx context.Context, b *bank, sessionID uint, categoryCode string) (domain.AnswerTotals, error) {
	items, err := r.games().sessionAnswers(ctx, sessionID)
	if err != nil {
		return domain.AnswerTotals{}, err
	}
	var totals domain.AnswerTotals
	for _, item := range items {
		if item.ChoiceID == nil {
			continue
		}
		if categoryCode != "" {
			q, ok := b.byID[item.QuestionID]
			if !ok || q.Category.Code != categoryCode {
				continue
			}
		}
		totals.Total++
		if item.IsCorrect {
			totals.Correct++
		}
	}
	return totals, nil
}

func (r *statsRepo) userTotals(userID uint, categoryCode string) (domain.AnswerTotals, error) {
	ctx := context.Background()
	sessions, err := r.games().userSessions(ctx, userID)
	if err != nil {
		return domain.AnswerTotals{}, err
	}
	b, err := r.s.bank(ctx)
	if err != nil {
		return domain.AnswerTotals{}, err
	}
	var totals domain.AnswerTotals
	for _, session := range sessions {
		t, err := r.answerTally(ctx, b, session.ID, categoryCode)
		if err != nil {
			return domain.AnswerTotals{}, err
		}
		totals.Total += t.Total
		totals.Correct += t.Correct
	}
	return totals, nil
}

func (r *statsRepo) OverallTotals(userID uint) (domain.AnswerTotals, error) {
	return r.userTotals(userID, "")
}

func (r *statsRepo) CategoryTotals(userID uint, categoryCode string) (domain.AnswerTotals, error) {
	return r.userTotals(userID, categoryCode)
}

func (r *statsRepo) RecentProgress(userID uint, limit int) ([]domain.ProgressEntry, error) {
	ctx := context.Background()
	sessions, err := r.games().userSessions(ctx, userID) // newest first
	if err != nil {
		return nil, err
	}
	b, err := r.s.bank(ctx)
	if err != nil {
		return nil, err
	}
	entries := make([]domain.ProgressEntry, 0, limit)
	for _, session := range sessions {
		if session.Status != "COMPLETED" {
			continue
		}
		if limit > 0 && len(entries) == limit {
			break
		}
		tally, err := r.answerTally(ctx, b, session.ID, "")
		if err != nil {
			return nil, err
		}
		entries = append(entries, domain.ProgressEntry{
			Date:     session.CreatedAt,
			Score:    session.Score,
			Answered: tally.Total,
			Correct:  tally.Correct,
		})
	}
	return entries, nil
}

func (r *statsRepo) WeakCategories(userID uint) ([]string, error) {
	ctx := context.Background()
	sessions, err := r.games().userSessions(ctx, userID)
	if err != nil {
		return nil, err
	}
	b, err := r.s.bank(ctx)
	if err != nil {
		return nil, err
	}
	totals := make(map[string]*domain.AnswerTotals)
	for _, session := range sessions {
		items, err := r.games().sessionAnswers(ctx, session.ID)
		if err != nil {
			return nil, err
		}
		for _, item := range items {
			if item.ChoiceID == nil {
				continue
			}
			q, ok := b.byID[item.QuestionID]
			if !ok {
				continue
			}
			t := totals[q.Category.Code]
			if t == nil {
				t = &domain.AnswerTotals{}
				totals[q.Category.Code] = t
			}
			t.Total++
			if item.IsCorrect {
				t.Correct++
			}
		}
	}
	var weak []string
	for code, t := range totals {
		if float64(t.Correct)*100.0/float64(t.Total) < 50 {
			weak = append(weak, code)
		}
	}
	// Most answered first, like the reference ORDER BY COUNT(*) DESC.
	sort.Slice(weak, func(i, j int) bool {
		return totals[weak[i]].Total > totals[weak[j]].Total
	})
	return weak, nil
}

func (r *statsRepo) CreateRecommendation(rec *domain.StudyRecommendation) error {
	ctx := context.Background()
	id, err := r.s.nextID(ctx, "rec")
	if err != nil {
		return err
	}
	rec.ID = id
	rec.CreatedAt = time.Now()
	return r.s.putItem(ctx, recommendationItem{
		PK: pkUser(rec.UserID), SK: skRec(id),
		ID: id, UserID: rec.UserID, Category: rec.Category, SubCategory: rec.SubCategory,
		Weakness: rec.Weakness, Description: rec.Description, Resources: rec.Resources,
		Priority: rec.Priority, CreatedAt: rec.CreatedAt,
	})
}

func (r *statsRepo) RecommendationsByUser(userID uint, limit int) ([]domain.StudyRecommendation, error) {
	raw, err := r.s.queryPrefix(context.Background(), pkUser(userID), prefixRec)
	if err != nil {
		return nil, err
	}
	recs := make([]domain.StudyRecommendation, 0, len(raw))
	for _, itemMap := range raw {
		var item recommendationItem
		if err := attributevalue.UnmarshalMap(itemMap, &item); err != nil {
			return nil, fmt.Errorf("unmarshal recommendation: %w", err)
		}
		recs = append(recs, item.toDomain())
	}
	sort.SliceStable(recs, func(i, j int) bool { return recs[i].Priority > recs[j].Priority })
	if limit > 0 && len(recs) > limit {
		recs = recs[:limit]
	}
	return recs, nil
}
