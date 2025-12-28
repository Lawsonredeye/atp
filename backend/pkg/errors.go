package pkg

import "errors"

var (
	ErrSubjectNotFound        = errors.New("subject not found")
	ErrQuestionNotFound       = errors.New("question not found")
	ErrQuestionOptionNotFound = errors.New("question option not found")
	ErrQuizNotFound           = errors.New("quiz not found")
)
