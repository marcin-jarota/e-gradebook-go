package service

import (
	"e-student/internal/app/domain"
	"e-student/internal/app/ports"
)

type StudentService struct {
	studentRepo ports.StudentRepository
}

func NewStudentService(repo ports.StudentRepository) *StudentService {
	return &StudentService{
		studentRepo: repo,
	}
}

func (s *StudentService) GetAll() ([]*domain.StudentResponse, error) {
	var studentsResponse []*domain.StudentResponse
	students, err := s.studentRepo.GetAll()

	if err != nil {
		return nil, err
	}

	for _, student := range students {
		studentsResponse = append(studentsResponse, &domain.StudentResponse{
			ID:        student.ID,
			Name:      student.User.Name,
			Surname:   student.User.Surname,
			FullName:  student.User.GetFullName(),
			Email:     student.User.Email,
			Active:    student.User.Active,
			Marks:     student.Marks,
			CreatedAt: student.CreatedAt,
			UpdatedAt: student.UpdatedAt,
			DeletedAt: student.DeletedAt,
		})
	}

	return studentsResponse, nil
}
