package student

import (
	"e-student/internal/app/domain"
	"e-student/internal/app/ports"
	"errors"
)

type StudentService struct {
	studentRepo ports.StudentRepository
}

func NewStudentService(repo ports.StudentRepository) *StudentService {
	return &StudentService{
		studentRepo: repo,
	}
}

func (s *StudentService) GetAll() ([]*ports.StudentOutput, error) {
	var studentsResponse []*ports.StudentOutput
	students, err := s.studentRepo.GetAll()

	if err != nil {
		return nil, err
	}

	for _, student := range students {
		studentsResponse = append(studentsResponse, &ports.StudentOutput{
			ID:        student.ID,
			Name:      student.User.Name,
			Surname:   student.User.Surname,
			FullName:  student.User.GetFullName(),
			Email:     student.User.Email,
			Active:    student.User.Active,
			Marks:     student.Marks,
			CreatedAt: student.CreatedAt,
			UpdatedAt: student.UpdatedAt,
		})
	}

	return studentsResponse, nil
}

func (s *StudentService) GetMarks(studentId int) ([]*ports.StudentMarkOutput, error) {
	var studentMarks []*ports.StudentMarkOutput
	marks, err := s.studentRepo.GetMarks(studentId)

	if err != nil {
		return nil, err
	}

	for _, mark := range marks {
		studentMarks = append(studentMarks, &ports.StudentMarkOutput{
			ID:      mark.ID,
			Value:   mark.Value,
			Subject: mark.Subject,
		})
	}

	return studentMarks, nil
}

func (s *StudentService) AddStudent(student *ports.StudentCreatePayload) error {
	// TODO: validate
	exists, err := s.studentRepo.ExistsByEmail(student.Email)

	if err != nil {
		return err
	}

	if exists {
		return errors.New("student with this email exists")
	}

	return s.studentRepo.AddStudent(&domain.Student{
		User: domain.User{
			Name:     student.Name,
			Surname:  student.Surname,
			Email:    student.Email,
			Password: student.Password,
			Role:     domain.StudentRole,
		},
	})
}
