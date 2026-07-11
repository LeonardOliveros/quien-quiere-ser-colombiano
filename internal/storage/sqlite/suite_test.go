package sqlite_test

import (
	"path/filepath"
	"testing"

	"quiz-app/internal/domain"
	"quiz-app/internal/storage/sqlite"
	"quiz-app/internal/storage/storagetest"
)

// TestConformance runs the shared storage suite against the SQLite adapter,
// which is the reference implementation the suite semantics were taken from.
func TestConformance(t *testing.T) {
	storagetest.RunStoreSuite(t, func(t *testing.T) domain.Store {
		st, err := sqlite.Open(filepath.Join(t.TempDir(), "quiz-test.db"))
		if err != nil {
			t.Fatalf("sqlite.Open: %v", err)
		}
		t.Cleanup(func() { st.Close() })
		return st
	})
}
