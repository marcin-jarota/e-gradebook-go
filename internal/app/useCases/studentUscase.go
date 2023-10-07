package usecases

import "e-student/internal/app/domain"

type StudentUscase interface {
	CreateStudent() (*domain.Student, error)
}
