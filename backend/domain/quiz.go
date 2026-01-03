package domain

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
