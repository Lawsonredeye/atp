package pkg

import "errors"

var (
	ErrSubjectNotFound            = errors.New("subject not found")
	ErrQuestionNotFound           = errors.New("question not found")
	ErrQuestionOptionNotFound     = errors.New("question option not found")
	ErrQuizNotFound               = errors.New("quiz not found")
	ErrInvalidName                = errors.New("invalid name")
	ErrInvalidEmail               = errors.New("invalid email")
	ErrInvalidPasswordHash        = errors.New("invalid password hash")
	ErrInvalidUserID              = errors.New("invalid User ID")
	ErrInternalServerError        = errors.New("internal server error")
	ErrQuestionTextNotFound       = errors.New("invalid / empty question text")
	ErrQuestionOptionTextNotFound = errors.New("invalid / empty question option text")
	ErrSubjectNameNotFound        = errors.New("invalid / empty subject name")
	ErrSubjectWithNameExists      = errors.New("subject with name already exists")
	ErrUserNotFound               = errors.New("user not found")
	ErrInvalidRole                = errors.New("invalid role")
	ErrInvalidPasswordLength      = errors.New("invalid password length should be greater or equal to 6")
	ErrUnauthorized               = errors.New("unauthorized")
)
