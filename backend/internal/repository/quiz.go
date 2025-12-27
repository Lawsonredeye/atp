package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

type Quiz struct {
	Text             string
	SubjectId        int64
	IsMultipleChoice bool
	QuestionOptions  []QuestionOptions
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type QuizRepository interface {
	CreateQuiz(ctx context.Context, quiz Quiz) (int64, error)
	CreateMultipleQuiz(ctx context.Context, quiz []Quiz) (int64, error)
	GetQuizById(ctx context.Context, id int) (*Quiz, error)
}

type quizRepository struct {
	db *sql.DB
}

func NewQuizRepository(db *sql.DB) *quizRepository {
	return &quizRepository{db: db}
}

func (qr *quizRepository) CreateQuiz(ctx context.Context, quiz Quiz) (int64, error) {
	query := "SELECT id FROM subjects WHERE id = $1"
	if err := qr.db.QueryRowContext(ctx, query, quiz.SubjectId).Scan(&quiz.SubjectId); err != nil {
		return 0, errors.New("subject not found")
	}

	query = "INSERT INTO questions (subject_id, text, created_at, updated_at) VALUES ($1, $2, $3, $4)"
	res, err := qr.db.ExecContext(ctx, query, quiz.SubjectId, quiz.Text, quiz.CreatedAt, quiz.UpdatedAt)
	if err != nil {
		return 0, err
	}
	// loop through the question options and create them
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	for _, option := range quiz.QuestionOptions {
		query = "INSERT INTO question_options (question_id, text, is_correct, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)"
		if _, err := qr.db.ExecContext(ctx, query, id, option.Text, option.IsCorrect, option.CreatedAt, option.UpdatedAt); err != nil {
			return 0, err
		}
	}
	return res.LastInsertId()
}

func (qr *quizRepository) CreateMultipleQuiz(ctx context.Context, quiz []Quiz) (int64, error) {
	for _, q := range quiz {
		qr.CreateQuiz(ctx, q)
	}
	return 0, nil
}

func (qr *quizRepository) GetQuizById(ctx context.Context, id int64) (*Quiz, error) {
	query := "SELECT text, subject_id, created_at, updated_at FROM questions WHERE id = $1"
	row := qr.db.QueryRowContext(ctx, query, id)
	var quiz Quiz
	err := row.Scan(&quiz.Text, &quiz.SubjectId, &quiz.CreatedAt, &quiz.UpdatedAt)
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
