package ports

import "e-student/internal/app/domain"

type (
	TeacherRepository interface {
		GetAll() ([]domain.Teacher, error)
		AddTeacher(teacher domain.Teacher) error
		GetTeacherByUserID(id int) (domain.Teacher, error)
		ExistsByEmail(email string) (bool, error)
		// GetAllByClassGroup(classGroupID uint) ([]domain.Teacher, error)
	}

	TeacherBaseOutput struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		Surname string `json:"surname"`
		Email   string `json:"email,omitempty"`
	}

	TeacherService interface {
		GetAll() ([]TeacherBaseOutput, error)
		AddTeacher(user UserCreatePayload) error
		GetTeacherByUserID(id int) (*TeacherBaseOutput, error)
		// GetAllByClassGroup(classGroupID uint)
	}
)
