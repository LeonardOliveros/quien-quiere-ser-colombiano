package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	// Load .env file
	godotenv.Load()

	// Initialize database
	initDB()

	// Setup Gin router
	r := gin.Default()
	
	// Configure CORS
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	r.Use(cors.New(config))

	// Serve Vue.js SPA static files
	r.Static("/assets", "./dist/assets")
	r.StaticFile("/favicon.ico", "./dist/favicon.ico")

	// Routes
	setupRoutes(r)

	// Serve Vue.js SPA for all non-API routes (must be last)
	r.NoRoute(func(c *gin.Context) {
		c.File("./dist/index.html")
	})

	// Seed database with questions if empty
	seedQuestions()

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

func initDB() {
	var err error
	db, err = gorm.Open(sqlite.Open("quiz.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto migrate models
	err = db.AutoMigrate(
		&Question{},
		&Choice{},
		&User{},
		&GameSession{},
		&GameAnswer{},
		&QuestionHistory{},
		&StudyRecommendation{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	log.Println("Database connected and migrated successfully")
}

func setupRoutes(r *gin.Engine) {
	// API routes
	api := r.Group("/api")
	{
		// Auth routes (public)
		api.POST("/register", registerUser)
		api.POST("/login", loginUser)

		// Protected routes (require authentication)
		protected := api.Group("")
		protected.Use(authRequired())
		{
			// Game routes
			protected.POST("/game/start", startGame)
			protected.GET("/game/:sessionId/question", getNextQuestion)
			protected.POST("/game/:sessionId/answer", submitAnswer)
			protected.POST("/game/:sessionId/flag", flagQuestion)
			protected.POST("/game/:sessionId/fifty-fifty", useFiftyFifty)
			protected.POST("/game/:sessionId/autosolve", useAutosolve)
			protected.GET("/game/:sessionId/results", getGameResults)
			protected.POST("/game/:sessionId/end", endGame)
			protected.POST("/game/:sessionId/pause", pauseGame)
			protected.GET("/game/paused/:mode", getPausedGame)
			protected.GET("/game/paused", getAnyPausedGame)

			// User statistics
			protected.GET("/user/:userId/stats", getUserStats)
			protected.GET("/user/:userId/weak-areas", getWeakAreas)
			protected.GET("/user/:userId/history", getGameHistory)
			protected.DELETE("/user/:userId/stats", resetUserStats)

			// Study recommendations
			protected.GET("/recommendations/:userId", getStudyRecommendations)
		}

		// Public question routes
		api.GET("/questions", getQuestions)
		api.GET("/questions/:id", getQuestion)
		api.GET("/questions/category/:category", getQuestionsByCategory)
		api.GET("/questions/count", getQuestionsCount)
	}
}
