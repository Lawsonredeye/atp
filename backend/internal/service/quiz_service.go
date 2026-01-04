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

type QuizRequest struct {
	QuizId           int64   `json:"quiz_id"`
	IsMultipleChoice bool    `json:"is_multiple_choice"`
	OptionIds        []int64 `json:"option_ids"`
}

type quizService struct {
	quizRepository     repository.QuizRepository
	subjectRepository  repository.SubjectRepository
	questionRepository repository.QuestionRepository
	scoreRepository    repository.ScoreRepository
}

type QuizService interface {
	GenerateQuizBySubjectID(ctx context.Context, subjectId int64, numOfQuestions int64) ([]repository.Quiz, error)
	SubmitQuiz(ctx context.Context, userID int64, quizRequest []QuizRequest) ([]domain.QuizResponse, int64, error)
	CalculateQuizScore(ctx context.Context, numOfQuestions int64, score int64) int64
}

func NewQuizService(quizRepository repository.QuizRepository, subjectRepository repository.SubjectRepository, questionRepository repository.QuestionRepository, scoreRepository repository.ScoreRepository) *quizService {
	return &quizService{quizRepository: quizRepository, subjectRepository: subjectRepository, questionRepository: questionRepository, scoreRepository: scoreRepository}
}

// GenerateQuizBySubjectID generates a quiz based on the subject ID and number of questions
// if subject is found then it returns the number of questions based on numOfQuestions.
// if subject is not found then it returns an error.
func (qs *quizService) GenerateQuizBySubjectID(ctx context.Context, subjectId int64, numOfQuestions int64) ([]repository.Quiz, error) {
	var quiz []repository.Quiz

	for i := 0; i < int(numOfQuestions); i++ {
		question, err := qs.questionRepository.GetRandomQuestion(ctx, subjectId)
		if err != nil {
			fmt.Println("error getting quiz: ", err)
			return nil, pkg.ErrSubjectNotFound
		}

		questionOption, err := qs.questionRepository.GetQuestionOptions(ctx, question.Id)
		if err != nil {
			// log error here
			return nil, pkg.ErrQuestionOptionNotFound
		}

		quiz = append(quiz, repository.Quiz{
			Question:         question.Question,
			SubjectId:        question.SubjectId,
			IsMultipleChoice: question.IsMultipleChoice,
			QuestionOptions:  questionOption,
			CreatedAt:        question.CreatedAt,
			UpdatedAt:        question.UpdatedAt,
		})
	}

	return quiz, nil
}

func (qs *quizService) GetQuizById(ctx context.Context, id int64) (*repository.Quiz, error) {
	return qs.quizRepository.GetQuizById(ctx, id)
}

// SubmitQuiz takes a list of quiz request and checks to see if the request questions
// has the correct options selected.
func (qs *quizService) SubmitQuiz(ctx context.Context, userID int64, quizRequest []QuizRequest) ([]domain.QuizResponse, int64, error) {
	score := int64(0)
	correctAnswers := int64(0)
	incorrectAnswers := int64(0)
	if len(quizRequest) == 0 {
		return nil, 0, nil
	}
	result := make([]domain.QuizResponse, 0)

	var subjectID int64

	for _, quiz := range quizRequest {
		question, err := qs.questionRepository.GetQuestionById(ctx, quiz.QuizId)
		if err != nil {
			fmt.Println("error getting quiz: ", err)
			break
		}

		// Capture subjectID from the first question found
		if subjectID == 0 {
			subjectID = question.SubjectId
		}

		answer, err := qs.questionRepository.GetAnswerById(ctx, quiz.QuizId)
		if err != nil {
			fmt.Println("error getting answer: ", err)
		}
		correctOption, err := qs.questionRepository.GetCorrectQuestionOptionByQuestionID(ctx, quiz.QuizId)
		if err != nil {
			fmt.Println("error getting question options: ", err)
		}

		if slices.Contains(quiz.OptionIds, correctOption.Id) {
			score++
			correctAnswers++
		} else {
			incorrectAnswers++
		}

		opts := make([]string, 0)
		for _, optionId := range quiz.OptionIds {
			questionOption, err := qs.questionRepository.GetQuestionOptionsById(ctx, optionId)
			if err != nil {
				fmt.Println("error getting question options: ", err)
			}
			opts = append(opts, questionOption.Option)
		}
		result = append(result, domain.QuizResponse{
			QuestionId:      question.Id,
			Question:        question.Question,
			SelectedOptions: opts,
			Answer:          correctOption.Option,
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
		// We log but maybe don't fail the request? Or should we?
		// For MVP, logging is okay, but ideally we return error.
		// Use ErrInternalServerError?
		// "SubmitQuiz" signature returns error. Let's return error.
		return result, score, err
	}

	return result, score, nil
}

// CalculateQuizScore takes a number of questions and a score and returns the percentage of the score.
func (qs *quizService) CalculateQuizScore(ctx context.Context, numOfQuestions int64, score int64) int64 {
	if numOfQuestions == 0 {
		return 0
	}
	return score * 100 / numOfQuestions
}
