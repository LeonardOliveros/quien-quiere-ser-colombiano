package main

import (
	"testing"
	"time"
)

func TestCalculateCurrentTimeElapsed(t *testing.T) {
	t.Run("paused session returns saved elapsed time", func(t *testing.T) {
		session := GameSession{
			Status:      "PAUSED",
			TimeElapsed: 120,
			StartTime:   time.Now().Add(-10 * time.Minute),
		}
		if got := calculateCurrentTimeElapsed(session); got != 120 {
			t.Errorf("expected 120, got %d", got)
		}
	})

	t.Run("active session measures from start time", func(t *testing.T) {
		session := GameSession{
			Status:    "ACTIVE",
			StartTime: time.Now().Add(-90 * time.Second),
		}
		got := calculateCurrentTimeElapsed(session)
		if got < 89 || got > 92 {
			t.Errorf("expected ~90s, got %d", got)
		}
	})
}

func TestGetTimeRemaining(t *testing.T) {
	t.Run("unlimited time returns -1", func(t *testing.T) {
		session := GameSession{TimeLimit: 0, Status: "ACTIVE", StartTime: time.Now()}
		if got := getTimeRemaining(session); got != -1 {
			t.Errorf("expected -1, got %d", got)
		}
	})

	t.Run("remaining time is limit minus elapsed", func(t *testing.T) {
		session := GameSession{
			TimeLimit: 3600,
			Status:    "ACTIVE",
			StartTime: time.Now().Add(-600 * time.Second),
		}
		got := getTimeRemaining(session)
		if got < 2998 || got > 3001 {
			t.Errorf("expected ~3000s, got %d", got)
		}
	})

	t.Run("expired time clamps to zero", func(t *testing.T) {
		session := GameSession{
			TimeLimit: 60,
			Status:    "ACTIVE",
			StartTime: time.Now().Add(-2 * time.Hour),
		}
		if got := getTimeRemaining(session); got != 0 {
			t.Errorf("expected 0, got %d", got)
		}
	})

	t.Run("paused session uses saved elapsed time", func(t *testing.T) {
		session := GameSession{
			TimeLimit:   3600,
			Status:      "PAUSED",
			TimeElapsed: 1000,
			StartTime:   time.Now().Add(-5 * time.Hour),
		}
		if got := getTimeRemaining(session); got != 2600 {
			t.Errorf("expected 2600, got %d", got)
		}
	})
}

func TestCheckIfPassed(t *testing.T) {
	cases := []struct {
		category   string
		percentage float64
		want       bool
	}{
		{"CONSTITUCION", 60.0, true},
		{"CONSTITUCION", 59.9, false},
		{"GEOGRAFIA", 55.0, true},
		{"GEOGRAFIA", 54.9, false},
		{"HISTORIA", 40.0, true},
		{"HISTORIA", 39.9, false},
		{"CULTURA", 40.0, true},
		{"CULTURA", 39.9, false},
		{"OTRA", 50.0, true},
		{"OTRA", 49.9, false},
	}
	for _, tc := range cases {
		if got := checkIfPassed(tc.category, tc.percentage); got != tc.want {
			t.Errorf("checkIfPassed(%q, %.1f) = %v, want %v", tc.category, tc.percentage, got, tc.want)
		}
	}
}

func TestCategoryHelpers(t *testing.T) {
	if got := getCategory("HISTORIA_Independencia"); got != "HISTORIA" {
		t.Errorf("getCategory: expected HISTORIA, got %q", got)
	}
	if got := getSubCategory("HISTORIA_Independencia"); got != "Independencia" {
		t.Errorf("getSubCategory: expected Independencia, got %q", got)
	}
	if got := getCategory("SINSEPARADOR"); got != "SINSEPARADOR" {
		t.Errorf("getCategory without separator: expected SINSEPARADOR, got %q", got)
	}
	if got := getSubCategory("SINSEPARADOR"); got != "" {
		t.Errorf("getSubCategory without separator: expected empty, got %q", got)
	}
}

func TestHideCorrectChoices(t *testing.T) {
	questions := []Question{
		{Choices: []Choice{{IsCorrect: true}, {IsCorrect: false}}},
		{Choices: []Choice{{IsCorrect: false}, {IsCorrect: true}}},
	}
	hideCorrectChoices(questions)
	for i, q := range questions {
		for j, c := range q.Choices {
			if c.IsCorrect {
				t.Errorf("question %d choice %d still exposes IsCorrect", i, j)
			}
		}
	}
}

func TestGenerateToken(t *testing.T) {
	a, b := generateToken(), generateToken()
	if len(a) != 32 {
		t.Errorf("expected 32 hex chars, got %d", len(a))
	}
	if a == b {
		t.Error("two generated tokens should not collide")
	}
}
