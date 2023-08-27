package domain

type ClassGroup struct {
	baseModel
	Students []Student `json:"students"`
}
