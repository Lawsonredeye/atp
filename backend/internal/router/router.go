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
	leaderboardHandler *handler.LeaderboardHandler,
	cfg *config.Config,
) {
	// Set up error handlers
	e.HTTPErrorHandler = middleware.CustomHTTPErrorHandler
	e.Use(middleware.RecoverMiddleware())

	// Set up custom validator
	e.Validator = middleware.NewValidator()

	// Handle 404 and 405 errors
	echo.NotFoundHandler = func(c echo.Context) error {
		return middleware.NotFoundHandler(c)
	}
	echo.MethodNotAllowedHandler = func(c echo.Context) error {
		return middleware.MethodNotAllowedHandler(c)
	}

	// User routes
	e.POST("/user/login", userHandler.Login)
	e.POST("/user/register", userHandler.CreateUser)
	e.POST("/admin/register", userHandler.CreateUserAdmin)
	e.POST("/admin/login", userHandler.AdminLogin)

	// Protected
	api := e.Group("/api/v1")
	api.Use(middleware.JWTAuthMiddleware(cfg.Server.JWTSecret))

	// User
	api.GET("/dashboard", userHandler.UserDashboard)
	api.PUT("/user/username", userHandler.UpdateUsername)
	api.PUT("/user/email", userHandler.UpdateEmail)
	api.PUT("/user/password", userHandler.UpdatePassword)
	api.DELETE("/user/account", userHandler.DeleteUserAccount)

	// Admin routes
	api.POST("/admin/questions/bulk/:subject_id", adminHandler.CreateBulkQuestions)
	api.POST("/admin/questions/single/:subject_id", adminHandler.UploadSingleQuestion)
	api.GET("/admin/questions", adminHandler.GetAllQuestions)
	api.GET("/admin/questions/:id", adminHandler.GetQuestionById)
	api.DELETE("/admin/questions/:id", adminHandler.DeleteQuestionById)

	// Subject routes
	api.GET("/admin/subject", adminHandler.GetAllSubjects)
	api.GET("/admin/subject/:id", adminHandler.GetSubjectById)
	api.POST("/admin/subject", adminHandler.CreateSubject)

	// Quiz routes
	api.POST("/quiz/create", quizHandler.CreateQuiz)
	api.POST("/quiz/submit", quizHandler.SubmitQuiz)

	// Leaderboard routes
	api.GET("/leaderboard", leaderboardHandler.GetLeaderboard)
	api.GET("/leaderboard/me", leaderboardHandler.GetMyRank)
	api.GET("/leaderboard/user/:user_id", leaderboardHandler.GetUserRank)
}
