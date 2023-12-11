package classgroup

import (
	"e-student/internal/app/domain"
	"e-student/internal/app/ports"
	"errors"
	"fmt"
	"log"
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

func (s *ClassGroupService) AddTeacherWithSubject(input ports.TeacherSubjectClassgroupID) error {
	err := s.repo.AddTeacherWithSubject(input.ClassGroupID, input.TeacherID, input.SubjectID)

	if err != nil {
		log.Println(err)
		return errors.New("classGroup.error.assignTeacherSubject")
	}
	return nil
}

func (s *ClassGroupService) GetTeachersWithSubject(classGroupID int) ([]ports.TeacherSubject, error) {
	var output []ports.TeacherSubject

	res, err := s.repo.GetTeachersWithSubject(classGroupID)

	if err != nil {
		return nil, err
	}

	for _, entity := range res {
		output = append(output, ports.TeacherSubject{
			Teacher: ports.TeacherBaseOutput{
				ID:      int(entity.Teacher.ID),
				Name:    entity.Teacher.User.Name,
				Surname: entity.Teacher.User.Surname,
				Email:   entity.Teacher.User.Email,
			},
			Subject: ports.SubjectBaseOutput{
				ID:   entity.Subject.ID,
				Name: entity.Subject.Name,
			},
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
	if err := s.repo.AddSubject(input.ClassGroupID, input.SubjectID); err != nil {
		log.Println(err)
		return errors.New("classGroup.error.assignSubject")
	}

	return nil
}

func (s *ClassGroupService) GetTeachers(classGroupID int) ([]ports.ClassGroupTeacher, error) {
	var output []ports.ClassGroupTeacher

	teachers, err := s.repo.GetTeachers(classGroupID)

	if err != nil {
		return nil, err
	}

	for _, teacher := range teachers {
		output = append(output, ports.ClassGroupTeacher{
			ID:      int(teacher.ID),
			Name:    teacher.User.Name,
			Surname: teacher.User.Surname,
			Email:   teacher.User.Email,
		})
	}

	return output, nil
}

func (s *ClassGroupService) AddTeacher(input ports.TeacherClassGroup) error {
	return s.repo.AddTeacher(input.ClassGroupID, input.TeacherID)
}
