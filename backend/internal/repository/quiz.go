package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type Quiz struct {
	Question         string
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

	query = "INSERT INTO questions (subject_id, question, is_multiple_choice, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)"
	res, err := qr.db.ExecContext(ctx, query, id, quiz.Question, quiz.IsMultipleChoice, quiz.CreatedAt, quiz.UpdatedAt)
	if err != nil {
		return 0, err
	}
	// loop through the question options and create them
	createdId, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	for _, option := range quiz.QuestionOptions {
		query = "INSERT INTO question_options (question_id, option, is_correct, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)"
		if _, err := qr.db.ExecContext(ctx, query, createdId, option.Option, option.IsCorrect, option.CreatedAt, option.UpdatedAt); err != nil {
			return 0, err
		}
		if option.IsCorrect {
			query = "INSERT INTO answers (question_id, answer, created_at, updated_at) VALUES ($1, $2, $3, $4)"
			if _, err := qr.db.ExecContext(ctx, query, createdId, option.Option, option.CreatedAt, option.UpdatedAt); err != nil {
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

// GetQuizById This generates a quick quiz question from db
func (qr *quizRepository) GetQuizById(ctx context.Context, id int64) (*Quiz, error) {
	query := "SELECT question, subject_id, is_multiple_choice, created_at, updated_at FROM questions WHERE id = $1"
	row := qr.db.QueryRowContext(ctx, query, id)
	var quiz Quiz
	err := row.Scan(&quiz.Question, &quiz.SubjectId, &quiz.IsMultipleChoice, &quiz.CreatedAt, &quiz.UpdatedAt)
	if err != nil {
		return nil, err
	}
	query = "SELECT id, question_id, option, is_correct, created_at, updated_at FROM question_options WHERE question_id = $1"
	rows, err := qr.db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var option QuestionOptions
		err := rows.Scan(&option.Id, &option.QuestionId, &option.Option, &option.IsCorrect, &option.CreatedAt, &option.UpdatedAt)
		if err != nil {
			return nil, err
		}
		quiz.QuestionOptions = append(quiz.QuestionOptions, option)
	}
	return &quiz, nil
}
