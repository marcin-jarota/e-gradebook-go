package service

import (
	"e-student/internal/app/domain"
	"e-student/internal/app/ports"
)

type StudentService struct {
	studentRepo ports.StudentRepository
	marksRepo   ports.MarkRepository
}

func NewStudentService(repo ports.StudentRepository, marksRepo ports.MarkRepository) *StudentService {
	return &StudentService{
		studentRepo: repo,
		marksRepo:   marksRepo,
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

func (s *StudentService) GetMarks(studentId int) ([]*domain.Mark, error) {
	var studentMarks []*domain.StudentMarkResponse
	marks, err := s.marksRepo.GetMarksByStudent(studentId)

	if err != nil {
		return marks, err
	}

	for _, mark := range marks {
		studentMarks = append(studentMarks, &domain.StudentMarkResponse{
			ID:      mark.ID,
			Subject: mark.Subject,
		})
	}

	return marks, nil
}
