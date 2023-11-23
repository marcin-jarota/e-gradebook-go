package mark

import (
	"e-student/internal/app/domain"
	"e-student/internal/app/ports"
	"errors"
	"fmt"
	"strconv"
	"time"
)

type markService struct {
	markRepository ports.MarkRepository
}

func NewMarkService(markRepository ports.MarkRepository) *markService {
	return &markService{
		markRepository: markRepository,
	}
}

func (s *markService) GetByStudent(studentID int) ([]ports.MarkOutput, error) {
	list := []ports.MarkOutput{}
	marks, err := s.markRepository.GetByStudent(studentID)

	if err != nil {
		return nil, errors.New("marks.getByStudent.fetchError")
	}

	for _, mark := range marks {
		list = append(list, ports.MarkOutput{
			Value:   mark.Value,
			Comment: mark.Comment,
			Subject: ports.MarkOutputSubject{
				ID:   int(mark.Subject.ID),
				Name: mark.Subject.Name,
			},
			Teacher: ports.MarkOutputTeacher{
				ID:      int(mark.Teacher.ID),
				Name:    mark.Teacher.User.Name,
				Surname: mark.Teacher.User.Surname,
			},
		})
	}

	return list, nil
}

func (s *markService) CreateMark(p ports.MarkCreatePayload) error {
	date := time.Now()

	if p.Date != "" {
		parsed, err := time.Parse("01-02-2006", p.Date)

		if err != nil {
			fmt.Println(err)
			return errors.New("mark.create.invalidDate")
		}

		date = parsed
	}

	return s.markRepository.AddMark(&domain.Mark{
		TeacherID: uint(p.TeacherID),
		StudentID: uint(p.StudentID),
		SubjectID: uint(p.SubjectID),
		Comment:   p.Comment,
		Value:     p.Value,
		Date:      &date,
	})
}

func (s *markService) GetByClassGroup(classGroupID int) ([]ports.SimpleMark, error) {
	output := []ports.SimpleMark{}
	marks, err := s.markRepository.GetByClassGroup(classGroupID)

	if err != nil {
		return output, err
	}

	for _, mark := range marks {
		output = append(output, ports.SimpleMark{
			ID:    int(mark.ID),
			Value: mark.Value,
		})
	}

	return output, nil
}

func (s *markService) CalculateAverage(marks []domain.Mark) float32 {
	var avgMarg float32

	for _, mark := range marks {
		avgMarg += mark.Value
	}

	avg := avgMarg / float32(len(marks))

	val, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", avg), 32)

	return float32(val)
}
