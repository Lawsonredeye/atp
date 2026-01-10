package domain

// LeaderboardEntry represents a single entry in the leaderboard
type LeaderboardEntry struct {
	Rank            int64   `json:"rank"`
	UserID          int64   `json:"user_id"`
	UserName        string  `json:"user_name"`
	TotalScore      int64   `json:"total_score"`
	TotalQuizzes    int64   `json:"total_quizzes"`
	CorrectAnswers  int64   `json:"correct_answers"`
	TotalQuestions  int64   `json:"total_questions"`
	AccuracyPercent float64 `json:"accuracy_percent"`
}

// LeaderboardResponse is the response for leaderboard requests
type LeaderboardResponse struct {
	SubjectId   *int64             `json:"subject_id,omitempty"`
	SubjectName string             `json:"subject_name,omitempty"`
	Period      string             `json:"period"` // "all_time", "weekly", "monthly"
	TotalUsers  int64              `json:"total_users"`
	Entries     []LeaderboardEntry `json:"entries"`
}

// UserRankResponse shows user's position on the leaderboard
type UserRankResponse struct {
	UserID          int64   `json:"user_id"`
	UserName        string  `json:"user_name"`
	Rank            int64   `json:"rank"`
	TotalScore      int64   `json:"total_score"`
	TotalQuizzes    int64   `json:"total_quizzes"`
	CorrectAnswers  int64   `json:"correct_answers"`
	TotalQuestions  int64   `json:"total_questions"`
	AccuracyPercent float64 `json:"accuracy_percent"`
	TotalUsers      int64   `json:"total_users"`
}

// LeaderboardQuery represents query parameters for leaderboard
type LeaderboardQuery struct {
	SubjectId *int64 `query:"subject_id"`
	Period    string `query:"period" validate:"omitempty,oneof=all_time weekly monthly"`
	Limit     int    `query:"limit" validate:"omitempty,gte=1,lte=100"`
	Offset    int    `query:"offset" validate:"omitempty,gte=0"`
}
