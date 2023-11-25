package teacher

import (
	"e-student/internal/app/domain"
	"e-student/internal/app/ports"
	"errors"
)

type TeacherService struct {
	repo ports.TeacherRepository
}

func NewTeacherService(repo ports.TeacherRepository) *TeacherService {
	return &TeacherService{
		repo: repo,
	}
}

func (s *TeacherService) AddTeacher(user ports.UserCreatePayload) error {

	exists, err := s.repo.ExistsByEmail(user.Email)

	if err != nil {
		return err
	}

	if exists {
		return errors.New("teacher.create.exists")
	}

	return s.repo.AddTeacher(domain.Teacher{
		User: domain.User{
			Name:    user.Name,
			Surname: user.Surname,
			Email:   user.Email,
			Role:    domain.TeacherRole,
		},
	})
}
