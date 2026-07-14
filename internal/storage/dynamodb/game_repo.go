package dynamodb

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

	"quiz-app/internal/domain"
)

type gameRepo struct{ s *Store }

// CreateSession writes the session under the user partition plus a pointer
// item (SESSION#<id>/META -> user_id) in one transaction, so SessionByID can
// resolve ownership without a GSI.
func (r *gameRepo) CreateSession(session *domain.GameSession) error {
	ctx := context.Background()
	id, err := r.s.nextID(ctx, "session")
	if err != nil {
		return err
	}
	now := time.Now()
	session.ID = id
	session.CreatedAt = now
	session.UpdatedAt = now

	item, err := attributevalue.MarshalMap(newSessionItem(*session))
	if err != nil {
		return fmt.Errorf("marshal session: %w", err)
	}
	pointerItem := sessionPointerItem{
		PK: pkSession(id), SK: skMeta, UserID: session.UserID, IsGuest: session.IsGuest,
	}
	if session.IsGuest {
		pointerItem.TTL = guestTTL(now)
	}
	pointer, err := attributevalue.MarshalMap(pointerItem)
	if err != nil {
		return fmt.Errorf("marshal session pointer: %w", err)
	}

	_, err = r.s.client.TransactWriteItems(ctx, &dynamodb.TransactWriteItemsInput{
		TransactItems: []types.TransactWriteItem{
			{Put: &types.Put{TableName: aws.String(r.s.table), Item: item}},
			{Put: &types.Put{TableName: aws.String(r.s.table), Item: pointer}},
		},
	})
	if err != nil {
		return fmt.Errorf("create session: %w", err)
	}
	r.s.guestSessions.Store(id, session.IsGuest)
	return nil
}

// sessionPointer resolves the pointer item of a session, caching its guest
// flag for the per-question TTL stamping in SaveAnswer/AddHistory.
func (r *gameRepo) sessionPointer(ctx context.Context, sessionID uint) (sessionPointerItem, error) {
	var pointer sessionPointerItem
	if err := r.s.getItem(ctx, pkSession(sessionID), skMeta, &pointer); err != nil {
		return pointer, err
	}
	r.s.guestSessions.Store(sessionID, pointer.IsGuest)
	return pointer, nil
}

// sessionOwner resolves the owner of a session via the pointer item.
func (r *gameRepo) sessionOwner(ctx context.Context, sessionID uint) (uint, error) {
	pointer, err := r.sessionPointer(ctx, sessionID)
	if err != nil {
		return 0, err
	}
	return pointer.UserID, nil
}

// isGuestSession reports whether the session belongs to a guest, from cache
// when possible. Lookup failures degrade to false: the write then simply
// lacks a TTL, which only delays cleanup, never breaks correctness.
func (r *gameRepo) isGuestSession(ctx context.Context, sessionID uint) bool {
	if v, ok := r.s.guestSessions.Load(sessionID); ok {
		return v.(bool)
	}
	pointer, err := r.sessionPointer(ctx, sessionID)
	if err != nil {
		return false
	}
	return pointer.IsGuest
}

func (r *gameRepo) SessionByID(id uint) (domain.GameSession, error) {
	ctx := context.Background()
	userID, err := r.sessionOwner(ctx, id)
	if err != nil {
		return domain.GameSession{}, err
	}
	var item sessionItem
	if err := r.s.getItem(ctx, pkUser(userID), skSession(id), &item); err != nil {
		return domain.GameSession{}, err
	}
	return item.toDomain(), nil
}

// SaveSession persists every field of the (already loaded) session, bumping
// updated_at. Nil pointer fields are omitted from the item, so pause/resume
// round-trips PausedAt back to nil like the SQLite Save does.
func (r *gameRepo) SaveSession(session *domain.GameSession) error {
	session.UpdatedAt = time.Now()
	return r.s.putItem(context.Background(), newSessionItem(*session))
}

// AddScore atomically increments correct_answers and score. It deliberately
// does NOT touch updated_at (the SQLite adapter uses UpdateColumn, which
// skips it) — LatestResumable ordering depends on that.
func (r *gameRepo) AddScore(sessionID uint, points int) error {
	ctx := context.Background()
	userID, err := r.sessionOwner(ctx, sessionID)
	if err != nil {
		return err
	}
	_, err = r.s.client.UpdateItem(ctx, &dynamodb.UpdateItemInput{
		TableName: aws.String(r.s.table),
		Key: map[string]types.AttributeValue{
			"PK": &types.AttributeValueMemberS{Value: pkUser(userID)},
			"SK": &types.AttributeValueMemberS{Value: skSession(sessionID)},
		},
		UpdateExpression: aws.String("ADD correct_answers :one, score :points"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":one":    &types.AttributeValueMemberN{Value: "1"},
			":points": &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", points)},
		},
	})
	if err != nil {
		return fmt.Errorf("add score: %w", err)
	}
	return nil
}

// userSessions returns all of the user's sessions, newest first (the SKs are
// zero-padded counter IDs, so descending key order == descending creation).
func (r *gameRepo) userSessions(ctx context.Context, userID uint) ([]sessionItem, error) {
	var sessions []sessionItem
	var startKey map[string]types.AttributeValue
	for {
		out, err := r.s.client.Query(ctx, &dynamodb.QueryInput{
			TableName:              aws.String(r.s.table),
			KeyConditionExpression: aws.String("PK = :pk AND begins_with(SK, :prefix)"),
			ExpressionAttributeValues: map[string]types.AttributeValue{
				":pk":     &types.AttributeValueMemberS{Value: pkUser(userID)},
				":prefix": &types.AttributeValueMemberS{Value: prefixSession},
			},
			ScanIndexForward:  aws.Bool(false),
			ConsistentRead:    aws.Bool(true),
			ExclusiveStartKey: startKey,
		})
		if err != nil {
			return nil, fmt.Errorf("query sessions of user %d: %w", userID, err)
		}
		for _, raw := range out.Items {
			var item sessionItem
			if err := attributevalue.UnmarshalMap(raw, &item); err != nil {
				return nil, fmt.Errorf("unmarshal session: %w", err)
			}
			sessions = append(sessions, item)
		}
		if out.LastEvaluatedKey == nil {
			break
		}
		startKey = out.LastEvaluatedKey
	}
	return sessions, nil
}

func (r *gameRepo) SessionsByUser(userID uint, limit int) ([]domain.GameSession, error) {
	items, err := r.userSessions(context.Background(), userID)
	if err != nil {
		return nil, err
	}
	sessions := make([]domain.GameSession, 0, len(items))
	for _, item := range items {
		if limit > 0 && len(sessions) == limit {
			break
		}
		sessions = append(sessions, item.toDomain())
	}
	return sessions, nil
}

func isResumable(item sessionItem, mode string) bool {
	if item.Status != "ACTIVE" && item.Status != "PAUSED" {
		return false
	}
	return mode == "" || item.Mode == mode
}

func (r *gameRepo) LatestResumable(userID uint, mode string) (domain.GameSession, error) {
	items, err := r.userSessions(context.Background(), userID)
	if err != nil {
		return domain.GameSession{}, err
	}
	var latest *sessionItem
	for i := range items {
		item := &items[i]
		if !isResumable(*item, mode) {
			continue
		}
		if latest == nil || item.UpdatedAt.After(latest.UpdatedAt) {
			latest = item
		}
	}
	if latest == nil {
		return domain.GameSession{}, domain.ErrNotFound
	}
	return latest.toDomain(), nil
}

func (r *gameRepo) CompleteResumable(userID uint, mode string, keepID uint) error {
	ctx := context.Background()
	items, err := r.userSessions(ctx, userID)
	if err != nil {
		return err
	}
	now := mustMarshalTime(time.Now())
	for _, item := range items {
		if !isResumable(item, mode) || item.ID == keepID {
			continue
		}
		_, err := r.s.client.UpdateItem(ctx, &dynamodb.UpdateItemInput{
			TableName: aws.String(r.s.table),
			Key: map[string]types.AttributeValue{
				"PK": &types.AttributeValueMemberS{Value: pkUser(userID)},
				"SK": &types.AttributeValueMemberS{Value: item.SK},
			},
			UpdateExpression: aws.String("SET #status = :completed, updated_at = :now"),
			ExpressionAttributeNames: map[string]string{
				"#status": "status",
			},
			ExpressionAttributeValues: map[string]types.AttributeValue{
				":completed": &types.AttributeValueMemberS{Value: "COMPLETED"},
				":now":       now,
			},
		})
		if err != nil {
			return fmt.Errorf("complete session %d: %w", item.ID, err)
		}
	}
	return nil
}

// SaveAnswer inserts (assigning a new ID) or overwrites the answer row for
// (session, question). The row is keyed by question, matching the app's
// one-answer-per-question invariant; filling a flag placeholder keeps its ID.
// The write is conditional so a racing duplicate submission (both readers
// passing HasAnswered before either writes) can't overwrite an already
// scored answer and double-count the session's score.
func (r *gameRepo) SaveAnswer(answer *domain.GameAnswer) error {
	ctx := context.Background()
	if answer.ID == 0 {
		id, err := r.s.nextID(ctx, "answer")
		if err != nil {
			return err
		}
		answer.ID = id
	}

	answerRow := newAnswerItem(*answer)
	if r.isGuestSession(ctx, answer.GameSessionID) {
		answerRow.TTL = guestTTL(time.Now())
	}
	item, err := attributevalue.MarshalMap(answerRow)
	if err != nil {
		return fmt.Errorf("marshal answer: %w", err)
	}
	_, err = r.s.client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName:           aws.String(r.s.table),
		Item:                item,
		ConditionExpression: aws.String("attribute_not_exists(PK) OR attribute_not_exists(choice_id)"),
	})
	var condErr *types.ConditionalCheckFailedException
	if errors.As(err, &condErr) {
		return domain.ErrAlreadyAnswered
	}
	if err != nil {
		return fmt.Errorf("save answer: %w", err)
	}
	return nil
}

func (r *gameRepo) answerRow(ctx context.Context, sessionID, questionID uint) (answerItem, error) {
	var item answerItem
	err := r.s.getItem(ctx, pkSession(sessionID), skAnswer(questionID), &item)
	return item, err
}

func (r *gameRepo) AnswerPlaceholder(sessionID, questionID uint) (domain.GameAnswer, error) {
	item, err := r.answerRow(context.Background(), sessionID, questionID)
	if err != nil {
		return domain.GameAnswer{}, err
	}
	if item.ChoiceID != nil { // answered: no placeholder to return
		return domain.GameAnswer{}, domain.ErrNotFound
	}
	return item.toDomain(), nil
}

func (r *gameRepo) HasAnswered(sessionID, questionID uint) (bool, error) {
	item, err := r.answerRow(context.Background(), sessionID, questionID)
	if err == domain.ErrNotFound {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return item.ChoiceID != nil, nil
}

func (r *gameRepo) ToggleFlag(sessionID, questionID uint) (bool, error) {
	ctx := context.Background()
	item, err := r.answerRow(ctx, sessionID, questionID)
	if err == domain.ErrNotFound {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	_, err = r.s.client.UpdateItem(ctx, &dynamodb.UpdateItemInput{
		TableName: aws.String(r.s.table),
		Key: map[string]types.AttributeValue{
			"PK": &types.AttributeValueMemberS{Value: pkSession(sessionID)},
			"SK": &types.AttributeValueMemberS{Value: skAnswer(questionID)},
		},
		UpdateExpression: aws.String("SET is_flagged = :v"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":v": &types.AttributeValueMemberBOOL{Value: !item.IsFlagged},
		},
	})
	if err != nil {
		return false, fmt.Errorf("toggle flag: %w", err)
	}
	return true, nil
}

// sessionAnswers returns all answer rows of the session (placeholders
// included), sorted by answer ID — SQLite returns insertion order.
func (r *gameRepo) sessionAnswers(ctx context.Context, sessionID uint) ([]answerItem, error) {
	raw, err := r.s.queryPrefix(ctx, pkSession(sessionID), prefixAnswer)
	if err != nil {
		return nil, err
	}
	answers := make([]answerItem, 0, len(raw))
	for _, itemMap := range raw {
		var item answerItem
		if err := attributevalue.UnmarshalMap(itemMap, &item); err != nil {
			return nil, fmt.Errorf("unmarshal answer: %w", err)
		}
		answers = append(answers, item)
	}
	sort.Slice(answers, func(i, j int) bool { return answers[i].ID < answers[j].ID })
	return answers, nil
}

func (r *gameRepo) AnswersBySession(sessionID uint) ([]domain.GameAnswer, error) {
	ctx := context.Background()
	items, err := r.sessionAnswers(ctx, sessionID)
	if err != nil {
		return nil, err
	}
	b, err := r.s.bank(ctx)
	if err != nil {
		return nil, err
	}
	answers := make([]domain.GameAnswer, 0, len(items))
	for _, item := range items {
		a := item.toDomain()
		if q, ok := b.byID[a.QuestionID]; ok {
			a.Question = copyQuestion(q)
		}
		if a.ChoiceID != nil {
			for _, c := range a.Question.Choices {
				if c.ID == *a.ChoiceID {
					choice := c
					a.Choice = &choice
					break
				}
			}
		}
		answers = append(answers, a)
	}
	return answers, nil
}

// IncorrectAnswers matches the SQLite filter `is_correct = false`, which also
// matches flag placeholders (choice_id null) — generateRecommendations relies
// on that. Question comes with taxonomy but without choices, and Choice stays
// nil, exactly like the reference preloads.
func (r *gameRepo) IncorrectAnswers(sessionID uint) ([]domain.GameAnswer, error) {
	ctx := context.Background()
	items, err := r.sessionAnswers(ctx, sessionID)
	if err != nil {
		return nil, err
	}
	b, err := r.s.bank(ctx)
	if err != nil {
		return nil, err
	}
	var answers []domain.GameAnswer
	for _, item := range items {
		if item.IsCorrect {
			continue
		}
		a := item.toDomain()
		if q, ok := b.byID[a.QuestionID]; ok {
			question := copyQuestion(q)
			question.Choices = nil
			a.Question = question
		}
		answers = append(answers, a)
	}
	return answers, nil
}

func (r *gameRepo) AnsweredCount(sessionID uint) (int64, error) {
	items, err := r.sessionAnswers(context.Background(), sessionID)
	if err != nil {
		return 0, err
	}
	var count int64
	for _, item := range items {
		if item.ChoiceID != nil {
			count++
		}
	}
	return count, nil
}

func (r *gameRepo) IncorrectCount(sessionID uint) (int64, error) {
	items, err := r.sessionAnswers(context.Background(), sessionID)
	if err != nil {
		return 0, err
	}
	var count int64
	for _, item := range items {
		if item.ChoiceID != nil && !item.IsCorrect {
			count++
		}
	}
	return count, nil
}

func (r *gameRepo) FlaggedCount(sessionID uint) (int64, error) {
	items, err := r.sessionAnswers(context.Background(), sessionID)
	if err != nil {
		return 0, err
	}
	var count int64
	for _, item := range items {
		if item.IsFlagged {
			count++
		}
	}
	return count, nil
}

func (r *gameRepo) FlaggedQuestionIDs(sessionID uint) ([]uint, error) {
	items, err := r.sessionAnswers(context.Background(), sessionID)
	if err != nil {
		return nil, err
	}
	var ids []uint
	for _, item := range items {
		if item.IsFlagged {
			ids = append(ids, item.QuestionID)
		}
	}
	return ids, nil
}

func (r *gameRepo) AnsweredQuestionIDs(sessionID uint) ([]uint, error) {
	items, err := r.sessionAnswers(context.Background(), sessionID)
	if err != nil {
		return nil, err
	}
	var ids []uint
	for _, item := range items {
		if item.ChoiceID != nil {
			ids = append(ids, item.QuestionID)
		}
	}
	return ids, nil
}

func (r *gameRepo) AnsweredCountByCategory(sessionID uint) (map[string]int64, error) {
	ctx := context.Background()
	items, err := r.sessionAnswers(ctx, sessionID)
	if err != nil {
		return nil, err
	}
	b, err := r.s.bank(ctx)
	if err != nil {
		return nil, err
	}
	counts := make(map[string]int64)
	for _, item := range items {
		if item.ChoiceID == nil {
			continue
		}
		if q, ok := b.byID[item.QuestionID]; ok {
			counts[q.Category.Code]++
		}
	}
	return counts, nil
}

func (r *gameRepo) AddHistory(sessionID, questionID uint) error {
	ctx := context.Background()
	now := time.Now()
	item := historyItem{
		PK: pkSession(sessionID), SK: skHistory(questionID),
		GameSessionID: sessionID, QuestionID: questionID, CreatedAt: now,
	}
	if r.isGuestSession(ctx, sessionID) {
		item.TTL = guestTTL(now)
	}
	return r.s.putItem(ctx, item)
}

func (r *gameRepo) UsedQuestionIDs(sessionID uint) ([]uint, error) {
	raw, err := r.s.queryPrefix(context.Background(), pkSession(sessionID), prefixHistory)
	if err != nil {
		return nil, err
	}
	var ids []uint
	for _, itemMap := range raw {
		var item historyItem
		if err := attributevalue.UnmarshalMap(itemMap, &item); err != nil {
			return nil, fmt.Errorf("unmarshal history: %w", err)
		}
		ids = append(ids, item.QuestionID)
	}
	return ids, nil
}
