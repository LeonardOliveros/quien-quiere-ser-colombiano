package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math"
	mrand "math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"quiz-app/internal/domain"
)

// This file is the HTTP (driving) adapter: it translates HTTP requests into
// calls against the domain.Store port. It must not know which storage
// technology backs the store.

// User authentication handlers
func registerUser(c *gin.Context) {
	var registerData struct {
		Username string `json:"username" binding:"required,min=3,max=50"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=8,max=72"`
	}
	if err := c.ShouldBindJSON(&registerData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerData.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	user := domain.User{
		Username: registerData.Username,
		Email:    registerData.Email,
		Password: string(hashedPassword),
	}

	if err := store.Users().Create(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username or email already exists"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "user_id": user.ID})
}

func loginUser(c *gin.Context) {
	var loginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := store.Users().ByUsername(loginData.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate session token with expiry
	token := generateToken()
	expiresAt := time.Now().Add(tokenTTL)

	if err := store.Users().SaveSessionToken(user.ID, token, expiresAt); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create session"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"user_id": user.ID,
		"token":   token,
	})
}

// Middleware for authentication
func authRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			token = c.Query("token")
		}

		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing authorization token"})
			c.Abort()
			return
		}

		user, err := store.Users().ByToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		if user.TokenExpiresAt == nil || time.Now().After(*user.TokenExpiresAt) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token expired"})
			c.Abort()
			return
		}

		// Store user ID in context
		c.Set("userID", user.ID)
		c.Next()
	}
}

// requireOwnedSession loads the session from the :sessionId param and verifies
// it belongs to the authenticated user.
func requireOwnedSession(c *gin.Context) (domain.GameSession, bool) {
	sessionID, err := strconv.Atoi(c.Param("sessionId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session ID"})
		return domain.GameSession{}, false
	}

	session, err := store.Games().SessionByID(uint(sessionID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Session not found"})
		return domain.GameSession{}, false
	}

	userID, exists := c.Get("userID")
	if !exists || session.UserID != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Session does not belong to the authenticated user"})
		return domain.GameSession{}, false
	}

	return session, true
}

// requireOwnUserID verifies the :userId param matches the authenticated user.
func requireOwnUserID(c *gin.Context) (uint, bool) {
	userID, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return 0, false
	}

	authUserID, exists := c.Get("userID")
	if !exists || authUserID.(uint) != uint(userID) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authorized to access this user's data"})
		return 0, false
	}

	return uint(userID), true
}

// Game handlers
func startGame(c *gin.Context) {
	var config domain.GameConfig
	if err := c.ShouldBindJSON(&config); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get user ID from context (set by authRequired middleware)
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// Complete any existing active or paused sessions for this user and mode
	store.Games().CompleteResumable(userID.(uint), config.Mode, 0)

	// Convert categories slice to comma-separated string
	categoriesStr := ""

	// For TIMED mode, force the 4 main categories, 80 questions (20 per category), and 1 hour time limit
	if config.Mode == "TIMED" {
		categoriesStr = "CONSTITUCION,HISTORIA,GEOGRAFIA,CULTURA"
		config.QuestionCount = 80
		config.TimeLimit = 3600 // 1 hour in seconds (60 minutes * 60 seconds)
		config.Categories = []string{"CONSTITUCION", "HISTORIA", "GEOGRAFIA", "CULTURA"}
	} else if len(config.Categories) > 0 {
		categoriesStr = strings.Join(config.Categories, ",")
	}

	// Verify that questions are available for the selected criteria
	availableCount, err := store.Questions().Count(splitCategories(categoriesStr))
	if err != nil || availableCount == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No questions available for the selected criteria"})
		return
	}

	session := domain.GameSession{
		UserID:         userID.(uint),
		Mode:           config.Mode,
		Categories:     categoriesStr,
		Status:         "ACTIVE",
		StartTime:      time.Now(),
		TimeLimit:      config.TimeLimit,
		TotalQuestions: config.QuestionCount,
	}

	if err := store.Games().CreateSession(&session); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create game session"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"session_id": session.ID,
		"config":     config,
		"message":    "Game started successfully",
	})
}

// splitCategories converts the session's comma-separated category codes into
// a slice ("" -> nil, meaning no filter).
func splitCategories(categoriesStr string) []string {
	if categoriesStr == "" {
		return nil
	}
	return strings.Split(categoriesStr, ",")
}

func getNextQuestion(c *gin.Context) {
	session, ok := requireOwnedSession(c)
	if !ok {
		return
	}
	games := store.Games()

	// If session is PAUSED, reactivate it and adjust start time
	if session.Status == "PAUSED" {
		// Adjust the start time to account for paused time
		// New start time = now - time_elapsed
		session.Status = "ACTIVE"
		session.StartTime = time.Now().Add(-time.Duration(session.TimeElapsed) * time.Second)
		session.PausedAt = nil
		games.SaveSession(&session)
	}

	// Get IDs of questions already presented in this session (from history)
	usedQuestionIDs, _ := games.UsedQuestionIDs(session.ID)

	// Get answered question IDs (flag placeholders don't count as answers)
	answeredIDs, _ := games.AnsweredQuestionIDs(session.ID)

	// Check if we've reached the maximum number of questions
	if len(answeredIDs) >= session.TotalQuestions {
		c.JSON(http.StatusNotFound, gin.H{"error": "No more questions available"})
		return
	}

	// Restrict the random pick to the session's categories
	var pickFrom []string

	// For TIMED mode, ensure 20 questions per category
	if session.Mode == "TIMED" {
		// Count questions answered per category
		categoryCountMap, _ := games.AnsweredCountByCategory(session.ID)

		// Find categories that haven't reached 20 questions yet
		availableCategories := []string{}
		mainCategories := []string{"CONSTITUCION", "HISTORIA", "GEOGRAFIA", "CULTURA"}
		for _, cat := range mainCategories {
			if categoryCountMap[cat] < 20 {
				availableCategories = append(availableCategories, cat)
			}
		}

		if len(availableCategories) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "No more questions available"})
			return
		}

		// Randomly select one of the available categories
		pickFrom = []string{availableCategories[mrand.Intn(len(availableCategories))]}
	} else {
		// Apply category filter if exists for other modes
		pickFrom = splitCategories(session.Categories)
	}

	// Get a random question not used yet in this session
	questionID, err := store.Questions().RandomID(pickFrom, usedQuestionIDs)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No more questions available"})
		return
	}

	// Load the selected question with choices
	question, err := store.Questions().ByID(questionID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Question not found"})
		return
	}

	// Register this question in the history to prevent it from being used again in active sessions
	games.AddHistory(session.ID, questionID)

	// Randomize the order of choices to prevent visual memorization
	mrand.Shuffle(len(question.Choices), func(i, j int) {
		question.Choices[i], question.Choices[j] = question.Choices[j], question.Choices[i]
	})

	// Hide correct answer from choices
	for i := range question.Choices {
		question.Choices[i].IsCorrect = false
	}

	// Calculate elapsed time and remaining time
	timeElapsed := calculateCurrentTimeElapsed(session)
	timeRemaining := getTimeRemaining(session)

	c.JSON(http.StatusOK, gin.H{
		"question":        question,
		"question_number": len(answeredIDs) + 1,
		"total_questions": session.TotalQuestions,
		"time_remaining":  timeRemaining,
		"time_elapsed":    timeElapsed,
	})
}

func submitAnswer(c *gin.Context) {
	session, ok := requireOwnedSession(c)
	if !ok {
		return
	}
	games := store.Games()

	var answerData struct {
		QuestionID uint `json:"question_id" binding:"required"`
		ChoiceID   uint `json:"choice_id" binding:"required"`
		TimeSpent  int  `json:"time_spent"`
	}

	if err := c.ShouldBindJSON(&answerData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if session.Status == "COMPLETED" {
		c.JSON(http.StatusConflict, gin.H{"error": "Game session is already completed"})
		return
	}

	// Enforce the time limit server-side
	if session.TimeLimit > 0 && getTimeRemaining(session) <= 0 {
		now := time.Now()
		session.Status = "COMPLETED"
		session.EndTime = &now
		games.SaveSession(&session)
		c.JSON(http.StatusConflict, gin.H{"error": "Time limit exceeded"})
		return
	}

	// Reject answering the same question twice in a session
	alreadyAnswered, _ := games.HasAnswered(session.ID, answerData.QuestionID)
	if alreadyAnswered {
		c.JSON(http.StatusConflict, gin.H{"error": "Question already answered"})
		return
	}

	// Check if answer is correct, validating the choice belongs to the question
	choice, err := store.Questions().ChoiceByID(answerData.ChoiceID)
	if err != nil || choice.QuestionID != answerData.QuestionID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Choice does not belong to the submitted question"})
		return
	}

	gameAnswer := domain.GameAnswer{
		GameSessionID: session.ID,
		QuestionID:    answerData.QuestionID,
		ChoiceID:      &answerData.ChoiceID,
		IsCorrect:     choice.IsCorrect,
		TimeSpent:     answerData.TimeSpent,
		AnsweredAt:    time.Now(),
	}

	// Fill in the flag placeholder row if the question was flagged first
	if placeholder, err := games.AnswerPlaceholder(session.ID, answerData.QuestionID); err == nil {
		gameAnswer.ID = placeholder.ID
		gameAnswer.IsFlagged = placeholder.IsFlagged
	}

	if err := games.SaveAnswer(&gameAnswer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save answer"})
		return
	}

	// Update session score if correct
	if choice.IsCorrect {
		games.AddScore(session.ID, 10)
	}

	// Get the correct choice to send back in the response
	correctChoices, _ := store.Questions().CorrectChoices([]uint{answerData.QuestionID})

	c.JSON(http.StatusOK, gin.H{
		"correct":           choice.IsCorrect,
		"choice_id":         choice.ID,
		"correct_choice_id": correctChoices[answerData.QuestionID].ID,
		"explanation":       getQuestionExplanation(answerData.QuestionID),
	})
}

func flagQuestion(c *gin.Context) {
	session, ok := requireOwnedSession(c)
	if !ok {
		return
	}

	var flagData struct {
		QuestionID uint `json:"question_id"`
	}

	if err := c.ShouldBindJSON(&flagData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Toggle the flag on the existing answer row, if any
	toggled, err := store.Games().ToggleFlag(session.ID, flagData.QuestionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// The question may not have been answered yet: create a flagged
	// placeholder row that submitAnswer fills in later
	if !toggled {
		placeholder := domain.GameAnswer{
			GameSessionID: session.ID,
			QuestionID:    flagData.QuestionID,
			IsFlagged:     true,
		}
		if err := store.Games().SaveAnswer(&placeholder); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to flag question"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Flag status toggled",
	})
}

func useFiftyFifty(c *gin.Context) {
	if _, ok := requireOwnedSession(c); !ok {
		return
	}

	var request struct {
		QuestionID uint `json:"question_id"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Load question with choices
	question, err := store.Questions().ByID(request.QuestionID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Question not found"})
		return
	}

	// Find incorrect choices
	var incorrectChoiceIDs []uint
	for _, choice := range question.Choices {
		if !choice.IsCorrect {
			incorrectChoiceIDs = append(incorrectChoiceIDs, choice.ID)
		}
	}

	// Randomly select 2 incorrect choices to remove
	if len(incorrectChoiceIDs) < 2 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not enough incorrect choices"})
		return
	}

	// Shuffle and take 2
	mrand.Shuffle(len(incorrectChoiceIDs), func(i, j int) {
		incorrectChoiceIDs[i], incorrectChoiceIDs[j] = incorrectChoiceIDs[j], incorrectChoiceIDs[i]
	})
	toRemove := incorrectChoiceIDs[:2]

	c.JSON(http.StatusOK, gin.H{
		"remove_choice_ids": toRemove,
	})
}

func useAutosolve(c *gin.Context) {
	if _, ok := requireOwnedSession(c); !ok {
		return
	}

	var request struct {
		QuestionID uint `json:"question_id"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Load question with choices
	question, err := store.Questions().ByID(request.QuestionID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Question not found"})
		return
	}

	// Find the correct choice
	var correctChoiceID uint
	for _, choice := range question.Choices {
		if choice.IsCorrect {
			correctChoiceID = choice.ID
			break
		}
	}

	if correctChoiceID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No correct answer found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"correct_choice_id": correctChoiceID,
	})
}

func endGame(c *gin.Context) {
	session, ok := requireOwnedSession(c)
	if !ok {
		return
	}

	endTime := time.Now()
	session.Status = "COMPLETED"
	session.EndTime = &endTime
	store.Games().SaveSession(&session)

	// Generate study recommendations
	generateRecommendations(session.ID)

	c.JSON(http.StatusOK, gin.H{"message": "Game ended successfully"})
}

func pauseGame(c *gin.Context) {
	session, ok := requireOwnedSession(c)
	if !ok {
		return
	}

	// Calculate and save the elapsed time before pausing
	// (start_time is adjusted on resume, so this spans only active play)
	currentElapsed := session.TimeElapsed
	if session.Status == "ACTIVE" {
		currentElapsed = int(time.Since(session.StartTime).Seconds())
	}

	now := time.Now()
	session.Status = "PAUSED"
	session.TimeElapsed = currentElapsed
	session.PausedAt = &now

	if err := store.Games().SaveSession(&session); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to pause game"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Game paused successfully"})
}

// respondPausedGame finds the most recent active/paused session for the
// authenticated user (optionally filtered by mode), completes any older
// leftover sessions and responds with the session's progress summary.
func respondPausedGame(c *gin.Context, mode string) {
	// Get user ID from context
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}
	games := store.Games()

	// Find the most recent active or paused game
	session, err := games.LatestResumable(userID.(uint), mode)
	if err != nil {
		// No active or paused game found
		c.JSON(http.StatusNotFound, gin.H{"error": "No paused game found"})
		return
	}

	// Clean up old active/paused games (keep only the most recent one)
	games.CompleteResumable(userID.(uint), mode, session.ID)

	// Get progress information (flag placeholders don't count as answers)
	answeredCount, _ := games.AnsweredCount(session.ID)

	// Count incorrect answers
	incorrectCount, _ := games.IncorrectCount(session.ID)

	// Count flagged questions
	flaggedCount, _ := games.FlaggedCount(session.ID)

	// Get flagged question IDs
	flaggedIDs, _ := games.FlaggedQuestionIDs(session.ID)

	// Calculate elapsed time
	timeElapsed := calculateCurrentTimeElapsed(session)
	timeRemaining := getTimeRemaining(session)

	c.JSON(http.StatusOK, gin.H{
		"session_id":         session.ID,
		"mode":               session.Mode,
		"categories":         session.Categories,
		"total_questions":    session.TotalQuestions,
		"answered_questions": answeredCount,
		"correct_answers":    session.CorrectAnswers,
		"incorrect_answers":  incorrectCount,
		"flagged_count":      flaggedCount,
		"flagged_questions":  flaggedIDs,
		"score":              session.Score,
		"time_limit":         session.TimeLimit,
		"time_elapsed":       timeElapsed,
		"time_remaining":     timeRemaining,
		"start_time":         session.StartTime,
		"status":             session.Status,
	})
}

func getAnyPausedGame(c *gin.Context) {
	respondPausedGame(c, "")
}

func getPausedGame(c *gin.Context) {
	respondPausedGame(c, c.Param("mode"))
}

func getGameResults(c *gin.Context) {
	session, ok := requireOwnedSession(c)
	if !ok {
		return
	}

	// Get all answers
	answers, err := store.Games().AnswersBySession(session.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load answers"})
		return
	}

	// Calculate category scores
	categoryScores := make(map[string]domain.CategoryScore)
	incorrectAnswers := []domain.IncorrectAnswer{}
	flaggedQuestions := []domain.Question{}

	// Prefetch the correct choice of every question in one query
	questionIDs := make([]uint, 0, len(answers))
	for _, answer := range answers {
		questionIDs = append(questionIDs, answer.QuestionID)
	}
	correctByQuestion, _ := store.Questions().CorrectChoices(questionIDs)

	for _, answer := range answers {
		// Flag placeholders were never answered: track the flag, skip scoring
		if answer.ChoiceID == nil {
			if answer.IsFlagged {
				flaggedQuestions = append(flaggedQuestions, answer.Question)
			}
			continue
		}

		category := answer.Question.Category.Code
		if _, exists := categoryScores[category]; !exists {
			categoryScores[category] = domain.CategoryScore{Category: category}
		}

		score := categoryScores[category]
		score.TotalQuestions++
		if answer.IsCorrect {
			score.CorrectAnswers++
		} else {
			// Add to incorrect answers
			correctChoice := correctByQuestion[answer.QuestionID]

			incorrectAnswers = append(incorrectAnswers, domain.IncorrectAnswer{
				Question:      answer.Question,
				UserChoice:    answer.Choice,
				CorrectChoice: correctChoice,
				Explanation:   answer.Question.Explanation,
			})
		}

		if answer.IsFlagged {
			flaggedQuestions = append(flaggedQuestions, answer.Question)
		}

		score.Percentage = float64(score.CorrectAnswers) / float64(score.TotalQuestions) * 100
		score.Passed = checkIfPassed(category, score.Percentage)
		categoryScores[category] = score
	}

	// Get recommendations
	var recommendations []string
	recs, _ := store.Stats().RecommendationsByUser(session.UserID, 5)
	for _, rec := range recs {
		recommendations = append(recommendations, rec.Description)
	}

	// Count only real answers (not flag placeholders)
	answeredTotal := 0
	for _, answer := range answers {
		if answer.ChoiceID != nil {
			answeredTotal++
		}
	}

	// Calculate percentage safely (avoid division by zero)
	percentage := 0.0
	if answeredTotal > 0 {
		percentage = float64(session.CorrectAnswers) / float64(answeredTotal) * 100
	}

	// EndTime is nil until the game is ended; fall back to elapsed time
	timeTaken := calculateCurrentTimeElapsed(session)
	if session.EndTime != nil {
		timeTaken = int(session.EndTime.Sub(session.StartTime).Seconds())
	}

	result := domain.GameResult{
		SessionID:        session.ID,
		TotalQuestions:   answeredTotal,
		CorrectAnswers:   session.CorrectAnswers,
		Score:            session.Score,
		Percentage:       percentage,
		TimeTaken:        timeTaken,
		CategoryScores:   categoryScores,
		IncorrectAnswers: incorrectAnswers,
		FlaggedQuestions: flaggedQuestions,
		Recommendations:  recommendations,
	}

	c.JSON(http.StatusOK, result)
}

// Question handlers

// hideCorrectChoices strips the answer key before questions leave the API.
func hideCorrectChoices(questions []domain.Question) {
	for i := range questions {
		for j := range questions[i].Choices {
			questions[i].Choices[j].IsCorrect = false
		}
	}
}

func getQuestions(c *gin.Context) {
	questions, err := store.Questions().List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load questions"})
		return
	}
	hideCorrectChoices(questions)
	c.JSON(http.StatusOK, questions)
}

func getQuestion(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	question, err := store.Questions().ByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Question not found"})
		return
	}
	hideCorrectChoices([]domain.Question{question})
	c.JSON(http.StatusOK, question)
}

func getQuestionsByCategory(c *gin.Context) {
	questions, err := store.Questions().ListByCategory(c.Param("category"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load questions"})
		return
	}
	hideCorrectChoices(questions)
	c.JSON(http.StatusOK, questions)
}

// getCategories returns the canonical taxonomy (categories with their subcategories).
func getCategories(c *gin.Context) {
	categories, err := store.Questions().Categories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load categories"})
		return
	}
	c.JSON(http.StatusOK, categories)
}

func getQuestionsCount(c *gin.Context) {
	// Get total count
	totalCount, _ := store.Questions().Count(nil)

	// Get count by category, ensuring the 4 main categories always appear
	categoryCount := map[string]int64{"CULTURA": 0, "GEOGRAFIA": 0, "HISTORIA": 0, "CONSTITUCION": 0}
	byCategory, _ := store.Questions().CountsByCategory()
	for code, count := range byCategory {
		categoryCount[code] = count
	}

	// Get count by subcategory (optional detailed breakdown)
	subcategoryCounts, _ := store.Questions().CountsBySubcategory()

	c.JSON(http.StatusOK, gin.H{
		"total":          totalCount,
		"by_category":    categoryCount,
		"by_subcategory": subcategoryCounts,
	})
}

// User statistics handlers
func getUserStats(c *gin.Context) {
	userID, ok := requireOwnUserID(c)
	if !ok {
		return
	}

	stats := calculateUserStats(userID)
	c.JSON(http.StatusOK, stats)
}

func getWeakAreas(c *gin.Context) {
	userID, ok := requireOwnUserID(c)
	if !ok {
		return
	}

	weakAreas := identifyWeakAreas(userID)
	c.JSON(http.StatusOK, gin.H{"weak_areas": weakAreas})
}

func getGameHistory(c *gin.Context) {
	userID, ok := requireOwnUserID(c)
	if !ok {
		return
	}

	sessions, err := store.Games().SessionsByUser(userID, 20)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load history"})
		return
	}
	c.JSON(http.StatusOK, sessions)
}

func getStudyRecommendations(c *gin.Context) {
	userID, ok := requireOwnUserID(c)
	if !ok {
		return
	}

	recommendations, err := store.Stats().RecommendationsByUser(userID, 0)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load recommendations"})
		return
	}
	c.JSON(http.StatusOK, recommendations)
}

func resetUserStats(c *gin.Context) {
	userID, ok := requireOwnUserID(c)
	if !ok {
		return
	}

	if err := store.ResetUserData(userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to reset statistics"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Statistics reset successfully"})
}

// tokenTTL is how long a login session token stays valid.
const tokenTTL = 7 * 24 * time.Hour

// Helper functions
func generateToken() string {
	b := make([]byte, 16)
	rand.Read(b)
	return hex.EncodeToString(b)
}

// calculateCurrentTimeElapsed returns the total time elapsed excluding paused time
func calculateCurrentTimeElapsed(session domain.GameSession) int {
	if session.Status == "PAUSED" {
		// Return the saved elapsed time when paused
		return session.TimeElapsed
	}

	// For ACTIVE sessions, calculate time since start
	return int(time.Since(session.StartTime).Seconds())
}

func getTimeRemaining(session domain.GameSession) int {
	if session.TimeLimit == 0 {
		return -1 // Unlimited time
	}

	elapsed := calculateCurrentTimeElapsed(session)
	remaining := session.TimeLimit - elapsed

	if remaining < 0 {
		return 0
	}
	return remaining
}

func getQuestionExplanation(questionID uint) string {
	question, err := store.Questions().ByID(questionID)
	if err != nil {
		return ""
	}
	return question.Explanation
}

func checkIfPassed(category string, percentage float64) bool {
	passingScores := map[string]float64{
		"CONSTITUCION": 60.0,
		"HISTORIA":     40.0,
		"GEOGRAFIA":    55.0,
		"CULTURA":      40.0,
	}

	if required, exists := passingScores[category]; exists {
		return percentage >= required
	}
	return percentage >= 50.0
}

func getUserWeakCategories(userID uint) []string {
	weakCategories, _ := store.Stats().WeakCategories(userID)
	return weakCategories
}

func calculateUserStats(userID uint) domain.UserStats {
	var stats domain.UserStats
	stats.UserID = userID
	statsRepo := store.Stats()

	// Get total games
	stats.TotalGames, _ = statsRepo.TotalSessions(userID)

	// Get overall performance from actually answered questions
	// (session.TotalQuestions is the planned count, which overstates
	// abandoned or partially played games)
	overall, _ := statsRepo.OverallTotals(userID)
	stats.TotalQuestions = overall.Total
	stats.CorrectAnswers = overall.Correct

	stats.BestScore, _ = statsRepo.BestScore(userID)

	if stats.TotalQuestions > 0 {
		stats.AverageScore = float64(stats.CorrectAnswers) / float64(stats.TotalQuestions) * 100
	}

	// Calculate category stats
	stats.CategoryStats = make(map[string]domain.CategoryStats)
	categories := []string{"CULTURA", "GEOGRAFIA", "HISTORIA", "CONSTITUCION"}

	for _, category := range categories {
		var catStats domain.CategoryStats
		catStats.Category = category

		// Get performance for this category
		totals, _ := statsRepo.CategoryTotals(userID, category)

		catStats.TotalQuestions = totals.Total
		catStats.CorrectAnswers = totals.Correct
		if catStats.TotalQuestions > 0 {
			catStats.AveragePercentage = float64(catStats.CorrectAnswers) / float64(catStats.TotalQuestions) * 100
		}

		stats.CategoryStats[category] = catStats

		// Identify weak and strong areas
		if catStats.AveragePercentage < 50 {
			stats.WeakAreas = append(stats.WeakAreas, category)
		} else if catStats.AveragePercentage > 75 {
			stats.StrongAreas = append(stats.StrongAreas, category)
		}
	}

	// Get recent progress, computing percentages from answered questions
	progress, _ := statsRepo.RecentProgress(userID, 10)
	for _, p := range progress {
		percentage := 0.0
		if p.Answered > 0 {
			percentage = float64(p.Correct) / float64(p.Answered) * 100
		}
		stats.RecentProgress = append(stats.RecentProgress, domain.ProgressPoint{
			Date:       p.Date,
			Score:      p.Score,
			Percentage: percentage,
		})
	}

	return stats
}

func identifyWeakAreas(userID uint) []string {
	stats := calculateUserStats(userID)
	return stats.WeakAreas
}

func generateRecommendations(sessionID uint) {
	session, err := store.Games().SessionByID(sessionID)
	if err != nil {
		return
	}

	// Analyze incorrect answers
	incorrectAnswers, err := store.Games().IncorrectAnswers(sessionID)
	if err != nil {
		return
	}

	// Group by category and subcategory
	weaknesses := make(map[string]int)
	for _, answer := range incorrectAnswers {
		key := fmt.Sprintf("%s_%s", answer.Question.Category.Code, answer.Question.SubCategory.Name)
		weaknesses[key]++
	}

	// Create recommendations
	for area, count := range weaknesses {
		if count >= 2 { // If failed 2+ questions in this area
			rec := domain.StudyRecommendation{
				UserID:      session.UserID,
				Category:    getCategory(area),
				SubCategory: getSubCategory(area),
				Weakness:    area,
				Description: generateStudyDescription(area, count),
				Resources:   generateResources(area),
				Priority:    int(math.Min(5, float64(count))),
			}
			store.Stats().CreateRecommendation(&rec)
		}
	}
}

func getCategory(area string) string {
	// Extract category from area string (e.g., "HISTORIA_Independencia" -> "HISTORIA")
	if len(area) > 0 {
		for i, c := range area {
			if c == '_' {
				return area[:i]
			}
		}
	}
	return area
}

func getSubCategory(area string) string {
	// Extract subcategory from area string
	for i, c := range area {
		if c == '_' {
			return area[i+1:]
		}
	}
	return ""
}

func generateStudyDescription(area string, count int) string {
	return fmt.Sprintf("Necesitas repasar %s. Fallaste %d preguntas en esta área.", area, count)
}

func generateResources(area string) string {
	// Return JSON string with study materials
	return `{"videos": [], "documents": ["COLOMBIA: NUESTRA CASA"], "exercises": []}`
}
