package ports

import (
	"e-student/internal/app/domain"
	"time"
)

type (
	StudentRepository interface {
		GetAll() ([]domain.Student, error)
		GetAllByClassGroup(classGroupID uint) ([]domain.Student, error)
		AddStudent(student *domain.Student) error
		ExistsByEmail(email string) (bool, error)
		SetClassGroup(studentID uint, classGroupID uint) error
		RemoveClassGroup(studentID uint) error
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

	StudentByClassGroup struct {
		ID      int     `json:"id"`
		Name    string  `json:"name"`
		Surname string  `json:"surname"`
		Email   string  `json:"email"`
		AvgMark float32 `json:"avgMark"`
	}

	SetClassGroupPayload struct {
		StudentID    int `json:"studentID"`
		ClassGroupID int `json:"classGroupID"`
	}

	StudentService interface {
		GetAll() ([]*StudentOutput, error)
		GetAllByClassGroup(classGroupID int) ([]StudentByClassGroup, error)
		// GetMarks(studentID int) ([]*StudentMarkOutput, error)
		AddStudent(user UserCreatePayload) error
		SetClassGroup(payload SetClassGroupPayload) error
	}
)
