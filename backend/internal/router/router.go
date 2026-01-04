package router

import (
	"github.com/labstack/echo/v4"
	"github.com/lawson/otterprep/internal/handler"
)

func NewRouter(
	e *echo.Echo,
	adminHandler *handler.AdminHandler,
	userHandler *handler.UserHandler,
	quizHandler *handler.QuizHandler,
) {
	// Admin routes
	e.POST("/admin/questions/bulk", adminHandler.CreateBulkQuestions)
	e.POST("/admin/questions/single", adminHandler.UploadSingleQuestion)
	e.GET("/admin/questions", adminHandler.GetAllQuestions)
	e.GET("/admin/questions/:question_id", adminHandler.GetQuestionById)

	// Subject routes
	e.GET("/admin/subject", adminHandler.GetAllSubjects)
	e.GET("/admin/subject/:subject_id", adminHandler.GetSubjectById)
	e.POST("/admin/subject", adminHandler.CreateSubject)

	// User routes
	e.POST("/user/login", userHandler.Login)
	e.POST("/user/register", userHandler.CreateUser)
	e.POST("/admin/register", userHandler.CreateUserAdmin)

	// Quiz routes
	e.POST("/quiz/create", quizHandler.CreateQuiz)
	e.GET("/quiz/submit", quizHandler.SubmitQuiz)
}
