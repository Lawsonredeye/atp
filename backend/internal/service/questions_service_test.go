package service

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/lawson/otterprep/domain"
	"github.com/lawson/otterprep/internal/repository"
	"github.com/stretchr/testify/assert"
)

func TestUserServiceCreateMultipleQuestionBySubjectID(t *testing.T) {
	pool := setUpDB(t)
	defer pool.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	questionRepository := repository.NewQuestionRepository(pool)
	subjectRepository := repository.NewSubjectRepository(pool)
	logger := log.New(os.Stdout, "questionService: ", log.LstdFlags)
	questionService := NewQuestionService(questionRepository, subjectRepository, logger)

	questions := []domain.QuestionsData{
		{
			Name:        "What is the capital of France?",
			Options:     []string{"Paris", "London", "Berlin", "Madrid"},
			Answer:      "Paris",
			Explanation: "Paris is the capital of France.",
		},
		{
			Name:        "What is the capital of Germany?",
			Options:     []string{"Paris", "London", "Berlin", "Madrid"},
			Answer:      "Berlin",
			Explanation: "Berlin is the capital of Germany.",
		},
	}
	subjectId, err := subjectRepository.CreateSubject(ctx, repository.Subject{
		Name: "General Knowledge",
	})
	assert.Nil(t, err)
	assert.Equal(t, int64(1), subjectId)

	err = questionService.CreateMultipleQuestionBySubjectID(ctx, subjectId, questions)
	assert.Nil(t, err)

	question, err := questionService.GetQuestionById(ctx, 1)
	assert.Nil(t, err)
	assert.Equal(t, question.Text, "What is the capital of France?")
	fmt.Printf("question data: %+v\n", question)
}

func TestCreateSingleQuestion(t *testing.T) {
	pool := setUpDB(t)
	defer pool.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	questionRepository := repository.NewQuestionRepository(pool)
	subjectRepository := repository.NewSubjectRepository(pool)
	logger := log.New(os.Stdout, "questionService: ", log.LstdFlags)
	questionService := NewQuestionService(questionRepository, subjectRepository, logger)

	subjectId, err := subjectRepository.CreateSubject(ctx, repository.Subject{
		Name: "General Knowledge",
	})
	assert.Nil(t, err)
	assert.Equal(t, int64(1), subjectId)

	newQuestion := domain.QuestionsData{
		Name:        "What is the capital of France?",
		Options:     []string{"Paris", "London", "Berlin", "Madrid"},
		Answer:      "Paris",
		Explanation: "Paris is the capital of France.",
	}

	id, err := questionService.CreateQuestion(ctx, subjectId, newQuestion)
	assert.Nil(t, err)
	assert.Equal(t, int64(1), id)

	question, err := questionService.GetQuestionById(ctx, id)
	assert.Nil(t, err)
	assert.Equal(t, question.Text, "What is the capital of France?")
	fmt.Printf("question data: %+v\n", question)
}

func TestDeleteQuestionById(t *testing.T) {
	pool := setUpDB(t)
	defer pool.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	questionRepository := repository.NewQuestionRepository(pool)
	subjectRepository := repository.NewSubjectRepository(pool)
	logger := log.New(os.Stdout, "questionService: ", log.LstdFlags)
	questionService := NewQuestionService(questionRepository, subjectRepository, logger)

	subjectId, err := subjectRepository.CreateSubject(ctx, repository.Subject{
		Name: "General Knowledge",
	})
	assert.Nil(t, err)
	assert.Equal(t, int64(1), subjectId)

	newQuestion := domain.QuestionsData{
		Name:        "What is the capital of France?",
		Options:     []string{"Paris", "London", "Berlin", "Madrid"},
		Answer:      "Paris",
		Explanation: "Paris is the capital of France.",
	}

	id, err := questionService.CreateQuestion(ctx, subjectId, newQuestion)
	assert.Nil(t, err)
	assert.Equal(t, int64(1), id)

	err = questionService.DeleteQuestionById(ctx, id)
	assert.Nil(t, err)

	question, err := questionService.GetQuestionById(ctx, id)
	assert.NotNil(t, err)
	assert.Nil(t, question)
	fmt.Printf("result [%+v] :: %+v\n", question, err)
}

func TestGetSubjectByName(t *testing.T) {
	pool := setUpDB(t)
	defer pool.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	subjectRepository := repository.NewSubjectRepository(pool)
	questionRepository := repository.NewQuestionRepository(pool)
	logger := log.New(os.Stdout, "questionService: ", log.LstdFlags)
	questionService := NewQuestionService(questionRepository, subjectRepository, logger)

	subjectId, err := subjectRepository.CreateSubject(ctx, repository.Subject{
		Name: "General Knowledge",
	})
	assert.Nil(t, err)
	assert.Equal(t, int64(1), subjectId)

	subject, err := questionService.GetSubjectByName(ctx, "General Knowledge")
	assert.Nil(t, err)
	assert.Equal(t, subjectId, subject.Id)
	fmt.Printf("subject data: %+v\n", subject)
}

func TestCreateSubjectWithSameName(t *testing.T) {
	pool := setUpDB(t)
	defer pool.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	subjectRepository := repository.NewSubjectRepository(pool)
	questionRepository := repository.NewQuestionRepository(pool)
	logger := log.New(os.Stdout, "questionService: ", log.LstdFlags)
	questionService := NewQuestionService(questionRepository, subjectRepository, logger)

	firstSubject, err := questionService.CreateSubject(ctx, "General Knowledge")
	assert.Nil(t, err)
	assert.Equal(t, int64(1), firstSubject)

	secondSubject, err := questionService.CreateSubject(ctx, "General Knowledge")
	assert.NotNil(t, err)
	assert.Equal(t, firstSubject, secondSubject)

	subject, err := questionService.GetSubjectByName(ctx, "General Knowledge")
	assert.Nil(t, err)
	assert.Equal(t, int64(1), subject.Id)
	fmt.Printf("subject data: %+v\n", subject)
}

func TestGetSubjects(t *testing.T) {
	pool := setUpDB(t)
	defer pool.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	subjectRepository := repository.NewSubjectRepository(pool)
	questionRepository := repository.NewQuestionRepository(pool)
	logger := log.New(os.Stdout, "questionService: ", log.LstdFlags)
	questionService := NewQuestionService(questionRepository, subjectRepository, logger)

	subjectNames := []string{
		"General Knowledge",
		"Mathematics",
	}

	for _, subjectName := range subjectNames {
		_, err := questionService.CreateSubject(ctx, subjectName)
		assert.Nil(t, err)
	}

	subjects, err := questionService.GetSubjects(ctx)
	assert.Nil(t, err)
	assert.Equal(t, int64(1), subjects[0].Id)
	fmt.Printf("subjects data: %+v\n", subjects)
}

func TestGetQuestionById(t *testing.T) {
	pool := setUpDB(t)
	defer pool.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	questionRepository := repository.NewQuestionRepository(pool)
	subjectRepository := repository.NewSubjectRepository(pool)
	logger := log.New(os.Stdout, "questionService: ", log.LstdFlags)
	questionService := NewQuestionService(questionRepository, subjectRepository, logger)

	subjectId, err := subjectRepository.CreateSubject(ctx, repository.Subject{
		Name: "General Knowledge",
	})
	assert.Nil(t, err)
	assert.Equal(t, int64(1), subjectId)

	newQuestion := domain.QuestionsData{
		Name:        "What is the capital of France?",
		Options:     []string{"Paris", "London", "Berlin", "Madrid"},
		Answer:      "Paris",
		Explanation: "Paris is the capital of France.",
	}

	id, err := questionService.CreateQuestion(ctx, subjectId, newQuestion)
	assert.Nil(t, err)
	assert.Equal(t, int64(1), id)

	question, err := questionService.GetQuestionById(ctx, id)
	assert.Nil(t, err)
	assert.Equal(t, question.Text, "What is the capital of France?")
	assert.Equal(t, question.Option, []string{"Paris", "London", "Berlin", "Madrid"})
	assert.Equal(t, question.Answer, "")
	assert.Equal(t, question.Explanation, "")
	fmt.Printf("question data: %+v\n", question)
}
