package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/lawson/otterprep/domain"
	"github.com/lawson/otterprep/internal/middleware"
	"github.com/lawson/otterprep/internal/service"
	"github.com/lawson/otterprep/pkg"
)

type AdminHandler struct {
	userService     service.UserServiceInterface
	questionService service.QuestionService
	logger          *log.Logger
}

func NewAdminHandler(userService service.UserServiceInterface, questionService service.QuestionService, logger *log.Logger) *AdminHandler {
	return &AdminHandler{
		userService:     userService,
		questionService: questionService,
		logger:          logger,
	}
}

// CreateBulkQuestions creates multiple questions and their options and answers.
// It returns an error if any.
func (ah *AdminHandler) CreateBulkQuestions(c echo.Context) error {
	var questions []domain.QuestionsData
	if err := c.Bind(&questions); err != nil {
		ah.logger.Println("error binding questions: ", err)
		return pkg.ErrorResponse(c, err, http.StatusBadRequest)
	}
	if err := c.Validate(&questions); err != nil {
		return err
	}
	subjectId := c.Param("subject_id")
	if subjectId == "" {
		ah.logger.Println("subject id is empty. Proceeding to return error.")
		return pkg.ErrorResponse(c, pkg.ErrSubjectNotFound, http.StatusBadRequest)
	}
	subjectIdInt, err := strconv.ParseInt(subjectId, 10, 64)
	if err != nil {
		ah.logger.Println("error parsing subject id: ", err)
		return pkg.ErrorResponse(c, err, http.StatusBadRequest)
	}
	err = ah.questionService.CreateMultipleQuestionBySubjectID(c.Request().Context(), subjectIdInt, questions)
	if err != nil {
		ah.logger.Println("error creating multiple questions: ", err)
		return pkg.ErrorResponse(c, err, http.StatusInternalServerError)
	}
	ah.logger.Println("Successfully created multiple questions. Proceeding to return success response.")
	return pkg.SuccessResponse(c, nil, http.StatusOK)
}

// UploadSingleQuestion uploads a single question and its options and answers.
// It returns an error if any.
func (ah *AdminHandler) UploadSingleQuestion(c echo.Context) error {
	userRole := c.Get("role").(string)
	if userRole != "admin" {
		ah.logger.Println("user is not admin. Proceeding to return error.")
		return pkg.ErrorResponse(c, pkg.ErrUnauthorized, http.StatusUnauthorized)
	}
	var question domain.QuestionsData
	if err := c.Bind(&question); err != nil {
		ah.logger.Println("error binding question: ", err)
		return pkg.ErrorResponse(c, err, http.StatusBadRequest)
	}
	if err := c.Validate(&question); err != nil {
		return err
	}
	subjectId := c.Param("subject_id")
	if subjectId == "" {
		ah.logger.Println("subject id is empty. Proceeding to return error.")
		return pkg.ErrorResponse(c, pkg.ErrSubjectNotFound, http.StatusBadRequest)
	}
	subjectIdInt, err := strconv.ParseInt(subjectId, 10, 64)
	if err != nil {
		ah.logger.Println("error parsing subject id: ", err)
		return pkg.ErrorResponse(c, err, http.StatusBadRequest)
	}
	_, err = ah.questionService.CreateQuestion(c.Request().Context(), subjectIdInt, question)
	if err != nil {
		ah.logger.Println("error creating question: ", err)
		return pkg.ErrorResponse(c, err, http.StatusInternalServerError)
	}
	ah.logger.Println("Successfully created question. Proceeding to return success response.")
	return pkg.SuccessResponse(c, nil, http.StatusCreated)
}

func (ah *AdminHandler) GetAllQuestions(c echo.Context) error {
	userRole := c.Get("role").(string)
	if userRole != "admin" {
		ah.logger.Println("user is not admin. Proceeding to return error.")
		return pkg.ErrorResponse(c, pkg.ErrUnauthorized, http.StatusUnauthorized)
	}
	questions, err := ah.questionService.GetAllQuestions(c.Request().Context())
	if err != nil {
		ah.logger.Println("error getting all questions: ", err)
		return pkg.ErrorResponse(c, err, http.StatusInternalServerError)
	}
	ah.logger.Println("Successfully got all questions. Proceeding to return success response.")
	return pkg.SuccessResponse(c, questions, http.StatusOK)
}

func (ah *AdminHandler) GetQuestionById(c echo.Context) error {
	//userRole := c.Get("role").(string)
	//if userRole != "admin" {
	//	ah.logger.Println("user is not admin. Proceeding to return error.")
	//	return pkg.ErrorResponse(c, pkg.ErrUnauthorized, http.StatusUnauthorized)
	//}
	questionId := c.Param("id")
	if questionId == "" {
		ah.logger.Println("question id is empty. Proceeding to return error.")
		return pkg.ErrorResponse(c, pkg.ErrQuestionNotFound, http.StatusBadRequest)
	}
	questionIdInt, err := strconv.ParseInt(questionId, 10, 64)
	if err != nil {
		ah.logger.Println("error parsing question id: ", err)
		return pkg.ErrorResponse(c, err, http.StatusBadRequest)
	}
	question, err := ah.questionService.GetQuestionById(c.Request().Context(), questionIdInt)
	if err != nil {
		ah.logger.Println("error getting question by id: ", err)
		return pkg.ErrorResponse(c, err, http.StatusInternalServerError)
	}
	ah.logger.Println("Successfully got question by id. Proceeding to return success response.")
	return pkg.SuccessResponse(c, question, http.StatusOK)
}

func (ah *AdminHandler) GetQuestionOptions(c echo.Context) error {
	userRole := c.Get("role").(string)
	if userRole != "admin" {
		ah.logger.Println("user is not admin. Proceeding to return error.")
		return pkg.ErrorResponse(c, pkg.ErrUnauthorized, http.StatusUnauthorized)
	}
	questionId := c.QueryParam("question_id")
	if questionId == "" {
		ah.logger.Println("question id is empty. Proceeding to return error.")
		return pkg.ErrorResponse(c, pkg.ErrQuestionNotFound, http.StatusBadRequest)
	}
	questionIdInt, err := strconv.ParseInt(questionId, 10, 64)
	if err != nil {
		ah.logger.Println("error parsing question id: ", err)
		return pkg.ErrorResponse(c, err, http.StatusBadRequest)
	}
	questionOptions, err := ah.questionService.GetQuestionOptions(c.Request().Context(), questionIdInt)
	if err != nil {
		ah.logger.Println("error getting question options: ", err)
		return pkg.ErrorResponse(c, err, http.StatusInternalServerError)
	}
	ah.logger.Println("Successfully got question options. Proceeding to return success response.")
	return pkg.SuccessResponse(c, questionOptions, http.StatusOK)
}

func (ah *AdminHandler) DeleteQuestionById(c echo.Context) error {
	userRole := c.Get("role").(string)
	if userRole != "admin" {
		ah.logger.Println("user is not admin. Proceeding to return error.")
		return pkg.ErrorResponse(c, pkg.ErrUnauthorized, http.StatusUnauthorized)
	}
	questionId := c.QueryParam("question_id")
	if questionId == "" {
		ah.logger.Println("question id is empty. Proceeding to return error.")
		return pkg.ErrorResponse(c, pkg.ErrQuestionNotFound, http.StatusBadRequest)
	}
	questionIdInt, err := strconv.ParseInt(questionId, 10, 64)
	if err != nil {
		ah.logger.Println("error parsing question id: ", err)
		return pkg.ErrorResponse(c, err, http.StatusBadRequest)
	}
	err = ah.questionService.DeleteQuestionById(c.Request().Context(), questionIdInt)
	if err != nil {
		ah.logger.Println("error deleting question by id: ", err)
		return pkg.ErrorResponse(c, err, http.StatusInternalServerError)
	}
	ah.logger.Println("Successfully deleted question by id. Proceeding to return success response.")
	return pkg.SuccessResponse(c, nil, http.StatusOK)
}

// CreateSubject creates a new subject.
// @Summary Create a new subject
// @Description Create a new subject
// @Tags Subject
// @Accept json
// @Produce json
// @Param subject body domain.Subject true "Subject"
// @Success 200 {object} pkg.SuccessResponse
// @Failure 400 {object} pkg.ErrorResponse
// @Failure 500 {object} pkg.ErrorResponse
// @Router /admin/subject [post]
func (ah *AdminHandler) CreateSubject(c echo.Context) error {
	userRole, ok := middleware.GetUserRole(c)
	if !ok {
		return pkg.ErrorResponse(c, pkg.ErrInvalidRole, http.StatusUnauthorized)
	}
	if userRole != "admin" {
		ah.logger.Println("user is not admin. Proceeding to return error.")
		return pkg.ErrorResponse(c, pkg.ErrUnauthorized, http.StatusUnauthorized)
	}
	var subject domain.Subject
	if err := c.Bind(&subject); err != nil {
		ah.logger.Println("error binding subject: ", err)
		return pkg.ErrorResponse(c, err, http.StatusBadRequest)
	}
	if err := c.Validate(&subject); err != nil {
		ah.logger.Println("error binding subject: ", err)
		return pkg.ErrorResponse(c, err, http.StatusBadRequest)
	}
	subjectId, err := ah.questionService.CreateSubject(c.Request().Context(), subject.Name)
	if err != nil {
		ah.logger.Println("error creating subject: ", err)
		return pkg.ErrorResponse(c, err, http.StatusInternalServerError)
	}
	ah.logger.Println("Successfully created subject with id: ", subjectId)
	result := map[string]interface{}{
		"subject_id": subjectId,
	}
	return pkg.SuccessResponse(c, result, http.StatusOK)
}

// GetSubjectById gets a subject by id.
// @Summary Get a subject by id
// @Description Get a subject by id
// @Tags Subject
// @Accept json
// @Produce json
// @Param subject_id query int true "Subject ID"
// @Success 200 {object} pkg.SuccessResponse
// @Failure 400 {object} pkg.ErrorResponse
// @Failure 500 {object} pkg.ErrorResponse
// @Router /admin/subject [get]
func (ah *AdminHandler) GetSubjectById(c echo.Context) error {
	//userRole, _ := middleware.GetUserRole(c)
	//if userRole != "admin" {
	//	ah.logger.Println("user is not admin. Proceeding to return error.")
	//	return pkg.ErrorResponse(c, pkg.ErrUnauthorized, http.StatusUnauthorized)
	//}
	subjectId := c.Param("id")
	if subjectId == "" {
		ah.logger.Println("subject id is empty. Proceeding to return error.")
		return pkg.ErrorResponse(c, pkg.ErrSubjectNotFound, http.StatusBadRequest)
	}
	subjectIdInt, err := strconv.ParseInt(subjectId, 10, 64)
	if err != nil {
		ah.logger.Println("error parsing subject id: ", err)
		return pkg.ErrorResponse(c, err, http.StatusBadRequest)
	}
	subject, err := ah.questionService.GetSubjectById(c.Request().Context(), subjectIdInt)
	if err != nil {
		ah.logger.Println("error getting subject by id: ", err)
		return pkg.ErrorResponse(c, err, http.StatusInternalServerError)
	}
	ah.logger.Println("Successfully got subject by id. Proceeding to return success response.")
	return pkg.SuccessResponse(c, subject, http.StatusOK)
}

// GetAllSubjects gets all subjects.
// @Summary Get all subjects
// @Description Get all subjects
// @Tags Subject
// @Accept json
// @Produce json
// @Success 200 {object} pkg.SuccessResponse
// @Failure 400 {object} pkg.ErrorResponse
// @Failure 500 {object} pkg.ErrorResponse
// @Router /admin/subject [get]
func (ah *AdminHandler) GetAllSubjects(c echo.Context) error {
	//userRole := c.Get("role").(string)
	////if userRole != "admin" {
	////	ah.logger.Println("user is not admin. Proceeding to return error.")
	////	return pkg.ErrorResponse(c, pkg.ErrUnauthorized, http.StatusUnauthorized)
	////}
	subjects, err := ah.questionService.GetAllSubjects(c.Request().Context())
	if err != nil {
		ah.logger.Println("error getting all subjects: ", err)
		return pkg.ErrorResponse(c, err, http.StatusInternalServerError)
	}
	ah.logger.Println("Successfully got all subjects. Proceeding to return success response.")
	return pkg.SuccessResponse(c, subjects, http.StatusOK)
}
