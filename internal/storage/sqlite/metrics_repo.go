package sqlite

import (
	"time"

	"gorm.io/gorm"

	"quiz-app/internal/domain"
)

type metricsRepo struct {
	db *gorm.DB
}

// Record* are no-ops: with SQL available, Totals and Daily are derived
// directly from the users/game_sessions tables, so nothing extra needs to be
// written when a user or game is created.
func (r *metricsRepo) RecordUserCreated(isGuest bool, day string) error { return nil }
func (r *metricsRepo) RecordGameStarted(userID uint, day string) error  { return nil }

func (r *metricsRepo) Totals() (domain.MetricsTotals, error) {
	var t domain.MetricsTotals
	if err := r.db.Model(&domain.User{}).Where("is_guest = ?", false).Count(&t.RegisteredUsers).Error; err != nil {
		return t, err
	}
	if err := r.db.Model(&domain.User{}).Where("is_guest = ?", true).Count(&t.GuestUsers).Error; err != nil {
		return t, err
	}
	err := r.db.Model(&domain.GameSession{}).Count(&t.TotalGames).Error
	return t, err
}

// cotDay buckets a stored timestamp into a Colombia-time (UTC-5) date; SQLite
// normalizes timezone-suffixed timestamps to UTC before applying modifiers.
const cotDay = "date(created_at, '-5 hours')"

func (r *metricsRepo) Daily(days int) ([]domain.DailyMetrics, error) {
	now := time.Now().In(domain.MetricsTimezone)
	oldest := domain.MetricsDay(now.AddDate(0, 0, -(days - 1)))

	byDay := make(map[string]*domain.DailyMetrics, days)
	result := make([]domain.DailyMetrics, days)
	for i := 0; i < days; i++ {
		day := domain.MetricsDay(now.AddDate(0, 0, -i))
		result[i] = domain.DailyMetrics{Date: day}
		byDay[day] = &result[i]
	}

	var userRows []struct {
		Day     string
		IsGuest bool
		N       int64
	}
	if err := r.db.Model(&domain.User{}).
		Select(cotDay+" AS day, is_guest, COUNT(*) AS n").
		Where(cotDay+" >= ?", oldest).
		Group("day, is_guest").Scan(&userRows).Error; err != nil {
		return nil, err
	}
	for _, row := range userRows {
		if m := byDay[row.Day]; m != nil {
			if row.IsGuest {
				m.NewGuests = row.N
			} else {
				m.NewUsers = row.N
			}
		}
	}

	var gameRows []struct {
		Day     string
		Games   int64
		Players int64
	}
	if err := r.db.Model(&domain.GameSession{}).
		Select(cotDay+" AS day, COUNT(*) AS games, COUNT(DISTINCT user_id) AS players").
		Where(cotDay+" >= ?", oldest).
		Group("day").Scan(&gameRows).Error; err != nil {
		return nil, err
	}
	for _, row := range gameRows {
		if m := byDay[row.Day]; m != nil {
			m.GamesStarted = row.Games
			m.ActiveUsers = row.Players
		}
	}

	return result, nil
}
