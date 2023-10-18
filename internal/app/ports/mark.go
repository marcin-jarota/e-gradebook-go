package ports

import "e-student/internal/app/domain"

type MarkRepository interface {
	AddMark(mark *domain.Mark) error
	GetMarksByStudent(studentId int) ([]*domain.Mark, error)
}
