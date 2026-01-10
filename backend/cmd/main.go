package main

import (
	"log"
	"os"

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

	// Getting all services
	subjectService := service.NewSubjectService(subjectRepository)
	userService := service.NewUserService(*userRepository, scoreRepository, logger)
	quizService := service.NewQuizService(quizRepository, subjectRepository, questionRepository, scoreRepository)
	questionService := service.NewQuestionService(questionRepository, subjectRepository, logger)

	// Getting all handlers
	adminHandler := handler.NewAdminHandler(userService, questionService)
	userHandler := handler.NewUserHandler(userService, logger, cfg.Server.JWTSecret)
	quizHandler := handler.NewQuizHandler(quizService, subjectService, logger)

	e := echo.New()
	router.NewRouter(e, adminHandler, userHandler, quizHandler, cfg)

	logger.Printf("Starting server on port %s", cfg.Server.Port)
	if err := e.Start(":" + cfg.Server.Port); err != nil {
		logger.Fatal("Failed to start server: ", err)
	}
}
