package handler

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/lawson/otterprep/domain"
	"github.com/lawson/otterprep/internal/middleware"
	"github.com/lawson/otterprep/internal/service"
	"github.com/lawson/otterprep/pkg"
)

type LeaderboardHandler struct {
	leaderboardService service.LeaderboardService
	logger             *log.Logger
}

func NewLeaderboardHandler(leaderboardService service.LeaderboardService, logger *log.Logger) *LeaderboardHandler {
	return &LeaderboardHandler{
		leaderboardService: leaderboardService,
		logger:             logger,
	}
}

// GetLeaderboard returns the leaderboard
// @Summary Get leaderboard
// @Description Get the leaderboard with optional filtering by subject and time period
// @Tags Leaderboard
// @Accept json
// @Produce json
// @Param subject_id query int false "Subject ID for subject-specific leaderboard"
// @Param period query string false "Time period: all_time, weekly, monthly" default(all_time)
// @Param limit query int false "Number of entries to return" default(10)
// @Param offset query int false "Offset for pagination" default(0)
// @Success 200 {object} domain.LeaderboardResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /leaderboard [get]
func (h *LeaderboardHandler) GetLeaderboard(c echo.Context) error {
	var query domain.LeaderboardQuery

	// Parse subject_id
	subjectIdStr := c.QueryParam("subject_id")
	if subjectIdStr != "" {
		subjectId, err := strconv.ParseInt(subjectIdStr, 10, 64)
		if err != nil {
			h.logger.Println("error parsing subject_id: ", err)
			return pkg.ErrorResponse(c, err, http.StatusBadRequest)
		}
		query.SubjectId = &subjectId
	}

	// Parse period
	query.Period = c.QueryParam("period")
	if query.Period == "" {
		query.Period = "all_time"
	}

	// Parse limit
	limitStr := c.QueryParam("limit")
	if limitStr != "" {
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			h.logger.Println("error parsing limit: ", err)
			return pkg.ErrorResponse(c, err, http.StatusBadRequest)
		}
		query.Limit = limit
	} else {
		query.Limit = 10
	}

	// Parse offset
	offsetStr := c.QueryParam("offset")
	if offsetStr != "" {
		offset, err := strconv.Atoi(offsetStr)
		if err != nil {
			h.logger.Println("error parsing offset: ", err)
			return pkg.ErrorResponse(c, err, http.StatusBadRequest)
		}
		query.Offset = offset
	}

	// Validate
	if err := c.Validate(&query); err != nil {
		return err
	}

	leaderboard, err := h.leaderboardService.GetLeaderboard(c.Request().Context(), query)
	if err != nil {
		h.logger.Println("error getting leaderboard: ", err)
		return pkg.ErrorResponse(c, err, http.StatusInternalServerError)
	}

	h.logger.Println("Successfully retrieved leaderboard")
	return pkg.SuccessResponse(c, leaderboard, http.StatusOK)
}

// GetMyRank returns the authenticated user's rank
// @Summary Get my rank
// @Description Get the authenticated user's rank on the leaderboard
// @Tags Leaderboard
// @Accept json
// @Produce json
// @Param subject_id query int false "Subject ID for subject-specific rank"
// @Success 200 {object} domain.UserRankResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /leaderboard/me [get]
func (h *LeaderboardHandler) GetMyRank(c echo.Context) error {
	userId, ok := middleware.GetUserID(c)
	if !ok {
		return pkg.ErrorResponse(c, pkg.ErrUnauthorized, http.StatusUnauthorized)
	}

	var subjectId *int64
	subjectIdStr := c.QueryParam("subject_id")
	if subjectIdStr != "" {
		id, err := strconv.ParseInt(subjectIdStr, 10, 64)
		if err != nil {
			h.logger.Println("error parsing subject_id: ", err)
			return pkg.ErrorResponse(c, err, http.StatusBadRequest)
		}
		subjectId = &id
	}

	rank, err := h.leaderboardService.GetUserRank(c.Request().Context(), userId, subjectId)
	if err != nil {
		h.logger.Println("error getting user rank: ", err)
		if err == pkg.ErrUserRankNotFound {
			return pkg.ErrorResponse(c, err, http.StatusNotFound)
		}
		return pkg.ErrorResponse(c, err, http.StatusInternalServerError)
	}

	h.logger.Printf("Successfully retrieved rank for user %d", userId)
	return pkg.SuccessResponse(c, rank, http.StatusOK)
}

// GetUserRank returns a specific user's rank (admin only or public profile)
// @Summary Get user rank
// @Description Get a specific user's rank on the leaderboard
// @Tags Leaderboard
// @Accept json
// @Produce json
// @Param user_id path int true "User ID"
// @Param subject_id query int false "Subject ID for subject-specific rank"
// @Success 200 {object} domain.UserRankResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /leaderboard/user/{user_id} [get]
func (h *LeaderboardHandler) GetUserRank(c echo.Context) error {
	userIdStr := c.Param("user_id")
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		h.logger.Println("error parsing user_id: ", err)
		return pkg.ErrorResponse(c, err, http.StatusBadRequest)
	}

	var subjectId *int64
	subjectIdStr := c.QueryParam("subject_id")
	if subjectIdStr != "" {
		id, err := strconv.ParseInt(subjectIdStr, 10, 64)
		if err != nil {
			h.logger.Println("error parsing subject_id: ", err)
			return pkg.ErrorResponse(c, err, http.StatusBadRequest)
		}
		subjectId = &id
	}

	rank, err := h.leaderboardService.GetUserRank(c.Request().Context(), userId, subjectId)
	if err != nil {
		h.logger.Println("error getting user rank: ", err)
		if errors.Is(err, pkg.ErrUserRankNotFound) {
			return pkg.ErrorResponse(c, err, http.StatusNotFound)
		}
		return pkg.ErrorResponse(c, err, http.StatusInternalServerError)
	}

	h.logger.Printf("Successfully retrieved rank for user %d", userId)
	return pkg.SuccessResponse(c, rank, http.StatusOK)
}
