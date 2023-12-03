package domain

import (
	"gorm.io/gorm"
)

type Teacher struct {
	gorm.Model
	User     User
	UserID   uint
	Marks    []Mark
	Subjects []Subject `gorm:"many2many:subject_teachers;"`
  ClassGroups []ClassGroup `gorm:"many2many:class_teachers;"`
}
