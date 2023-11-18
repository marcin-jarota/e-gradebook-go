package classgroup

import (
	"e-student/internal/app/domain"
	"e-student/internal/app/ports"
	"errors"
	"fmt"
)

type ClassGroupService struct {
	repo ports.ClassGroupRepository
}

func NewClassGroupService(repo ports.ClassGroupRepository) *ClassGroupService {
	return &ClassGroupService{
		repo: repo,
	}
}

// ListStudents(classGroupID uint) ([]*ListStudentsOutput, error)
// AddStudent(studentID uint, classGroupID uint) error
// AddClassGroup(input *AddClassGroupInput) error

func (s *ClassGroupService) GetAll() ([]*ports.ListClassGroupsOutput, error) {
	var output []*ports.ListClassGroupsOutput
	classGroups, err := s.repo.GetAll()

	if err != nil {
		fmt.Println(err)
		return nil, errors.New("classGroup.error.fetch")
	}

	for _, c := range classGroups {
		output = append(output, &ports.ListClassGroupsOutput{
			Name:          c.Name,
			StudentsCount: len(c.Students),
		})
	}

	return output, nil
}

func (s *ClassGroupService) ListStudents(classGroupID uint) ([]*ports.ListStudentsOutput, error) {
	var output []*ports.ListStudentsOutput

	students, err := s.repo.GetStudents(classGroupID)

	if err != nil {
		fmt.Println(err)
		return nil, errors.New("classGroup.error.studentsFetch")
	}

	for _, s := range students {
		output = append(output, &ports.ListStudentsOutput{
			Name:    s.User.Name,
			Surname: s.User.Surname,
		})
	}

	return output, nil
}

func (s *ClassGroupService) AddClassGroup(input *ports.AddClassGroupInput) error {
	err := s.repo.AddClassGroup(&domain.ClassGroup{
		Name: input.Name,
	})

	return err
}
