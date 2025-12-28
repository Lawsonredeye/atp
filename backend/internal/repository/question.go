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
	GetQuestionById(ctx context.Context, id int64) (*Question, error)
	GetRandomQuestion(ctx context.Context, subjectId int64) (*Question, error)
	CreateQuestion(ctx context.Context, question Question) (int64, error)
	CreateQuestionOption(ctx context.Context, option QuestionOptions) (int64, error)
	GetQuestionOptions(ctx context.Context, questionId int64) ([]QuestionOptions, error)
	CreateAnswer(ctx context.Context, answer Answer) (int64, error)
	GetAnswerById(ctx context.Context, id int64) (*Answer, error)
}

type Question struct {
	Id               int       `json:"id"`
	SubjectId        int       `json:"subject_id"`
	Text             string    `json:"text"`
	IsMultipleChoice bool      `json:"is_multiple_choice"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
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

func (qr *questionRepository) CreateQuestion(ctx context.Context, question Question) (int64, error) {
	query := "INSERT INTO questions (subject_id, text, is_multiple_choice, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)"
	res, err := qr.db.ExecContext(ctx, query, question.SubjectId, question.Text, question.IsMultipleChoice, question.CreatedAt, question.UpdatedAt)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func (qr *questionRepository) GetQuestionById(ctx context.Context, id int64) (*Question, error) {
	query := "SELECT id, subject_id, text, is_multiple_choice FROM questions WHERE id = $1"
	row := qr.db.QueryRowContext(ctx, query, id)
	var question Question
	err := row.Scan(&question.Id, &question.SubjectId, &question.Text, &question.IsMultipleChoice)
	if err != nil {
		return nil, err
	}
	return &question, nil
}

func (qr *questionRepository) GetRandomQuestion(ctx context.Context, subjectId int64) (*Question, error) {
	query := "SELECT id, subject_id, text, is_multiple_choice FROM questions WHERE subject_id = $1 ORDER BY random() LIMIT 1"
	row := qr.db.QueryRowContext(ctx, query, subjectId)
	var question Question
	err := row.Scan(&question.Id, &question.SubjectId, &question.Text, &question.IsMultipleChoice)
	if err != nil {
		return nil, err
	}
	return &question, nil
}

func (qr *questionRepository) CreateQuestionOption(ctx context.Context, option QuestionOptions) (int64, error) {
	query := "INSERT INTO question_options (question_id, text, created_at, updated_at) VALUES ($1, $2, $3, $4)"
	res, err := qr.db.ExecContext(ctx, query, option.QuestionId, option.Text, option.CreatedAt, option.UpdatedAt)
	if err != nil {
		return 0, err
	}
	if option.IsCorrect {
		qr.CreateAnswer(ctx, Answer{QuestionId: option.QuestionId,
			Text:      option.Text,
			CreatedAt: option.CreatedAt,
			UpdatedAt: option.UpdatedAt})
	}
	return res.LastInsertId()
}

func (qr *questionRepository) GetQuestionOptions(ctx context.Context, questionId int64) ([]QuestionOptions, error) {
	query := "SELECT id, question_id, text FROM question_options WHERE question_id = $1"
	rows, err := qr.db.QueryContext(ctx, query, questionId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var options []QuestionOptions
	for rows.Next() {
		var option QuestionOptions
		err := rows.Scan(&option.Id, &option.QuestionId, &option.Text)
		if err != nil {
			return nil, err
		}
		options = append(options, option)
	}
	return options, nil
}

func (qr *questionRepository) CreateAnswer(ctx context.Context, answer Answer) (int64, error) {
	query := "INSERT INTO answers (question_id, text, created_at, updated_at) VALUES ($1, $2, $3, $4)"
	res, err := qr.db.ExecContext(ctx, query, answer.QuestionId, answer.Text, answer.CreatedAt, answer.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, errors.New("question not found")
		}
		if errors.Is(err, context.Canceled) {
			return 0, errors.New("context canceled")
		}
		if errors.Is(err, context.DeadlineExceeded) {
			return 0, errors.New("context deadline exceeded")
		}
		return 0, err
	}
	return res.LastInsertId()
}

func (qr *questionRepository) GetAnswerById(ctx context.Context, id int64) (*Answer, error) {
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

func (qr *questionRepository) GetAllQuestions(ctx context.Context) ([]Question, error) {
	query := "SELECT id, subject_id, text FROM questions"
	rows, err := qr.db.QueryContext(ctx, query)
	if err != nil {
		fmt.Println("errors :", err)
		return nil, err
	}
	defer rows.Close()
	var questions []Question
	for rows.Next() {
		var question Question
		err := rows.Scan(&question.Id, &question.SubjectId, &question.Text)
		if err != nil {
			fmt.Println("error storing values: ", err)
			return nil, err
		}
		questions = append(questions, question)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	fmt.Println("all questions:", len(questions))
	return questions, nil
}
