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

func (sr *scoreRepository) StoreUserScore(ctx context.Context, userScore domain.UserScore) (*domain.UserScore, error) {
	query := "INSERT INTO scores (user_id, score, mode, correct_answers, incorrect_answers, total_questions, time_taken_seconds, subject_id, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	result, err := sr.db.ExecContext(ctx, query, userScore.UserID, userScore.Score, userScore.Mode, userScore.CorrectAnswers, userScore.IncorrectAnswers, userScore.TotalQuestions, userScore.TimeTakenSeconds, userScore.SubjectID, userScore.CreatedAt, userScore.UpdatedAt)
	if err != nil {
		return nil, err
	}
	userScore.ID, err = result.LastInsertId()
	return &userScore, nil
}

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

// GetUserOverallScoreStats returns the overall score stats for a user
// It returns the total number of quizzes taken, total correct answers, total incorrect answers, and total questions answered inside a UserStats struct.
func (sr *scoreRepository) GetUserOverallScoreStats(ctx context.Context, userID int64) (*domain.UserStats, error) {
	// "CREATE TABLE IF NOT EXISTS scores (id INTEGER PRIMARY KEY AUTOINCREMENT, user_id BIGINT, score BIGINT, mode VARCHAR(255), correct_answers BIGINT, incorrect_answers BIGINT, total_questions BIGINT, time_taken_seconds BIGINT, subject_id BIGINT, created_at TIMESTAMP, updated_at TIMESTAMP)",

	query := "SELECT user_id, total_questions, correct_answers, incorrect_answers FROM scores WHERE user_id = ?"
	rows, err := sr.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}

	//type UserScore struct {
	//	ID               int64     `json:"id"`
	//	UserID           int64     `json:"user_id"`
	//	SubjectID        int64     `json:"subject_id"`
	//	Score            int64     `json:"score"`
	//	CorrectAnswers   int64     `json:"correct_answers"`
	//	IncorrectAnswers int64     `json:"incorrect_answers"`
	//	TotalQuestions   int64     `json:"total_questions"`
	//	TimeTakenSeconds int64     `json:"time_taken_seconds"`
	//	Mode             string    `json:"mode"`
	//	CreatedAt        time.Time `json:"created_at"`
	//	UpdatedAt        time.Time `json:"updated_at"`
	//}
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
