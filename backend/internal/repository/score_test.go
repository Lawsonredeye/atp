package repository

import (
	"context"
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
