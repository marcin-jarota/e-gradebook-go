package domain

import (
	// "time"

	"gorm.io/gorm"
)

type Mark struct {
	gorm.Model
	Value     float32
	Subject   Subject `gorm:"foreignKey:subject_id;references:id"`
	StudentID uint    `gorm:"foreignKey:student_id;references:id"`
	SubjectID uint
	TeacherID uint
}
