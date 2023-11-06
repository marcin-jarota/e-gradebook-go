package domain

import (
	// "time"

	"time"

	"gorm.io/gorm"
)

type Mark struct {
	gorm.Model
	Value     float32
	Subject   Subject `gorm:"foreignKey:subject_id;references:id"`
	StudentID uint    `gorm:"foreignKey:student_id;references:id"`
	Comment   string
	Date      *time.Time
	SubjectID uint
	TeacherID uint
}
