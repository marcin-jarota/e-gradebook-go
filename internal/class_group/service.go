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

func (s *ClassGroupService) GetAll() ([]*ports.ListClassGroupsOutput, error) {
	var output []*ports.ListClassGroupsOutput
	classGroups, err := s.repo.GetAll()

	if err != nil {
		fmt.Println(err)
		return nil, errors.New("classGroup.error.fetch")
	}

	for _, c := range classGroups {
		output = append(output, &ports.ListClassGroupsOutput{
			ID:            int(c.ID),
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

	for _, student := range students {
		output = append(output, &ports.ListStudentsOutput{
			Name:    student.User.Name,
			Surname: student.User.Surname,
			Email:   student.User.Email,
			AvgMark: s.calculateAverageMark(student.Marks),
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

func (s *ClassGroupService) calculateAverageMark(marks []domain.Mark) float32 {
	var avgMark float32
	for _, mark := range marks {
		avgMark = avgMark + mark.Value
	}
	return avgMark
}
