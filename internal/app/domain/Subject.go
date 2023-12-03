package domain

import (
	"gorm.io/gorm"
)

type Subject struct {
	gorm.Model
	Name        string
	Marks       []Mark
	ClassGroups []ClassGroup `gorm:"many2many:subject_classes;"`
	Teachers    []Teacher    `gorm:"many2many:subject_teachers;"`
}
