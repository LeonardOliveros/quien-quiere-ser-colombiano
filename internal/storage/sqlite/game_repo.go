package sqlite

import (
	"gorm.io/gorm"

	"quiz-app/internal/domain"
)

type gameRepo struct {
	db *gorm.DB
}

func (r *gameRepo) CreateSession(session *domain.GameSession) error {
	return r.db.Create(session).Error
}

func (r *gameRepo) SessionByID(id uint) (domain.GameSession, error) {
	var session domain.GameSession
	err := r.db.First(&session, id).Error
	return session, translate(err)
}

func (r *gameRepo) SaveSession(session *domain.GameSession) error {
	// Save writes every field, so nil pointers (e.g. paused_at on resume)
	// are persisted back to NULL
	return r.db.Save(session).Error
}

func (r *gameRepo) AddScore(sessionID uint, points int) error {
	return r.db.Model(&domain.GameSession{}).Where("id = ?", sessionID).
		UpdateColumn("correct_answers", gorm.Expr("correct_answers + 1")).
		UpdateColumn("score", gorm.Expr("score + ?", points)).Error
}

func (r *gameRepo) SessionsByUser(userID uint, limit int) ([]domain.GameSession, error) {
	var sessions []domain.GameSession
	err := r.db.Where("user_id = ?", userID).Order("created_at desc").Limit(limit).Find(&sessions).Error
	return sessions, err
}

// resumable scopes a query to the user's ACTIVE/PAUSED sessions.
func (r *gameRepo) resumable(userID uint, mode string) *gorm.DB {
	q := r.db.Where("user_id = ? AND status IN ?", userID, []string{"ACTIVE", "PAUSED"})
	if mode != "" {
		q = q.Where("mode = ?", mode)
	}
	return q
}

func (r *gameRepo) LatestResumable(userID uint, mode string) (domain.GameSession, error) {
	var session domain.GameSession
	err := r.resumable(userID, mode).Order("updated_at DESC").First(&session).Error
	return session, translate(err)
}

func (r *gameRepo) CompleteResumable(userID uint, mode string, keepID uint) error {
	q := r.resumable(userID, mode).Model(&domain.GameSession{})
	if keepID != 0 {
		q = q.Where("id != ?", keepID)
	}
	return q.Update("status", "COMPLETED").Error
}

func (r *gameRepo) SaveAnswer(answer *domain.GameAnswer) error {
	return r.db.Save(answer).Error
}

func (r *gameRepo) AnswerPlaceholder(sessionID, questionID uint) (domain.GameAnswer, error) {
	var placeholder domain.GameAnswer
	err := r.db.Where("game_session_id = ? AND question_id = ? AND choice_id IS NULL",
		sessionID, questionID).First(&placeholder).Error
	return placeholder, translate(err)
}

func (r *gameRepo) HasAnswered(sessionID, questionID uint) (bool, error) {
	var count int64
	err := r.db.Model(&domain.GameAnswer{}).
		Where("game_session_id = ? AND question_id = ? AND choice_id IS NOT NULL", sessionID, questionID).
		Count(&count).Error
	return count > 0, err
}

func (r *gameRepo) ToggleFlag(sessionID, questionID uint) (bool, error) {
	result := r.db.Model(&domain.GameAnswer{}).
		Where("game_session_id = ? AND question_id = ?", sessionID, questionID).
		Update("is_flagged", gorm.Expr("NOT is_flagged"))
	return result.RowsAffected > 0, result.Error
}

func (r *gameRepo) AnswersBySession(sessionID uint) ([]domain.GameAnswer, error) {
	var answers []domain.GameAnswer
	err := r.db.Preload("Question.Category").Preload("Question.SubCategory").Preload("Question.Choices").
		Preload("Question").Preload("Choice").
		Where("game_session_id = ?", sessionID).Find(&answers).Error
	return answers, err
}

func (r *gameRepo) IncorrectAnswers(sessionID uint) ([]domain.GameAnswer, error) {
	var answers []domain.GameAnswer
	err := r.db.Preload("Question.Category").Preload("Question.SubCategory").Preload("Question").
		Where("game_session_id = ? AND is_correct = false", sessionID).Find(&answers).Error
	return answers, err
}

// answered scopes a query to real answers (flag placeholders excluded).
func (r *gameRepo) answered(sessionID uint) *gorm.DB {
	return r.db.Model(&domain.GameAnswer{}).
		Where("game_session_id = ? AND choice_id IS NOT NULL", sessionID)
}

func (r *gameRepo) AnsweredCount(sessionID uint) (int64, error) {
	var count int64
	err := r.answered(sessionID).Count(&count).Error
	return count, err
}

func (r *gameRepo) IncorrectCount(sessionID uint) (int64, error) {
	var count int64
	err := r.answered(sessionID).Where("is_correct = ?", false).Count(&count).Error
	return count, err
}

func (r *gameRepo) FlaggedCount(sessionID uint) (int64, error) {
	var count int64
	err := r.db.Model(&domain.GameAnswer{}).
		Where("game_session_id = ? AND is_flagged = ?", sessionID, true).Count(&count).Error
	return count, err
}

func (r *gameRepo) FlaggedQuestionIDs(sessionID uint) ([]uint, error) {
	var ids []uint
	err := r.db.Model(&domain.GameAnswer{}).
		Where("game_session_id = ? AND is_flagged = ?", sessionID, true).
		Pluck("question_id", &ids).Error
	return ids, err
}

func (r *gameRepo) AnsweredQuestionIDs(sessionID uint) ([]uint, error) {
	var ids []uint
	err := r.answered(sessionID).Pluck("question_id", &ids).Error
	return ids, err
}

func (r *gameRepo) AnsweredCountByCategory(sessionID uint) (map[string]int64, error) {
	var rows []struct {
		Category string
		Count    int64
	}
	err := r.db.Table("game_answers").
		Select("categories.code as category, COUNT(*) as count").
		Joins("JOIN questions ON questions.id = game_answers.question_id").
		Joins("JOIN categories ON categories.id = questions.category_id").
		Where("game_answers.game_session_id = ? AND game_answers.choice_id IS NOT NULL", sessionID).
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

func (r *gameRepo) AddHistory(sessionID, questionID uint) error {
	return r.db.Create(&domain.QuestionHistory{
		GameSessionID: sessionID,
		QuestionID:    questionID,
	}).Error
}

func (r *gameRepo) UsedQuestionIDs(sessionID uint) ([]uint, error) {
	var ids []uint
	err := r.db.Model(&domain.QuestionHistory{}).
		Where("game_session_id = ?", sessionID).
		Pluck("question_id", &ids).Error
	return ids, err
}
