package domain

type Subject struct {
	baseModel
	Name  string `json:"name"`
	Marks []Mark `json:"marks,omitempty"`
}
