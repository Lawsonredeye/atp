package service

import (
	"context"
	"fmt"

	"github.com/lawson/otterprep/internal/repository"
	"github.com/lawson/otterprep/pkg"
)

type quizService struct {
	quizRepository     repository.QuizRepository
	subjectRepository  repository.SubjectRepository
	questionRepository repository.QuestionRepository
}

type QuizService interface {
	GenerateQuizBySubjectID(ctx context.Context, subjectId int64, numOfQuestions int64) ([]repository.Quiz, error)
}

func NewQuizService(quizRepository repository.QuizRepository, subjectRepository repository.SubjectRepository, questionRepository repository.QuestionRepository) *quizService {
	return &quizService{quizRepository: quizRepository, subjectRepository: subjectRepository, questionRepository: questionRepository}
}

func (qs *quizService) GenerateQuizBySubjectID(ctx context.Context, subjectId int64, numOfQuestions int64) ([]repository.Quiz, error) {
	var quiz []repository.Quiz

	for i := 0; i < int(numOfQuestions); i++ {
		question, err := qs.questionRepository.GetRandomQuestion(ctx, subjectId)
		if err != nil {
			// log error here
			fmt.Println("error getting quiz: ", err)
			return nil, pkg.ErrQuestionNotFound
		}

		questionOption, err := qs.questionRepository.GetQuestionOptions(ctx, int64(question.Id))
		if err != nil {
			// log error here
			return nil, pkg.ErrQuestionOptionNotFound
		}

		quiz = append(quiz, repository.Quiz{
			Text:             question.Text,
			SubjectId:        int64(question.SubjectId),
			IsMultipleChoice: question.IsMultipleChoice,
			QuestionOptions:  questionOption,
			CreatedAt:        question.CreatedAt,
			UpdatedAt:        question.UpdatedAt,
		})
	}

	return quiz, nil
}
