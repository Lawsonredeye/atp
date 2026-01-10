package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/lawson/otterprep/domain"
	"github.com/lawson/otterprep/pkg"
)

type LeaderboardRepository interface {
	GetGlobalLeaderboard(ctx context.Context, limit, offset int) ([]domain.LeaderboardEntry, int64, error)
	GetSubjectLeaderboard(ctx context.Context, subjectId int64, limit, offset int) ([]domain.LeaderboardEntry, int64, error)
	GetWeeklyLeaderboard(ctx context.Context, limit, offset int) ([]domain.LeaderboardEntry, int64, error)
	GetMonthlyLeaderboard(ctx context.Context, limit, offset int) ([]domain.LeaderboardEntry, int64, error)
	GetUserRank(ctx context.Context, userId int64) (*domain.UserRankResponse, error)
	GetUserSubjectRank(ctx context.Context, userId, subjectId int64) (*domain.UserRankResponse, error)
}

type leaderboardRepository struct {
	db *sql.DB
}

func NewLeaderboardRepository(db *sql.DB) LeaderboardRepository {
	return &leaderboardRepository{db: db}
}

// GetGlobalLeaderboard returns the global leaderboard (all time, all subjects)
func (lr *leaderboardRepository) GetGlobalLeaderboard(ctx context.Context, limit, offset int) ([]domain.LeaderboardEntry, int64, error) {
	// Get total count of users with scores
	var totalUsers int64
	countQuery := `SELECT COUNT(DISTINCT user_id) FROM scores`
	if err := lr.db.QueryRowContext(ctx, countQuery).Scan(&totalUsers); err != nil {
		return nil, 0, err
	}

	query := `
		SELECT 
			u.id as user_id,
			u.name as user_name,
			COALESCE(SUM(s.score), 0) as total_score,
			COUNT(s.id) as total_quizzes,
			COALESCE(SUM(s.correct_answers), 0) as correct_answers,
			COALESCE(SUM(s.total_questions), 0) as total_questions
		FROM users u
		INNER JOIN scores s ON u.id = s.user_id
		GROUP BY u.id, u.name
		ORDER BY total_score DESC, correct_answers DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := lr.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var entries []domain.LeaderboardEntry
	rank := int64(offset + 1)

	for rows.Next() {
		var entry domain.LeaderboardEntry
		err := rows.Scan(
			&entry.UserID,
			&entry.UserName,
			&entry.TotalScore,
			&entry.TotalQuizzes,
			&entry.CorrectAnswers,
			&entry.TotalQuestions,
		)
		if err != nil {
			return nil, 0, err
		}
		entry.Rank = rank
		if entry.TotalQuestions > 0 {
			entry.AccuracyPercent = float64(entry.CorrectAnswers) / float64(entry.TotalQuestions) * 100
		}
		entries = append(entries, entry)
		rank++
	}

	if err = rows.Err(); err != nil {
		return nil, 0, err
	}

	return entries, totalUsers, nil
}

// GetSubjectLeaderboard returns the leaderboard for a specific subject
func (lr *leaderboardRepository) GetSubjectLeaderboard(ctx context.Context, subjectId int64, limit, offset int) ([]domain.LeaderboardEntry, int64, error) {
	var totalUsers int64
	countQuery := `SELECT COUNT(DISTINCT user_id) FROM scores WHERE subject_id = $1`
	if err := lr.db.QueryRowContext(ctx, countQuery, subjectId).Scan(&totalUsers); err != nil {
		return nil, 0, err
	}

	query := `
		SELECT 
			u.id as user_id,
			u.name as user_name,
			COALESCE(SUM(s.score), 0) as total_score,
			COUNT(s.id) as total_quizzes,
			COALESCE(SUM(s.correct_answers), 0) as correct_answers,
			COALESCE(SUM(s.total_questions), 0) as total_questions
		FROM users u
		INNER JOIN scores s ON u.id = s.user_id
		WHERE s.subject_id = $1
		GROUP BY u.id, u.name
		ORDER BY total_score DESC, correct_answers DESC
		LIMIT $2 OFFSET $3
	`

	rows, err := lr.db.QueryContext(ctx, query, subjectId, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var entries []domain.LeaderboardEntry
	rank := int64(offset + 1)

	for rows.Next() {
		var entry domain.LeaderboardEntry
		err := rows.Scan(
			&entry.UserID,
			&entry.UserName,
			&entry.TotalScore,
			&entry.TotalQuizzes,
			&entry.CorrectAnswers,
			&entry.TotalQuestions,
		)
		if err != nil {
			return nil, 0, err
		}
		entry.Rank = rank
		if entry.TotalQuestions > 0 {
			entry.AccuracyPercent = float64(entry.CorrectAnswers) / float64(entry.TotalQuestions) * 100
		}
		entries = append(entries, entry)
		rank++
	}

	if err = rows.Err(); err != nil {
		return nil, 0, err
	}

	return entries, totalUsers, nil
}

// GetWeeklyLeaderboard returns the leaderboard for the current week
func (lr *leaderboardRepository) GetWeeklyLeaderboard(ctx context.Context, limit, offset int) ([]domain.LeaderboardEntry, int64, error) {
	startOfWeek := time.Now().AddDate(0, 0, -7)

	var totalUsers int64
	countQuery := `SELECT COUNT(DISTINCT user_id) FROM scores WHERE created_at >= $1`
	if err := lr.db.QueryRowContext(ctx, countQuery, startOfWeek).Scan(&totalUsers); err != nil {
		return nil, 0, err
	}

	query := `
		SELECT 
			u.id as user_id,
			u.name as user_name,
			COALESCE(SUM(s.score), 0) as total_score,
			COUNT(s.id) as total_quizzes,
			COALESCE(SUM(s.correct_answers), 0) as correct_answers,
			COALESCE(SUM(s.total_questions), 0) as total_questions
		FROM users u
		INNER JOIN scores s ON u.id = s.user_id
		WHERE s.created_at >= $1
		GROUP BY u.id, u.name
		ORDER BY total_score DESC, correct_answers DESC
		LIMIT $2 OFFSET $3
	`

	rows, err := lr.db.QueryContext(ctx, query, startOfWeek, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var entries []domain.LeaderboardEntry
	rank := int64(offset + 1)

	for rows.Next() {
		var entry domain.LeaderboardEntry
		err := rows.Scan(
			&entry.UserID,
			&entry.UserName,
			&entry.TotalScore,
			&entry.TotalQuizzes,
			&entry.CorrectAnswers,
			&entry.TotalQuestions,
		)
		if err != nil {
			return nil, 0, err
		}
		entry.Rank = rank
		if entry.TotalQuestions > 0 {
			entry.AccuracyPercent = float64(entry.CorrectAnswers) / float64(entry.TotalQuestions) * 100
		}
		entries = append(entries, entry)
		rank++
	}

	if err = rows.Err(); err != nil {
		return nil, 0, err
	}

	return entries, totalUsers, nil
}

// GetMonthlyLeaderboard returns the leaderboard for the current month
func (lr *leaderboardRepository) GetMonthlyLeaderboard(ctx context.Context, limit, offset int) ([]domain.LeaderboardEntry, int64, error) {
	startOfMonth := time.Now().AddDate(0, -1, 0)

	var totalUsers int64
	countQuery := `SELECT COUNT(DISTINCT user_id) FROM scores WHERE created_at >= $1`
	if err := lr.db.QueryRowContext(ctx, countQuery, startOfMonth).Scan(&totalUsers); err != nil {
		return nil, 0, err
	}

	query := `
		SELECT 
			u.id as user_id,
			u.name as user_name,
			COALESCE(SUM(s.score), 0) as total_score,
			COUNT(s.id) as total_quizzes,
			COALESCE(SUM(s.correct_answers), 0) as correct_answers,
			COALESCE(SUM(s.total_questions), 0) as total_questions
		FROM users u
		INNER JOIN scores s ON u.id = s.user_id
		WHERE s.created_at >= $1
		GROUP BY u.id, u.name
		ORDER BY total_score DESC, correct_answers DESC
		LIMIT $2 OFFSET $3
	`

	rows, err := lr.db.QueryContext(ctx, query, startOfMonth, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var entries []domain.LeaderboardEntry
	rank := int64(offset + 1)

	for rows.Next() {
		var entry domain.LeaderboardEntry
		err := rows.Scan(
			&entry.UserID,
			&entry.UserName,
			&entry.TotalScore,
			&entry.TotalQuizzes,
			&entry.CorrectAnswers,
			&entry.TotalQuestions,
		)
		if err != nil {
			return nil, 0, err
		}
		entry.Rank = rank
		if entry.TotalQuestions > 0 {
			entry.AccuracyPercent = float64(entry.CorrectAnswers) / float64(entry.TotalQuestions) * 100
		}
		entries = append(entries, entry)
		rank++
	}

	if err = rows.Err(); err != nil {
		return nil, 0, err
	}

	return entries, totalUsers, nil
}

// GetUserRank returns the user's rank on the global leaderboard
func (lr *leaderboardRepository) GetUserRank(ctx context.Context, userId int64) (*domain.UserRankResponse, error) {
	query := `
		WITH ranked_users AS (
			SELECT 
				u.id as user_id,
				u.name as user_name,
				COALESCE(SUM(s.score), 0) as total_score,
				COUNT(s.id) as total_quizzes,
				COALESCE(SUM(s.correct_answers), 0) as correct_answers,
				COALESCE(SUM(s.total_questions), 0) as total_questions,
				RANK() OVER (ORDER BY COALESCE(SUM(s.score), 0) DESC, COALESCE(SUM(s.correct_answers), 0) DESC) as rank
			FROM users u
			INNER JOIN scores s ON u.id = s.user_id
			GROUP BY u.id, u.name
		)
		SELECT user_id, user_name, total_score, total_quizzes, correct_answers, total_questions, rank
		FROM ranked_users
		WHERE user_id = $1
	`

	var userRank domain.UserRankResponse
	err := lr.db.QueryRowContext(ctx, query, userId).Scan(
		&userRank.UserID,
		&userRank.UserName,
		&userRank.TotalScore,
		&userRank.TotalQuizzes,
		&userRank.CorrectAnswers,
		&userRank.TotalQuestions,
		&userRank.Rank,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, pkg.ErrUserRankNotFound
		}
		return nil, err
	}

	if userRank.TotalQuestions > 0 {
		userRank.AccuracyPercent = float64(userRank.CorrectAnswers) / float64(userRank.TotalQuestions) * 100
	}

	// Get total users count
	countQuery := `SELECT COUNT(DISTINCT user_id) FROM scores`
	if err := lr.db.QueryRowContext(ctx, countQuery).Scan(&userRank.TotalUsers); err != nil {
		return nil, err
	}

	return &userRank, nil
}

// GetUserSubjectRank returns the user's rank on a subject-specific leaderboard
func (lr *leaderboardRepository) GetUserSubjectRank(ctx context.Context, userId, subjectId int64) (*domain.UserRankResponse, error) {
	query := `
		WITH ranked_users AS (
			SELECT 
				u.id as user_id,
				u.name as user_name,
				COALESCE(SUM(s.score), 0) as total_score,
				COUNT(s.id) as total_quizzes,
				COALESCE(SUM(s.correct_answers), 0) as correct_answers,
				COALESCE(SUM(s.total_questions), 0) as total_questions,
				RANK() OVER (ORDER BY COALESCE(SUM(s.score), 0) DESC, COALESCE(SUM(s.correct_answers), 0) DESC) as rank
			FROM users u
			INNER JOIN scores s ON u.id = s.user_id
			WHERE s.subject_id = $1
			GROUP BY u.id, u.name
		)
		SELECT user_id, user_name, total_score, total_quizzes, correct_answers, total_questions, rank
		FROM ranked_users
		WHERE user_id = $2
	`

	var userRank domain.UserRankResponse
	err := lr.db.QueryRowContext(ctx, query, subjectId, userId).Scan(
		&userRank.UserID,
		&userRank.UserName,
		&userRank.TotalScore,
		&userRank.TotalQuizzes,
		&userRank.CorrectAnswers,
		&userRank.TotalQuestions,
		&userRank.Rank,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, pkg.ErrUserRankNotFound
		}
		return nil, err
	}

	if userRank.TotalQuestions > 0 {
		userRank.AccuracyPercent = float64(userRank.CorrectAnswers) / float64(userRank.TotalQuestions) * 100
	}

	// Get total users count for this subject
	countQuery := `SELECT COUNT(DISTINCT user_id) FROM scores WHERE subject_id = $1`
	if err := lr.db.QueryRowContext(ctx, countQuery, subjectId).Scan(&userRank.TotalUsers); err != nil {
		return nil, err
	}

	return &userRank, nil
}
