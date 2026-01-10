package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lawson/otterprep/domain"
	"github.com/lawson/otterprep/internal/service"
	"github.com/lawson/otterprep/pkg"
)

type QuizHandler struct {
	quizService    service.QuizService
	subjectService service.SubjectService
	logger         *log.Logger
}

func NewQuizHandler(quizService service.QuizService, subjectService service.SubjectService, logger *log.Logger) *QuizHandler {
	return &QuizHandler{
		quizService:    quizService,
		subjectService: subjectService,
		logger:         logger,
	}
}

// =========================================================
// 		Quiz Handler
// =========================================================

// CreateQuiz creates a new quiz
// @Summary Create a new quiz
// @Tags Quizzes
// @Accept JSON
// @Produce JSON
// @Param quiz body domain.Quiz true "Quiz"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /quizzes [post]
func (h *QuizHandler) CreateQuiz(c echo.Context) error {
	var quizRequest domain.QuizRequest
	if err := c.Bind(&quizRequest); err != nil {
		return pkg.ErrorResponse(c, err, http.StatusBadRequest)
	}
	if err := c.Validate(&quizRequest); err != nil {
		return err
	}
	subjectId := quizRequest.SubjectId
	// check if subject with id exists
	_, err := h.subjectService.GetSubjectById(c.Request().Context(), subjectId)
	if err != nil {
		h.logger.Println("error getting subject: ", err)
		return pkg.ErrorResponse(c, err, http.StatusInternalServerError)
	}
	quiz, err := h.quizService.GenerateQuizBySubjectID(c.Request().Context(), quizRequest.SubjectId, quizRequest.NumOfQuestions)
	if err != nil {
		h.logger.Println("error creating quiz: ", err)
		return pkg.ErrorResponse(c, err, http.StatusInternalServerError)
	}
	return pkg.SuccessResponse(c, quiz, http.StatusOK)
}

// SubmitQuiz submits a quiz
// @Summary Submit a quiz
// @Tags Quizzes
// @Accept JSON
// @Produce JSON
// @Param quiz body []domain.SubmitQuizRequest true "Quiz Submission"
// @Success 200 {object} domain.QuizSubmitResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /quizzes/submit [post]
func (h *QuizHandler) SubmitQuiz(c echo.Context) error {
	var quizRequest []domain.SubmitQuizRequest
	if err := c.Bind(&quizRequest); err != nil {
		return pkg.ErrorResponse(c, err, http.StatusBadRequest)
	}
	if err := c.Validate(&quizRequest); err != nil {
		return err
	}

	userId := c.Get("user_id").(int64)
	result, err := h.quizService.SubmitQuiz(c.Request().Context(), userId, quizRequest)
	if err != nil {
		h.logger.Println("error submitting quiz: ", err)
		return pkg.ErrorResponse(c, err, http.StatusInternalServerError)
	}
	return pkg.SuccessResponse(c, result, http.StatusOK)
}
