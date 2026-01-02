package repository

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/lawson/otterprep/domain"
	"github.com/stretchr/testify/assert"
)

func TestStoreUserScore(t *testing.T) {
	pool := setUpDB(t)
	ss := NewScoreRepository(pool)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	userScore := domain.UserScore{
		UserID:           int64(1),
		Score:            10,
		Mode:             domain.ModeSingle,
		SubjectID:        int64(1),
		CorrectAnswers:   int64(10),
		IncorrectAnswers: int64(0),
		TotalQuestions:   int64(10),
		TimeTakenSeconds: int64(10),
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	newScore, err := ss.StoreUserScore(ctx, userScore)
	assert.NoError(t, err)
	assert.NotNil(t, newScore)
}

func TestGetUserScoreById(t *testing.T) {
	pool := setUpDB(t)
	ss := NewScoreRepository(pool)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	userNewScore := domain.UserScore{
		UserID:           int64(1),
		Score:            10,
		Mode:             domain.ModeSingle,
		SubjectID:        int64(1),
		CorrectAnswers:   int64(10),
		IncorrectAnswers: int64(0),
		TotalQuestions:   int64(10),
		TimeTakenSeconds: int64(10),
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	userScore, err := ss.StoreUserScore(ctx, userNewScore)
	assert.NoError(t, err)

	userScore, err = ss.GetUserScoreById(ctx, userScore.ID)
	assert.NoError(t, err)
	assert.Equal(t, userScore.Score, int64(10))
}

func TestGetOverallScoreByUserID(t *testing.T) {
	pool := setUpDB(t)
	ss := NewScoreRepository(pool)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	userNewScore := []domain.UserScore{
		{
			UserID:           int64(1),
			Score:            10,
			Mode:             domain.ModeSingle,
			SubjectID:        int64(1),
			CorrectAnswers:   int64(10),
			IncorrectAnswers: int64(0),
			TotalQuestions:   int64(10),
			TimeTakenSeconds: int64(10),
			CreatedAt:        time.Now(),
			UpdatedAt:        time.Now(),
		},
		{
			UserID:           int64(1),
			Score:            10,
			Mode:             domain.ModeSingle,
			SubjectID:        int64(1),
			CorrectAnswers:   int64(10),
			IncorrectAnswers: int64(0),
			TotalQuestions:   int64(10),
			TimeTakenSeconds: int64(10),
			CreatedAt:        time.Now(),
			UpdatedAt:        time.Now(),
		},
	}

	for _, score := range userNewScore {
		userScore, err := ss.StoreUserScore(ctx, score)
		assert.NoError(t, err)
		assert.NotNil(t, userScore)
		fmt.Printf("user inputted score: %+v\n", userScore)
	}

	stats, err := ss.GetUserOverallScoreStats(ctx, userNewScore[0].UserID)
	assert.NoError(t, err)
	assert.Equal(t, stats.TotalCorrectAnswers, int64(20))
	assert.Equal(t, stats.TotalIncorrectAnswers, int64(0))
	assert.Equal(t, stats.TotalQuestionsAnswered, int64(20))
	fmt.Printf("user stats: %+v\n", stats)
}
