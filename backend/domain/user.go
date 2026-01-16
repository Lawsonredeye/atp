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

type UserResponse struct {
	ID        int64     `json:"id"`
	Name      string    `json:"full_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type RegisterUser struct {
	Name     string `json:"full_name" validate:"required,min=2,max=100"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,password"`
}

type LoginUser struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UpdateUsername struct {
	NewUsername string `json:"new_username" validate:"required,min=2,max=100"`
}

type UpdateEmail struct {
	NewEmail string `json:"new_email" validate:"required,email"`
}

type UpdatePassword struct {
	NewPassword string `json:"new_password" validate:"required,password"`
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

// User roles
var (
	UserAdmin = "admin"
	UserUser  = "user"
)

// Quiz modes
var (
	ModeSingle   = "single"
	ModeMultiple = "multiple"
)

// User Dashboard details, including scores and other details
type UserDashboard struct {
	UserResponse
	UserStats
	Roles []string `json:"roles"`
}

// User stats
type UserStats struct {
	UserID                 int64 `json:"user_id"`
	TotalQuizzesTaken      int64 `json:"total_quizzes_taken"`
	TotalCorrectAnswers    int64 `json:"total_correct_answers"`
	TotalIncorrectAnswers  int64 `json:"total_incorrect_answers"`
	TotalQuestionsAnswered int64 `json:"total_questions_answered"`
}

// ForgotPasswordRequest is the request body for initiating a password reset
type ForgotPasswordRequest struct {
	Email string `json:"email" validate:"required,email"`
}

// ResetPasswordRequest is the request body for resetting password with token
type ResetPasswordRequest struct {
	Token       string `json:"token" validate:"required"`
	NewPassword string `json:"new_password" validate:"required,password"`
}

// ValidateResetTokenRequest is the request to validate a reset token
type ValidateResetTokenRequest struct {
	Token string `json:"token" validate:"required"`
}
