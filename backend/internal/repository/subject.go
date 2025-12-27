package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

type SubjectRepository interface {
	GetSubjectById(ctx context.Context, id int) (*Subject, error)
	CreateSubject(ctx context.Context, subject Subject) error
}

type Subject struct {
	Id        int       `json:"id"`
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

func (sr *subjectRepository) GetSubjectById(ctx context.Context, id int) (*Subject, error) {
	query := "SELECT id, name FROM subjects WHERE id = $1"
	row := sr.db.QueryRowContext(ctx, query, id)
	var subject Subject
	err := row.Scan(&subject.Id, &subject.Name, &subject.CreatedAt, &subject.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &subject, nil
}

func (sr *subjectRepository) CreateSubject(ctx context.Context, subject Subject) error {
	query := "INSERT INTO subjects (name, created_at, updated_at) VALUES ($1, $2, $3)"
	_, err := sr.db.ExecContext(ctx, query, subject.Name, subject.CreatedAt, subject.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (sr *subjectRepository) UpdateSubjectById(ctx context.Context, id int, subject Subject) (*Subject, error) {
	query := "UPDATE subjects SET name = $2, updated_at = $3 WHERE id = $1 RETURNING id, name"
	row := sr.db.QueryRowContext(ctx, query, id, subject.Name, subject.UpdatedAt)
	if row == nil {
		return nil, errors.New("subject not found")
	}
	var resp Subject
	err := row.Scan(&resp.Id, &resp.Name)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
