package subject

import (
	"e-student/internal/app/domain"
	"e-student/internal/app/ports"
	"errors"
)

type SubjectService struct {
	repo ports.SubjectRepository
}

func NewSubjectService(repo ports.SubjectRepository) *SubjectService {
	return &SubjectService{
		repo,
	}
}

func (s *SubjectService) GetAll() ([]*ports.SubjectOutput, error) {
	var response []*ports.SubjectOutput
	subjects, err := s.repo.GetAll()

	if err != nil {
		return nil, err
	}

	for _, subject := range subjects {
		response = append(response, &ports.SubjectOutput{ID: subject.ID, Name: subject.Name})
	}

	return response, nil
}

func (s *SubjectService) AddSubject(name string) error {
	if name == "" {
		return errors.New("subject.missingName")
	}

	exists, err := s.repo.Exists(name)

	if err != nil {
		return err
	}

	if exists {
		return errors.New("subject.exists")
	}

	err = s.repo.AddSubject(&domain.Subject{Name: name})

	if err != nil {
		return err
	}

	return nil
}
