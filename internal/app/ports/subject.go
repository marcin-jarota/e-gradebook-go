package ports

import "e-student/internal/app/domain"

type SubjectRepository interface {
	AddSubject(subject *domain.Subject) error
	GetAll() ([]*domain.Subject, error)
}
