package domain

import "gorm.io/gorm"

type ClassGroup struct {
	gorm.Model
	Name     string
	Students []Student `gorm:"foreignkey:ClassGroupID"`
	Subjects []Subject `gorm:"many2many:subject_classes;"`
  Teachers []Teacher `gorm:"many2many:class_teachers;"`
}
