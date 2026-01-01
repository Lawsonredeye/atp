package domain

type QuizResponse struct {
	QuestionId      int64
	Question        string
	SelectedOptions []string
	Answer          string
	Explanation     string
}
