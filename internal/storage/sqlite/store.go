// Package sqlite is the SQLite/GORM adapter of the domain.Store port.
// It is the default driver, meant for local development and single-node
// deployments (DB_DRIVER=sqlite or unset).
package sqlite

import (
	"errors"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"quiz-app/internal/domain"
)

// Store implements domain.Store on top of a SQLite database.
type Store struct {
	db *gorm.DB
}

var _ domain.Store = (*Store)(nil)

// Open connects to the SQLite database at path and runs schema migrations.
func Open(path string) (*Store, error) {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(
		&domain.Category{},
		&domain.SubCategory{},
		&domain.Question{},
		&domain.Choice{},
		&domain.User{},
		&domain.GameSession{},
		&domain.GameAnswer{},
		&domain.QuestionHistory{},
		&domain.StudyRecommendation{},
	); err != nil {
		return nil, err
	}

	return &Store{db: db}, nil
}

func (s *Store) Users() domain.UserRepository         { return &userRepo{db: s.db} }
func (s *Store) Questions() domain.QuestionRepository { return &questionRepo{db: s.db} }
func (s *Store) Games() domain.GameRepository         { return &gameRepo{db: s.db} }
func (s *Store) Stats() domain.StatsRepository        { return &statsRepo{db: s.db} }

func (s *Store) ResetUserData(userID uint) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec("DELETE FROM game_answers WHERE game_session_id IN (SELECT id FROM game_sessions WHERE user_id = ?)", userID).Error; err != nil {
			return err
		}
		if err := tx.Exec("DELETE FROM question_histories WHERE game_session_id IN (SELECT id FROM game_sessions WHERE user_id = ?)", userID).Error; err != nil {
			return err
		}
		if err := tx.Where("user_id = ?", userID).Delete(&domain.GameSession{}).Error; err != nil {
			return err
		}
		return tx.Where("user_id = ?", userID).Delete(&domain.StudyRecommendation{}).Error
	})
}

func (s *Store) Close() error {
	sqlDB, err := s.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

// translate maps GORM errors to domain sentinels.
func translate(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return domain.ErrNotFound
	}
	return err
}
