package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type SubjectRepository interface {
	GetSubjectById(ctx context.Context, id int64) (*Subject, error)
	CreateSubject(ctx context.Context, subject Subject) (int64, error)
	UpdateSubjectById(ctx context.Context, id int64, subject Subject) (*Subject, error)
}

type Subject struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewSubjectRepository(db *sql.DB) *subjectRepository {
	return &subjectRepository{db: db}
}

type subjectRepository struct {
	db *sql.DB
}

func (sr *subjectRepository) GetSubjectById(ctx context.Context, id int64) (*Subject, error) {
	query := "SELECT id, name, created_at, updated_at FROM subjects WHERE id = $1"
	row := sr.db.QueryRowContext(ctx, query, id)
	var subject Subject
	err := row.Scan(&subject.Id, &subject.Name, &subject.CreatedAt, &subject.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &subject, nil
}

func (sr *subjectRepository) CreateSubject(ctx context.Context, subject Subject) (int64, error) {
	query := "INSERT INTO subjects (name, created_at, updated_at) VALUES ($1, $2, $3)"
	result, err := sr.db.ExecContext(ctx, query, subject.Name, subject.CreatedAt, subject.UpdatedAt)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (sr *subjectRepository) UpdateSubjectById(ctx context.Context, id int64, subject Subject) (*Subject, error) {
	query := fmt.Sprintf("UPDATE subjects SET name = '%s', updated_at = '%s' WHERE id = %d", subject.Name, subject.UpdatedAt, id)
	_, err := sr.db.ExecContext(ctx, query)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("subject not found")
		}
		if errors.Is(err, context.Canceled) {
			return nil, errors.New("context canceled")
		}
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, errors.New("context deadline exceeded")
		}
		return nil, err
	}

	resp, err := sr.GetSubjectById(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("subject not found")
		}
		if errors.Is(err, context.Canceled) {
			return nil, errors.New("context canceled")
		}
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, errors.New("context deadline exceeded")
		}
		return nil, err
	}
	return resp, nil
}
