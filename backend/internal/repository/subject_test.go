package repository

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateSubject(t *testing.T) {
	pool := setUP(t)
	repo := NewSubjectRepository(pool)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	id, err := repo.CreateSubject(ctx, Subject{
		Name:      "test",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	assert.Nil(t, err)
	assert.Equal(t, id, int64(1))
}

func TestGetSubjectById(t *testing.T) {
	pool := setUP(t)
	repo := NewSubjectRepository(pool)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// create the subject first and attempt to get it
	id, err := repo.CreateSubject(ctx, Subject{
		Name:      "test",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	assert.Nil(t, err)
	subject, err := repo.GetSubjectById(ctx, id)
	assert.Nil(t, err)
	assert.Equal(t, subject.Name, "test")
}

func TestUpdateSubjectById(t *testing.T) {
	pool := setUP(t)
	repo := NewSubjectRepository(pool)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// create the subject first and attempt to update it
	id, err := repo.CreateSubject(ctx, Subject{
		Name:      "test",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	assert.Nil(t, err)
	subject, err := repo.GetSubjectById(ctx, id)
	assert.Nil(t, err)
	assert.Equal(t, subject.Name, "test")

	subject.Name = "updated"
	subject.UpdatedAt = time.Now()
	updatedSubject, err := repo.UpdateSubjectById(ctx, id, *subject)
	assert.Nil(t, err)
	assert.Equal(t, updatedSubject.Name, "updated")
}
