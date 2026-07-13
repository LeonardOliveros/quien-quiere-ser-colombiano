package main

import (
	"context"
	"embed"
	"flag"
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

	seedOnly := flag.Bool("seed", false, "sync the question bank into the store and exit")
	flag.Parse()

	// Initialize the storage adapter
	var err error
	store, err = openStore()
	if err != nil {
		log.Fatal("Failed to initialize storage: ", err)
	}
	defer store.Close()

	// One-off seeding (e.g. against DynamoDB Local or a real AWS table).
	if *seedOnly {
		if err := syncQuestionBank(); err != nil {
			log.Fatal("Failed to sync question bank: ", err)
		}
		return
	}

	// Sync the embedded question bank into the store, unless disabled
	// (SEED_ON_START=false in Lambda: seeding happens at deploy time there,
	// not on every cold start).
	if seedOnStart() {
		if err := syncQuestionBank(); err != nil {
			log.Fatal("Failed to sync question bank: ", err)
		}
	}

	// Inside AWS Lambda the router is driven by API Gateway events instead of
	// a listening socket, and the SPA is served from S3/CloudFront.
	if os.Getenv("AWS_LAMBDA_FUNCTION_NAME") != "" {
		runLambda(buildRouter(false))
		return
	}

	r := buildRouter(true)

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

// buildRouter assembles the Gin engine. serveSPA additionally mounts the Vue
// build from ./dist (local single-binary mode); in Lambda the SPA lives on
// S3 behind CloudFront and only the API routes are needed.
func buildRouter(serveSPA bool) *gin.Engine {
	r := gin.Default()

	// Configure CORS: restrict to ALLOWED_ORIGINS when set. In release mode
	// (Lambda/production) ALLOWED_ORIGINS is required — fail fast instead of
	// silently falling back to a wildcard that would let any site call the
	// API with a stolen/leaked token. Local dev (debug mode) may still leave
	// it unset and get an open CORS policy for convenience.
	config := cors.DefaultConfig()
	config.AllowMethods = []string{"GET", "POST", "DELETE"}
	config.AllowHeaders = []string{"Content-Type", "Authorization"}
	config.AllowCredentials = false

	if origins := os.Getenv("ALLOWED_ORIGINS"); origins != "" {
		list := strings.Split(origins, ",")
		for i, o := range list {
			list[i] = strings.TrimSpace(o)
		}
		config.AllowOrigins = list
	} else if gin.Mode() == gin.ReleaseMode {
		log.Fatal("ALLOWED_ORIGINS must be set in release mode")
	} else {
		config.AllowAllOrigins = true
	}
	r.Use(cors.New(config))

	if serveSPA {
		// Serve Vue.js SPA static files
		r.Static("/assets", "./dist/assets")
		r.StaticFile("/favicon.ico", "./dist/favicon.ico")
	}

	// Routes
	setupRoutes(r)

	if serveSPA {
		// Serve Vue.js SPA for all non-API routes (must be last)
		r.NoRoute(func(c *gin.Context) {
			c.File("./dist/index.html")
		})
	}

	return r
}

// seedOnStart reports whether the startup question-bank sync is enabled.
// Defaults to true so the local `go run .` flow keeps working unchanged.
func seedOnStart() bool {
	switch strings.ToLower(os.Getenv("SEED_ON_START")) {
	case "false", "0", "no":
		return false
	}
	return true
}

// syncQuestionBank loads the embedded seed files and upserts them.
func syncQuestionBank() error {
	taxonomy, seeds, err := seed.Load(seedFS)
	if err != nil {
		return fmt.Errorf("invalid seed data: %w", err)
	}
	return store.SyncQuestionBank(taxonomy, seeds)
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
