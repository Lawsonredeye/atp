package domain

type Question struct {
	ID          int64    `json:"id"`
	Text        string   `json:"text"`
	Option      []string `json:"option"`
	Answer      string   `json:"answer"`
	Explanation string   `json:"explanation"`
}
