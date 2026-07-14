package dynamodb

import (
	"time"

	"quiz-app/internal/domain"
)

// Item shapes stored in the table. Conversions to/from domain structs live
// next to them so every attribute name appears in exactly one place.

type categoryItem struct {
	PK   string `dynamodbav:"PK"`
	SK   string `dynamodbav:"SK"`
	ID   uint   `dynamodbav:"id"`
	Code string `dynamodbav:"code"`
	Name string `dynamodbav:"name"`
}

type subCategoryItem struct {
	PK         string `dynamodbav:"PK"`
	SK         string `dynamodbav:"SK"`
	ID         uint   `dynamodbav:"id"`
	CategoryID uint   `dynamodbav:"category_id"`
	Code       string `dynamodbav:"code"`
	Name       string `dynamodbav:"name"`
}

type choiceItem struct {
	ID        uint   `dynamodbav:"id"`
	Text      string `dynamodbav:"text"`
	IsCorrect bool   `dynamodbav:"is_correct"`
	Order     int    `dynamodbav:"order"`
}

type questionItem struct {
	PK              string       `dynamodbav:"PK"`
	SK              string       `dynamodbav:"SK"`
	ID              uint         `dynamodbav:"id"`
	Key             string       `dynamodbav:"key"`
	CategoryCode    string       `dynamodbav:"category_code"`
	SubCategoryCode string       `dynamodbav:"subcategory_code"`
	Text            string       `dynamodbav:"text"`
	Difficulty      int          `dynamodbav:"difficulty"`
	Points          int          `dynamodbav:"points"`
	Hint            string       `dynamodbav:"hint"`
	Explanation     string       `dynamodbav:"explanation"`
	CreatedAt       time.Time    `dynamodbav:"created_at"`
	Choices         []choiceItem `dynamodbav:"choices"`
}

type userItem struct {
	PK             string     `dynamodbav:"PK"`
	SK             string     `dynamodbav:"SK"`
	ID             uint       `dynamodbav:"id"`
	Username       string     `dynamodbav:"username"`
	Password       string     `dynamodbav:"password"`
	Token          string     `dynamodbav:"token"`
	TokenExpiresAt *time.Time `dynamodbav:"token_expires_at,omitempty"`
	IsGuest        bool       `dynamodbav:"guest,omitempty"`
	LastActivityAt *time.Time `dynamodbav:"last_activity_at,omitempty"`
	CreatedAt      time.Time  `dynamodbav:"created_at"`
	UpdatedAt      time.Time  `dynamodbav:"updated_at"`
	TTL            int64      `dynamodbav:"ttl,omitempty"` // guests only
}

func (i userItem) toDomain() domain.User {
	return domain.User{
		ID: i.ID, Username: i.Username, Password: i.Password,
		Token: i.Token, TokenExpiresAt: i.TokenExpiresAt,
		IsGuest: i.IsGuest, LastActivityAt: i.LastActivityAt,
		CreatedAt: i.CreatedAt, UpdatedAt: i.UpdatedAt,
	}
}

// guestTTL is the item expiry stamped on everything a guest owns: their data
// self-destructs domain.GuestDataTTL after it was last written/touched.
func guestTTL(now time.Time) int64 { return now.Add(domain.GuestDataTTL).Unix() }

// uniqItem reserves a unique value (username) for a user.
type uniqItem struct {
	PK     string `dynamodbav:"PK"`
	SK     string `dynamodbav:"SK"`
	UserID uint   `dynamodbav:"user_id"`
}

// tokenItem makes auth-token lookups a strongly consistent GetItem. TTL lets
// DynamoDB garbage-collect expired tokens (expiry is still enforced by the
// HTTP layer, since TTL deletion can lag by up to ~48h).
type tokenItem struct {
	PK        string    `dynamodbav:"PK"`
	SK        string    `dynamodbav:"SK"`
	UserID    uint      `dynamodbav:"user_id"`
	ExpiresAt time.Time `dynamodbav:"expires_at"`
	TTL       int64     `dynamodbav:"ttl"`
}

type sessionItem struct {
	PK             string     `dynamodbav:"PK"`
	SK             string     `dynamodbav:"SK"`
	ID             uint       `dynamodbav:"id"`
	UserID         uint       `dynamodbav:"user_id"`
	Mode           string     `dynamodbav:"mode"`
	Categories     string     `dynamodbav:"categories"`
	Status         string     `dynamodbav:"status"`
	StartTime      time.Time  `dynamodbav:"start_time"`
	EndTime        *time.Time `dynamodbav:"end_time,omitempty"`
	PausedAt       *time.Time `dynamodbav:"paused_at,omitempty"`
	TimeElapsed    int        `dynamodbav:"time_elapsed"`
	TimeLimit      int        `dynamodbav:"time_limit"`
	TotalQuestions int        `dynamodbav:"total_questions"`
	CorrectAnswers int        `dynamodbav:"correct_answers"`
	Score          int        `dynamodbav:"score"`
	IsGuest        bool       `dynamodbav:"guest,omitempty"`
	CreatedAt      time.Time  `dynamodbav:"created_at"`
	UpdatedAt      time.Time  `dynamodbav:"updated_at"`
	TTL            int64      `dynamodbav:"ttl,omitempty"` // guests only
}

func newSessionItem(s domain.GameSession) sessionItem {
	item := sessionItem{
		PK: pkUser(s.UserID), SK: skSession(s.ID),
		ID: s.ID, UserID: s.UserID, Mode: s.Mode, Categories: s.Categories,
		Status: s.Status, StartTime: s.StartTime, EndTime: s.EndTime,
		PausedAt: s.PausedAt, TimeElapsed: s.TimeElapsed, TimeLimit: s.TimeLimit,
		TotalQuestions: s.TotalQuestions, CorrectAnswers: s.CorrectAnswers,
		Score: s.Score, IsGuest: s.IsGuest, CreatedAt: s.CreatedAt, UpdatedAt: s.UpdatedAt,
	}
	if s.IsGuest {
		item.TTL = guestTTL(time.Now())
	}
	return item
}

func (i sessionItem) toDomain() domain.GameSession {
	return domain.GameSession{
		ID: i.ID, UserID: i.UserID, Mode: i.Mode, Categories: i.Categories,
		Status: i.Status, StartTime: i.StartTime, EndTime: i.EndTime,
		PausedAt: i.PausedAt, TimeElapsed: i.TimeElapsed, TimeLimit: i.TimeLimit,
		TotalQuestions: i.TotalQuestions, CorrectAnswers: i.CorrectAnswers,
		Score: i.Score, IsGuest: i.IsGuest, CreatedAt: i.CreatedAt, UpdatedAt: i.UpdatedAt,
	}
}

// sessionPointerItem maps a session ID to its owner, so SessionByID resolves
// with two strongly consistent GetItems and no GSI.
type sessionPointerItem struct {
	PK      string `dynamodbav:"PK"`
	SK      string `dynamodbav:"SK"`
	UserID  uint   `dynamodbav:"user_id"`
	IsGuest bool   `dynamodbav:"guest,omitempty"`
	TTL     int64  `dynamodbav:"ttl,omitempty"` // guests only
}

type answerItem struct {
	PK            string    `dynamodbav:"PK"`
	SK            string    `dynamodbav:"SK"`
	ID            uint      `dynamodbav:"id"`
	GameSessionID uint      `dynamodbav:"game_session_id"`
	QuestionID    uint      `dynamodbav:"question_id"`
	ChoiceID      *uint     `dynamodbav:"choice_id,omitempty"`
	IsCorrect     bool      `dynamodbav:"is_correct"`
	IsFlagged     bool      `dynamodbav:"is_flagged"`
	TimeSpent     int       `dynamodbav:"time_spent"`
	AnsweredAt    time.Time `dynamodbav:"answered_at"`
	TTL           int64     `dynamodbav:"ttl,omitempty"` // guest sessions only
}

func newAnswerItem(a domain.GameAnswer) answerItem {
	return answerItem{
		PK: pkSession(a.GameSessionID), SK: skAnswer(a.QuestionID),
		ID: a.ID, GameSessionID: a.GameSessionID, QuestionID: a.QuestionID,
		ChoiceID: a.ChoiceID, IsCorrect: a.IsCorrect, IsFlagged: a.IsFlagged,
		TimeSpent: a.TimeSpent, AnsweredAt: a.AnsweredAt,
	}
}

func (i answerItem) toDomain() domain.GameAnswer {
	return domain.GameAnswer{
		ID: i.ID, GameSessionID: i.GameSessionID, QuestionID: i.QuestionID,
		ChoiceID: i.ChoiceID, IsCorrect: i.IsCorrect, IsFlagged: i.IsFlagged,
		TimeSpent: i.TimeSpent, AnsweredAt: i.AnsweredAt,
	}
}

type historyItem struct {
	PK            string    `dynamodbav:"PK"`
	SK            string    `dynamodbav:"SK"`
	GameSessionID uint      `dynamodbav:"game_session_id"`
	QuestionID    uint      `dynamodbav:"question_id"`
	CreatedAt     time.Time `dynamodbav:"created_at"`
	TTL           int64     `dynamodbav:"ttl,omitempty"` // guest sessions only
}

type recommendationItem struct {
	PK          string    `dynamodbav:"PK"`
	SK          string    `dynamodbav:"SK"`
	ID          uint      `dynamodbav:"id"`
	UserID      uint      `dynamodbav:"user_id"`
	Category    string    `dynamodbav:"category"`
	SubCategory string    `dynamodbav:"subcategory"`
	Weakness    string    `dynamodbav:"weakness"`
	Description string    `dynamodbav:"description"`
	Resources   string    `dynamodbav:"resources"`
	Priority    int       `dynamodbav:"priority"`
	CreatedAt   time.Time `dynamodbav:"created_at"`
	TTL         int64     `dynamodbav:"ttl,omitempty"` // guests only
}

// dayMetricsItem is the per-day usage rollup read by the admin dashboard.
// Counters are bumped with ADD expressions; the struct exists for reads.
type dayMetricsItem struct {
	PK           string `dynamodbav:"PK"`
	SK           string `dynamodbav:"SK"`
	GamesStarted int64  `dynamodbav:"games_started"`
	ActiveUsers  int64  `dynamodbav:"active_users"`
	NewGuests    int64  `dynamodbav:"new_guests"`
	NewUsers     int64  `dynamodbav:"new_users"`
	TTL          int64  `dynamodbav:"ttl,omitempty"`
}

func (i recommendationItem) toDomain() domain.StudyRecommendation {
	return domain.StudyRecommendation{
		ID: i.ID, UserID: i.UserID, Category: i.Category, SubCategory: i.SubCategory,
		Weakness: i.Weakness, Description: i.Description, Resources: i.Resources,
		Priority: i.Priority, CreatedAt: i.CreatedAt,
	}
}
