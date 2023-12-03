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

func (s *TeacherService) GetAll() ([]ports.TeacherBaseOutput, error) {
	var output []ports.TeacherBaseOutput

	teachers, err := s.repo.GetAll()

	if err != nil {
		return nil, err
	}

	for _, t := range teachers {
		output = append(output, ports.TeacherBaseOutput{ID: int(t.ID), Name: t.User.Name, Surname: t.User.Surname})
	}

	return output, nil
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
