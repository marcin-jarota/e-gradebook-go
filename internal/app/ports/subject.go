package ports

import "e-student/internal/app/domain"

type (
	SubjectRepository interface {
		AddSubject(subject *domain.Subject) error
		GetAll() ([]*domain.Subject, error)
		GetOneByName(name string) (*domain.Subject, error)
		Exists(name string) (bool, error)
		DeleteByID(id uint) error
	}

	SubjectOutput struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	}

	SubjectAddPayload struct {
		Name string `json:"name"`
	}

	SubjectService interface {
		AddSubject(name string) error
		GetAll() ([]*SubjectOutput, error)
		Delete(id uint) error
	}
)
