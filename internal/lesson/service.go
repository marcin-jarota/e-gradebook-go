package lesson

import (
	"e-student/internal/app/domain"
	"e-student/internal/app/ports"
	"errors"
	"time"
)

type LessonService struct {
	lessonRepo ports.LessonRepository
}

func NewLessonService(lessonRepo ports.LessonRepository) *LessonService {
	return &LessonService{
		lessonRepo: lessonRepo,
	}
}

func (s *LessonService) Create(payload *ports.CreateLessonPayload) error {
	var lesson domain.Lesson
	parsedStart, err := time.Parse("15:04", payload.StartTime)

	if err != nil {
		return errors.New("lesson.create.invalidStartTime")
	}

	parsedEnd, err := time.Parse("15:04", payload.EndTime)

	if err != nil {
		return errors.New("lesson.create.invalidEndTime")

	}

	if !lesson.ValidateDayOfWeek(payload.DayOfWeek) {
		return errors.New("lesson.create.invalidDayOfWeek")
	}

	lesson.DayOfWeek = payload.DayOfWeek
	lesson.Start = parsedStart
	lesson.End = parsedEnd
	lesson.TeacherID = uint(payload.TeacherID)
	lesson.SubjectID = uint(payload.SubjectID)
	lesson.ClassGroupID = uint(payload.ClassGroupID)

	if err := s.lessonRepo.Create(&lesson); err != nil {
		return errors.New("lesson.create.internalError")
	}

	return nil
}

func (s *LessonService) GetByClassGroup(classGroupID int) ([]*ports.ClassGroupLesson, error) {
	var output []*ports.ClassGroupLesson

	lessons, err := s.lessonRepo.GetByClassGroup(classGroupID)

	if err != nil {
		return nil, errors.New("lessons.classGroup.fetchInternal")
	}

	for _, l := range lessons {
		output = append(output, &ports.ClassGroupLesson{
			ID:        int(l.ID),
			Subject:   l.Subject.Name,
			Teacher:   l.Teacher.User.GetFullName(),
			Start:     l.Start.Format("15:04"),
			End:       l.End.Format("15:04"),
			DayOfWeek: l.DayOfWeek,
		})
	}

	return output, nil
}
