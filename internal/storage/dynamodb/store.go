// Package dynamodb is the (pending) DynamoDB adapter of the domain.Store port,
// meant for cloud deployments (DB_DRIVER=dynamodb).
//
// # Status
//
// This is a compile-checked skeleton: every repository method returns
// ErrNotImplemented. The interface assertion below guarantees that when the
// implementation is written, it satisfies exactly the same port the SQLite
// adapter does — the HTTP layer will not need any change.
//
// # Implementation notes
//
// Dependencies: github.com/aws/aws-sdk-go-v2 (config, dynamodb, attributevalue).
//
// Suggested single-table design (env DYNAMODB_TABLE, on-demand billing):
//
//	Entity            PK                      SK
//	Category          TAXONOMY                CAT#<code>
//	SubCategory       TAXONOMY                CAT#<code>#SUB#<code>
//	Question          QUESTION#<key>          META            (choices embedded as a list attribute)
//	User              USER#<username>         META
//	GameSession       USER#<id>               SESSION#<id>
//	GameAnswer        SESSION#<id>            ANSWER#<questionID>
//	QuestionHistory   SESSION#<id>            HISTORY#<questionID>
//	StudyRecommend.   USER#<id>               REC#<id>
//
// GSIs:
//   - gsi_category (question lookups by category): PK=CAT#<code>, SK=QUESTION#<key>
//   - gsi_token (auth): PK=TOKEN#<token>
//
// Porting guidance per port method:
//   - QuestionRepository.RandomID: query gsi_category (or scan all keys),
//     filter excludeIDs in memory, pick with math/rand. The question bank is
//     ~700 items, so loading the key list is cheap; cache it in the adapter.
//   - Choices travel embedded in the question item; ChoiceByID/CorrectChoices
//     resolve from the parent question (choice IDs are "<questionID>#<order>").
//   - StatsRepository: DynamoDB has no GROUP BY. Query the user's sessions
//     (PK=USER#<id>), then their answers per session, and aggregate in Go.
//     Volumes are small (a user's own history); no need for precomputed
//     counters yet. If they grow, maintain running totals on an item
//     USER#<id>/STATS updated transactionally from SaveAnswer.
//   - SyncQuestionBank: BatchWriteItem in chunks of 25; same upsert-by-key
//     semantics as the SQLite adapter (read existing keys first, never delete).
//   - Numeric IDs (sessions, answers, recommendations): generate with a
//     counter item (UpdateItem ADD) or switch the domain ID fields to strings
//     with ULIDs — decide when implementing.
//
// Local testing: DynamoDB Local via Docker
// (amazon/dynamodb-local, endpoint override http://localhost:8000).
package dynamodb

import (
	"errors"
	"fmt"
	"time"

	"quiz-app/internal/domain"
)

// ErrNotImplemented is returned by every method until the adapter is written.
var ErrNotImplemented = errors.New("dynamodb adapter not implemented yet")

// Store will implement domain.Store on top of DynamoDB.
type Store struct{}

// The build breaks here if Store drifts from the port contract.
var _ domain.Store = (*Store)(nil)

// Open validates configuration and connects to DynamoDB.
func Open(table string) (*Store, error) {
	return nil, fmt.Errorf("DB_DRIVER=dynamodb: %w (see internal/storage/dynamodb/store.go for the implementation plan)", ErrNotImplemented)
}

func (s *Store) Users() domain.UserRepository         { return userRepo{} }
func (s *Store) Questions() domain.QuestionRepository { return questionRepo{} }
func (s *Store) Games() domain.GameRepository         { return gameRepo{} }
func (s *Store) Stats() domain.StatsRepository        { return statsRepo{} }

func (s *Store) SyncQuestionBank([]domain.SeedCategory, []domain.SeedQuestion) error {
	return ErrNotImplemented
}
func (s *Store) ResetUserData(uint) error { return ErrNotImplemented }
func (s *Store) Close() error             { return nil }

type userRepo struct{}

func (userRepo) Create(*domain.User) error                      { return ErrNotImplemented }
func (userRepo) ByUsername(string) (domain.User, error)         { return domain.User{}, ErrNotImplemented }
func (userRepo) ByToken(string) (domain.User, error)            { return domain.User{}, ErrNotImplemented }
func (userRepo) SaveSessionToken(uint, string, time.Time) error { return ErrNotImplemented }

type questionRepo struct{}

func (questionRepo) ByID(uint) (domain.Question, error) { return domain.Question{}, ErrNotImplemented }
func (questionRepo) List() ([]domain.Question, error)   { return nil, ErrNotImplemented }
func (questionRepo) ListByCategory(string) ([]domain.Question, error) {
	return nil, ErrNotImplemented
}
func (questionRepo) RandomID([]string, []uint) (uint, error) { return 0, ErrNotImplemented }
func (questionRepo) Count([]string) (int64, error)           { return 0, ErrNotImplemented }
func (questionRepo) CountsByCategory() (map[string]int64, error) {
	return nil, ErrNotImplemented
}
func (questionRepo) CountsBySubcategory() ([]domain.SubcategoryCount, error) {
	return nil, ErrNotImplemented
}
func (questionRepo) ChoiceByID(uint) (domain.Choice, error) {
	return domain.Choice{}, ErrNotImplemented
}
func (questionRepo) CorrectChoices([]uint) (map[uint]domain.Choice, error) {
	return nil, ErrNotImplemented
}
func (questionRepo) Categories() ([]domain.Category, error) { return nil, ErrNotImplemented }

type gameRepo struct{}

func (gameRepo) CreateSession(*domain.GameSession) error { return ErrNotImplemented }
func (gameRepo) SessionByID(uint) (domain.GameSession, error) {
	return domain.GameSession{}, ErrNotImplemented
}
func (gameRepo) SaveSession(*domain.GameSession) error { return ErrNotImplemented }
func (gameRepo) AddScore(uint, int) error              { return ErrNotImplemented }
func (gameRepo) SessionsByUser(uint, int) ([]domain.GameSession, error) {
	return nil, ErrNotImplemented
}
func (gameRepo) LatestResumable(uint, string) (domain.GameSession, error) {
	return domain.GameSession{}, ErrNotImplemented
}
func (gameRepo) CompleteResumable(uint, string, uint) error { return ErrNotImplemented }
func (gameRepo) SaveAnswer(*domain.GameAnswer) error        { return ErrNotImplemented }
func (gameRepo) AnswerPlaceholder(uint, uint) (domain.GameAnswer, error) {
	return domain.GameAnswer{}, ErrNotImplemented
}
func (gameRepo) HasAnswered(uint, uint) (bool, error) { return false, ErrNotImplemented }
func (gameRepo) ToggleFlag(uint, uint) (bool, error)  { return false, ErrNotImplemented }
func (gameRepo) AnswersBySession(uint) ([]domain.GameAnswer, error) {
	return nil, ErrNotImplemented
}
func (gameRepo) IncorrectAnswers(uint) ([]domain.GameAnswer, error) {
	return nil, ErrNotImplemented
}
func (gameRepo) AnsweredCount(uint) (int64, error)        { return 0, ErrNotImplemented }
func (gameRepo) IncorrectCount(uint) (int64, error)       { return 0, ErrNotImplemented }
func (gameRepo) FlaggedCount(uint) (int64, error)         { return 0, ErrNotImplemented }
func (gameRepo) FlaggedQuestionIDs(uint) ([]uint, error)  { return nil, ErrNotImplemented }
func (gameRepo) AnsweredQuestionIDs(uint) ([]uint, error) { return nil, ErrNotImplemented }
func (gameRepo) AnsweredCountByCategory(uint) (map[string]int64, error) {
	return nil, ErrNotImplemented
}
func (gameRepo) AddHistory(uint, uint) error          { return ErrNotImplemented }
func (gameRepo) UsedQuestionIDs(uint) ([]uint, error) { return nil, ErrNotImplemented }

type statsRepo struct{}

func (statsRepo) TotalSessions(uint) (int64, error) { return 0, ErrNotImplemented }
func (statsRepo) BestScore(uint) (int, error)       { return 0, ErrNotImplemented }
func (statsRepo) OverallTotals(uint) (domain.AnswerTotals, error) {
	return domain.AnswerTotals{}, ErrNotImplemented
}
func (statsRepo) CategoryTotals(uint, string) (domain.AnswerTotals, error) {
	return domain.AnswerTotals{}, ErrNotImplemented
}
func (statsRepo) RecentProgress(uint, int) ([]domain.ProgressEntry, error) {
	return nil, ErrNotImplemented
}
func (statsRepo) WeakCategories(uint) ([]string, error) { return nil, ErrNotImplemented }
func (statsRepo) CreateRecommendation(*domain.StudyRecommendation) error {
	return ErrNotImplemented
}
func (statsRepo) RecommendationsByUser(uint, int) ([]domain.StudyRecommendation, error) {
	return nil, ErrNotImplemented
}
