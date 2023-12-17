package domain

import (
	"time"

	"gorm.io/gorm"
)

type Lesson struct {
	gorm.Model
	Subject      Subject
	SubjectID    uint
	Teacher      Teacher
	TeacherID    uint
	ClassGroup   ClassGroup
	ClassGroupID uint
	Start        time.Time
	End          time.Time
	DayOfWeek    int
	SchoolYear   SchoolYear
	SchoolYearID *uint
}

func (l *Lesson) ValidateDayOfWeek(dayOfWeek int) bool {
	return dayOfWeek >= 0 && dayOfWeek <= 6
}
