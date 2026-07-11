// Package storagetest is a conformance suite for domain.Store implementations.
// It encodes the exact semantics of the reference SQLite adapter so that any
// other adapter (DynamoDB, ...) can be verified against the same contract.
//
// Usage from an adapter package:
//
//	func TestConformance(t *testing.T) {
//	    storagetest.RunStoreSuite(t, func(t *testing.T) domain.Store {
//	        return openFreshEmptyStore(t) // one isolated store per subtest
//	    })
//	}
package storagetest

import (
	"slices"
	"testing"
	"time"

	"quiz-app/internal/domain"
)

// RunStoreSuite exercises the full domain.Store contract. open must return a
// fresh, empty store for each call (each subtest gets its own).
func RunStoreSuite(t *testing.T, open func(t *testing.T) domain.Store) {
	t.Run("SyncQuestionBank", func(t *testing.T) { testSyncQuestionBank(t, open(t)) })
	t.Run("Users", func(t *testing.T) { testUsers(t, open(t)) })
	t.Run("QuestionQueries", func(t *testing.T) { testQuestionQueries(t, open(t)) })
	t.Run("RandomID", func(t *testing.T) { testRandomID(t, open(t)) })
	t.Run("SessionLifecycle", func(t *testing.T) { testSessionLifecycle(t, open(t)) })
	t.Run("AnswerLifecycle", func(t *testing.T) { testAnswerLifecycle(t, open(t)) })
	t.Run("Scoring", func(t *testing.T) { testScoring(t, open(t)) })
	t.Run("Resumable", func(t *testing.T) { testResumable(t, open(t)) })
	t.Run("PauseResumeRoundTrip", func(t *testing.T) { testPauseResumeRoundTrip(t, open(t)) })
	t.Run("SessionQueries", func(t *testing.T) { testSessionQueries(t, open(t)) })
	t.Run("Stats", func(t *testing.T) { testStats(t, open(t)) })
	t.Run("Recommendations", func(t *testing.T) { testRecommendations(t, open(t)) })
	t.Run("ResetUserData", func(t *testing.T) { testResetUserData(t, open(t)) })
}

// --- fixture -----------------------------------------------------------------

func fixtureTaxonomy() []domain.SeedCategory {
	return []domain.SeedCategory{
		{Code: "GEO", Name: "Geografía", SubCategories: []domain.SeedSubCategory{
			{Code: "RIOS", Name: "Ríos"},
			{Code: "MONTANAS", Name: "Montañas"},
		}},
		{Code: "HIS", Name: "Historia", SubCategories: []domain.SeedSubCategory{
			{Code: "PROCERES", Name: "Próceres"},
		}},
	}
}

func fixtureQuestions() []domain.SeedQuestion {
	choices := func() []domain.SeedChoice {
		return []domain.SeedChoice{
			{Text: "opción 1", IsCorrect: false, Order: 1},
			{Text: "opción 2", IsCorrect: true, Order: 2},
			{Text: "opción 3", IsCorrect: false, Order: 3},
			{Text: "opción 4", IsCorrect: false, Order: 4},
		}
	}
	q := func(key, cat, sub, text string) domain.SeedQuestion {
		return domain.SeedQuestion{
			Key: key, Category: cat, SubCategory: sub, Text: text,
			Difficulty: 2, Points: 10, Hint: "pista", Explanation: "explicación",
			Choices: choices(),
		}
	}
	return []domain.SeedQuestion{
		q("GEO-0001", "GEO", "RIOS", "¿Río más largo?"),
		q("GEO-0002", "GEO", "RIOS", "¿Río más caudaloso?"),
		q("GEO-0003", "GEO", "MONTANAS", "¿Pico más alto?"),
		q("HIS-0001", "HIS", "PROCERES", "¿Libertador?"),
		q("HIS-0002", "HIS", "PROCERES", "¿Primer presidente?"),
		q("HIS-0003", "HIS", "PROCERES", "¿Grito de independencia?"),
	}
}

// --- helpers -----------------------------------------------------------------

func mustSync(t *testing.T, st domain.Store) {
	t.Helper()
	if err := st.SyncQuestionBank(fixtureTaxonomy(), fixtureQuestions()); err != nil {
		t.Fatalf("SyncQuestionBank: %v", err)
	}
}

func questionsByKey(t *testing.T, st domain.Store) map[string]domain.Question {
	t.Helper()
	list, err := st.Questions().List()
	if err != nil {
		t.Fatalf("Questions().List: %v", err)
	}
	byKey := make(map[string]domain.Question, len(list))
	for _, q := range list {
		byKey[q.Key] = q
	}
	return byKey
}

func createUser(t *testing.T, st domain.Store, username string) domain.User {
	t.Helper()
	u := domain.User{Username: username, Password: "hash-" + username, Email: username + "@example.com"}
	if err := st.Users().Create(&u); err != nil {
		t.Fatalf("Users().Create(%s): %v", username, err)
	}
	if u.ID == 0 {
		t.Fatalf("Users().Create(%s): ID not assigned", username)
	}
	return u
}

func createSession(t *testing.T, st domain.Store, userID uint, mode, status string) domain.GameSession {
	t.Helper()
	s := domain.GameSession{
		UserID: userID, Mode: mode, Status: status,
		StartTime: time.Now(), TotalQuestions: 10,
	}
	if err := st.Games().CreateSession(&s); err != nil {
		t.Fatalf("CreateSession: %v", err)
	}
	if s.ID == 0 {
		t.Fatal("CreateSession: ID not assigned")
	}
	return s
}

// choiceOf returns the correct or an incorrect choice of the question.
func choiceOf(t *testing.T, q domain.Question, correct bool) domain.Choice {
	t.Helper()
	for _, c := range q.Choices {
		if c.IsCorrect == correct {
			return c
		}
	}
	t.Fatalf("question %s has no choice with IsCorrect=%v", q.Key, correct)
	return domain.Choice{}
}

// answerQuestion simulates submitAnswer: saves a real answer row.
func answerQuestion(t *testing.T, st domain.Store, sessionID uint, q domain.Question, correct bool) domain.GameAnswer {
	t.Helper()
	choice := choiceOf(t, q, correct)
	ans := domain.GameAnswer{
		GameSessionID: sessionID, QuestionID: q.ID,
		ChoiceID: &choice.ID, IsCorrect: correct,
		TimeSpent: 5, AnsweredAt: time.Now(),
	}
	if err := st.Games().SaveAnswer(&ans); err != nil {
		t.Fatalf("SaveAnswer(%s): %v", q.Key, err)
	}
	if ans.ID == 0 {
		t.Fatal("SaveAnswer: ID not assigned")
	}
	return ans
}

// flagPlaceholder simulates flagQuestion on a not-yet-answered question.
func flagPlaceholder(t *testing.T, st domain.Store, sessionID uint, q domain.Question) domain.GameAnswer {
	t.Helper()
	ans := domain.GameAnswer{GameSessionID: sessionID, QuestionID: q.ID, IsFlagged: true}
	if err := st.Games().SaveAnswer(&ans); err != nil {
		t.Fatalf("SaveAnswer placeholder(%s): %v", q.Key, err)
	}
	return ans
}

// sleep guarantees distinguishable timestamps between writes.
func sleep() { time.Sleep(15 * time.Millisecond) }

// --- subtests ----------------------------------------------------------------

func testSyncQuestionBank(t *testing.T, st domain.Store) {
	mustSync(t, st)

	first := questionsByKey(t, st)
	if len(first) != 6 {
		t.Fatalf("List: got %d questions, want 6", len(first))
	}
	for key, q := range first {
		if q.ID == 0 {
			t.Errorf("question %s: ID not assigned", key)
		}
		if q.Category.Code == "" || q.SubCategory.Name == "" {
			t.Errorf("question %s: Category/SubCategory not loaded (cat=%q sub=%q)", key, q.Category.Code, q.SubCategory.Name)
		}
		if len(q.Choices) != 4 {
			t.Errorf("question %s: got %d choices, want 4", key, len(q.Choices))
		}
		correct := 0
		for _, c := range q.Choices {
			if c.ID == 0 {
				t.Errorf("question %s: choice ID not assigned", key)
			}
			if c.QuestionID != q.ID {
				t.Errorf("question %s: choice.QuestionID=%d, want %d", key, c.QuestionID, q.ID)
			}
			if c.IsCorrect {
				correct++
			}
		}
		if correct != 1 {
			t.Errorf("question %s: %d correct choices, want 1", key, correct)
		}
	}

	// Idempotency: a second sync must not create duplicates or change IDs.
	mustSync(t, st)
	second := questionsByKey(t, st)
	if len(second) != 6 {
		t.Fatalf("List after re-sync: got %d questions, want 6", len(second))
	}
	for key, q := range first {
		if second[key].ID != q.ID {
			t.Errorf("question %s: ID changed on re-sync (%d -> %d)", key, q.ID, second[key].ID)
		}
	}

	// Updating a question and its choices keeps IDs stable.
	seeds := fixtureQuestions()
	seeds[0].Text = "¿Río más largo? (v2)"
	seeds[0].Choices[0].Text = "opción 1 (v2)"
	if err := st.SyncQuestionBank(fixtureTaxonomy(), seeds); err != nil {
		t.Fatalf("SyncQuestionBank update: %v", err)
	}
	updated := questionsByKey(t, st)["GEO-0001"]
	if updated.ID != first["GEO-0001"].ID {
		t.Errorf("GEO-0001: ID changed on update (%d -> %d)", first["GEO-0001"].ID, updated.ID)
	}
	if updated.Text != "¿Río más largo? (v2)" {
		t.Errorf("GEO-0001: text not updated, got %q", updated.Text)
	}
	oldChoice := first["GEO-0001"].Choices[0]
	var newChoice domain.Choice
	for _, c := range updated.Choices {
		if c.Order == oldChoice.Order {
			newChoice = c
		}
	}
	if newChoice.ID != oldChoice.ID {
		t.Errorf("GEO-0001 choice order %d: ID changed on update (%d -> %d)", oldChoice.Order, oldChoice.ID, newChoice.ID)
	}
	if newChoice.Text != "opción 1 (v2)" {
		t.Errorf("GEO-0001 choice: text not updated, got %q", newChoice.Text)
	}

	// Categories: ordered by ID, subcategories loaded.
	cats, err := st.Questions().Categories()
	if err != nil {
		t.Fatalf("Categories: %v", err)
	}
	if len(cats) != 2 {
		t.Fatalf("Categories: got %d, want 2", len(cats))
	}
	if !slices.IsSortedFunc(cats, func(a, b domain.Category) int { return int(a.ID) - int(b.ID) }) {
		t.Error("Categories: not ordered by ID")
	}
	subCounts := map[string]int{"GEO": 2, "HIS": 1}
	for _, c := range cats {
		if len(c.SubCategories) != subCounts[c.Code] {
			t.Errorf("category %s: got %d subcategories, want %d", c.Code, len(c.SubCategories), subCounts[c.Code])
		}
	}
}

func testUsers(t *testing.T, st domain.Store) {
	mustSync(t, st)
	users := st.Users()

	u := createUser(t, st, "ana")

	got, err := users.ByUsername("ana")
	if err != nil {
		t.Fatalf("ByUsername: %v", err)
	}
	if got.ID != u.ID || got.Email != u.Email || got.Password != u.Password {
		t.Errorf("ByUsername: got %+v, want id=%d email=%s", got, u.ID, u.Email)
	}

	if _, err := users.ByUsername("nadie"); err != domain.ErrNotFound {
		t.Errorf("ByUsername(missing): got %v, want ErrNotFound", err)
	}

	// Duplicate username and duplicate email must fail.
	dupU := domain.User{Username: "ana", Password: "x", Email: "otra@example.com"}
	if err := users.Create(&dupU); err == nil {
		t.Error("Create duplicate username: expected error, got nil")
	}
	dupE := domain.User{Username: "otra", Password: "x", Email: u.Email}
	if err := users.Create(&dupE); err == nil {
		t.Error("Create duplicate email: expected error, got nil")
	}

	// Token save + lookup.
	exp := time.Now().Add(24 * time.Hour).UTC().Truncate(time.Second)
	if err := users.SaveSessionToken(u.ID, "token-1", exp); err != nil {
		t.Fatalf("SaveSessionToken: %v", err)
	}
	byTok, err := users.ByToken("token-1")
	if err != nil {
		t.Fatalf("ByToken: %v", err)
	}
	if byTok.ID != u.ID {
		t.Errorf("ByToken: got user %d, want %d", byTok.ID, u.ID)
	}
	if byTok.TokenExpiresAt == nil || !byTok.TokenExpiresAt.Equal(exp) {
		t.Errorf("ByToken: TokenExpiresAt=%v, want %v", byTok.TokenExpiresAt, exp)
	}
	byName, err := users.ByUsername("ana")
	if err != nil {
		t.Fatalf("ByUsername after token: %v", err)
	}
	if byName.Token != "token-1" {
		t.Errorf("ByUsername: Token=%q, want token-1", byName.Token)
	}

	// Rotation: the new token works, the old one is invalidated.
	if err := users.SaveSessionToken(u.ID, "token-2", exp); err != nil {
		t.Fatalf("SaveSessionToken rotate: %v", err)
	}
	if got, err := users.ByToken("token-2"); err != nil || got.ID != u.ID {
		t.Errorf("ByToken(new): got (%v, %v), want user %d", got.ID, err, u.ID)
	}
	if _, err := users.ByToken("token-1"); err != domain.ErrNotFound {
		t.Errorf("ByToken(old after rotation): got %v, want ErrNotFound", err)
	}
	if _, err := users.ByToken("no-such-token"); err != domain.ErrNotFound {
		t.Errorf("ByToken(missing): got %v, want ErrNotFound", err)
	}
}

func testQuestionQueries(t *testing.T, st domain.Store) {
	mustSync(t, st)
	qs := st.Questions()
	byKey := questionsByKey(t, st)

	// ByID with relations loaded.
	q, err := qs.ByID(byKey["HIS-0001"].ID)
	if err != nil {
		t.Fatalf("ByID: %v", err)
	}
	if q.Key != "HIS-0001" || q.Category.Code != "HIS" || q.SubCategory.Name != "Próceres" || len(q.Choices) != 4 {
		t.Errorf("ByID: incomplete question %+v", q)
	}
	if _, err := qs.ByID(99999); err != domain.ErrNotFound {
		t.Errorf("ByID(missing): got %v, want ErrNotFound", err)
	}

	// ListByCategory.
	geo, err := qs.ListByCategory("GEO")
	if err != nil {
		t.Fatalf("ListByCategory: %v", err)
	}
	if len(geo) != 3 {
		t.Errorf("ListByCategory(GEO): got %d, want 3", len(geo))
	}
	for _, q := range geo {
		if q.Category.Code != "GEO" {
			t.Errorf("ListByCategory(GEO): question %s has category %s", q.Key, q.Category.Code)
		}
	}

	// Counts.
	if n, err := qs.Count(nil); err != nil || n != 6 {
		t.Errorf("Count(nil): got (%d, %v), want 6", n, err)
	}
	if n, err := qs.Count([]string{"HIS"}); err != nil || n != 3 {
		t.Errorf("Count(HIS): got (%d, %v), want 3", n, err)
	}
	if n, err := qs.Count([]string{"GEO", "HIS"}); err != nil || n != 6 {
		t.Errorf("Count(GEO,HIS): got (%d, %v), want 6", n, err)
	}

	counts, err := qs.CountsByCategory()
	if err != nil {
		t.Fatalf("CountsByCategory: %v", err)
	}
	if counts["GEO"] != 3 || counts["HIS"] != 3 || len(counts) != 2 {
		t.Errorf("CountsByCategory: got %v", counts)
	}

	subCounts, err := qs.CountsBySubcategory()
	if err != nil {
		t.Fatalf("CountsBySubcategory: %v", err)
	}
	want := map[string]int64{"GEO/Ríos": 2, "GEO/Montañas": 1, "HIS/Próceres": 3}
	gotSub := make(map[string]int64, len(subCounts))
	for _, row := range subCounts {
		gotSub[row.Category+"/"+row.SubCategory] = row.Count
	}
	for k, v := range want {
		if gotSub[k] != v {
			t.Errorf("CountsBySubcategory[%s]: got %d, want %d (all: %v)", k, gotSub[k], v, gotSub)
		}
	}

	// Choice lookups.
	someQ := byKey["GEO-0002"]
	correct := choiceOf(t, someQ, true)
	c, err := qs.ChoiceByID(correct.ID)
	if err != nil {
		t.Fatalf("ChoiceByID: %v", err)
	}
	if c.QuestionID != someQ.ID || !c.IsCorrect {
		t.Errorf("ChoiceByID: got %+v, want correct choice of question %d", c, someQ.ID)
	}
	if _, err := qs.ChoiceByID(9999999); err != domain.ErrNotFound {
		t.Errorf("ChoiceByID(missing): got %v, want ErrNotFound", err)
	}

	ids := []uint{byKey["GEO-0001"].ID, byKey["HIS-0003"].ID}
	correctMap, err := qs.CorrectChoices(ids)
	if err != nil {
		t.Fatalf("CorrectChoices: %v", err)
	}
	if len(correctMap) != 2 {
		t.Fatalf("CorrectChoices: got %d entries, want 2", len(correctMap))
	}
	for _, id := range ids {
		if ch, ok := correctMap[id]; !ok || !ch.IsCorrect || ch.QuestionID != id {
			t.Errorf("CorrectChoices[%d]: got %+v", id, ch)
		}
	}
	if empty, err := qs.CorrectChoices(nil); err != nil || len(empty) != 0 {
		t.Errorf("CorrectChoices(nil): got (%v, %v), want empty map", empty, err)
	}

	// Results must be safe to mutate (the HTTP layer shuffles choices and
	// strips IsCorrect); a second read must not see those mutations.
	fresh, err := qs.ByID(someQ.ID)
	if err != nil {
		t.Fatalf("ByID: %v", err)
	}
	for i := range fresh.Choices {
		fresh.Choices[i].IsCorrect = false
		fresh.Choices[i].Text = "mutated"
	}
	again, err := qs.ByID(someQ.ID)
	if err != nil {
		t.Fatalf("ByID: %v", err)
	}
	if choiceOf(t, again, true).Text == "mutated" {
		t.Error("ByID: mutation of a previous result leaked into a fresh read")
	}
}

func testRandomID(t *testing.T, st domain.Store) {
	mustSync(t, st)
	qs := st.Questions()
	byKey := questionsByKey(t, st)

	geoIDs := map[uint]bool{}
	var geoList []uint
	for _, key := range []string{"GEO-0001", "GEO-0002", "GEO-0003"} {
		geoIDs[byKey[key].ID] = true
		geoList = append(geoList, byKey[key].ID)
	}

	// Category restriction holds over repeated draws.
	for i := 0; i < 20; i++ {
		id, err := qs.RandomID([]string{"GEO"}, nil)
		if err != nil {
			t.Fatalf("RandomID(GEO): %v", err)
		}
		if !geoIDs[id] {
			t.Fatalf("RandomID(GEO): returned non-GEO question %d", id)
		}
	}

	// Exclusions hold.
	exclude := geoList[:2]
	for i := 0; i < 10; i++ {
		id, err := qs.RandomID([]string{"GEO"}, exclude)
		if err != nil {
			t.Fatalf("RandomID(GEO, exclude): %v", err)
		}
		if id != geoList[2] {
			t.Fatalf("RandomID(GEO, exclude 2 of 3): got %d, want %d", id, geoList[2])
		}
	}

	// Exhaustion -> ErrNotFound.
	if _, err := qs.RandomID([]string{"GEO"}, geoList); err != domain.ErrNotFound {
		t.Errorf("RandomID(GEO, all excluded): got %v, want ErrNotFound", err)
	}
	if _, err := qs.RandomID([]string{"NOPE"}, nil); err != domain.ErrNotFound {
		t.Errorf("RandomID(unknown category): got %v, want ErrNotFound", err)
	}

	// Unrestricted draw returns a valid question.
	id, err := qs.RandomID(nil, nil)
	if err != nil {
		t.Fatalf("RandomID(nil): %v", err)
	}
	if _, err := qs.ByID(id); err != nil {
		t.Errorf("RandomID(nil): returned unknown id %d", id)
	}
}

func testSessionLifecycle(t *testing.T, st domain.Store) {
	mustSync(t, st)
	u := createUser(t, st, "ana")
	games := st.Games()

	s := createSession(t, st, u.ID, "PRACTICE", "ACTIVE")

	got, err := games.SessionByID(s.ID)
	if err != nil {
		t.Fatalf("SessionByID: %v", err)
	}
	if got.UserID != u.ID || got.Mode != "PRACTICE" || got.Status != "ACTIVE" || got.TotalQuestions != 10 {
		t.Errorf("SessionByID: got %+v", got)
	}
	if _, err := games.SessionByID(99999); err != domain.ErrNotFound {
		t.Errorf("SessionByID(missing): got %v, want ErrNotFound", err)
	}

	// SaveSession writes every field.
	got.Status = "COMPLETED"
	got.TimeElapsed = 300
	got.Score = 70
	now := time.Now()
	got.EndTime = &now
	if err := games.SaveSession(&got); err != nil {
		t.Fatalf("SaveSession: %v", err)
	}
	reloaded, err := games.SessionByID(s.ID)
	if err != nil {
		t.Fatalf("SessionByID after save: %v", err)
	}
	if reloaded.Status != "COMPLETED" || reloaded.TimeElapsed != 300 || reloaded.Score != 70 || reloaded.EndTime == nil {
		t.Errorf("SaveSession round-trip: got %+v", reloaded)
	}
}

func testAnswerLifecycle(t *testing.T, st domain.Store) {
	mustSync(t, st)
	u := createUser(t, st, "ana")
	games := st.Games()
	byKey := questionsByKey(t, st)
	s := createSession(t, st, u.ID, "PRACTICE", "ACTIVE")
	q := byKey["GEO-0001"]

	// Flag before answering -> placeholder row.
	placeholder := flagPlaceholder(t, st, s.ID, q)

	gotPH, err := games.AnswerPlaceholder(s.ID, q.ID)
	if err != nil {
		t.Fatalf("AnswerPlaceholder: %v", err)
	}
	if gotPH.ChoiceID != nil || !gotPH.IsFlagged {
		t.Errorf("AnswerPlaceholder: got %+v", gotPH)
	}
	if answered, err := games.HasAnswered(s.ID, q.ID); err != nil || answered {
		t.Errorf("HasAnswered(placeholder): got (%v, %v), want false", answered, err)
	}
	if n, err := games.AnsweredCount(s.ID); err != nil || n != 0 {
		t.Errorf("AnsweredCount(placeholder only): got (%d, %v), want 0", n, err)
	}
	if n, err := games.FlaggedCount(s.ID); err != nil || n != 1 {
		t.Errorf("FlaggedCount: got (%d, %v), want 1", n, err)
	}
	// Reference quirk: IncorrectAnswers includes placeholders (is_correct=false,
	// choice_id null), while IncorrectCount only counts real answers.
	if inc, err := games.IncorrectAnswers(s.ID); err != nil || len(inc) != 1 {
		t.Errorf("IncorrectAnswers(placeholder): got (%d, %v), want 1", len(inc), err)
	}
	if n, err := games.IncorrectCount(s.ID); err != nil || n != 0 {
		t.Errorf("IncorrectCount(placeholder): got (%d, %v), want 0", n, err)
	}

	// Fill the placeholder (what submitAnswer does): same row, ID preserved.
	choice := choiceOf(t, q, false)
	filled := gotPH
	filled.ChoiceID = &choice.ID
	filled.IsCorrect = false
	filled.TimeSpent = 12
	filled.AnsweredAt = time.Now()
	if err := games.SaveAnswer(&filled); err != nil {
		t.Fatalf("SaveAnswer(fill placeholder): %v", err)
	}

	if answered, err := games.HasAnswered(s.ID, q.ID); err != nil || !answered {
		t.Errorf("HasAnswered(filled): got (%v, %v), want true", answered, err)
	}
	if _, err := games.AnswerPlaceholder(s.ID, q.ID); err != domain.ErrNotFound {
		t.Errorf("AnswerPlaceholder(filled): got %v, want ErrNotFound", err)
	}
	if n, err := games.AnsweredCount(s.ID); err != nil || n != 1 {
		t.Errorf("AnsweredCount(filled): got (%d, %v), want 1", n, err)
	}
	if n, err := games.IncorrectCount(s.ID); err != nil || n != 1 {
		t.Errorf("IncorrectCount(filled wrong): got (%d, %v), want 1", n, err)
	}

	answers, err := games.AnswersBySession(s.ID)
	if err != nil {
		t.Fatalf("AnswersBySession: %v", err)
	}
	if len(answers) != 1 {
		t.Fatalf("AnswersBySession: got %d rows, want 1 (placeholder filled in place)", len(answers))
	}
	a := answers[0]
	if a.ID != placeholder.ID {
		t.Errorf("filled answer: ID changed (%d -> %d)", placeholder.ID, a.ID)
	}
	if a.Question.ID != q.ID || a.Question.Category.Code != "GEO" || len(a.Question.Choices) != 4 {
		t.Errorf("AnswersBySession: Question not fully loaded: %+v", a.Question)
	}
	if a.Choice == nil || a.Choice.ID != choice.ID {
		t.Errorf("AnswersBySession: Choice not loaded: %+v", a.Choice)
	}
	if !a.IsFlagged {
		t.Error("filled answer: IsFlagged lost")
	}

	// ToggleFlag flips; missing row -> (false, nil).
	if on, err := games.ToggleFlag(s.ID, q.ID); err != nil || !on {
		t.Errorf("ToggleFlag: got (%v, %v), want (true, nil)", on, err)
	}
	if n, _ := games.FlaggedCount(s.ID); n != 0 {
		t.Errorf("FlaggedCount after unflag: got %d, want 0", n)
	}
	if on, err := games.ToggleFlag(s.ID, q.ID); err != nil || !on {
		t.Errorf("ToggleFlag again: got (%v, %v), want (true, nil)", on, err)
	}
	if n, _ := games.FlaggedCount(s.ID); n != 1 {
		t.Errorf("FlaggedCount after re-flag: got %d, want 1", n)
	}
	if ok, err := games.ToggleFlag(s.ID, byKey["HIS-0001"].ID); err != nil || ok {
		t.Errorf("ToggleFlag(no row): got (%v, %v), want (false, nil)", ok, err)
	}

	// Direct answer (no placeholder) on another question.
	answerQuestion(t, st, s.ID, byKey["GEO-0002"], true)
	if n, _ := games.AnsweredCount(s.ID); n != 2 {
		t.Errorf("AnsweredCount: got %d, want 2", n)
	}
	gotIDs, err := games.AnsweredQuestionIDs(s.ID)
	if err != nil {
		t.Fatalf("AnsweredQuestionIDs: %v", err)
	}
	wantIDs := []uint{q.ID, byKey["GEO-0002"].ID}
	slices.Sort(gotIDs)
	slices.Sort(wantIDs)
	if !slices.Equal(gotIDs, wantIDs) {
		t.Errorf("AnsweredQuestionIDs: got %v, want %v", gotIDs, wantIDs)
	}
	flagged, err := games.FlaggedQuestionIDs(s.ID)
	if err != nil {
		t.Fatalf("FlaggedQuestionIDs: %v", err)
	}
	if !slices.Equal(flagged, []uint{q.ID}) {
		t.Errorf("FlaggedQuestionIDs: got %v, want [%d]", flagged, q.ID)
	}
}

func testScoring(t *testing.T, st domain.Store) {
	mustSync(t, st)
	u := createUser(t, st, "ana")
	games := st.Games()
	s := createSession(t, st, u.ID, "PRACTICE", "ACTIVE")

	before, err := games.SessionByID(s.ID)
	if err != nil {
		t.Fatalf("SessionByID: %v", err)
	}

	sleep()
	if err := games.AddScore(s.ID, 10); err != nil {
		t.Fatalf("AddScore: %v", err)
	}
	if err := games.AddScore(s.ID, 15); err != nil {
		t.Fatalf("AddScore: %v", err)
	}

	after, err := games.SessionByID(s.ID)
	if err != nil {
		t.Fatalf("SessionByID: %v", err)
	}
	if after.CorrectAnswers != 2 || after.Score != 25 {
		t.Errorf("AddScore: got correct=%d score=%d, want 2/25", after.CorrectAnswers, after.Score)
	}
	// AddScore must NOT bump updated_at (LatestResumable ordering depends on it).
	if !after.UpdatedAt.Equal(before.UpdatedAt) {
		t.Errorf("AddScore bumped updated_at: %v -> %v", before.UpdatedAt, after.UpdatedAt)
	}

	// SaveSession DOES bump updated_at.
	sleep()
	if err := games.SaveSession(&after); err != nil {
		t.Fatalf("SaveSession: %v", err)
	}
	bumped, err := games.SessionByID(s.ID)
	if err != nil {
		t.Fatalf("SessionByID: %v", err)
	}
	if !bumped.UpdatedAt.After(before.UpdatedAt) {
		t.Errorf("SaveSession did not bump updated_at: %v -> %v", before.UpdatedAt, bumped.UpdatedAt)
	}
}

func testResumable(t *testing.T, st domain.Store) {
	mustSync(t, st)
	u := createUser(t, st, "ana")
	other := createUser(t, st, "otro")
	games := st.Games()

	a := createSession(t, st, u.ID, "PRACTICE", "ACTIVE")
	sleep()
	b := createSession(t, st, u.ID, "TIMED", "PAUSED")
	sleep()
	createSession(t, st, u.ID, "PRACTICE", "COMPLETED")
	sleep()
	createSession(t, st, other.ID, "PRACTICE", "ACTIVE") // other user's, must never surface

	// Latest resumable across modes is the most recently updated ACTIVE/PAUSED.
	latest, err := games.LatestResumable(u.ID, "")
	if err != nil {
		t.Fatalf("LatestResumable: %v", err)
	}
	if latest.ID != b.ID {
		t.Errorf("LatestResumable: got session %d, want %d", latest.ID, b.ID)
	}

	// Mode filter.
	latestPractice, err := games.LatestResumable(u.ID, "PRACTICE")
	if err != nil {
		t.Fatalf("LatestResumable(PRACTICE): %v", err)
	}
	if latestPractice.ID != a.ID {
		t.Errorf("LatestResumable(PRACTICE): got %d, want %d", latestPractice.ID, a.ID)
	}
	if _, err := games.LatestResumable(u.ID, "WEAK_AREAS"); err != domain.ErrNotFound {
		t.Errorf("LatestResumable(no match): got %v, want ErrNotFound", err)
	}

	// Touching A via SaveSession makes it the latest.
	sleep()
	aLoaded, _ := games.SessionByID(a.ID)
	if err := games.SaveSession(&aLoaded); err != nil {
		t.Fatalf("SaveSession: %v", err)
	}
	latest, err = games.LatestResumable(u.ID, "")
	if err != nil {
		t.Fatalf("LatestResumable after touch: %v", err)
	}
	if latest.ID != a.ID {
		t.Errorf("LatestResumable after touch: got %d, want %d", latest.ID, a.ID)
	}

	// CompleteResumable with keepID completes everything else.
	if err := games.CompleteResumable(u.ID, "", b.ID); err != nil {
		t.Fatalf("CompleteResumable: %v", err)
	}
	if got, _ := games.SessionByID(a.ID); got.Status != "COMPLETED" {
		t.Errorf("CompleteResumable: session A status=%s, want COMPLETED", got.Status)
	}
	if got, _ := games.SessionByID(b.ID); got.Status != "PAUSED" {
		t.Errorf("CompleteResumable: kept session B status=%s, want PAUSED", got.Status)
	}
	// Other user's session untouched.
	if lr, err := games.LatestResumable(other.ID, ""); err != nil || lr.UserID != other.ID {
		t.Errorf("other user's resumable affected: (%+v, %v)", lr, err)
	}

	// Completing the rest empties the resumable set.
	if err := games.CompleteResumable(u.ID, "", 0); err != nil {
		t.Fatalf("CompleteResumable(all): %v", err)
	}
	if _, err := games.LatestResumable(u.ID, ""); err != domain.ErrNotFound {
		t.Errorf("LatestResumable after completing all: got %v, want ErrNotFound", err)
	}
}

func testPauseResumeRoundTrip(t *testing.T, st domain.Store) {
	mustSync(t, st)
	u := createUser(t, st, "ana")
	games := st.Games()
	s := createSession(t, st, u.ID, "PRACTICE", "ACTIVE")

	// Pause: PausedAt set.
	loaded, _ := games.SessionByID(s.ID)
	pausedAt := time.Now()
	loaded.Status = "PAUSED"
	loaded.PausedAt = &pausedAt
	loaded.TimeElapsed = 120
	if err := games.SaveSession(&loaded); err != nil {
		t.Fatalf("SaveSession(pause): %v", err)
	}
	paused, _ := games.SessionByID(s.ID)
	if paused.PausedAt == nil {
		t.Fatal("pause: PausedAt not persisted")
	}
	if paused.Status != "PAUSED" || paused.TimeElapsed != 120 {
		t.Errorf("pause: got %+v", paused)
	}

	// Resume: PausedAt must round-trip back to nil.
	paused.Status = "ACTIVE"
	paused.PausedAt = nil
	if err := games.SaveSession(&paused); err != nil {
		t.Fatalf("SaveSession(resume): %v", err)
	}
	resumed, _ := games.SessionByID(s.ID)
	if resumed.PausedAt != nil {
		t.Errorf("resume: PausedAt=%v, want nil", resumed.PausedAt)
	}
	if resumed.EndTime != nil {
		t.Errorf("resume: EndTime=%v, want nil", resumed.EndTime)
	}
	if resumed.Status != "ACTIVE" {
		t.Errorf("resume: Status=%s, want ACTIVE", resumed.Status)
	}
}

func testSessionQueries(t *testing.T, st domain.Store) {
	mustSync(t, st)
	u := createUser(t, st, "ana")
	games := st.Games()
	byKey := questionsByKey(t, st)

	s1 := createSession(t, st, u.ID, "PRACTICE", "COMPLETED")
	sleep()
	s2 := createSession(t, st, u.ID, "TIMED", "COMPLETED")
	sleep()
	s3 := createSession(t, st, u.ID, "PRACTICE", "ACTIVE")

	// SessionsByUser: newest first, limit respected.
	sessions, err := games.SessionsByUser(u.ID, 2)
	if err != nil {
		t.Fatalf("SessionsByUser: %v", err)
	}
	if len(sessions) != 2 || sessions[0].ID != s3.ID || sessions[1].ID != s2.ID {
		ids := make([]uint, len(sessions))
		for i, s := range sessions {
			ids[i] = s.ID
		}
		t.Errorf("SessionsByUser(limit 2): got %v, want [%d %d]", ids, s3.ID, s2.ID)
	}
	all, err := games.SessionsByUser(u.ID, 10)
	if err != nil || len(all) != 3 {
		t.Errorf("SessionsByUser(limit 10): got (%d, %v), want 3", len(all), err)
	}
	_ = s1

	// AnsweredCountByCategory counts real answers per category code.
	answerQuestion(t, st, s3.ID, byKey["GEO-0001"], true)
	answerQuestion(t, st, s3.ID, byKey["GEO-0002"], false)
	answerQuestion(t, st, s3.ID, byKey["HIS-0001"], true)
	flagPlaceholder(t, st, s3.ID, byKey["GEO-0003"]) // placeholder: not counted

	counts, err := games.AnsweredCountByCategory(s3.ID)
	if err != nil {
		t.Fatalf("AnsweredCountByCategory: %v", err)
	}
	if counts["GEO"] != 2 || counts["HIS"] != 1 {
		t.Errorf("AnsweredCountByCategory: got %v, want GEO:2 HIS:1", counts)
	}

	// Question history.
	if err := games.AddHistory(s3.ID, byKey["GEO-0001"].ID); err != nil {
		t.Fatalf("AddHistory: %v", err)
	}
	if err := games.AddHistory(s3.ID, byKey["HIS-0002"].ID); err != nil {
		t.Fatalf("AddHistory: %v", err)
	}
	used, err := games.UsedQuestionIDs(s3.ID)
	if err != nil {
		t.Fatalf("UsedQuestionIDs: %v", err)
	}
	wantUsed := []uint{byKey["GEO-0001"].ID, byKey["HIS-0002"].ID}
	slices.Sort(used)
	slices.Sort(wantUsed)
	if !slices.Equal(used, wantUsed) {
		t.Errorf("UsedQuestionIDs: got %v, want %v", used, wantUsed)
	}
	if other, err := games.UsedQuestionIDs(s2.ID); err != nil || len(other) != 0 {
		t.Errorf("UsedQuestionIDs(other session): got (%v, %v), want empty", other, err)
	}
}

func testStats(t *testing.T, st domain.Store) {
	mustSync(t, st)
	u := createUser(t, st, "ana")
	games := st.Games()
	stats := st.Stats()
	byKey := questionsByKey(t, st)

	// No data yet.
	if n, err := stats.TotalSessions(u.ID); err != nil || n != 0 {
		t.Errorf("TotalSessions(empty): got (%d, %v), want 0", n, err)
	}
	if best, err := stats.BestScore(u.ID); err != nil || best != 0 {
		t.Errorf("BestScore(empty): got (%d, %v), want 0", best, err)
	}
	if totals, err := stats.OverallTotals(u.ID); err != nil || totals.Total != 0 {
		t.Errorf("OverallTotals(empty): got (%+v, %v)", totals, err)
	}

	// Session 1 (COMPLETED, score 30): GEO 1/2 correct, HIS 1/1 correct.
	s1 := createSession(t, st, u.ID, "PRACTICE", "COMPLETED")
	answerQuestion(t, st, s1.ID, byKey["GEO-0001"], true)
	answerQuestion(t, st, s1.ID, byKey["GEO-0002"], false)
	answerQuestion(t, st, s1.ID, byKey["HIS-0001"], true)
	flagPlaceholder(t, st, s1.ID, byKey["HIS-0002"]) // must not count anywhere
	s1L, _ := games.SessionByID(s1.ID)
	s1L.Score = 30
	if err := games.SaveSession(&s1L); err != nil {
		t.Fatal(err)
	}

	sleep()
	// Session 2 (COMPLETED, score 10): GEO 0/2 correct.
	s2 := createSession(t, st, u.ID, "TIMED", "COMPLETED")
	answerQuestion(t, st, s2.ID, byKey["GEO-0001"], false)
	answerQuestion(t, st, s2.ID, byKey["GEO-0003"], false)
	s2L, _ := games.SessionByID(s2.ID)
	s2L.Score = 10
	if err := games.SaveSession(&s2L); err != nil {
		t.Fatal(err)
	}

	sleep()
	// Session 3 (ACTIVE, score 99): must not count toward BestScore/RecentProgress.
	s3 := createSession(t, st, u.ID, "PRACTICE", "ACTIVE")
	answerQuestion(t, st, s3.ID, byKey["HIS-0003"], true)
	s3L, _ := games.SessionByID(s3.ID)
	s3L.Score = 99
	if err := games.SaveSession(&s3L); err != nil {
		t.Fatal(err)
	}

	if n, err := stats.TotalSessions(u.ID); err != nil || n != 3 {
		t.Errorf("TotalSessions: got (%d, %v), want 3", n, err)
	}
	if best, err := stats.BestScore(u.ID); err != nil || best != 30 {
		t.Errorf("BestScore: got (%d, %v), want 30 (COMPLETED only)", best, err)
	}

	totals, err := stats.OverallTotals(u.ID)
	if err != nil {
		t.Fatalf("OverallTotals: %v", err)
	}
	if totals.Total != 6 || totals.Correct != 3 {
		t.Errorf("OverallTotals: got %+v, want {6 3}", totals)
	}

	geoTotals, err := stats.CategoryTotals(u.ID, "GEO")
	if err != nil {
		t.Fatalf("CategoryTotals: %v", err)
	}
	if geoTotals.Total != 4 || geoTotals.Correct != 1 {
		t.Errorf("CategoryTotals(GEO): got %+v, want {4 1}", geoTotals)
	}
	hisTotals, err := stats.CategoryTotals(u.ID, "HIS")
	if err != nil {
		t.Fatalf("CategoryTotals: %v", err)
	}
	if hisTotals.Total != 2 || hisTotals.Correct != 2 {
		t.Errorf("CategoryTotals(HIS): got %+v, want {2 2}", hisTotals)
	}

	// RecentProgress: COMPLETED only, newest first, real answers only.
	progress, err := stats.RecentProgress(u.ID, 5)
	if err != nil {
		t.Fatalf("RecentProgress: %v", err)
	}
	if len(progress) != 2 {
		t.Fatalf("RecentProgress: got %d entries, want 2", len(progress))
	}
	if progress[0].Score != 10 || progress[0].Answered != 2 || progress[0].Correct != 0 {
		t.Errorf("RecentProgress[0]: got %+v, want score=10 answered=2 correct=0", progress[0])
	}
	if progress[1].Score != 30 || progress[1].Answered != 3 || progress[1].Correct != 2 {
		t.Errorf("RecentProgress[1]: got %+v, want score=30 answered=3 correct=2", progress[1])
	}
	if limited, err := stats.RecentProgress(u.ID, 1); err != nil || len(limited) != 1 || limited[0].Score != 10 {
		t.Errorf("RecentProgress(limit 1): got (%+v, %v)", limited, err)
	}

	// WeakCategories: accuracy < 50%, most answered first.
	// GEO: 1/4 = 25% (weak). HIS: 3/3 = 100% (not weak).
	weak, err := stats.WeakCategories(u.ID)
	if err != nil {
		t.Fatalf("WeakCategories: %v", err)
	}
	if !slices.Equal(weak, []string{"GEO"}) {
		t.Errorf("WeakCategories: got %v, want [GEO]", weak)
	}
}

func testRecommendations(t *testing.T, st domain.Store) {
	mustSync(t, st)
	u := createUser(t, st, "ana")
	stats := st.Stats()

	for i, prio := range []int{2, 5, 3} {
		rec := domain.StudyRecommendation{
			UserID: u.ID, Category: "GEO", SubCategory: "RIOS",
			Weakness: "ríos", Description: "estudiar", Resources: "[]",
			Priority: prio,
		}
		if err := stats.CreateRecommendation(&rec); err != nil {
			t.Fatalf("CreateRecommendation(%d): %v", i, err)
		}
		if rec.ID == 0 {
			t.Fatal("CreateRecommendation: ID not assigned")
		}
	}

	recs, err := stats.RecommendationsByUser(u.ID, 0)
	if err != nil {
		t.Fatalf("RecommendationsByUser(0): %v", err)
	}
	if len(recs) != 3 {
		t.Fatalf("RecommendationsByUser(0=all): got %d, want 3", len(recs))
	}
	if recs[0].Priority != 5 || recs[1].Priority != 3 || recs[2].Priority != 2 {
		t.Errorf("RecommendationsByUser: priorities %d,%d,%d, want 5,3,2",
			recs[0].Priority, recs[1].Priority, recs[2].Priority)
	}

	limited, err := stats.RecommendationsByUser(u.ID, 2)
	if err != nil || len(limited) != 2 || limited[0].Priority != 5 {
		t.Errorf("RecommendationsByUser(limit 2): got (%d entries, %v)", len(limited), err)
	}

	if none, err := stats.RecommendationsByUser(99999, 0); err != nil || len(none) != 0 {
		t.Errorf("RecommendationsByUser(unknown user): got (%v, %v), want empty", none, err)
	}
}

func testResetUserData(t *testing.T, st domain.Store) {
	mustSync(t, st)
	u := createUser(t, st, "ana")
	other := createUser(t, st, "otro")
	games := st.Games()
	stats := st.Stats()
	byKey := questionsByKey(t, st)

	exp := time.Now().Add(time.Hour)
	if err := st.Users().SaveSessionToken(u.ID, "tok-ana", exp); err != nil {
		t.Fatal(err)
	}

	s := createSession(t, st, u.ID, "PRACTICE", "ACTIVE")
	answerQuestion(t, st, s.ID, byKey["GEO-0001"], true)
	if err := games.AddHistory(s.ID, byKey["GEO-0001"].ID); err != nil {
		t.Fatal(err)
	}
	rec := domain.StudyRecommendation{UserID: u.ID, Category: "GEO", Priority: 1}
	if err := stats.CreateRecommendation(&rec); err != nil {
		t.Fatal(err)
	}

	os := createSession(t, st, other.ID, "PRACTICE", "ACTIVE")
	answerQuestion(t, st, os.ID, byKey["HIS-0001"], true)

	if err := st.ResetUserData(u.ID); err != nil {
		t.Fatalf("ResetUserData: %v", err)
	}

	// User's game data is gone.
	if _, err := games.SessionByID(s.ID); err != domain.ErrNotFound {
		t.Errorf("session after reset: got %v, want ErrNotFound", err)
	}
	if sessions, _ := games.SessionsByUser(u.ID, 10); len(sessions) != 0 {
		t.Errorf("SessionsByUser after reset: got %d, want 0", len(sessions))
	}
	if n, _ := stats.TotalSessions(u.ID); n != 0 {
		t.Errorf("TotalSessions after reset: got %d, want 0", n)
	}
	if answers, _ := games.AnswersBySession(s.ID); len(answers) != 0 {
		t.Errorf("answers after reset: got %d, want 0", len(answers))
	}
	if used, _ := games.UsedQuestionIDs(s.ID); len(used) != 0 {
		t.Errorf("history after reset: got %d, want 0", len(used))
	}
	if recs, _ := stats.RecommendationsByUser(u.ID, 0); len(recs) != 0 {
		t.Errorf("recommendations after reset: got %d, want 0", len(recs))
	}

	// Account and token survive.
	if got, err := st.Users().ByUsername("ana"); err != nil || got.ID != u.ID {
		t.Errorf("account after reset: (%+v, %v)", got, err)
	}
	if got, err := st.Users().ByToken("tok-ana"); err != nil || got.ID != u.ID {
		t.Errorf("token after reset: (%+v, %v)", got, err)
	}

	// Other user's data untouched.
	if _, err := games.SessionByID(os.ID); err != nil {
		t.Errorf("other user's session after reset: %v", err)
	}
	if totals, _ := stats.OverallTotals(other.ID); totals.Total != 1 {
		t.Errorf("other user's totals after reset: %+v", totals)
	}
}
