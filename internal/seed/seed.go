// Package seed loads and validates the embedded question bank
// (data/taxonomy.json + data/questions/*.json) into domain seed types,
// ready to be handed to any Store implementation via SyncQuestionBank.
package seed

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"

	"quiz-app/internal/domain"
)

// Load parses and validates the seed data files from fsys, which must contain
// data/taxonomy.json and data/questions/*.json (one file per category, named
// after the category code in lowercase).
func Load(fsys fs.FS) ([]domain.SeedCategory, []domain.SeedQuestion, error) {
	rawTaxonomy, err := fs.ReadFile(fsys, "data/taxonomy.json")
	if err != nil {
		return nil, nil, err
	}
	var taxonomy []domain.SeedCategory
	if err := json.Unmarshal(rawTaxonomy, &taxonomy); err != nil {
		return nil, nil, fmt.Errorf("parsing data/taxonomy.json: %w", err)
	}

	validSubs := make(map[string]bool) // "CATEGORY/SUBCATEGORY"
	for _, cat := range taxonomy {
		for _, sub := range cat.SubCategories {
			validSubs[cat.Code+"/"+sub.Code] = true
		}
	}

	files, err := fs.ReadDir(fsys, "data/questions")
	if err != nil {
		return nil, nil, err
	}

	var seeds []domain.SeedQuestion
	seenKeys := make(map[string]string)
	for _, f := range files {
		categoryCode := strings.ToUpper(strings.TrimSuffix(f.Name(), filepath.Ext(f.Name())))
		raw, err := fs.ReadFile(fsys, "data/questions/"+f.Name())
		if err != nil {
			return nil, nil, err
		}
		var fileSeeds []domain.SeedQuestion
		if err := json.Unmarshal(raw, &fileSeeds); err != nil {
			return nil, nil, fmt.Errorf("parsing data/questions/%s: %w", f.Name(), err)
		}
		for _, s := range fileSeeds {
			s.Category = categoryCode
			if s.Key == "" {
				return nil, nil, fmt.Errorf("%s: question %q has no key", f.Name(), s.Text)
			}
			if prev, dup := seenKeys[s.Key]; dup {
				return nil, nil, fmt.Errorf("duplicate question key %s (in %s and %s)", s.Key, prev, f.Name())
			}
			seenKeys[s.Key] = f.Name()
			if !validSubs[categoryCode+"/"+s.SubCategory] {
				return nil, nil, fmt.Errorf("%s: question %s references unknown subcategory %q", f.Name(), s.Key, s.SubCategory)
			}
			correct := 0
			seenOrders := make(map[int]bool, len(s.Choices))
			for _, c := range s.Choices {
				if c.IsCorrect {
					correct++
				}
				if seenOrders[c.Order] {
					return nil, nil, fmt.Errorf("%s: question %s has duplicate choice order %d", f.Name(), s.Key, c.Order)
				}
				seenOrders[c.Order] = true
			}
			if len(s.Choices) < 2 || correct != 1 {
				return nil, nil, fmt.Errorf("%s: question %s must have 2+ choices and exactly 1 correct (has %d/%d)", f.Name(), s.Key, correct, len(s.Choices))
			}
			seeds = append(seeds, s)
		}
	}
	return taxonomy, seeds, nil
}
