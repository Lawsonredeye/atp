package repository

import (
	"context"
	"database/sql"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func setUP(t *testing.T) *sql.DB {
	t.Helper()
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	queries := []string{
		"CREATE TABLE question_options (id integer primary key autoincrement, question_id integer, text text, is_correct boolean, created_at timestamp, updated_at timestamp)",
		"CREATE TABLE questions (id integer primary key autoincrement, subject_id integer, text text, is_multiple_choice boolean, created_at timestamp, updated_at timestamp)",
		"CREATE TABLE answers (id integer primary key autoincrement, question_id integer, text text, created_at timestamp, updated_at timestamp)",
		"CREATE TABLE subjects (id integer primary key autoincrement, name text, created_at timestamp, updated_at timestamp)",
	}
	for _, query := range queries {
		if _, err := db.Exec(query); err != nil {
			t.Fatal(err)
		}
	}
	db.SetMaxOpenConns(1)
	return db
}

func TestGetQuestionById(t *testing.T) {
	pool := setUP(t)
	repo := NewQuestionRepository(pool)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	createdQuestionId, err := repo.CreateQuestion(ctx, Question{
		SubjectId:        1,
		Text:             "test",
		IsMultipleChoice: false,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	})
	assert.Nil(t, err)
	assert.Equal(t, createdQuestionId, int64(1))
}

func TestGetRandomQuestion(t *testing.T) {
	pool := setUP(t)
	repo := NewQuestionRepository(pool)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	createdQuestionId, err := repo.CreateQuestion(ctx, Question{
		SubjectId:        1,
		Text:             "test",
		IsMultipleChoice: false,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	})
	assert.Nil(t, err)
	assert.Equal(t, createdQuestionId, int64(1))
	createdQuestion, err := repo.GetRandomQuestion(ctx, 1)
	assert.Nil(t, err)
	assert.Equal(t, createdQuestion.Id, createdQuestionId)
}

func TestCreateQuestion(t *testing.T) {
	pool := setUP(t)
	repo := NewQuestionRepository(pool)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	id, err := repo.CreateQuestion(ctx, Question{
		SubjectId:        1,
		Text:             "test",
		IsMultipleChoice: false,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	})
	assert.Nil(t, err)
	assert.Equal(t, id, int64(1))
}

func TestCreateQuestionOption(t *testing.T) {
	pool := setUP(t)
	repo := NewQuestionRepository(pool)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	createdOptionId, err := repo.CreateQuestionOption(ctx, QuestionOptions{
		QuestionId: 1,
		Text:       "test",
		IsCorrect:  true,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	})
	assert.Nil(t, err)
	assert.Equal(t, createdOptionId, int64(1))
}

func TestGetQuestionOptions(t *testing.T) {
	pool := setUP(t)
	repo := NewQuestionRepository(pool)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := repo.CreateQuestionOption(ctx, QuestionOptions{
		QuestionId: 1,
		Text:       "test",
		IsCorrect:  true,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	})
	assert.Nil(t, err)
	createdQuestion, err := repo.GetQuestionOptions(ctx, 1)
	assert.Nil(t, err)
	assert.NotNil(t, createdQuestion, "should have created question")
}

func TestCreateAnswer(t *testing.T) {
	pool := setUP(t)
	repo := NewQuestionRepository(pool)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	createdOptionId, err := repo.CreateAnswer(ctx, Answer{
		QuestionId: 1,
		Text:       "test",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	})
	assert.Nil(t, err)
	assert.Equal(t, createdOptionId, int64(1))
}

func TestGetAnswerById(t *testing.T) {
	pool := setUP(t)
	repo := NewQuestionRepository(pool)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	createdOptionId, err := repo.CreateAnswer(ctx, Answer{
		QuestionId: 1,
		Text:       "test",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	})
	assert.Nil(t, err)
	assert.Equal(t, createdOptionId, int64(1))

	createdAnswer, err := repo.GetAnswerById(ctx, 1)
	assert.Nil(t, err)
	assert.NotNil(t, createdAnswer, "should have created answer")
}

func TestUpdateAnswerByID(t *testing.T) {
	pool := setUP(t)
	repo := NewQuestionRepository(pool)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	createdOptionId, err := repo.CreateAnswer(ctx, Answer{
		QuestionId: 1,
		Text:       "test",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	})
	assert.Nil(t, err)
	assert.Equal(t, createdOptionId, int64(1))

	createdAnswer, err := repo.GetAnswerById(ctx, 1)
	assert.Nil(t, err)
	assert.NotNil(t, createdAnswer, "should have created answer")

	updatedAnswer, err := repo.UpdateAnswerByID(ctx, Answer{
		Id:         1,
		QuestionId: 1,
		Text:       "test updated",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	})
	assert.Nil(t, err)
	assert.NotNil(t, updatedAnswer, "should have updated answer")
}
