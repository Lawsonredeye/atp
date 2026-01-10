package service

import (
	"context"
	"log"
	"time"

	"github.com/lawson/otterprep/domain"
	"github.com/lawson/otterprep/internal/repository"
	"github.com/lawson/otterprep/pkg"
)

type QuestionService interface {
	CreateQuestion(ctx context.Context, subjectId int64, question domain.QuestionsData) (int64, error)
	CreateQuestionOption(ctx context.Context, questionOption repository.QuestionOptions) (int64, error)
	CreateMultipleQuestionBySubjectID(ctx context.Context, subjectId int64, questions []domain.QuestionsData) error
	GetQuestionById(ctx context.Context, id int64) (*domain.Question, error)
	GetQuestionOptions(ctx context.Context, questionId int64) ([]repository.QuestionOptions, error)
	GetAllQuestions(ctx context.Context) ([]repository.Questions, error)
	DeleteQuestionById(ctx context.Context, id int64) error
	CreateSubject(ctx context.Context, subject string) (int64, error)
	GetSubjectById(ctx context.Context, id int64) (*domain.Subject, error)
	GetAllSubjects(ctx context.Context) ([]repository.Subject, error)
}

type questionService struct {
	questionRepository repository.QuestionRepository
	subjectRepository  repository.SubjectRepository
	logger             *log.Logger
}

func (qs *questionService) GetAllSubjects(ctx context.Context) ([]repository.Subject, error) {
	qs.logger.Println("Getting all subjects.")
	result, err := qs.subjectRepository.GetSubjects(ctx)
	if err != nil {
		qs.logger.Println("Failed to get subjects: ", err)
		return nil, err
	}
	qs.logger.Println("Successfully got subjects. Proceeding to return result.")
	return result, nil
}

func NewQuestionService(questionRepository repository.QuestionRepository, subjectRepository repository.SubjectRepository, logger *log.Logger) *questionService {
	return &questionService{questionRepository: questionRepository, subjectRepository: subjectRepository, logger: logger}
}

// CreateQuestion creates a new question and its options and answer.
// It returns the id of the created question and an error if any.
func (qs *questionService) CreateQuestion(ctx context.Context, subjectId int64, question domain.QuestionsData) (int64, error) {

	if question.Name == "" {
		qs.logger.Println("Question name is empty. Proceeding to return error.")
		return 0, pkg.ErrQuestionTextNotFound
	}

	if subjectId == 0 {
		qs.logger.Println("Subject id is 0. Proceeding to return error.")
		return 0, pkg.ErrSubjectNotFound
	}
	qs.logger.Println("check if subject exists.")
	_, err := qs.subjectRepository.GetSubjectById(ctx, subjectId)
	if err != nil {
		qs.logger.Println("Failed to get subject by id: ", err)
		return 0, err
	}
	qs.logger.Println("Successfully got subject by id. Proceeding to create question.")

	id, err := qs.questionRepository.CreateQuestion(ctx, repository.Questions{
		SubjectId: subjectId,
		Question:  question.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		qs.logger.Println("Failed to create question: ", err)
		return 0, err
	}
	qs.logger.Println("Successfully created question. Proceeding to create options.")
	for _, option := range question.Options {
		_, err = qs.CreateQuestionOption(ctx, repository.QuestionOptions{
			QuestionId: id,
			Option:     option,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
			IsCorrect:  option == question.Answer,
		})
		if err != nil {
			qs.logger.Println("Failed to create question option: ", err)
			return 0, err
		}
	}
	qs.logger.Println("Successfully created question options. Proceeding to create answer.")
	_, err = qs.questionRepository.CreateAnswer(ctx, repository.Answers{
		QuestionId: id,
		Answer:     question.Explanation,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	})
	if err != nil {
		qs.logger.Println("Failed to create answer: ", err)
		return 0, err
	}
	qs.logger.Println("Successfully created answer. Proceeding to return id.")
	return id, nil
}

// CreateQuestionOption creates a question option.
// It returns the id of the created question option and an error if any.
func (qs *questionService) CreateQuestionOption(ctx context.Context, questionOption repository.QuestionOptions) (int64, error) {
	if questionOption.QuestionId == 0 {
		qs.logger.Println("Question id is 0. Proceeding to return error.")
		return 0, pkg.ErrQuestionNotFound
	}
	qs.logger.Println("Successfully created question option. Proceeding to return id.")
	result, err := qs.questionRepository.CreateQuestionOption(ctx, questionOption)
	if err != nil {
		qs.logger.Println("Failed to create question option: ", err)
		return 0, err
	}
	qs.logger.Println("Successfully created question option. Proceeding to return id.")
	return result, nil
}

// GetAllQuestions gets all questions.
// It returns the questions and an error if any.
func (qs *questionService) GetAllQuestions(ctx context.Context) ([]repository.Questions, error) {
	qs.logger.Println("Successfully got all questions. Proceeding to return success response.")
	return qs.questionRepository.GetAllQuestions(ctx)
}

// GetQuestionById gets a question by id.
// It returns the question and an error if any.
func (qs *questionService) GetQuestionById(ctx context.Context, id int64) (*domain.Question, error) {
	if id == 0 {
		qs.logger.Println("Question id is 0. Proceeding to return error.")
		return nil, pkg.ErrQuestionNotFound
	}
	qs.logger.Println("Successfully got question by id. Proceeding to get question.")
	result, err := qs.questionRepository.GetQuestionById(ctx, id)
	if err != nil {
		qs.logger.Println("Failed to get question by id: ", err)
		return nil, err
	}
	qs.logger.Println("Successfully got question by id. Proceeding to get the question options.")
	questionOptions, err := qs.questionRepository.GetQuestionOptions(ctx, id)
	if err != nil {
		qs.logger.Println("Failed to get question options: ", err)
		return nil, err
	}
	options := make([]string, len(questionOptions))
	for i, option := range questionOptions {
		options[i] = option.Option
	}
	domainQuestion := domain.Question{
		ID:          result.Id,
		Text:        result.Question,
		Option:      options,
		Answer:      "",
		Explanation: "",
	}
	qs.logger.Println("Successfully got question options. Proceeding to return result.")
	return &domainQuestion, nil
}

// GetQuestionOptions gets the options of a question by id.
// It returns the options and an error if any.
func (qs *questionService) GetQuestionOptions(ctx context.Context, questionId int64) ([]repository.QuestionOptions, error) {
	if questionId == 0 {
		qs.logger.Println("Question id is 0. Proceeding to return error.")
		return nil, pkg.ErrQuestionNotFound
	}
	qs.logger.Println("Successfully got question options. Proceeding to get question options.")
	result, err := qs.questionRepository.GetQuestionOptions(ctx, questionId)
	if err != nil {
		qs.logger.Println("Failed to get question options: ", err)
		return nil, err
	}
	qs.logger.Println("Successfully got question options. Proceeding to return result.")
	return result, nil
}

// CreateMultipleQuestionBySubjectID creates multiple questions and their options and answers.
// It returns an error if any.
func (qs *questionService) CreateMultipleQuestionBySubjectID(ctx context.Context, subjectId int64, questions []domain.QuestionsData) error {
	if subjectId == 0 {
		qs.logger.Println("Subject id is 0. Proceeding to return error.")
		return pkg.ErrSubjectNotFound
	}
	for _, question := range questions {
		id, err := qs.CreateQuestion(ctx, subjectId, question)
		if err != nil {
			qs.logger.Println("Failed to create question: ", err)
			break
		}
		for _, option := range question.Options {
			isCorrect := false
			if option == question.Answer {
				isCorrect = true
			}
			now := time.Now()
			_, err = qs.CreateQuestionOption(ctx, repository.QuestionOptions{
				QuestionId: id,
				Option:     option,
				CreatedAt:  now,
				UpdatedAt:  now,
				IsCorrect:  isCorrect,
			})
			if err != nil {
				qs.logger.Println("Failed to create question option: ", err)
				break
			}
		}
		now := time.Now()
		_, err = qs.questionRepository.CreateAnswer(ctx, repository.Answers{
			QuestionId: id,
			Answer:     question.Explanation,
			CreatedAt:  now,
			UpdatedAt:  now,
		})
		if err != nil {
			qs.logger.Println("Failed to create answer: ", err)
			break
		}
	}
	qs.logger.Println("Successfully created questions")
	return nil
}

// DeleteQuestionById deletes a question by id.
func (qs *questionService) DeleteQuestionById(ctx context.Context, id int64) error {
	if id > 1 {
		qs.logger.Println("Question id is greater than 1. Proceeding to delete question.")
		return pkg.ErrQuestionNotFound
	}
	return qs.questionRepository.DeleteQuestionById(ctx, id)
}

// CreateSubject creates a subject.
func (qs *questionService) CreateSubject(ctx context.Context, subjectName string) (int64, error) {
	if subjectName == "" {
		qs.logger.Println("Subject name is empty. Proceeding to return error.")
		return 0, pkg.ErrSubjectNameNotFound
	}
	// check if subject already exists
	subject, err := qs.GetSubjectByName(ctx, subjectName)
	if err == nil {
		qs.logger.Println("Subject already exists. Proceeding to return error.")
		return subject.Id, pkg.ErrSubjectWithNameExists
	}
	now := time.Now()
	subjectId, err := qs.subjectRepository.CreateSubject(ctx, repository.Subject{
		Name:      subjectName,
		CreatedAt: now,
		UpdatedAt: now,
	})
	if err != nil {
		qs.logger.Println("Failed to create subject: ", err)
		return 0, err
	}
	qs.logger.Println("Successfully created subject. Proceeding to return id.")
	return subjectId, nil
}

// GetSubjectById gets a subject by id from the subject repository.
func (qs *questionService) GetSubjectById(ctx context.Context, id int64) (*domain.Subject, error) {
	if id > 1 {
		qs.logger.Println("Subject id is greater than 1. Proceeding to return error.")
		return nil, pkg.ErrSubjectNotFound
	}
	qs.logger.Println("Successfully got subject by id. Proceeding to get subject.")
	result, err := qs.subjectRepository.GetSubjectById(ctx, id)
	if err != nil {
		qs.logger.Println("Failed to get subject by id: ", err)
		return nil, pkg.ErrSubjectNotFound
	}
	qs.logger.Println("Successfully got subject by id. Proceeding to return result.")
	domainSubject := domain.Subject{
		Id:   result.Id,
		Name: result.Name,
	}
	return &domainSubject, nil
}

// GetSubjectByName gets a subject by name from the subject repository.
func (qs *questionService) GetSubjectByName(ctx context.Context, subjectName string) (*repository.Subject, error) {
	if subjectName == "" {
		qs.logger.Println("Subject name is empty. Proceeding to return error.")
		return nil, pkg.ErrSubjectNameNotFound
	}
	qs.logger.Println("Successfully got subject by name. Proceeding to get subject.")
	result, err := qs.subjectRepository.GetSubjectByName(ctx, subjectName)
	if err != nil {
		qs.logger.Println("Failed to get subject by name: ", err)
		return nil, err
	}
	qs.logger.Println("Successfully got subject by name. Proceeding to return result.")
	return result, nil
}

// GetSubjects gets all subjects from the subject repository.
func (qs *questionService) GetSubjects(ctx context.Context) ([]repository.Subject, error) {
	qs.logger.Println("Successfully got subjects. Proceeding to get subjects.")
	result, err := qs.subjectRepository.GetSubjects(ctx)
	if err != nil {
		qs.logger.Println("Failed to get subjects: ", err)
		return nil, err
	}
	qs.logger.Println("Successfully got subjects. Proceeding to return result.")
	return result, nil
}
