package main

import (
	"time"
)

// User represents a quiz participant
type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"unique;not null" json:"username"`
	Password  string    `gorm:"column:password" json:"password"`
	Email     string    `gorm:"unique" json:"email"`
	Token     string    `gorm:"column:token" json:"token,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Question represents a quiz question
type Question struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	Category   string    `gorm:"index" json:"category"` // CULTURA, GEOGRAFIA, HISTORIA, CONSTITUCION
	SubCategory string   `json:"subcategory"` // e.g., "Símbolos Patrios", "Región Caribe", etc.
	Text       string    `gorm:"not null" json:"text"`
	Difficulty int       `json:"difficulty"` // 1-5 scale
	Points     int       `json:"points"`
	Hint       string    `json:"hint"`
	Explanation string   `json:"explanation"`
	CreatedAt  time.Time `json:"created_at"`
	Choices    []Choice  `json:"choices"`
}

// Choice represents an answer option
type Choice struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	QuestionID uint   `gorm:"index" json:"question_id"`
	Text       string `json:"text"`
	IsCorrect  bool   `json:"is_correct,omitempty"`
	Order      int    `json:"order"` // Display order (A, B, C, D)
}

// GameSession represents a quiz game session
type GameSession struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	UserID        uint      `gorm:"index" json:"user_id"`
	User          User      `json:"user,omitempty"`
	Mode          string    `json:"mode"` // PRACTICE, TIMED, WEAK_AREAS, FULL_TEST, CATEGORY
	Categories    string    `json:"categories"` // Comma-separated list of categories for filtering
	QuestionSequence string `gorm:"type:text" json:"question_sequence"` // Comma-separated list of question IDs in order
	Status        string    `json:"status"` // ACTIVE, PAUSED, COMPLETED
	StartTime     time.Time `json:"start_time"`
	EndTime       *time.Time `json:"end_time"`
	TimeLimit     int       `json:"time_limit"` // in seconds, 0 for unlimited
	TotalQuestions int      `json:"total_questions"`
	CorrectAnswers int      `json:"correct_answers"`
	Score         int       `json:"score"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// GameAnswer represents a user's answer in a game session
type GameAnswer struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	GameSessionID uint      `gorm:"index" json:"game_session_id"`
	QuestionID    uint      `gorm:"index" json:"question_id"`
	Question      Question  `json:"question,omitempty"`
	ChoiceID      *uint     `json:"choice_id"`
	Choice        *Choice   `json:"choice,omitempty"`
	IsCorrect     bool      `json:"is_correct"`
	IsFlagged     bool      `json:"is_flagged"` // User marked as unsure
	TimeSpent     int       `json:"time_spent"` // seconds
	AnsweredAt    time.Time `json:"answered_at"`
}

// StudyRecommendation represents study suggestions based on performance
type StudyRecommendation struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	UserID      uint   `gorm:"index" json:"user_id"`
	Category    string `json:"category"`
	SubCategory string `json:"subcategory"`
	Weakness    string `json:"weakness"` // Specific topic to study
	Description string `json:"description"`
	Resources   string `json:"resources"` // JSON string of study materials
	Priority    int    `json:"priority"` // 1-5, higher is more urgent
	CreatedAt   time.Time `json:"created_at"`
}

// GameResult represents the summary of a completed game
type GameResult struct {
	SessionID      uint                  `json:"session_id"`
	TotalQuestions int                   `json:"total_questions"`
	CorrectAnswers int                   `json:"correct_answers"`
	Score          int                   `json:"score"`
	Percentage     float64               `json:"percentage"`
	TimeTaken      int                   `json:"time_taken"` // in seconds
	CategoryScores map[string]CategoryScore `json:"category_scores"`
	IncorrectAnswers []IncorrectAnswer   `json:"incorrect_answers"`
	FlaggedQuestions []Question          `json:"flagged_questions"`
	Recommendations []string             `json:"recommendations"`
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
	UserID           uint                     `json:"user_id"`
	TotalGames       int64                    `json:"total_games"`
	TotalQuestions   int                      `json:"total_questions"`
	CorrectAnswers   int                      `json:"correct_answers"`
	AverageScore     float64                  `json:"average_score"`
	BestScore        int                      `json:"best_score"`
	CategoryStats    map[string]CategoryStats `json:"category_stats"`
	WeakAreas        []string                 `json:"weak_areas"`
	StrongAreas      []string                 `json:"strong_areas"`
	RecentProgress   []ProgressPoint          `json:"recent_progress"`
}

// CategoryStats represents statistics for a specific category
type CategoryStats struct {
	Category         string  `json:"category"`
	TotalQuestions   int     `json:"total_questions"`
	CorrectAnswers   int     `json:"correct_answers"`
	AveragePercentage float64 `json:"average_percentage"`
	Improvement      float64 `json:"improvement"` // % change over last 5 games
}

// ProgressPoint represents a point in the user's progress timeline
type ProgressPoint struct {
	Date       time.Time `json:"date"`
	Score      int       `json:"score"`
	Percentage float64   `json:"percentage"`
}

// GameConfig represents configuration for different game modes
type GameConfig struct {
	Mode           string `json:"mode"`
	QuestionCount  int    `json:"question_count"`
	TimeLimit      int    `json:"time_limit"` // seconds
	Categories     []string `json:"categories"`
	Difficulty     string `json:"difficulty"` // EASY, MEDIUM, HARD, MIXED
	FocusWeakAreas bool   `json:"focus_weak_areas"`
}
