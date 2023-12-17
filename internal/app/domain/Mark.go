package domain

import (
	"time"

	"gorm.io/gorm"
)

type Mark struct {
	gorm.Model
	Value        float32
	Subject      Subject
	StudentID    uint
	Comment      string
	Date         *time.Time
	SubjectID    uint
	Student      Student
	Teacher      Teacher
	TeacherID    uint
	SchoolYear   SchoolYear
	SchoolYearID *uint
}
