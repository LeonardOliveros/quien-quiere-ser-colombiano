package main

import (
	"context"
	"embed"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"quiz-app/internal/domain"
	"quiz-app/internal/seed"
	"quiz-app/internal/storage/dynamodb"
	"quiz-app/internal/storage/sqlite"
)

//go:embed data/taxonomy.json data/questions/*.json
var seedFS embed.FS

// store is the persistence port used by the HTTP handlers. The concrete
// adapter (SQLite locally, DynamoDB in the cloud) is chosen by DB_DRIVER.
var store domain.Store

func main() {
	// Load .env file
	godotenv.Load()

	// Initialize the storage adapter
	var err error
	store, err = openStore()
	if err != nil {
		log.Fatal("Failed to initialize storage: ", err)
	}
	defer store.Close()

	// Setup Gin router
	r := gin.Default()

	// Configure CORS: restrict to ALLOWED_ORIGINS when set (production),
	// otherwise allow all origins (development)
	config := cors.DefaultConfig()
	if origins := os.Getenv("ALLOWED_ORIGINS"); origins != "" {
		config.AllowOrigins = strings.Split(origins, ",")
	} else {
		config.AllowAllOrigins = true
	}
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

	// Sync the embedded question bank into the store
	taxonomy, seeds, err := seed.Load(seedFS)
	if err != nil {
		log.Fatal("Invalid seed data: ", err)
	}
	if err := store.SyncQuestionBank(taxonomy, seeds); err != nil {
		log.Fatal("Failed to sync question bank: ", err)
	}

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	go func() {
		log.Printf("Server starting on port %s", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Failed to start server:", err)
		}
	}()

	// Graceful shutdown on SIGINT/SIGTERM
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	log.Println("Server exited")
}

// openStore selects the storage adapter from the DB_DRIVER env var:
// "sqlite" (default) for local development, "dynamodb" for the cloud.
func openStore() (domain.Store, error) {
	driver := strings.ToLower(os.Getenv("DB_DRIVER"))
	switch driver {
	case "", "sqlite":
		dbPath := os.Getenv("DATABASE_PATH")
		if dbPath == "" {
			dbPath = "quiz.db"
		}
		log.Printf("Storage: sqlite (%s)", dbPath)
		return sqlite.Open(dbPath)
	case "dynamodb":
		log.Printf("Storage: dynamodb (table %s)", os.Getenv("DYNAMODB_TABLE"))
		return dynamodb.Open(os.Getenv("DYNAMODB_TABLE"))
	default:
		return nil, fmt.Errorf("unknown DB_DRIVER %q (supported: sqlite, dynamodb)", driver)
	}
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

			// Question routes (answers are stripped from responses)
			protected.GET("/questions", getQuestions)
			protected.GET("/questions/:id", getQuestion)
			protected.GET("/questions/category/:category", getQuestionsByCategory)
		}

		// Public question routes
		api.GET("/questions/count", getQuestionsCount)
		api.GET("/categories", getCategories)
	}
}
