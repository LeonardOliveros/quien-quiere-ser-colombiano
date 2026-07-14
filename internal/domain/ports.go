package domain

import (
	"errors"
	"time"
)

// ErrNotFound is returned by any repository method when the requested record
// does not exist. Adapters must translate their native "not found" errors
// (gorm.ErrRecordNotFound, dynamodb item missing, ...) into this sentinel.
var ErrNotFound = errors.New("record not found")

// ErrAlreadyAnswered is returned by SaveAnswer when a real answer (choice_id
// set) already exists for (game_session_id, question_id). Adapters must
// enforce this at the data layer (unique index / conditional write) and
// translate the resulting constraint violation into this sentinel, so a
// racing duplicate submission cannot double-score a session.
var ErrAlreadyAnswered = errors.New("question already answered")

// Store is the persistence port of the application. The HTTP adapter talks
// only to this interface; swapping SQLite for DynamoDB means providing
// another implementation and selecting it at startup (DB_DRIVER).
//
// Contract notes for implementors:
//   - Methods returning a single record return ErrNotFound when missing.
//   - Question results must come with Category, SubCategory and Choices loaded.
//   - No method exposes SQL semantics: aggregations may be computed natively
//     (SQL) or in application code (DynamoDB), as long as results match.
type Store interface {
	Users() UserRepository
	Questions() QuestionRepository
	Games() GameRepository
	Stats() StatsRepository
	Metrics() MetricsRepository

	// SyncQuestionBank idempotently upserts the taxonomy and question bank:
	// categories/subcategories keyed by code, questions keyed by Key,
	// choices keyed by (question, order). It never deletes records.
	SyncQuestionBank(taxonomy []SeedCategory, questions []SeedQuestion) error

	// ResetUserData deletes the user's sessions, answers, question history
	// and study recommendations (the user account itself is kept).
	ResetUserData(userID uint) error

	Close() error
}

// UserRepository handles user accounts and session tokens.
type UserRepository interface {
	Create(user *User) error
	ByUsername(username string) (User, error)
	ByToken(token string) (User, error)
	SaveSessionToken(userID uint, token string, expiresAt time.Time) error

	// TouchGuest extends a guest's expiry (profile and current token) and
	// records the activity time. Best-effort semantics: callers may ignore
	// the error. Only ever called for users with IsGuest set.
	TouchGuest(userID uint, token string, expiresAt time.Time) error
}

// MetricsRepository records and reads aggregate usage counters. Handlers call
// RecordUserCreated right after Users().Create and RecordGameStarted right
// after Games().CreateSession; implementations may rely on that pairing
// (DynamoDB keeps explicit counters, SQLite derives everything with SQL and
// implements the Record* methods as no-ops).
type MetricsRepository interface {
	RecordUserCreated(isGuest bool, day string) error
	RecordGameStarted(userID uint, day string) error

	Totals() (MetricsTotals, error)
	// Daily returns the last `days` days including today, newest first,
	// zero-filled for days without activity.
	Daily(days int) ([]DailyMetrics, error)
}

// SubcategoryCount is one row of the question-bank breakdown.
type SubcategoryCount struct {
	Category    string `json:"category"`
	SubCategory string `json:"subcategory"`
	Count       int64  `json:"count"`
}

// QuestionRepository handles the question bank (questions, choices, taxonomy).
type QuestionRepository interface {
	ByID(id uint) (Question, error)
	List() ([]Question, error)
	ListByCategory(categoryCode string) ([]Question, error)

	// RandomID picks a random question restricted to categoryCodes (empty =
	// all) and excluding excludeIDs. Returns ErrNotFound when none is left.
	RandomID(categoryCodes []string, excludeIDs []uint) (uint, error)

	Count(categoryCodes []string) (int64, error)
	CountsByCategory() (map[string]int64, error)
	CountsBySubcategory() ([]SubcategoryCount, error)

	ChoiceByID(id uint) (Choice, error)
	// CorrectChoices returns the correct choice of each given question.
	CorrectChoices(questionIDs []uint) (map[uint]Choice, error)

	// Categories returns the taxonomy with subcategories loaded.
	Categories() ([]Category, error)
}

// GameRepository handles game sessions, answers and question history.
type GameRepository interface {
	CreateSession(session *GameSession) error
	SessionByID(id uint) (GameSession, error)
	// SaveSession persists every field of an existing session (status,
	// times, counters). The session must have been loaded first.
	SaveSession(session *GameSession) error
	// AddScore atomically increments correct_answers by 1 and score by points.
	AddScore(sessionID uint, points int) error
	SessionsByUser(userID uint, limit int) ([]GameSession, error)

	// LatestResumable returns the most recently updated ACTIVE/PAUSED session
	// for the user, optionally filtered by mode ("" = any).
	LatestResumable(userID uint, mode string) (GameSession, error)
	// CompleteResumable marks ACTIVE/PAUSED sessions as COMPLETED, optionally
	// filtered by mode ("" = any), skipping keepID (0 = none).
	CompleteResumable(userID uint, mode string, keepID uint) error

	// SaveAnswer inserts the answer, or updates it when ID is set (used to
	// fill in a flag placeholder created before the question was answered).
	SaveAnswer(answer *GameAnswer) error
	// AnswerPlaceholder returns the flag placeholder row (no choice yet) for
	// the question, or ErrNotFound.
	AnswerPlaceholder(sessionID, questionID uint) (GameAnswer, error)
	HasAnswered(sessionID, questionID uint) (bool, error)
	// ToggleFlag flips is_flagged on the session's answer for the question,
	// returning false when no row exists yet.
	ToggleFlag(sessionID, questionID uint) (bool, error)

	// AnswersBySession returns all answer rows (including flag placeholders)
	// with Question (taxonomy included) and Choice loaded.
	AnswersBySession(sessionID uint) ([]GameAnswer, error)
	// IncorrectAnswers returns wrong answers with Question loaded.
	IncorrectAnswers(sessionID uint) ([]GameAnswer, error)
	AnsweredCount(sessionID uint) (int64, error)
	IncorrectCount(sessionID uint) (int64, error)
	FlaggedCount(sessionID uint) (int64, error)
	FlaggedQuestionIDs(sessionID uint) ([]uint, error)
	AnsweredQuestionIDs(sessionID uint) ([]uint, error)
	// AnsweredCountByCategory maps category code -> answered questions in the
	// session (used by TIMED mode to balance categories).
	AnsweredCountByCategory(sessionID uint) (map[string]int64, error)

	// AddHistory records that a question was served in the session.
	AddHistory(sessionID, questionID uint) error
	// UsedQuestionIDs returns every question already served in the session.
	UsedQuestionIDs(sessionID uint) ([]uint, error)
}

// ProgressEntry is one completed session in the user's recent history.
type ProgressEntry struct {
	Date     time.Time
	Score    int
	Answered int
	Correct  int
}

// AnswerTotals aggregates answered questions and how many were correct.
type AnswerTotals struct {
	Total   int
	Correct int
}

// StatsRepository handles cross-session aggregations and study recommendations.
type StatsRepository interface {
	TotalSessions(userID uint) (int64, error)
	BestScore(userID uint) (int, error)
	// OverallTotals aggregates over every answered question of the user.
	OverallTotals(userID uint) (AnswerTotals, error)
	// CategoryTotals aggregates the user's answers within one category.
	CategoryTotals(userID uint, categoryCode string) (AnswerTotals, error)
	// RecentProgress returns the user's last completed sessions, newest first.
	RecentProgress(userID uint, limit int) ([]ProgressEntry, error)
	// WeakCategories returns category codes with accuracy below 50%, most
	// answered first.
	WeakCategories(userID uint) ([]string, error)

	CreateRecommendation(rec *StudyRecommendation) error
	// RecommendationsByUser returns recommendations by descending priority
	// (limit 0 = all).
	RecommendationsByUser(userID uint, limit int) ([]StudyRecommendation, error)
}
