package repository

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateQuiz(t *testing.T) {
	pool := setUP(t)
	repo := NewQuizRepository(pool)
	subjectRepo := NewSubjectRepository(pool)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	subjectId, err := subjectRepo.CreateSubject(ctx, Subject{
		Name:      "test",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	assert.Nil(t, err)
	assert.Equal(t, subjectId, int64(1))

	quizId, err := repo.CreateQuiz(ctx, Quiz{
		Text:             "test",
		SubjectId:        subjectId,
		IsMultipleChoice: true,
		QuestionOptions: []QuestionOptions{
			{
				Text:      "test",
				IsCorrect: true,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			{
				Text:      "test",
				IsCorrect: false,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			{
				Text:      "test",
				IsCorrect: false,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			{
				Text:      "test",
				IsCorrect: false,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	assert.Nil(t, err)
	assert.Equal(t, quizId, int64(1))
}

func TestCreateMultipleQuiz(t *testing.T) {
	pool := setUP(t)
	repo := NewQuizRepository(pool)
	subjectRepo := NewSubjectRepository(pool)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	subjectId, err := subjectRepo.CreateSubject(ctx, Subject{
		Name:      "test",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	assert.Nil(t, err)
	assert.Equal(t, subjectId, int64(1))

	quiz := []Quiz{
		{
			Text:             "test",
			SubjectId:        subjectId,
			IsMultipleChoice: true,
			QuestionOptions: []QuestionOptions{
				{
					Text:      "test",
					IsCorrect: true,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
				{
					Text:      "test",
					IsCorrect: false,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
				{
					Text:      "test",
					IsCorrect: false,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
				{
					Text:      "test",
					IsCorrect: false,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
			},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Text:             "test",
			SubjectId:        subjectId,
			IsMultipleChoice: true,
			QuestionOptions: []QuestionOptions{
				{
					Text:      "test",
					IsCorrect: true,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
				{
					Text:      "test",
					IsCorrect: false,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
				{
					Text:      "test",
					IsCorrect: false,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
				{
					Text:      "test",
					IsCorrect: false,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
			},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	quizId, err := repo.CreateMultipleQuiz(ctx, quiz)
	assert.Nil(t, err)
	assert.Equal(t, quizId, int64(0))
}

func TestGetQuizById(t *testing.T) {
	pool := setUP(t)
	repo := NewQuizRepository(pool)
	subjectRepo := NewSubjectRepository(pool)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	subjectId, err := subjectRepo.CreateSubject(ctx, Subject{
		Name:      "test",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	assert.Nil(t, err)
	assert.Equal(t, subjectId, int64(1))
	quizId, err := repo.CreateQuiz(ctx, Quiz{
		Text:             "test",
		SubjectId:        subjectId,
		IsMultipleChoice: true,
		QuestionOptions: []QuestionOptions{
			{
				Text:      "test is the right answer",
				IsCorrect: true,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			{
				Text:      "test is not the right answer",
				IsCorrect: false,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			{
				Text:      "test is not the right answer",
				IsCorrect: false,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			{
				Text:      "test",
				IsCorrect: false,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	assert.Nil(t, err)
	assert.Equal(t, quizId, int64(1))
	quiz, err := repo.GetQuizById(ctx, quizId)
	assert.Nil(t, err)
	assert.Equal(t, quiz.Text, "test")
	fmt.Printf("fetched quiz: %+v\n", quiz)
}

// func TestGenerateQuizBySubjectID(t *testing.T) {
// 	pool := setUP(t)
// 	repo := NewQuizRepository(pool)
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()
// 	subjectRepo := NewSubjectRepository(pool)
// 	subjectId, err := subjectRepo.CreateSubject(ctx, Subject{
// 		Name:      "test",
// 		CreatedAt: time.Now(),
// 		UpdatedAt: time.Now(),
// 	})
// 	assert.Nil(t, err)
// 	assert.Equal(t, subjectId, int64(1))
// 	quizTest, err := repo.GenerateQuizBySubjectID(ctx, subjectId)
// 	assert.Nil(t, err)
// 	assert.NotNil(t, quizTest)
// 	fmt.Printf("fetched quiz: %+v\n", quizTest)
// }
