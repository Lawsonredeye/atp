package repository

import (
	"context"
	"database/sql"

	"github.com/lawson/otterprep/domain"
)

type ScoreRepository interface {
	StoreUserScore(ctx context.Context, userScore domain.UserScore) (*domain.UserScore, error)
	GetUserScoreById(ctx context.Context, id int64) (*domain.UserScore, error)
	GetUserOverallScoreStats(ctx context.Context, userID int64) (*domain.UserStats, error)
}

type scoreRepository struct {
	db *sql.DB
}

func NewScoreRepository(db *sql.DB) ScoreRepository {
	return &scoreRepository{db: db}
}

// StoreUserScore stores a user's score.
func (sr *scoreRepository) StoreUserScore(ctx context.Context, userScore domain.UserScore) (*domain.UserScore, error) {
	query := "INSERT INTO scores (user_id, score, mode, correct_answers, incorrect_answers, total_questions, time_taken_seconds, subject_id, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	result, err := sr.db.ExecContext(ctx, query, userScore.UserID, userScore.Score, userScore.Mode, userScore.CorrectAnswers, userScore.IncorrectAnswers, userScore.TotalQuestions, userScore.TimeTakenSeconds, userScore.SubjectID, userScore.CreatedAt, userScore.UpdatedAt)
	if err != nil {
		return nil, err
	}
	userScore.ID, err = result.LastInsertId()
	return &userScore, nil
}

// GetUserScoreById returns a user's score by id.
func (sr *scoreRepository) GetUserScoreById(ctx context.Context, id int64) (*domain.UserScore, error) {
	query := "SELECT id, user_id, score, mode, correct_answers, incorrect_answers, total_questions, time_taken_seconds, subject_id, created_at, updated_at FROM scores WHERE id = ?"
	row := sr.db.QueryRowContext(ctx, query, id)
	var userScore domain.UserScore
	err := row.Scan(&userScore.ID, &userScore.UserID, &userScore.Score, &userScore.Mode, &userScore.CorrectAnswers, &userScore.IncorrectAnswers, &userScore.TotalQuestions, &userScore.TimeTakenSeconds, &userScore.SubjectID, &userScore.CreatedAt, &userScore.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &userScore, nil
}

// GetUserOverallScoreStats returns the overall score stats for a user.
// It returns the total number of quizzes taken, total correct answers, total incorrect answers, and total questions answered inside a UserStats struct.
func (sr *scoreRepository) GetUserOverallScoreStats(ctx context.Context, userID int64) (*domain.UserStats, error) {
	query := "SELECT user_id, total_questions, correct_answers, incorrect_answers FROM scores WHERE user_id = ?"
	rows, err := sr.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var userStats domain.UserStats
	var (
		totalQuizzesTaken     int64
		totalCorrectAnswers   int64
		totalIncorrectAnswers int64
		totalQuestions        int64
	)
	for rows.Next() {
		err = rows.Scan(&userStats.UserID, &totalQuestions, &totalCorrectAnswers, &totalIncorrectAnswers)
		if err != nil {
			return nil, err
		}
		totalQuizzesTaken++
		userStats.TotalCorrectAnswers += totalCorrectAnswers
		userStats.TotalIncorrectAnswers += totalIncorrectAnswers
		userStats.TotalQuestionsAnswered += totalQuestions
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	userStats.TotalQuizzesTaken = totalQuizzesTaken
	return &userStats, nil
}
