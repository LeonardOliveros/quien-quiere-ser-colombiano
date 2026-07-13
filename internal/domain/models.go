// Package domain holds the core entities and persistence ports of the app.
// It has no dependencies on any storage technology: the gorm struct tags are
// plain metadata read by the SQLite adapter and ignored by everything else.
package domain

import (
	"encoding/json"
	"time"
)

// Category represents a top-level question category (CULTURA, GEOGRAFIA, ...)
type Category struct {
	ID            uint          `gorm:"primaryKey" json:"id"`
	Code          string        `gorm:"uniqueIndex;not null" json:"code"`
	Name          string        `json:"name"`
	SubCategories []SubCategory `json:"subcategories,omitempty"`
}

// SubCategory represents a curated topic within a category
type SubCategory struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	CategoryID uint   `gorm:"uniqueIndex:idx_subcat_cat_code" json:"category_id"`
	Code       string `gorm:"uniqueIndex:idx_subcat_cat_code;not null" json:"code"`
	Name       string `json:"name"`
}

// User represents a quiz participant
type User struct {
	ID             uint       `gorm:"primaryKey" json:"id"`
	Username       string     `gorm:"unique;not null" json:"username"`
	Password       string     `gorm:"column:password" json:"-"`
	Email          string     `gorm:"unique" json:"email"`
	Token          string     `gorm:"column:token" json:"-"`
	TokenExpiresAt *time.Time `json:"-"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
}

// Question represents a quiz question
type Question struct {
	ID            uint        `gorm:"primaryKey" json:"id"`
	Key           string      `gorm:"uniqueIndex;not null" json:"key"` // Stable seed key (e.g. CUL-0001)
	CategoryID    uint        `gorm:"index" json:"-"`
	Category      Category    `json:"-"`
	SubCategoryID uint        `gorm:"index" json:"-"`
	SubCategory   SubCategory `json:"-"`
	Text          string      `gorm:"not null" json:"text"`
	Difficulty    int         `json:"difficulty"` // 1-5 scale
	Points        int         `json:"points"`
	Hint          string      `json:"hint"`
	Explanation   string      `json:"explanation"`
	CreatedAt     time.Time   `json:"created_at"`
	Choices       []Choice    `json:"choices"`
}

// MarshalJSON keeps the API contract the frontend expects: "category" is the
// category code (used for filtering) and "subcategory" the display name.
// Both require the Category/SubCategory relations to be loaded.
func (q Question) MarshalJSON() ([]byte, error) {
	type alias Question
	return json.Marshal(struct {
		alias
		CategoryCode    string `json:"category"`
		SubCategoryName string `json:"subcategory"`
	}{
		alias:           alias(q),
		CategoryCode:    q.Category.Code,
		SubCategoryName: q.SubCategory.Name,
	})
}

// Choice represents an answer option
type Choice struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	QuestionID uint   `gorm:"uniqueIndex:idx_choice_question_order" json:"question_id"`
	Text       string `json:"text"`
	IsCorrect  bool   `json:"is_correct,omitempty"`
	Order      int    `gorm:"uniqueIndex:idx_choice_question_order" json:"order"` // Display order (A, B, C, D)
}

// GameSession represents a quiz game session
type GameSession struct {
	ID             uint       `gorm:"primaryKey" json:"id"`
	UserID         uint       `gorm:"index" json:"user_id"`
	User           User       `json:"user,omitempty"`
	Mode           string     `json:"mode"`       // PRACTICE, TIMED, WEAK_AREAS, FULL_TEST, CATEGORY
	Categories     string     `json:"categories"` // Comma-separated list of category codes for filtering
	Status         string     `json:"status"`     // ACTIVE, PAUSED, COMPLETED
	StartTime      time.Time  `json:"start_time"`
	EndTime        *time.Time `json:"end_time"`
	PausedAt       *time.Time `json:"paused_at"`    // When the game was paused
	TimeElapsed    int        `json:"time_elapsed"` // Total time elapsed in seconds (excluding paused time)
	TimeLimit      int        `json:"time_limit"`   // in seconds, 0 for unlimited
	TotalQuestions int        `json:"total_questions"`
	CorrectAnswers int        `json:"correct_answers"`
	Score          int        `json:"score"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
}

// GameAnswer represents a user's answer in a game session
type GameAnswer struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	GameSessionID uint      `gorm:"uniqueIndex:idx_answer_session_question" json:"game_session_id"`
	QuestionID    uint      `gorm:"uniqueIndex:idx_answer_session_question" json:"question_id"`
	Question      Question  `json:"question,omitempty"`
	ChoiceID      *uint     `json:"choice_id"`
	Choice        *Choice   `json:"choice,omitempty"`
	IsCorrect     bool      `json:"is_correct"`
	IsFlagged     bool      `json:"is_flagged"` // User marked as unsure
	TimeSpent     int       `json:"time_spent"` // seconds
	AnsweredAt    time.Time `json:"answered_at"`
}

// QuestionHistory represents the history of questions used in a game session
type QuestionHistory struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	GameSessionID uint      `gorm:"index" json:"game_session_id"`
	QuestionID    uint      `gorm:"index" json:"question_id"`
	CreatedAt     time.Time `json:"created_at"`
}

// StudyRecommendation represents study suggestions based on performance
type StudyRecommendation struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	UserID      uint      `gorm:"index" json:"user_id"`
	Category    string    `json:"category"`
	SubCategory string    `json:"subcategory"`
	Weakness    string    `json:"weakness"` // Specific topic to study
	Description string    `json:"description"`
	Resources   string    `json:"resources"` // JSON string of study materials
	Priority    int       `json:"priority"`  // 1-5, higher is more urgent
	CreatedAt   time.Time `json:"created_at"`
}

// GameResult represents the summary of a completed game
type GameResult struct {
	SessionID        uint                     `json:"session_id"`
	TotalQuestions   int                      `json:"total_questions"`
	CorrectAnswers   int                      `json:"correct_answers"`
	Score            int                      `json:"score"`
	Percentage       float64                  `json:"percentage"`
	TimeTaken        int                      `json:"time_taken"` // in seconds
	CategoryScores   map[string]CategoryScore `json:"category_scores"`
	IncorrectAnswers []IncorrectAnswer        `json:"incorrect_answers"`
	FlaggedQuestions []Question               `json:"flagged_questions"`
	Recommendations  []string                 `json:"recommendations"`
}

// CategoryScore represents performance in a specific category
type CategoryScore struct {
	Category       string  `json:"category"`
	TotalQuestions int     `json:"total_questions"`
	CorrectAnswers int     `json:"correct_answers"`
	Percentage     float64 `json:"percentage"`
	Passed         bool    `json:"passed"` // Based on minimum required percentage
}

// IncorrectAnswer represents a wrong answer for review
type IncorrectAnswer struct {
	Question      Question `json:"question"`
	UserChoice    *Choice  `json:"user_choice"`
	CorrectChoice Choice   `json:"correct_choice"`
	Explanation   string   `json:"explanation"`
}

// UserStats represents overall user statistics
type UserStats struct {
	UserID         uint                     `json:"user_id"`
	TotalGames     int64                    `json:"total_games"`
	TotalQuestions int                      `json:"total_questions"`
	CorrectAnswers int                      `json:"correct_answers"`
	AverageScore   float64                  `json:"average_score"`
	BestScore      int                      `json:"best_score"`
	CategoryStats  map[string]CategoryStats `json:"category_stats"`
	WeakAreas      []string                 `json:"weak_areas"`
	StrongAreas    []string                 `json:"strong_areas"`
	RecentProgress []ProgressPoint          `json:"recent_progress"`
}

// CategoryStats represents statistics for a specific category
type CategoryStats struct {
	Category          string  `json:"category"`
	TotalQuestions    int     `json:"total_questions"`
	CorrectAnswers    int     `json:"correct_answers"`
	AveragePercentage float64 `json:"average_percentage"`
	Improvement       float64 `json:"improvement"` // % change over last 5 games
}

// ProgressPoint represents a point in the user's progress timeline
type ProgressPoint struct {
	Date       time.Time `json:"date"`
	Score      int       `json:"score"`
	Percentage float64   `json:"percentage"`
}

// GameConfig represents configuration for different game modes
type GameConfig struct {
	Mode           string   `json:"mode"`
	QuestionCount  int      `json:"question_count"`
	TimeLimit      int      `json:"time_limit"` // seconds
	Categories     []string `json:"categories"`
	Difficulty     string   `json:"difficulty"` // EASY, MEDIUM, HARD, MIXED
	FocusWeakAreas bool     `json:"focus_weak_areas"`
}

// SeedCategory is a category entry from data/taxonomy.json.
type SeedCategory struct {
	Code          string            `json:"code"`
	Name          string            `json:"name"`
	SubCategories []SeedSubCategory `json:"subcategories"`
}

// SeedSubCategory is a subcategory entry from data/taxonomy.json.
type SeedSubCategory struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

// SeedChoice is an answer option from data/questions/*.json.
type SeedChoice struct {
	Text      string `json:"text"`
	IsCorrect bool   `json:"is_correct"`
	Order     int    `json:"order"`
}

// SeedQuestion is a question entry from data/questions/*.json.
type SeedQuestion struct {
	Key         string       `json:"key"`
	Category    string       `json:"-"` // Derived from the file name
	SubCategory string       `json:"subcategory"`
	Text        string       `json:"text"`
	Difficulty  int          `json:"difficulty"`
	Points      int          `json:"points"`
	Hint        string       `json:"hint"`
	Explanation string       `json:"explanation"`
	Choices     []SeedChoice `json:"choices"`
}
