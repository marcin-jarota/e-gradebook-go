package ports

import "e-student/internal/app/domain"

type (
	TeacherRepository interface {
		AddTeacher(teacher domain.Teacher) error
		ExistsByEmail(email string) (bool, error)
	}

	TeacherService interface {
		AddTeacher(user UserCreatePayload) error
	}
)
