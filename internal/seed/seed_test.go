package seed

import (
	"os"
	"testing"
)

// TestLoad guards the real data files at the repo root: valid taxonomy
// references, unique keys and exactly one correct choice per question.
func TestLoad(t *testing.T) {
	taxonomy, seeds, err := Load(os.DirFS("../.."))
	if err != nil {
		t.Fatalf("seed data is invalid: %v", err)
	}
	if len(taxonomy) != 4 {
		t.Errorf("expected 4 categories, got %d", len(taxonomy))
	}
	if len(seeds) < 700 {
		t.Errorf("expected 700+ questions, got %d", len(seeds))
	}
	byCategory := make(map[string]int)
	for _, s := range seeds {
		byCategory[s.Category]++
	}
	for _, code := range []string{"CULTURA", "GEOGRAFIA", "HISTORIA", "CONSTITUCION"} {
		if byCategory[code] == 0 {
			t.Errorf("category %s has no questions", code)
		}
	}
}
