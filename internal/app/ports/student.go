package ports

import (
	"e-student/internal/app/domain"
	"time"
)

type (
	StudentRepository interface {
		GetAll() ([]*domain.Student, error)
		AddStudent(student *domain.Student) error
		GetMarks(studentID int) ([]domain.Mark, error)
	}

	StudentOutput struct {
		ID        uint          `json:"id"`
		Name      string        `json:"name"`
		Surname   string        `json:"surname"`
		FullName  string        `json:"fullName"`
		Email     string        `json:"email"`
		Active    bool          `json:"active"`
		Marks     []domain.Mark `json:"marks"`
		CreatedAt time.Time     `json:"createdAt"`
		UpdatedAt time.Time     `json:"updatedAt"`
	}

	StudentMarkOutput struct {
		ID      uint           `json:"id"`
		Value   float32        `json:"value"`
		Subject domain.Subject `json:"subject"`
	}

	StudentService interface {
		GetAll() ([]*StudentOutput, error)
		GetMarks(studentID int) ([]*StudentMarkOutput, error)
	}
)