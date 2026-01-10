package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/lawson/otterprep/config"
	"github.com/lawson/otterprep/internal/handler"
	"github.com/lawson/otterprep/internal/repository"
	"github.com/lawson/otterprep/internal/router"
	"github.com/lawson/otterprep/internal/service"

	_ "github.com/lib/pq"
)

func main() {
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	cfg, err := config.Load()
	if err != nil {
		logger.Fatal("Failed to load config: ", err)
	}
	logger.Println("Connecting to database...")
	dbConn := cfg.Database.PostgresInit()
	logger.Println("Database connected successfully")

	// Getting all repositories
	subjectRepository := repository.NewSubjectRepository(dbConn)
	userRepository := repository.NewUserRepository(dbConn)
	quizRepository := repository.NewQuizRepository(dbConn)
	scoreRepository := repository.NewScoreRepository(dbConn)
	questionRepository := repository.NewQuestionRepository(dbConn)
	leaderboardRepository := repository.NewLeaderboardRepository(dbConn)

	// Getting all services
	subjectService := service.NewSubjectService(subjectRepository)
	userService := service.NewUserService(*userRepository, scoreRepository, logger)
	quizService := service.NewQuizService(quizRepository, subjectRepository, questionRepository, scoreRepository)
	questionService := service.NewQuestionService(questionRepository, subjectRepository, logger)
	leaderboardService := service.NewLeaderboardService(leaderboardRepository, subjectRepository)

	// Getting all handlers
	adminHandler := handler.NewAdminHandler(userService, questionService, logger)
	userHandler := handler.NewUserHandler(userService, logger, cfg.Server.JWTSecret)
	quizHandler := handler.NewQuizHandler(quizService, subjectService, logger)
	leaderboardHandler := handler.NewLeaderboardHandler(leaderboardService, logger)

	e := echo.New()
	router.NewRouter(e, adminHandler, userHandler, quizHandler, leaderboardHandler, cfg)

	// Start server in a goroutine
	go func() {
		logger.Printf("Starting server on port %s", cfg.Server.Port)
		if err := e.Start(":" + cfg.Server.Port); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Fatal("Failed to start server: ", err)
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Println("Shutting down server...")

	// Create a deadline for the shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Shutdown the Echo server
	if err := e.Shutdown(ctx); err != nil {
		logger.Printf("Error during server shutdown: %v", err)
	}

	// Close database connection
	if err := dbConn.Close(); err != nil {
		logger.Printf("Error closing database connection: %v", err)
	}

	logger.Println("Server gracefully stopped")
}
