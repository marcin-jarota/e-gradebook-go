package domain

type Student struct {
	baseModel
	Marks  []Mark `json:"marks"`
	UserID uint   `json:"-"`
	User   User   `json:"user"`
}
