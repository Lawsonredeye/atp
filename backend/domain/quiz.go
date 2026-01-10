package domain

import "errors"

// SubmitQuizRequest is used when submitting quiz answers
type SubmitQuizRequest struct {
	QuestionId       int64   `json:"question_id" validate:"required,gt=0"`
	IsMultipleChoice bool    `json:"is_multiple_choice"`
	OptionIds        []int64 `json:"option_ids" validate:"required,min=1,dive,gt=0"`
}

// QuizRequest is used when requesting to generate a quiz
type QuizRequest struct {
	SubjectId      int64 `json:"subject_id" validate:"required,gt=0"`
	NumOfQuestions int64 `json:"num_of_questions" validate:"required,gte=1,lte=100"`
}

// QuizOptionResponse represents an option without revealing if it's correct
type QuizOptionResponse struct {
	Id     int64  `json:"id"`
	Option string `json:"option"`
}

// QuizQuestionResponse represents a question in a generated quiz (for frontend)
type QuizQuestionResponse struct {
	QuestionId       int64                `json:"question_id"`
	Question         string               `json:"question"`
	SubjectId        int64                `json:"subject_id"`
	IsMultipleChoice bool                 `json:"is_multiple_choice"`
	Options          []QuizOptionResponse `json:"options"`
}

// GeneratedQuizResponse is the response when generating a quiz
type GeneratedQuizResponse struct {
	SubjectId  int64                  `json:"subject_id"`
	TotalCount int                    `json:"total_count"`
	Questions  []QuizQuestionResponse `json:"questions"`
}

// QuizResultResponse is the response after submitting a quiz (reveals answers)
type QuizResultResponse struct {
	QuestionId      int64    `json:"question_id"`
	Question        string   `json:"question"`
	SelectedOptions []string `json:"selected_options"`
	CorrectAnswer   string   `json:"correct_answer"`
	IsCorrect       bool     `json:"is_correct"`
	Explanation     string   `json:"explanation"`
}

// QuizSubmitResponse is the full response after submitting a quiz
type QuizSubmitResponse struct {
	UserId           int64                `json:"user_id"`
	SubjectId        int64                `json:"subject_id"`
	TotalQuestions   int64                `json:"total_questions"`
	CorrectAnswers   int64                `json:"correct_answers"`
	IncorrectAnswers int64                `json:"incorrect_answers"`
	Score            int64                `json:"score"`
	Results          []QuizResultResponse `json:"results"`
}

type QuestionsData struct {
	Name        string   `json:"name" validate:"required,min=1"`
	Options     []string `json:"options" validate:"required,min=2,dive,required"`
	Answer      string   `json:"answer" validate:"required"`
	Explanation string   `json:"explanation" validate:"required"`
}

func (qd *QuestionsData) Validate() error {
	if qd.Name == "" {
		return errors.New("question name is empty")
	}
	if len(qd.Options) == 0 {
		return errors.New("question options are empty")
	}
	if qd.Answer == "" {
		return errors.New("question answer is empty")
	}
	if qd.Explanation == "" {
		return errors.New("question explanation is empty")
	}
	return nil
}
