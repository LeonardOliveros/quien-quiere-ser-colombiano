package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math"
	mrand "math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"golang.org/x/crypto/bcrypt"
)

// User authentication handlers
func registerUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	plainPassword := user.Password

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	user.Password = string(hashedPassword)

	if err := db.Create(&user).Error; err != nil {
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

	var user User
	if err := db.Where("username = ?", loginData.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate session token
	token := generateToken()

	// Save token to database
	if err := db.Model(&user).Update("token", token).Error; err != nil {
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

		var user User
		if err := db.Where("token = ?", token).First(&user).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Store user ID in context
		c.Set("userID", user.ID)
		c.Next()
	}
}

// Game handlers
func startGame(c *gin.Context) {
	var config GameConfig
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

	// Convert categories slice to comma-separated string
	categoriesStr := ""
	if len(config.Categories) > 0 {
		categoriesStr = config.Categories[0]
		for i := 1; i < len(config.Categories); i++ {
			categoriesStr += "," + config.Categories[i]
		}
	}

	// Pre-generate the question sequence
	questionSequence := generateQuestionSequence(userID.(uint), config.Mode, categoriesStr, config.QuestionCount)
	if questionSequence == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No questions available for the selected criteria"})
		return
	}

	session := GameSession{
		UserID:        userID.(uint),
		Mode:          config.Mode,
		Categories:    categoriesStr,
		QuestionSequence: questionSequence,
		Status:        "ACTIVE",
		StartTime:     time.Now(),
		TimeLimit:     config.TimeLimit,
		TotalQuestions: config.QuestionCount,
	}

	if err := db.Create(&session).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create game session"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"session_id": session.ID,
		"config":     config,
		"message":    "Game started successfully",
	})
}

func getNextQuestion(c *gin.Context) {
	sessionID, _ := strconv.Atoi(c.Param("sessionId"))

	var session GameSession
	if err := db.First(&session, sessionID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Session not found"})
		return
	}

	// Parse the pre-generated question sequence
	questionIDs := parseQuestionSequence(session.QuestionSequence)
	if len(questionIDs) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No questions available"})
		return
	}

	// Get answered question IDs
	var answeredIDs []uint
	db.Model(&GameAnswer{}).
		Where("game_session_id = ?", sessionID).
		Pluck("question_id", &answeredIDs)

	// Check if we've reached the maximum number of questions
	if len(answeredIDs) >= session.TotalQuestions || len(answeredIDs) >= len(questionIDs) {
		c.JSON(http.StatusNotFound, gin.H{"error": "No more questions available"})
		return
	}

	// Find the next unanswered question in the sequence
	var questionID uint
	answeredMap := make(map[uint]bool)
	for _, id := range answeredIDs {
		answeredMap[id] = true
	}

	for _, qid := range questionIDs {
		if !answeredMap[qid] {
			questionID = qid
			break
		}
	}

	if questionID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No more questions available"})
		return
	}

	// Load the selected question with choices
	var question Question
	if err := db.Preload("Choices").First(&question, questionID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Question not found"})
		return
	}

	// Randomize the order of choices to prevent visual memorization
	mrand.Seed(time.Now().UnixNano())
	mrand.Shuffle(len(question.Choices), func(i, j int) {
		question.Choices[i], question.Choices[j] = question.Choices[j], question.Choices[i]
	})

	// Hide correct answer from choices
	for i := range question.Choices {
		question.Choices[i].IsCorrect = false
	}

	c.JSON(http.StatusOK, gin.H{
		"question":         question,
		"question_number":  len(answeredIDs) + 1,
		"total_questions":  session.TotalQuestions,
		"time_remaining":   getTimeRemaining(session),
	})
}

func submitAnswer(c *gin.Context) {
	sessionID, _ := strconv.Atoi(c.Param("sessionId"))
	
	var answerData struct {
		QuestionID uint `json:"question_id"`
		ChoiceID   uint `json:"choice_id"`
		TimeSpent  int  `json:"time_spent"`
	}

	if err := c.ShouldBindJSON(&answerData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if answer is correct
	var choice Choice
	db.First(&choice, answerData.ChoiceID)

	gameAnswer := GameAnswer{
		GameSessionID: uint(sessionID),
		QuestionID:    answerData.QuestionID,
		ChoiceID:      &answerData.ChoiceID,
		IsCorrect:     choice.IsCorrect,
		TimeSpent:     answerData.TimeSpent,
		AnsweredAt:    time.Now(),
	}

	if err := db.Create(&gameAnswer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save answer"})
		return
	}

	// Update session score if correct
	if choice.IsCorrect {
		db.Model(&GameSession{}).Where("id = ?", sessionID).
			UpdateColumn("correct_answers", db.Raw("correct_answers + ?", 1)).
			UpdateColumn("score", db.Raw("score + ?", 10))
	}

	// Get the correct choice to send back in the response
	var correctChoice Choice
	db.Where("question_id = ? AND is_correct = true", answerData.QuestionID).First(&correctChoice)

	c.JSON(http.StatusOK, gin.H{
		"correct":           choice.IsCorrect,
		"choice_id":         choice.ID,
		"correct_choice_id": correctChoice.ID,
		"explanation":       getQuestionExplanation(answerData.QuestionID),
	})
}

func flagQuestion(c *gin.Context) {
	sessionID, _ := strconv.Atoi(c.Param("sessionId"))

	var flagData struct {
		QuestionID uint `json:"question_id"`
	}

	if err := c.ShouldBindJSON(&flagData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Ejecuta el toggle y devuelve cuántas filas fueron afectadas
	result := db.Model(&GameAnswer{}).
		Where("game_session_id = ? AND question_id = ?", sessionID, flagData.QuestionID).
		Update("is_flagged", gorm.Expr("NOT is_flagged"))

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Flag status toggled",
	})
}

func endGame(c *gin.Context) {
	sessionID, _ := strconv.Atoi(c.Param("sessionId"))
	
	endTime := time.Now()
	db.Model(&GameSession{}).Where("id = ?", sessionID).Updates(map[string]interface{}{
		"end_time": endTime,
		"status":   "COMPLETED",
	})

	// Generate study recommendations
	generateRecommendations(uint(sessionID))

	c.JSON(http.StatusOK, gin.H{"message": "Game ended successfully"})
}

func getGameResults(c *gin.Context) {
	sessionID, _ := strconv.Atoi(c.Param("sessionId"))
	
	var session GameSession
	if err := db.Preload("User").First(&session, sessionID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Session not found"})
		return
	}

	// Get all answers
	var answers []GameAnswer
	db.Preload("Question").Preload("Choice").
		Where("game_session_id = ?", sessionID).Find(&answers)

	// Calculate category scores
	categoryScores := make(map[string]CategoryScore)
	incorrectAnswers := []IncorrectAnswer{}
	flaggedQuestions := []Question{}

	for _, answer := range answers {
		category := answer.Question.Category
		if _, exists := categoryScores[category]; !exists {
			categoryScores[category] = CategoryScore{Category: category}
		}

		score := categoryScores[category]
		score.TotalQuestions++
		if answer.IsCorrect {
			score.CorrectAnswers++
		} else {
			// Add to incorrect answers
			var correctChoice Choice
			db.Where("question_id = ? AND is_correct = true", answer.QuestionID).First(&correctChoice)
			
			incorrectAnswers = append(incorrectAnswers, IncorrectAnswer{
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
	var recs []StudyRecommendation
	db.Where("user_id = ?", session.UserID).Order("priority desc").Limit(5).Find(&recs)
	for _, rec := range recs {
		recommendations = append(recommendations, rec.Description)
	}

	result := GameResult{
		SessionID:        session.ID,
		TotalQuestions:   len(answers),
		CorrectAnswers:   session.CorrectAnswers,
		Score:            session.Score,
		Percentage:       float64(session.CorrectAnswers) / float64(len(answers)) * 100,
		TimeTaken:        int(session.EndTime.Sub(session.StartTime).Seconds()),
		CategoryScores:   categoryScores,
		IncorrectAnswers: incorrectAnswers,
		FlaggedQuestions: flaggedQuestions,
		Recommendations:  recommendations,
	}

	c.JSON(http.StatusOK, result)
}

// Question handlers
func getQuestions(c *gin.Context) {
	var questions []Question
	db.Preload("Choices").Find(&questions)
	c.JSON(http.StatusOK, questions)
}

func getQuestion(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var question Question
	if err := db.Preload("Choices").First(&question, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Question not found"})
		return
	}
	c.JSON(http.StatusOK, question)
}

func getQuestionsByCategory(c *gin.Context) {
	category := c.Param("category")
	var questions []Question
	db.Preload("Choices").Where("category = ?", category).Find(&questions)
	c.JSON(http.StatusOK, questions)
}

func getQuestionsCount(c *gin.Context) {
	// Get total count
	var totalCount int64
	db.Model(&Question{}).Count(&totalCount)

	// Get count by category
	categories := []string{"CULTURA", "GEOGRAFIA", "HISTORIA", "CONSTITUCION"}
	categoryCount := make(map[string]int64)

	for _, category := range categories {
		var count int64
		db.Model(&Question{}).Where("category = ?", category).Count(&count)
		categoryCount[category] = count
	}

	// Get count by subcategory (optional detailed breakdown)
	var subcategoryCounts []struct {
		Category    string
		SubCategory string
		Count       int64
	}

	db.Model(&Question{}).
		Select("category, sub_category, COUNT(*) as count").
		Group("category, sub_category").
		Scan(&subcategoryCounts)

	c.JSON(http.StatusOK, gin.H{
		"total":             totalCount,
		"by_category":       categoryCount,
		"by_subcategory":    subcategoryCounts,
	})
}

// User statistics handlers
func getUserStats(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Param("userId"))
	
	stats := calculateUserStats(uint(userID))
	c.JSON(http.StatusOK, stats)
}

func getWeakAreas(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Param("userId"))
	
	weakAreas := identifyWeakAreas(uint(userID))
	c.JSON(http.StatusOK, gin.H{"weak_areas": weakAreas})
}

func getGameHistory(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Param("userId"))
	
	var sessions []GameSession
	db.Where("user_id = ?", userID).Order("created_at desc").Limit(20).Find(&sessions)
	c.JSON(http.StatusOK, sessions)
}

func getStudyRecommendations(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Param("userId"))

	var recommendations []StudyRecommendation
	db.Where("user_id = ?", userID).Order("priority desc").Find(&recommendations)
	c.JSON(http.StatusOK, recommendations)
}

func resetUserStats(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Param("userId"))

	// Get user ID from auth context to verify authorization
	authUserID, exists := c.Get("userID")
	if !exists || authUserID.(uint) != uint(userID) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized to reset these statistics"})
		return
	}

	// Delete all game answers for user's sessions
	db.Exec("DELETE FROM game_answers WHERE game_session_id IN (SELECT id FROM game_sessions WHERE user_id = ?)", userID)

	// Delete all game sessions for the user
	db.Where("user_id = ?", userID).Delete(&GameSession{})

	// Delete all study recommendations for the user
	db.Where("user_id = ?", userID).Delete(&StudyRecommendation{})

	c.JSON(http.StatusOK, gin.H{"message": "Statistics reset successfully"})
}

// Helper functions
func generateToken() string {
	b := make([]byte, 16)
	rand.Read(b)
	return hex.EncodeToString(b)
}

func getTimeRemaining(session GameSession) int {
	if session.TimeLimit == 0 {
		return -1 // Unlimited time
	}
	elapsed := time.Since(session.StartTime).Seconds()
	remaining := float64(session.TimeLimit) - elapsed
	if remaining < 0 {
		return 0
	}
	return int(remaining)
}

func getQuestionExplanation(questionID uint) string {
	var question Question
	db.First(&question, questionID)
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
	// Analyze recent performance to identify weak categories
	var weakCategories []string
	
	query := `
		SELECT q.category 
		FROM game_answers ga
		JOIN questions q ON ga.question_id = q.id
		JOIN game_sessions gs ON ga.game_session_id = gs.id
		WHERE gs.user_id = ?
		GROUP BY q.category
		HAVING (SUM(CASE WHEN ga.is_correct THEN 1 ELSE 0 END) * 100.0 / COUNT(*)) < 50
		ORDER BY COUNT(*) DESC
	`
	
	db.Raw(query, userID).Pluck("category", &weakCategories)
	return weakCategories
}

func calculateUserStats(userID uint) UserStats {
	var stats UserStats
	stats.UserID = userID
	
	// Get total games
	db.Model(&GameSession{}).Where("user_id = ?", userID).Count(&stats.TotalGames)
	
	// Get overall performance
	var sessions []GameSession
	db.Where("user_id = ? AND status = ?", userID, "COMPLETED").Find(&sessions)
	
	for _, session := range sessions {
		stats.TotalQuestions += session.TotalQuestions
		stats.CorrectAnswers += session.CorrectAnswers
		if session.Score > stats.BestScore {
			stats.BestScore = session.Score
		}
	}
	
	if stats.TotalQuestions > 0 {
		stats.AverageScore = float64(stats.CorrectAnswers) / float64(stats.TotalQuestions) * 100
	}
	
	// Calculate category stats
	stats.CategoryStats = make(map[string]CategoryStats)
	categories := []string{"CULTURA", "GEOGRAFIA", "HISTORIA", "CONSTITUCION"}
	
	for _, category := range categories {
		var catStats CategoryStats
		catStats.Category = category
		
		// Get performance for this category
		query := `
			SELECT 
				COUNT(*) as total,
				SUM(CASE WHEN ga.is_correct THEN 1 ELSE 0 END) as correct
			FROM game_answers ga
			JOIN questions q ON ga.question_id = q.id
			JOIN game_sessions gs ON ga.game_session_id = gs.id
			WHERE gs.user_id = ? AND q.category = ?
		`
		
		var result struct {
			Total   int
			Correct int
		}
		db.Raw(query, userID, category).Scan(&result)
		
		catStats.TotalQuestions = result.Total
		catStats.CorrectAnswers = result.Correct
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
	
	// Get recent progress
	recentSessions := []GameSession{}
	db.Where("user_id = ? AND status = ?", userID, "COMPLETED").
		Order("created_at desc").Limit(10).Find(&recentSessions)
	
	for _, session := range recentSessions {
		point := ProgressPoint{
			Date:       session.CreatedAt,
			Score:      session.Score,
			Percentage: float64(session.CorrectAnswers) / float64(session.TotalQuestions) * 100,
		}
		stats.RecentProgress = append(stats.RecentProgress, point)
	}
	
	return stats
}

func identifyWeakAreas(userID uint) []string {
	stats := calculateUserStats(userID)
	return stats.WeakAreas
}

func generateRecommendations(sessionID uint) {
	var session GameSession
	db.First(&session, sessionID)
	
	// Analyze incorrect answers
	var incorrectAnswers []GameAnswer
	db.Preload("Question").
		Where("game_session_id = ? AND is_correct = false", sessionID).
		Find(&incorrectAnswers)
	
	// Group by category and subcategory
	weaknesses := make(map[string]int)
	for _, answer := range incorrectAnswers {
		key := fmt.Sprintf("%s_%s", answer.Question.Category, answer.Question.SubCategory)
		weaknesses[key]++
	}
	
	// Create recommendations
	for area, count := range weaknesses {
		if count >= 2 { // If failed 2+ questions in this area
			rec := StudyRecommendation{
				UserID:      session.UserID,
				Category:    getCategory(area),
				SubCategory: getSubCategory(area),
				Weakness:    area,
				Description: generateStudyDescription(area, count),
				Resources:   generateResources(area),
				Priority:    int(math.Min(5, float64(count))),
			}
			db.Create(&rec)
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

func splitString(s string, sep string) []string {
	if s == "" {
		return []string{}
	}
	result := []string{}
	current := ""
	for _, char := range s {
		if string(char) == sep {
			if current != "" {
				result = append(result, current)
				current = ""
			}
		} else {
			current += string(char)
		}
	}
	if current != "" {
		result = append(result, current)
	}
	return result
}

func generateQuestionSequence(userID uint, mode string, categoriesStr string, questionCount int) string {
	// Build query to get all available questions
	query := db.Model(&Question{})

	// Apply category filter
	if categoriesStr != "" {
		categories := splitString(categoriesStr, ",")
		if len(categories) > 0 {
			query = query.Where("category IN ?", categories)
		}
	} else if mode == "WEAK_AREAS" {
		weakCategories := getUserWeakCategories(userID)
		if len(weakCategories) > 0 {
			query = query.Where("category IN ?", weakCategories)
		}
	}

	// Get all matching question IDs, ordered by ID for consistency
	var questionIDs []uint
	query.Order("id ASC").Pluck("id", &questionIDs)

	if len(questionIDs) == 0 {
		return ""
	}

	// Limit to the requested number of questions
	if len(questionIDs) > questionCount {
		questionIDs = questionIDs[:questionCount]
	}

	// Convert to comma-separated string
	if len(questionIDs) == 0 {
		return ""
	}
	result := fmt.Sprintf("%d", questionIDs[0])
	for i := 1; i < len(questionIDs); i++ {
		result += fmt.Sprintf(",%d", questionIDs[i])
	}
	return result
}

func parseQuestionSequence(sequence string) []uint {
	if sequence == "" {
		return []uint{}
	}

	parts := splitString(sequence, ",")
	questionIDs := make([]uint, 0, len(parts))

	for _, part := range parts {
		if id, err := strconv.ParseUint(part, 10, 32); err == nil {
			questionIDs = append(questionIDs, uint(id))
		}
	}

	return questionIDs
}
