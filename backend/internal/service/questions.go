package service

import (
	"context"

	"github.com/lawson/otterprep/internal/repository"
)

type QuestionService interface {
	CreateQuestion(ctx context.Context, question repository.Question) (int64, error)
	CreateQuestionOption(ctx context.Context, questionOption repository.QuestionOptions) (int64, error)
	GetQuestionById(ctx context.Context, id int64) (*repository.Question, error)
	GetQuestionOptions(ctx context.Context, questionId int64) ([]repository.QuestionOptions, error)
}

type questionService struct {
	questionRepository repository.QuestionRepository
}

func NewQuestionService(questionRepository repository.QuestionRepository) *questionService {
	return &questionService{questionRepository: questionRepository}
}

func (qs *questionService) CreateQuestion(ctx context.Context, question repository.Question) (int64, error) {
	return qs.questionRepository.CreateQuestion(ctx, question)
}

func (qs *questionService) CreateQuestionOption(ctx context.Context, questionOption repository.QuestionOptions) (int64, error) {
	return qs.questionRepository.CreateQuestionOption(ctx, questionOption)
}

func (qs *questionService) GetQuestionById(ctx context.Context, id int64) (*repository.Question, error) {
	return qs.questionRepository.GetQuestionById(ctx, id)
}

func (qs *questionService) GetQuestionOptions(ctx context.Context, questionId int64) ([]repository.QuestionOptions, error) {
	return qs.questionRepository.GetQuestionOptions(ctx, questionId)
}
