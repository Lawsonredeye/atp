package service

import (
	"context"
	"fmt"
	"slices"
	"time"

	"github.com/lawson/otterprep/domain"
	"github.com/lawson/otterprep/internal/repository"
	"github.com/lawson/otterprep/pkg"
)

type quizService struct {
	quizRepository     repository.QuizRepository
	subjectRepository  repository.SubjectRepository
	questionRepository repository.QuestionRepository
	scoreRepository    repository.ScoreRepository
}

type QuizService interface {
	GenerateQuizBySubjectID(ctx context.Context, subjectId int64, numOfQuestions int64) (*domain.GeneratedQuizResponse, error)
	SubmitQuiz(ctx context.Context, userID int64, quizRequest []domain.SubmitQuizRequest) (*domain.QuizSubmitResponse, error)
	CalculateQuizScore(ctx context.Context, numOfQuestions int64, score int64) int64
}

func NewQuizService(quizRepository repository.QuizRepository, subjectRepository repository.SubjectRepository, questionRepository repository.QuestionRepository, scoreRepository repository.ScoreRepository) *quizService {
	return &quizService{quizRepository: quizRepository, subjectRepository: subjectRepository, questionRepository: questionRepository, scoreRepository: scoreRepository}
}

// GenerateQuizBySubjectID generates a quiz based on the subject ID and number of questions
// if subject is found then it returns the number of questions based on numOfQuestions.
// if subject is not found then it returns an error.
func (qs *quizService) GenerateQuizBySubjectID(ctx context.Context, subjectId int64, numOfQuestions int64) (*domain.GeneratedQuizResponse, error) {
	var questions []domain.QuizQuestionResponse
	usedQuestionIds := make(map[int64]bool)

	for i := 0; i < int(numOfQuestions); i++ {
		// Keep trying to get a unique question
		var question *repository.Questions
		var err error
		maxRetries := 10
		for retry := 0; retry < maxRetries; retry++ {
			question, err = qs.questionRepository.GetRandomQuestion(ctx, subjectId)
			if err != nil {
				fmt.Println("error getting quiz: ", err)
				return nil, pkg.ErrSubjectNotFound
			}
			if !usedQuestionIds[question.Id] {
				usedQuestionIds[question.Id] = true
				break
			}
		}

		questionOptions, err := qs.questionRepository.GetQuestionOptions(ctx, question.Id)
		if err != nil {
			return nil, pkg.ErrQuestionOptionNotFound
		}

		// Convert options without exposing is_correct
		options := make([]domain.QuizOptionResponse, len(questionOptions))
		for j, opt := range questionOptions {
			options[j] = domain.QuizOptionResponse{
				Id:     opt.Id,
				Option: opt.Option,
			}
		}

		questions = append(questions, domain.QuizQuestionResponse{
			QuestionId:       question.Id,
			Question:         question.Question,
			SubjectId:        question.SubjectId,
			IsMultipleChoice: question.IsMultipleChoice,
			Options:          options,
		})
	}

	return &domain.GeneratedQuizResponse{
		SubjectId:  subjectId,
		TotalCount: len(questions),
		Questions:  questions,
	}, nil
}

func (qs *quizService) GetQuizById(ctx context.Context, id int64) (*repository.Quiz, error) {
	return qs.quizRepository.GetQuizById(ctx, id)
}

// SubmitQuiz takes a list of quiz request and checks to see if the request questions
// has the correct options selected.
func (qs *quizService) SubmitQuiz(ctx context.Context, userID int64, quizRequest []domain.SubmitQuizRequest) (*domain.QuizSubmitResponse, error) {
	score := int64(0)
	correctAnswers := int64(0)
	incorrectAnswers := int64(0)
	if len(quizRequest) == 0 {
		return nil, nil
	}
	results := make([]domain.QuizResultResponse, 0)

	var subjectID int64

	for _, quiz := range quizRequest {
		question, err := qs.questionRepository.GetQuestionById(ctx, quiz.QuestionId)
		if err != nil {
			fmt.Println("error getting quiz: ", err)
			continue
		}

		// Capture subjectID from the first question found
		if subjectID == 0 {
			subjectID = question.SubjectId
		}

		answer, err := qs.questionRepository.GetAnswerById(ctx, quiz.QuestionId)
		if err != nil {
			fmt.Println("error getting answer: ", err)
		}
		correctOption, err := qs.questionRepository.GetCorrectQuestionOptionByQuestionID(ctx, quiz.QuestionId)
		if err != nil {
			fmt.Println("error getting question options: ", err)
		}

		isCorrect := slices.Contains(quiz.OptionIds, correctOption.Id)
		if isCorrect {
			score++
			correctAnswers++
		} else {
			incorrectAnswers++
		}

		// Get selected option texts
		selectedOpts := make([]string, 0)
		for _, optionId := range quiz.OptionIds {
			questionOption, err := qs.questionRepository.GetQuestionOptionsById(ctx, optionId)
			if err != nil {
				fmt.Println("error getting question options: ", err)
				continue
			}
			selectedOpts = append(selectedOpts, questionOption.Option)
		}

		results = append(results, domain.QuizResultResponse{
			QuestionId:      question.Id,
			Question:        question.Question,
			SelectedOptions: selectedOpts,
			CorrectAnswer:   correctOption.Option,
			IsCorrect:       isCorrect,
			Explanation:     answer.Answer,
		})
	}

	// Persist the score
	_, err := qs.scoreRepository.StoreUserScore(ctx, domain.UserScore{
		UserID:           userID,
		Score:            score,
		Mode:             "practice", // Default mode
		CorrectAnswers:   correctAnswers,
		IncorrectAnswers: incorrectAnswers,
		TotalQuestions:   int64(len(quizRequest)),
		TimeTakenSeconds: 0, // Not tracked yet
		SubjectID:        subjectID,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	})
	if err != nil {
		fmt.Println("error storing score: ", err)
		return nil, err
	}

	return &domain.QuizSubmitResponse{
		UserId:           userID,
		SubjectId:        subjectID,
		TotalQuestions:   int64(len(quizRequest)),
		CorrectAnswers:   correctAnswers,
		IncorrectAnswers: incorrectAnswers,
		Score:            score,
		Results:          results,
	}, nil
}

// CalculateQuizScore takes a number of questions and a score and returns the percentage of the score.
func (qs *quizService) CalculateQuizScore(ctx context.Context, numOfQuestions int64, score int64) int64 {
	if numOfQuestions == 0 {
		return 0
	}
	return score * 100 / numOfQuestions
}
