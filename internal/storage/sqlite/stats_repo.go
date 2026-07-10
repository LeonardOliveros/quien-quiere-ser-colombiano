package sqlite

import (
	"time"

	"gorm.io/gorm"

	"quiz-app/internal/domain"
)

type statsRepo struct {
	db *gorm.DB
}

func (r *statsRepo) TotalSessions(userID uint) (int64, error) {
	var count int64
	err := r.db.Model(&domain.GameSession{}).Where("user_id = ?", userID).Count(&count).Error
	return count, err
}

func (r *statsRepo) BestScore(userID uint) (int, error) {
	var best int
	err := r.db.Model(&domain.GameSession{}).
		Where("user_id = ? AND status = ?", userID, "COMPLETED").
		Select("COALESCE(MAX(score), 0)").Scan(&best).Error
	return best, err
}

func (r *statsRepo) OverallTotals(userID uint) (domain.AnswerTotals, error) {
	var totals domain.AnswerTotals
	err := r.db.Raw(`
		SELECT
			COUNT(*) as total,
			COALESCE(SUM(CASE WHEN ga.is_correct THEN 1 ELSE 0 END), 0) as correct
		FROM game_answers ga
		JOIN game_sessions gs ON ga.game_session_id = gs.id
		WHERE gs.user_id = ? AND ga.choice_id IS NOT NULL
	`, userID).Scan(&totals).Error
	return totals, err
}

func (r *statsRepo) CategoryTotals(userID uint, categoryCode string) (domain.AnswerTotals, error) {
	var totals domain.AnswerTotals
	err := r.db.Raw(`
		SELECT
			COUNT(*) as total,
			COALESCE(SUM(CASE WHEN ga.is_correct THEN 1 ELSE 0 END), 0) as correct
		FROM game_answers ga
		JOIN questions q ON ga.question_id = q.id
		JOIN categories c ON q.category_id = c.id
		JOIN game_sessions gs ON ga.game_session_id = gs.id
		WHERE gs.user_id = ? AND c.code = ? AND ga.choice_id IS NOT NULL
	`, userID, categoryCode).Scan(&totals).Error
	return totals, err
}

func (r *statsRepo) RecentProgress(userID uint, limit int) ([]domain.ProgressEntry, error) {
	var rows []struct {
		CreatedAt time.Time
		Score     int
		Answered  int
		Correct   int
	}
	err := r.db.Raw(`
		SELECT
			gs.created_at,
			gs.score,
			COUNT(ga.id) as answered,
			COALESCE(SUM(CASE WHEN ga.is_correct THEN 1 ELSE 0 END), 0) as correct
		FROM game_sessions gs
		LEFT JOIN game_answers ga ON ga.game_session_id = gs.id AND ga.choice_id IS NOT NULL
		WHERE gs.user_id = ? AND gs.status = 'COMPLETED'
		GROUP BY gs.id
		ORDER BY gs.created_at DESC
		LIMIT ?
	`, userID, limit).Scan(&rows).Error
	if err != nil {
		return nil, err
	}
	entries := make([]domain.ProgressEntry, 0, len(rows))
	for _, row := range rows {
		entries = append(entries, domain.ProgressEntry{
			Date:     row.CreatedAt,
			Score:    row.Score,
			Answered: row.Answered,
			Correct:  row.Correct,
		})
	}
	return entries, nil
}

func (r *statsRepo) WeakCategories(userID uint) ([]string, error) {
	var codes []string
	err := r.db.Raw(`
		SELECT c.code as category
		FROM game_answers ga
		JOIN questions q ON ga.question_id = q.id
		JOIN categories c ON q.category_id = c.id
		JOIN game_sessions gs ON ga.game_session_id = gs.id
		WHERE gs.user_id = ? AND ga.choice_id IS NOT NULL
		GROUP BY c.code
		HAVING (SUM(CASE WHEN ga.is_correct THEN 1 ELSE 0 END) * 100.0 / COUNT(*)) < 50
		ORDER BY COUNT(*) DESC
	`, userID).Pluck("category", &codes).Error
	return codes, err
}

func (r *statsRepo) CreateRecommendation(rec *domain.StudyRecommendation) error {
	return r.db.Create(rec).Error
}

func (r *statsRepo) RecommendationsByUser(userID uint, limit int) ([]domain.StudyRecommendation, error) {
	q := r.db.Where("user_id = ?", userID).Order("priority desc")
	if limit > 0 {
		q = q.Limit(limit)
	}
	var recs []domain.StudyRecommendation
	err := q.Find(&recs).Error
	return recs, err
}
