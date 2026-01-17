package router

import (
	"net/http"

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

	// Set up CORS
	var corsConfig middleware.CORSConfig
	if cfg.Server.Env == "production" {
		corsConfig = middleware.ProductionCORSConfig(cfg.Server.AllowOrigins)
	} else {
		corsConfig = middleware.DefaultCORSConfig()
	}
	e.Use(middleware.CORSMiddleware(corsConfig))

	// Set up custom validator
	e.Validator = middleware.NewValidator()

	// Handle 404 and 405 errors
	echo.NotFoundHandler = func(c echo.Context) error {
		return middleware.NotFoundHandler(c)
	}
	echo.MethodNotAllowedHandler = func(c echo.Context) error {
		return middleware.MethodNotAllowedHandler(c)
	}

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "OK")
	}, middleware.RateLimitMiddleware(middleware.HealthCheckLimiter))
	// Public routes with rate limiting
	// Auth routes - stricter rate limits
	authGroup := e.Group("")

	// Login endpoints - 5 attempts per minute
	e.POST("/user/login", userHandler.Login, middleware.RateLimitMiddleware(middleware.LoginRateLimiter))
	e.POST("/admin/login", userHandler.AdminLogin, middleware.RateLimitMiddleware(middleware.LoginRateLimiter))

	// Register endpoints - 3 attempts per minute
	e.POST("/user/register", userHandler.CreateUser, middleware.RateLimitMiddleware(middleware.RegisterRateLimiter))
	e.POST("/admin/register", userHandler.CreateUserAdmin, middleware.RateLimitMiddleware(middleware.RegisterRateLimiter))

	// Refresh token - 10 attempts per minute
	authGroup.POST("/auth/refresh", userHandler.RefreshToken, middleware.RateLimitMiddleware(middleware.RefreshTokenRateLimiter))

	// Password reset routes - rate limited (3 attempts per 5 minutes for email sending)
	e.POST("/auth/forgot-password", userHandler.ForgotPassword, middleware.RateLimitMiddleware(middleware.PasswordResetRateLimiter))
	e.POST("/auth/validate-reset-token", userHandler.ValidateResetToken)
	e.POST("/auth/reset-password", userHandler.ResetPassword, middleware.RateLimitMiddleware(middleware.LoginRateLimiter))

	// Protected routes with general API rate limiting
	api := e.Group("/api/v1")
	api.Use(middleware.JWTAuthMiddleware(cfg.Server.JWTSecret))
	api.Use(middleware.RateLimitMiddleware(middleware.APIRateLimiter))

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
