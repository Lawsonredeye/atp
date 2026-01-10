package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/lawson/otterprep/pkg"
)

// This handles the database operations for fetching questions and its answers.

type QuestionRepository interface {
	GetQuestionById(ctx context.Context, id int64) (*Questions, error)
	GetCorrectQuestionOptionByQuestionID(ctx context.Context, questionId int64) (*QuestionOptions, error)
	GetRandomQuestion(ctx context.Context, subjectId int64) (*Questions, error)
	CreateQuestion(ctx context.Context, question Questions) (int64, error)
	CreateQuestionOption(ctx context.Context, option QuestionOptions) (int64, error)
	GetQuestionOptions(ctx context.Context, questionId int64) ([]QuestionOptions, error)
	GetQuestionOptionsById(ctx context.Context, id int64) (*QuestionOptions, error)
	CreateAnswer(ctx context.Context, answer Answers) (int64, error)
	GetAnswerById(ctx context.Context, id int64) (*Answers, error)
	UpdateAnswerById(ctx context.Context, answer Answers) (*Answers, error)
	GetAllQuestions(ctx context.Context) ([]Questions, error)
	DeleteQuestionById(ctx context.Context, id int64) error
}

type Questions struct {
	Id               int64     `json:"id"`
	SubjectId        int64     `json:"subject_id"`
	Question         string    `json:"question"`
	IsMultipleChoice bool      `json:"is_multiple_choice"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type Answers struct {
	Id         int64     `json:"id"`
	Answer     string    `json:"answer"`
	QuestionId int64     `json:"question_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type QuestionOptions struct {
	Id         int64     `json:"id"`
	QuestionId int64     `json:"question_id"`
	Option     string    `json:"option"`
	IsCorrect  bool      `json:"is_correct"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type questionRepository struct {
	db *sql.DB
}

func NewQuestionRepository(db *sql.DB) QuestionRepository {
	return &questionRepository{db: db}
}

func (qr *questionRepository) CreateQuestion(ctx context.Context, question Questions) (int64, error) {
	query := "INSERT INTO questions (subject_id, question, is_multiple_choice, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	var id int64
	err := qr.db.QueryRowContext(ctx, query, question.SubjectId, question.Question, question.IsMultipleChoice, question.CreatedAt, question.UpdatedAt).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, pkg.ErrQuestionAlreadyExist
	}
	return id, nil
}

func (qr *questionRepository) GetQuestionById(ctx context.Context, id int64) (*Questions, error) {
	query := "SELECT id, subject_id, question, is_multiple_choice FROM questions WHERE id = $1"
	row := qr.db.QueryRowContext(ctx, query, id)
	var question Questions
	err := row.Scan(&question.Id, &question.SubjectId, &question.Question, &question.IsMultipleChoice)
	if err != nil {
		return nil, err
	}
	return &question, nil
}

func (qr *questionRepository) GetRandomQuestion(ctx context.Context, subjectId int64) (*Questions, error) {
	query := "SELECT id, subject_id, question, is_multiple_choice FROM questions WHERE subject_id = $1 ORDER BY random() LIMIT 1"
	row := qr.db.QueryRowContext(ctx, query, subjectId)
	var question Questions
	err := row.Scan(&question.Id, &question.SubjectId, &question.Question, &question.IsMultipleChoice)
	if err != nil {
		return nil, err
	}
	return &question, nil
}

func (qr *questionRepository) CreateQuestionOption(ctx context.Context, option QuestionOptions) (int64, error) {
	query := "INSERT INTO options (question_id, option, is_correct, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	var id int64
	err := qr.db.QueryRowContext(ctx, query, option.QuestionId, option.Option, option.IsCorrect, option.CreatedAt, option.UpdatedAt).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (qr *questionRepository) GetQuestionOptions(ctx context.Context, questionId int64) ([]QuestionOptions, error) {
	query := "SELECT id, question_id, option, is_correct FROM options WHERE question_id = $1"
	rows, err := qr.db.QueryContext(ctx, query, questionId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var options []QuestionOptions
	for rows.Next() {
		var option QuestionOptions
		err := rows.Scan(&option.Id, &option.QuestionId, &option.Option, &option.IsCorrect)
		if err != nil {
			return nil, err
		}
		options = append(options, option)
	}
	return options, nil
}

// GetCorrectQuestionOptionByQuestionID returns the correct option for a question without returning the entire options with the question id
func (qr *questionRepository) GetCorrectQuestionOptionByQuestionID(ctx context.Context, questionId int64) (*QuestionOptions, error) {
	query := "SELECT id, question_id, option, is_correct FROM options WHERE question_id = $1 AND is_correct = true"
	row := qr.db.QueryRowContext(ctx, query, questionId)
	var option QuestionOptions
	err := row.Scan(&option.Id, &option.QuestionId, &option.Option, &option.IsCorrect)
	if err != nil {
		return nil, err
	}
	return &option, nil
}

func (qr *questionRepository) CreateAnswer(ctx context.Context, answer Answers) (int64, error) {
	query := "INSERT INTO answers (question_id, answer, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING id"
	var id int64
	err := qr.db.QueryRowContext(ctx, query, answer.QuestionId, answer.Answer, answer.CreatedAt, answer.UpdatedAt).Scan(&id)
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
	return id, nil
}

// GetAnswerById returns the answers based on the selected question id.
func (qr *questionRepository) GetAnswerById(ctx context.Context, id int64) (*Answers, error) {
	query := "SELECT id, question_id, answer FROM answers WHERE question_id = $1"
	row := qr.db.QueryRowContext(ctx, query, id)
	var answer Answers
	err := row.Scan(&answer.Id, &answer.QuestionId, &answer.Answer)
	if err != nil {
		return nil, err
	}
	return &answer, nil
}

func (qr *questionRepository) UpdateAnswerById(ctx context.Context, answer Answers) (*Answers, error) {
	fmt.Printf("DEBUG: Updating Answers ID: %d\n", answer.Id)

	var count int
	if err := qr.db.QueryRowContext(ctx, "SELECT count(*) FROM answers WHERE id = $1", answer.Id).Scan(&count); err != nil {
		fmt.Printf("DEBUG: Error checking count: %v\n", err)
	}
	fmt.Printf("DEBUG: Count of ID %d: %d\n", answer.Id, count)

	query := fmt.Sprintf("UPDATE answers SET answer = $1, updated_at = $2 WHERE id = %d", answer.Id)
	res, err := qr.db.ExecContext(ctx, query, answer.Answer, answer.UpdatedAt)
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

	query = "SELECT id, question_id, answer, created_at, updated_at FROM answers WHERE id = $1"
	row := qr.db.QueryRowContext(ctx, query, answer.Id)
	var resp Answers
	err = row.Scan(&resp.Id, &resp.QuestionId, &resp.Answer, &resp.CreatedAt, &resp.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (qr *questionRepository) GetQuestionOptionsById(ctx context.Context, id int64) (*QuestionOptions, error) {
	query := "SELECT id, question_id, option, is_correct FROM options WHERE id = $1"
	row := qr.db.QueryRowContext(ctx, query, id)
	var option QuestionOptions
	err := row.Scan(&option.Id, &option.QuestionId, &option.Option, &option.IsCorrect)
	if err != nil {
		return nil, err
	}
	return &option, nil
}

// GetAllQuestions returns all the questions created on the database.
func (qr *questionRepository) GetAllQuestions(ctx context.Context) ([]Questions, error) {
	query := "SELECT id, subject_id, question FROM questions"
	rows, err := qr.db.QueryContext(ctx, query)
	if err != nil {
		fmt.Println("errors :", err)
		return nil, err
	}
	defer rows.Close()
	var questions []Questions
	for rows.Next() {
		var question Questions
		err := rows.Scan(&question.Id, &question.SubjectId, &question.Question)
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

// DeleteQuestionById deletes a question and its associated options and answers.
func (qr *questionRepository) DeleteQuestionById(ctx context.Context, id int64) error {
	query := "DELETE FROM questions WHERE id = $1"
	_, err := qr.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	_, err = qr.db.ExecContext(ctx, "DELETE FROM options WHERE question_id = $1", id)
	if err != nil {
		return err
	}
	_, err = qr.db.ExecContext(ctx, "DELETE FROM answers WHERE question_id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
