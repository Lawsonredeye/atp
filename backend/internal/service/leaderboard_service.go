package service

import (
	"context"

	"github.com/lawson/otterprep/domain"
	"github.com/lawson/otterprep/internal/repository"
)

type LeaderboardService interface {
	GetLeaderboard(ctx context.Context, query domain.LeaderboardQuery) (*domain.LeaderboardResponse, error)
	GetUserRank(ctx context.Context, userId int64, subjectId *int64) (*domain.UserRankResponse, error)
}

type leaderboardService struct {
	leaderboardRepo repository.LeaderboardRepository
	subjectRepo     repository.SubjectRepository
}

func NewLeaderboardService(leaderboardRepo repository.LeaderboardRepository, subjectRepo repository.SubjectRepository) LeaderboardService {
	return &leaderboardService{
		leaderboardRepo: leaderboardRepo,
		subjectRepo:     subjectRepo,
	}
}

// GetLeaderboard returns the leaderboard based on query parameters
func (ls *leaderboardService) GetLeaderboard(ctx context.Context, query domain.LeaderboardQuery) (*domain.LeaderboardResponse, error) {
	// Set defaults
	limit := query.Limit
	if limit == 0 {
		limit = 10
	}
	offset := query.Offset
	period := query.Period
	if period == "" {
		period = "all_time"
	}

	var entries []domain.LeaderboardEntry
	var totalUsers int64
	var err error
	var subjectName string

	// If subject_id is provided, get subject-specific leaderboard
	if query.SubjectId != nil && *query.SubjectId > 0 {
		// Get subject name
		subject, err := ls.subjectRepo.GetSubjectById(ctx, *query.SubjectId)
		if err != nil {
			return nil, err
		}
		subjectName = subject.Name

		entries, totalUsers, err = ls.leaderboardRepo.GetSubjectLeaderboard(ctx, *query.SubjectId, limit, offset)
		if err != nil {
			return nil, err
		}
	} else {
		// Global leaderboard based on period
		switch period {
		case "weekly":
			entries, totalUsers, err = ls.leaderboardRepo.GetWeeklyLeaderboard(ctx, limit, offset)
		case "monthly":
			entries, totalUsers, err = ls.leaderboardRepo.GetMonthlyLeaderboard(ctx, limit, offset)
		default: // "all_time"
			entries, totalUsers, err = ls.leaderboardRepo.GetGlobalLeaderboard(ctx, limit, offset)
		}
		if err != nil {
			return nil, err
		}
	}

	return &domain.LeaderboardResponse{
		SubjectId:   query.SubjectId,
		SubjectName: subjectName,
		Period:      period,
		TotalUsers:  totalUsers,
		Entries:     entries,
	}, nil
}

// GetUserRank returns the user's rank on the leaderboard
func (ls *leaderboardService) GetUserRank(ctx context.Context, userId int64, subjectId *int64) (*domain.UserRankResponse, error) {
	if subjectId != nil && *subjectId > 0 {
		return ls.leaderboardRepo.GetUserSubjectRank(ctx, userId, *subjectId)
	}
	return ls.leaderboardRepo.GetUserRank(ctx, userId)
}
