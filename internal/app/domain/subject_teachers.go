package domain

import "gorm.io/gorm"

type SubjectTeacher struct {
	gorm.Model
	SubjectID uint
	TeacherID uint
}
