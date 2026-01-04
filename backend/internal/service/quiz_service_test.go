package service

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/lawson/otterprep/internal/repository"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

// populateDBWithSubjectID is used to create a fake db populated data
// for testing with.
func populateDBWithSubjectID(subjectId int64) []repository.Quiz {
	return []repository.Quiz{
		{
			Text:             "test",
			SubjectId:        subjectId,
			IsMultipleChoice: true,
			QuestionOptions: []repository.QuestionOptions{
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
			},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Text:             "test",
			SubjectId:        subjectId,
			IsMultipleChoice: true,
			QuestionOptions: []repository.QuestionOptions{
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
			},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Text:             "test",
			SubjectId:        subjectId,
			IsMultipleChoice: true,
			QuestionOptions: []repository.QuestionOptions{
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
			},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Text:             "test",
			SubjectId:        subjectId,
			IsMultipleChoice: true,
			QuestionOptions: []repository.QuestionOptions{
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
			},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
}

func setUpDB(t *testing.T) *sql.DB {
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
		"CREATE TABLE users (id integer primary key autoincrement, name text, email text, password_hash text, created_at timestamp, updated_at timestamp)",
		"CREATE TABLE scores (id integer primary key autoincrement, user_id integer, score integer, mode text, correct_answers integer, incorrect_answers integer, total_questions integer, time_taken_seconds integer, subject_id integer, created_at timestamp, updated_at timestamp)",
		"CREATE TABLE user_roles (id integer primary key autoincrement, user_id integer, role text, created_at timestamp, updated_at timestamp)",
	}
	for _, query := range queries {
		if _, err := db.Exec(query); err != nil {
			t.Fatal(err)
		}
	}
	db.SetMaxOpenConns(1)
	return db
}

func TestGenerateQuizBySubjectID(t *testing.T) {
	pool := setUpDB(t)
	defer pool.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	qr := repository.NewQuizRepository(pool)
	questionRepo := repository.NewQuestionRepository(pool)
	subjectRepo := repository.NewSubjectRepository(pool)
	scoreRepo := repository.NewScoreRepository(pool)
	qs := NewQuizService(qr, subjectRepo, questionRepo, scoreRepo)

	subjectId, err := subjectRepo.CreateSubject(ctx, repository.Subject{
		Name: "use of english",
	})
	data := populateDBWithSubjectID(subjectId)
	assert.Equal(t, int64(1), subjectId)
	if _, err := qr.CreateMultipleQuiz(ctx, data); err != nil {
		t.Fatal("failed to create quiz")
	}

	createdQuiz, err := qr.GetQuizById(ctx, 1)
	if err != nil {
		t.Fatal("failed to get quiz")
	}
	fmt.Println("created quiz: ", createdQuiz)

	questions, err := questionRepo.GetAllQuestions(ctx)
	if err != nil {
		t.Fatal("failed to get questions")
	}
	fmt.Println("all created questions: ", questions)

	quiz, err := qs.GenerateQuizBySubjectID(ctx, 1, 3)
	assert.Nil(t, err)
	//assert.Equal(t, 3, len(quiz))
	fmt.Println("quiz: ", quiz)
}

func TestSubmitQuiz(t *testing.T) {
	pool := setUpDB(t)
	defer pool.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	qr := repository.NewQuizRepository(pool)
	questionRepo := repository.NewQuestionRepository(pool)
	subjectRepo := repository.NewSubjectRepository(pool)
	scoreRepo := repository.NewScoreRepository(pool)
	qs := NewQuizService(qr, subjectRepo, questionRepo, scoreRepo)

	subjectId, err := subjectRepo.CreateSubject(ctx, repository.Subject{
		Name: "use of english",
	})
	data := populateDBWithSubjectID(subjectId)
	assert.Equal(t, int64(1), subjectId)
	if _, err := qr.CreateMultipleQuiz(ctx, data); err != nil {
		t.Fatal("failed to create quiz")
	}

	createdQuiz, err := qr.GetQuizById(ctx, 1)
	if err != nil {
		t.Fatal("failed to get quiz")
	}
	fmt.Println("created quiz: ", createdQuiz)

	questions, err := questionRepo.GetAllQuestions(ctx)
	if err != nil {
		t.Fatal("failed to get questions")
	}
	fmt.Println("all created questions: ", questions)

	quiz, err := qs.GenerateQuizBySubjectID(ctx, 1, 3)
	assert.Nil(t, err)
	//assert.Equal(t, 3, len(quiz))
	fmt.Println("quiz: ", quiz)

	quizRequest := []QuizRequest{
		{
			QuizId:           1,
			IsMultipleChoice: true,
			OptionIds:        []int64{1},
		},
		{
			QuizId:           2,
			IsMultipleChoice: true,
			OptionIds:        []int64{2},
		},
		{
			QuizId:           3,
			IsMultipleChoice: true,
			OptionIds:        []int64{1},
		},
	}
	userID := int64(1)
	result, score, err := qs.SubmitQuiz(ctx, userID, quizRequest)
	assert.Nil(t, err)
	assert.Equal(t, int64(1), score)
	fmt.Println("result: ", result)

	// Verify score is saved
	scoreStats, err := scoreRepo.GetUserOverallScoreStats(ctx, userID)
	assert.Nil(t, err)
	assert.Equal(t, int64(3), scoreStats.TotalQuestionsAnswered)
	assert.Equal(t, int64(1), scoreStats.TotalCorrectAnswers)
	assert.Equal(t, int64(2), scoreStats.TotalIncorrectAnswers)
	fmt.Printf("user score stats: %+v\n", scoreStats)
}

func TestCalculateQuizScore(t *testing.T) {
	pool := setUpDB(t)
	defer pool.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	qr := repository.NewQuizRepository(pool)
	questionRepo := repository.NewQuestionRepository(pool)
	subjectRepo := repository.NewSubjectRepository(pool)
	scoreRepo := repository.NewScoreRepository(pool)
	qs := NewQuizService(qr, subjectRepo, questionRepo, scoreRepo)
	subjectId, err := subjectRepo.CreateSubject(ctx, repository.Subject{
		Name:      "use of english",
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	})
	data := populateDBWithSubjectID(subjectId)
	assert.Equal(t, int64(1), subjectId)
	if _, err := qr.CreateMultipleQuiz(ctx, data); err != nil {
		t.Fatal("failed to create quiz")
	}
	createdQuiz, err := qr.GetQuizById(ctx, 1)
	assert.Nil(t, err)
	fmt.Println("created quiz: ", createdQuiz)
	quizRequest := []QuizRequest{
		{
			QuizId:           1,
			IsMultipleChoice: true,
			OptionIds:        []int64{1},
		},
		{
			QuizId:           2,
			IsMultipleChoice: true,
			OptionIds:        []int64{2},
		},
		{
			QuizId:           3,
			IsMultipleChoice: true,
			OptionIds:        []int64{1},
		},
	}
	userID := int64(1)
	result, score, err := qs.SubmitQuiz(ctx, userID, quizRequest)
	assert.Nil(t, err)
	assert.Equal(t, int64(1), score)
	fmt.Println("result: ", result)

	numOfQuestions := len(quizRequest)
	finalScore := qs.CalculateQuizScore(ctx, int64(numOfQuestions), score)
	assert.Equal(t, int64(33), finalScore)
}
