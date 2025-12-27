package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

// This handles the database operations for fetching questions and its answers.

type QuestionRepository interface {
	GetQuestionById(ctx context.Context, id int) (*Question, error)
	CreateQuestion(ctx context.Context, question Question) error
	CreateQuestionOption(ctx context.Context, option QuestionOptions) error
	GetQuestionOptions(ctx context.Context, questionId int) ([]*QuestionOptions, error)
	CreateAnswer(ctx context.Context, answer Answer) error
	GetAnswerById(ctx context.Context, id int) (*Answer, error)
}

type Question struct {
	Id        int       `json:"id"`
	SubjectId int       `json:"subject_id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Answer struct {
	Id         int       `json:"id"`
	Text       string    `json:"text"`
	QuestionId int       `json:"question_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type QuestionOptions struct {
	Id         int       `json:"id"`
	QuestionId int       `json:"question_id"`
	Text       string    `json:"text"`
	IsCorrect  bool      `json:"is_correct"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type questionRepository struct {
	db *sql.DB
}

func NewQuestionRepository(db *sql.DB) *questionRepository {
	return &questionRepository{db: db}
}

func (qr *questionRepository) CreateQuestion(ctx context.Context, question Question) error {
	query := "INSERT INTO questions (subject_id, text, created_at, updated_at) VALUES ($1, $2, $3, $4)"
	_, err := qr.db.ExecContext(ctx, query, question.SubjectId, question.Text, question.CreatedAt, question.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (qr *questionRepository) GetQuestionById(ctx context.Context, id int) (*Question, error) {
	query := "SELECT id, subject_id, text FROM questions WHERE id = $1"
	row := qr.db.QueryRowContext(ctx, query, id)
	var question Question
	err := row.Scan(&question.Id, &question.SubjectId, &question.Text)
	if err != nil {
		return nil, err
	}
	return &question, nil
}

func (qr *questionRepository) CreateQuestionOption(ctx context.Context, option QuestionOptions) error {
	query := "INSERT INTO question_options (question_id, text, created_at, updated_at) VALUES ($1, $2, $3, $4)"
	_, err := qr.db.ExecContext(ctx, query, option.QuestionId, option.Text, option.CreatedAt, option.UpdatedAt)
	if err != nil {
		return err
	}
	if option.IsCorrect {
		query = "Insert into answers "
	}
	return nil
}

func (qr *questionRepository) GetQuestionOptions(ctx context.Context, questionId int) ([]*QuestionOptions, error) {
	query := "SELECT id, question_id, text FROM question_options WHERE question_id = $1"
	rows, err := qr.db.QueryContext(ctx, query, questionId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var options []*QuestionOptions
	for rows.Next() {
		var option QuestionOptions
		err := rows.Scan(&option.Id, &option.QuestionId, &option.Text)
		if err != nil {
			return nil, err
		}
		options = append(options, &option)
	}
	return options, nil
}

func (qr *questionRepository) CreateAnswer(ctx context.Context, answer Answer) error {
	query := "INSERT INTO answers (question_id, text, created_at, updated_at) VALUES ($1, $2, $3, $4)"
	res, err := qr.db.ExecContext(ctx, query, answer.QuestionId, answer.Text, answer.CreatedAt, answer.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("question not found")
		}
		if errors.Is(err, context.Canceled) {
			return errors.New("context canceled")
		}
		if errors.Is(err, context.DeadlineExceeded) {
			return errors.New("context deadline exceeded")
		}
		return err
	}
	id, _ := res.LastInsertId()
	fmt.Printf("DEBUG: Created Answer ID: %d\n", id)
	return nil
}

func (qr *questionRepository) GetAnswerById(ctx context.Context, id int) (*Answer, error) {
	query := "SELECT id, question_id, text FROM answers WHERE id = $1"
	row := qr.db.QueryRowContext(ctx, query, id)
	var answer Answer
	err := row.Scan(&answer.Id, &answer.QuestionId, &answer.Text)
	if err != nil {
		return nil, err
	}
	return &answer, nil
}

func (qr *questionRepository) UpdateAnswerByID(ctx context.Context, answer Answer) (*Answer, error) {
	fmt.Printf("DEBUG: Updating Answer ID: %d\n", answer.Id)

	var count int
	if err := qr.db.QueryRowContext(ctx, "SELECT count(*) FROM answers WHERE id = $1", answer.Id).Scan(&count); err != nil {
		fmt.Printf("DEBUG: Error checking count: %v\n", err)
	}
	fmt.Printf("DEBUG: Count of ID %d: %d\n", answer.Id, count)

	query := fmt.Sprintf("UPDATE answers SET text = $1, updated_at = $2 WHERE id = %d", answer.Id)
	res, err := qr.db.ExecContext(ctx, query, answer.Text, answer.UpdatedAt)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}
	fmt.Printf("DEBUG: RowsAffected: %d\n", rowsAffected)
	// if rowsAffected == 0 {
	// 	return nil, errors.New("answer not found")
	// }

	query = "SELECT id, question_id, text, created_at, updated_at FROM answers WHERE id = $1"
	row := qr.db.QueryRowContext(ctx, query, answer.Id)
	var resp Answer
	err = row.Scan(&resp.Id, &resp.QuestionId, &resp.Text, &resp.CreatedAt, &resp.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
