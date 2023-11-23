package student

import (
	"e-student/internal/app/domain"
	"e-student/internal/app/ports"
	"errors"
	"fmt"
)

type StudentService struct {
	studentRepo ports.StudentRepository
	markService ports.MarkService
}

func NewStudentService(repo ports.StudentRepository, markService ports.MarkService) *StudentService {
	return &StudentService{
		studentRepo: repo,
		markService: markService,
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

func (s *StudentService) SetClassGroup(payload ports.SetClassGroupPayload) error {
	return s.studentRepo.SetClassGroup(uint(payload.StudentID), uint(payload.ClassGroupID))
}

func (s *StudentService) GetAllByClassGroup(classGroupID int) ([]ports.StudentByClassGroup, error) {
	var list []ports.StudentByClassGroup

	students, err := s.studentRepo.GetAllByClassGroup(uint(classGroupID))

	if err != nil {
		fmt.Println(err)
		return nil, errors.New("students.classGroup.error")
	}

	for _, student := range students {
		list = append(list, ports.StudentByClassGroup{
			ID:      int(student.ID),
			Name:    student.User.Name,
			Surname: student.User.Surname,
			Email:   student.User.Email,
			AvgMark: s.markService.CalculateAverage(student.Marks),
		})
	}

	return list, nil
}

// func (s *StudentService) GetMarks(studentId int) ([]*ports.StudentMarkOutput, error) {
// 	var studentMarks []*ports.StudentMarkOutput
// 	marks, err := s.studentRepo.GetMarks(studentId)
//
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	for _, mark := range marks {
// 		studentMarks = append(studentMarks, &ports.StudentMarkOutput{
// 			ID:      mark.ID,
// 			Value:   mark.Value,
// 			Subject: mark.Subject,
// 		})
// 	}
//
// 	return studentMarks, nil
// }

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
