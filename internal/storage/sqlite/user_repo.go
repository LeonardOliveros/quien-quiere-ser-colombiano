package sqlite

import (
	"time"

	"gorm.io/gorm"

	"quiz-app/internal/domain"
)

type userRepo struct {
	db *gorm.DB
}

func (r *userRepo) Create(user *domain.User) error {
	return r.db.Create(user).Error
}

func (r *userRepo) ByUsername(username string) (domain.User, error) {
	var user domain.User
	err := r.db.Where("username = ?", username).First(&user).Error
	return user, translate(err)
}

func (r *userRepo) ByToken(token string) (domain.User, error) {
	var user domain.User
	err := r.db.Where("token = ?", token).First(&user).Error
	return user, translate(err)
}

func (r *userRepo) SaveSessionToken(userID uint, token string, expiresAt time.Time) error {
	return r.db.Model(&domain.User{}).Where("id = ?", userID).Updates(map[string]interface{}{
		"token":            token,
		"token_expires_at": expiresAt,
	}).Error
}

func (r *userRepo) TouchGuest(userID uint, token string, expiresAt time.Time) error {
	now := time.Now()
	return r.db.Model(&domain.User{}).Where("id = ? AND token = ?", userID, token).Updates(map[string]interface{}{
		"token_expires_at": expiresAt,
		"last_activity_at": now,
	}).Error
}
