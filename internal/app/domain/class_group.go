package domain

import "gorm.io/gorm"

type ClassGroup struct {
	gorm.Model
	Name          string
	EducationYear int
	SchoolYears   []SchoolYear `gorm:"many2many:school_year_class_group;"`
	Students      []Student    `gorm:"foreignkey:ClassGroupID"`
}
