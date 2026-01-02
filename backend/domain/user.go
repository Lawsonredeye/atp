package domain

import "time"

type User struct {
	ID           int64     `json:"id"`
	Name         string    `json:"full_name"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"password_hash"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type UserScore struct {
	ID               int64     `json:"id"`
	UserID           int64     `json:"user_id"`
	SubjectID        int64     `json:"subject_id"`
	Score            int64     `json:"score"`
	CorrectAnswers   int64     `json:"correct_answers"`
	IncorrectAnswers int64     `json:"incorrect_answers"`
	TotalQuestions   int64     `json:"total_questions"`
	TimeTakenSeconds int64     `json:"time_taken_seconds"`
	Mode             string    `json:"mode"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

var (
	ModeSingle   = "single"
	ModeMultiple = "multiple"
)

type UserStats struct {
	UserID                 int64 `json:"user_id"`
	TotalQuizzesTaken      int64 `json:"total_quizzes_taken"`
	TotalCorrectAnswers    int64 `json:"total_correct_answers"`
	TotalIncorrectAnswers  int64 `json:"total_incorrect_answers"`
	TotalQuestionsAnswered int64 `json:"total_questions_answered"`
}
