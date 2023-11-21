package mark

import (
	"e-student/internal/app/domain"
	"e-student/internal/app/ports"
	"errors"
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
	var list []ports.MarkOutput
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

func (s *markService) CalculateAverage(marks []domain.Mark) float32 {
	var avgMarg float32

	for _, mark := range marks {
		avgMarg += mark.Value
	}

	return avgMarg
}
