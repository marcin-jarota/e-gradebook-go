package subject

import (
	"e-student/internal/app/domain"
	"e-student/internal/app/ports"
	"errors"
	"fmt"
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
		var teachers []ports.SubjectTeacher

		for _, t := range subject.Teachers {
			teachers = append(teachers, ports.SubjectTeacher{
				ID:      int(t.ID),
				Name:    t.User.Name,
				Surname: t.User.Surname,
			})
		}
		response = append(response, &ports.SubjectOutput{ID: subject.ID, Name: subject.Name, Teachers: teachers})
	}

	return response, nil
}

func (s *SubjectService) Delete(id uint) error {
	err := s.repo.DeleteByID(id)

	if err != nil {
		fmt.Println(err)
		return errors.New("subject.error.cantDelete")
	}

	return nil
}

func (s *SubjectService) AddSubject(name string) error {
	if name == "" {
		return errors.New("subject.error.missingName")
	}

	exists, err := s.repo.Exists(name)

	if err != nil {
		return err
	}

	if exists {
		return errors.New("subject.error.exists")
	}

	err = s.repo.AddSubject(&domain.Subject{Name: name})

	if err != nil {
		return err
	}

	return nil
}

func (s *SubjectService) AddTeacher(payload ports.TeacherSubjectID) error {
	if err := s.repo.AddTeacher(payload.TeacherID, payload.SubjectID); err != nil {
		return errors.New("subject.error.assignTeacher")
	}

	return nil
}
