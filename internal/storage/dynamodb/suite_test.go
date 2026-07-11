package dynamodb_test

import (
	"fmt"
	"os"
	"sync/atomic"
	"testing"
	"time"

	"quiz-app/internal/domain"
	"quiz-app/internal/storage/dynamodb"
	"quiz-app/internal/storage/storagetest"
)

var tableSeq atomic.Int64

// TestConformance runs the shared storage suite against the DynamoDB adapter
// using DynamoDB Local. It is skipped unless DYNAMODB_TEST_ENDPOINT is set:
//
//	docker run --rm -d -p 8000:8000 amazon/dynamodb-local
//	DYNAMODB_TEST_ENDPOINT=http://localhost:8000 go test ./internal/storage/dynamodb/
func TestConformance(t *testing.T) {
	endpoint := os.Getenv("DYNAMODB_TEST_ENDPOINT")
	if endpoint == "" {
		t.Skip("DYNAMODB_TEST_ENDPOINT not set; start DynamoDB Local to run this suite")
	}
	os.Setenv("DYNAMODB_ENDPOINT", endpoint)

	storagetest.RunStoreSuite(t, func(t *testing.T) domain.Store {
		table := fmt.Sprintf("quiz-conformance-%d-%d", time.Now().UnixNano(), tableSeq.Add(1))
		st, err := dynamodb.Open(table)
		if err != nil {
			t.Fatalf("dynamodb.Open: %v", err)
		}
		t.Cleanup(func() { st.Close() })
		return st
	})
}
