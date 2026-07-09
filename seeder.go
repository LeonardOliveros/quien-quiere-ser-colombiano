package main

import (
	_ "embed"
	"encoding/json"
	"log"
)

//go:embed data/questions.json
var questionsJSON []byte

type seedChoice struct {
	Text      string `json:"text"`
	IsCorrect bool   `json:"is_correct"`
	Order     int    `json:"order"`
}

type seedQuestion struct {
	Category    string       `json:"category"`
	SubCategory string       `json:"subcategory"`
	Text        string       `json:"text"`
	Difficulty  int          `json:"difficulty"`
	Points      int          `json:"points"`
	Hint        string       `json:"hint"`
	Explanation string       `json:"explanation"`
	Choices     []seedChoice `json:"choices"`
}

func seedQuestions() {
	// Check if questions already exist
	var count int64
	db.Model(&Question{}).Count(&count)
	if count > 0 {
		log.Printf("Database already contains %d questions", count)
		return
	}

	log.Println("Seeding database with initial questions...")

	var seeds []seedQuestion
	if err := json.Unmarshal(questionsJSON, &seeds); err != nil {
		log.Fatalf("Failed to parse embedded data/questions.json: %v", err)
	}

	questions := make([]Question, 0, len(seeds))
	for _, s := range seeds {
		q := Question{
			Category:    s.Category,
			SubCategory: s.SubCategory,
			Text:        s.Text,
			Difficulty:  s.Difficulty,
			Points:      s.Points,
			Hint:        s.Hint,
			Explanation: s.Explanation,
		}
		for _, c := range s.Choices {
			q.Choices = append(q.Choices, Choice{
				Text:      c.Text,
				IsCorrect: c.IsCorrect,
				Order:     c.Order,
			})
		}
		questions = append(questions, q)
	}

	if err := db.CreateInBatches(questions, 100).Error; err != nil {
		log.Printf("Error seeding questions: %v", err)
		return
	}

	log.Printf("Successfully seeded %d questions", len(questions))
}
