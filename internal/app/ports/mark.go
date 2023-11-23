package ports

import "e-student/internal/app/domain"

type (
	MarkRepository interface {
		AddMark(mark domain.Mark) error
		GetByStudent(studentID int) ([]domain.Mark, error)
		GetByClassGroup(classGroupID int) ([]domain.Mark, error)
	}

	MarkOutputSubject struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	MarkOutputTeacher struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		Surname string `json:"surname"`
	}

	MarkOutput struct {
		Comment string            `json:"comment,omitempty"`
		Value   float32           `json:"value"`
		Subject MarkOutputSubject `json:"subject"`
		Teacher MarkOutputTeacher `json:"teacher"`
	}

	SimpleMark struct {
		ID    int     `json:"id"`
		Value float32 `json:"value"`
	}

	MarkService interface {
		CalculateAverage(marks []domain.Mark) float32
		GetByStudent(studentID int) ([]MarkOutput, error)
		GetByClassGroup(classGroupID int) ([]SimpleMark, error)
	}
)
