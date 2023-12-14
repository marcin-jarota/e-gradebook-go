package ports

import "e-student/internal/app/domain"

type (
	LessonRepository interface {
		Create(lesson *domain.Lesson) error
		GetByClassGroup(classGroupID int) ([]*domain.Lesson, error)
	}

	CreateLessonPayload struct {
		TeacherID    int    `json:"teacherID"`
		SubjectID    int    `json:"subjectID"`
		ClassGroupID int    `json:"classGroupID"`
		StartTime    string `json:"startTime"`
		EndTime      string `json:"endTime"`
		DayOfWeek    int    `json:"DayOfWeek"`
	}

	ClassGroupLesson struct {
		ID        int    `json:"id"`
		Subject   string `json:"subject"`
		Teacher   string `json:"teacher"`
		Start     string `json:"start"`
		End       string `json:"end"`
		DayOfWeek int    `json:"dayOfWeek"`
	}

	LessonService interface {
		Create(payload *CreateLessonPayload) error
		GetByClassGroup(classGroupID int) ([]*ClassGroupLesson, error)
	}
)
