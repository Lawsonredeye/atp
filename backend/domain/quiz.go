package domain

import "errors"

type SubmitQuizRequest struct {
	QuizId           int64   `json:"quiz_id"`
	IsMultipleChoice bool    `json:"is_multiple_choice"`
	OptionIds        []int64 `json:"option_ids"`
}

type QuizRequest struct {
	SubjectId      int64 `json:"subject_id"`
	NumOfQuestions int64 `json:"num_of_questions"`
}

type QuizResponse struct {
	QuestionId      int64
	Question        string
	SelectedOptions []string
	Answer          string
	Explanation     string
}

type QuestionsData struct {
	Name        string   `json:"name"`
	Options     []string `json:"options"`
	Answer      string   `json:"answer"`
	Explanation string   `json:"explanation"`
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
