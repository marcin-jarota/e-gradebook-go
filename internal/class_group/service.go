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

func (s *ClassGroupService) GetAll() ([]ports.ClassGroupOutput, error) {
	var output []ports.ClassGroupOutput
	classGroups, err := s.repo.GetAll()

	if err != nil {
		fmt.Println(err)
		return nil, errors.New("classGroup.error.fetch")
	}

	for _, c := range classGroups {
		output = append(output, ports.ClassGroupOutput{
			ID:            int(c.ID),
			Name:          c.Name,
			StudentsCount: len(c.Students),
		})
	}

	return output, nil
}

func (s *ClassGroupService) GetOneByID(classGroupID int) (ports.ClassGroupOutput, error) {
	var output ports.ClassGroupOutput
	classGroup, err := s.repo.GetOneByID(classGroupID)

	if err != nil {
		return output, errors.New("classGroup.error.fetchOne")
	}

	output = ports.ClassGroupOutput{
		ID:            int(classGroup.ID),
		Name:          classGroup.Name,
		StudentsCount: len(classGroup.Students),
	}

	return output, nil
}

func (s *ClassGroupService) AddClassGroup(input ports.AddClassGroupInput) error {
	err := s.repo.AddClassGroup(&domain.ClassGroup{
		Name: input.Name,
	})

	return err
}

func (s *ClassGroupService) AddSubject(input ports.AddSubjectToClassGroupPayload) error {
	return s.repo.AddSubject(input.ClassGroupID, input.SubjectID)
}
