package domain

import "errors"

type SubmitQuizRequest struct {
	QuizId           int64   `json:"quiz_id" validate:"required,gt=0"`
	IsMultipleChoice bool    `json:"is_multiple_choice"`
	OptionIds        []int64 `json:"option_ids" validate:"required,min=1,dive,gt=0"`
}

type QuizRequest struct {
	SubjectId      int64 `json:"subject_id" validate:"required,gt=0"`
	NumOfQuestions int64 `json:"num_of_questions" validate:"required,gte=1,lte=100"`
}

type QuizResponse struct {
	QuestionId      int64
	Question        string
	SelectedOptions []string
	Answer          string
	Explanation     string
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
