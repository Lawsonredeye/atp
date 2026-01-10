package domain

type Subject struct {
	Id   int64  `json:"id"`
	Name string `json:"name" validate:"required,min=1,max=100"`
}
