package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type Quiz struct {
	Text             string
	SubjectId        int64
	IsMultipleChoice bool `json:"is_multiple_choice"`
	QuestionOptions  []QuestionOptions
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type QuizRequest struct {
	QuizId           int64   `json:"quiz_id"`
	IsMultipleChoice bool    `json:"is_multiple_choice"`
	OptionIds        []int64 `json:"option_ids"`
}

type QuizRepository interface {
	CreateQuiz(ctx context.Context, quiz Quiz) (int64, error)
	CreateMultipleQuiz(ctx context.Context, quiz []Quiz) (int64, error)
	GetQuizById(ctx context.Context, id int64) (*Quiz, error)
	SubmitQuiz(ctx context.Context, quizRequest []QuizRequest) (int64, []QuestionOptions, []Question, error)
}

type quizRepository struct {
	db *sql.DB
}

func NewQuizRepository(db *sql.DB) *quizRepository {
	return &quizRepository{db: db}
}

func (qr *quizRepository) CreateQuiz(ctx context.Context, quiz Quiz) (int64, error) {
	query := "SELECT id FROM subjects WHERE id = $1"
	var id int64
	if err := qr.db.QueryRowContext(ctx, query, quiz.SubjectId).Scan(&id); err != nil {
		return 0, errors.New("subject not found")
	}

	query = "INSERT INTO questions (subject_id, text, is_multiple_choice, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)"
	res, err := qr.db.ExecContext(ctx, query, id, quiz.Text, quiz.IsMultipleChoice, quiz.CreatedAt, quiz.UpdatedAt)
	if err != nil {
		return 0, err
	}
	// loop through the question options and create them
	createdId, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	for _, option := range quiz.QuestionOptions {
		query = "INSERT INTO question_options (question_id, text, is_correct, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)"
		if _, err := qr.db.ExecContext(ctx, query, createdId, option.Text, option.IsCorrect, option.CreatedAt, option.UpdatedAt); err != nil {
			return 0, err
		}
		if option.IsCorrect {
			query = "INSERT INTO answers (question_id, text, created_at, updated_at) VALUES ($1, $2, $3, $4)"
			if _, err := qr.db.ExecContext(ctx, query, createdId, option.Text, option.CreatedAt, option.UpdatedAt); err != nil {
				return 0, err
			}
		}
	}
	return createdId, nil
}

func (qr *quizRepository) CreateMultipleQuiz(ctx context.Context, quiz []Quiz) (int64, error) {
	for _, q := range quiz {
		_, err := qr.CreateQuiz(ctx, q)
		if err != nil {
			fmt.Println("error creating quiz: ", err)
			//return 0, err
		}
	}
	return 0, nil
}

func (qr *quizRepository) GetQuizById(ctx context.Context, id int64) (*Quiz, error) {
	query := "SELECT text, subject_id, is_multiple_choice, created_at, updated_at FROM questions WHERE id = $1"
	row := qr.db.QueryRowContext(ctx, query, id)
	var quiz Quiz
	err := row.Scan(&quiz.Text, &quiz.SubjectId, &quiz.IsMultipleChoice, &quiz.CreatedAt, &quiz.UpdatedAt)
	if err != nil {
		return nil, err
	}
	query = "SELECT id, question_id, text, is_correct, created_at, updated_at FROM question_options WHERE question_id = $1"
	rows, err := qr.db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var option QuestionOptions
		err := rows.Scan(&option.Id, &option.QuestionId, &option.Text, &option.IsCorrect, &option.CreatedAt, &option.UpdatedAt)
		if err != nil {
			return nil, err
		}
		quiz.QuestionOptions = append(quiz.QuestionOptions, option)
	}
	return &quiz, nil
}

// SubmitQuiz takes the request and checks if the users submitted quiz is correct and scores the quiz.
func (qr *quizRepository) SubmitQuiz(ctx context.Context, quizRequest []QuizRequest) (int64, []QuestionOptions, []Question, error) {
	var score int64
	var userOptions []QuestionOptions
	var questions []Question

	for _, req := range quizRequest {
		query := "SELECT id, subject_id, text, is_multiple_choice, created_at, updated_at FROM questions WHERE id = $1"
		var q Question
		err := qr.db.QueryRowContext(ctx, query, req.QuizId).Scan(&q.Id, &q.SubjectId, &q.Text, &q.IsMultipleChoice, &q.CreatedAt, &q.UpdatedAt)
		if err != nil {
			return 0, nil, nil, err
		}
		questions = append(questions, q)

		var currentOptions []QuestionOptions
		for _, optionId := range req.OptionIds {
			query = "SELECT id, question_id, text, is_correct, created_at, updated_at FROM question_options WHERE id = $1"
			var option QuestionOptions
			err := qr.db.QueryRowContext(ctx, query, optionId).Scan(&option.Id, &option.QuestionId, &option.Text, &option.IsCorrect, &option.CreatedAt, &option.UpdatedAt)
			if err != nil {
				return 0, nil, nil, err
			}
			currentOptions = append(currentOptions, option)
		}
		userOptions = append(userOptions, currentOptions...)

		query = "SELECT text FROM answers WHERE question_id = $1"
		rows, err := qr.db.QueryContext(ctx, query, req.QuizId)
		if err != nil {
			return 0, nil, nil, err
		}

		correctAnswers := make(map[string]bool)
		for rows.Next() {
			var text string
			if err := rows.Scan(&text); err != nil {
				rows.Close()
				return 0, nil, nil, err
			}
			correctAnswers[text] = true
		}
		rows.Close()

		if err = rows.Err(); err != nil {
			return 0, nil, nil, err
		}

		isCorrect := true
		if len(currentOptions) != len(correctAnswers) {
			isCorrect = false
		} else {
			for _, opt := range currentOptions {
				if !correctAnswers[opt.Text] {
					isCorrect = false
					break
				}
			}
		}

		if isCorrect {
			score++
		}
	}
	return score, userOptions, questions, nil
}
