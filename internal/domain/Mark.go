package domain

type Mark struct {
	baseModel
	Value     float32 `json:"value"`
	SubjectID uint    `json:"-"`
	StudentID uint    `json:"studentId"`
}
