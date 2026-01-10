package router

import (
	"github.com/labstack/echo/v4"
	"github.com/lawson/otterprep/config"
	"github.com/lawson/otterprep/internal/handler"
	"github.com/lawson/otterprep/internal/middleware"
)

func NewRouter(
	e *echo.Echo,
	adminHandler *handler.AdminHandler,
	userHandler *handler.UserHandler,
	quizHandler *handler.QuizHandler,
	cfg *config.Config,
) {

	// User routes
	e.POST("/user/login", userHandler.Login)
	e.POST("/user/register", userHandler.CreateUser)
	e.POST("/admin/register", userHandler.CreateUserAdmin)
	e.POST("/admin/login", userHandler.AdminLogin)

	// Protected
	api := e.Group("/api/v1")
	api.Use(middleware.JWTAuthMiddleware(cfg.Server.JWTSecret))

	// Admin routes
	api.POST("/admin/questions/bulk", adminHandler.CreateBulkQuestions)
	api.POST("/admin/questions/single", adminHandler.UploadSingleQuestion)
	api.GET("/admin/questions", adminHandler.GetAllQuestions)
	api.GET("/admin/questions/:question_id", adminHandler.GetQuestionById)

	// Subject routes
	api.GET("/admin/subject", adminHandler.GetAllSubjects)
	api.GET("/admin/subject/:subject_id", adminHandler.GetSubjectById)
	api.POST("/admin/subject", adminHandler.CreateSubject)

	// Quiz routes
	api.POST("/quiz/create", quizHandler.CreateQuiz)
	api.GET("/quiz/submit", quizHandler.SubmitQuiz)
}
