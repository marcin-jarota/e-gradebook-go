package domain

import "gorm.io/gorm"

type ClassGroup struct {
	gorm.Model
	Name     string
	Students []Student `gorm:"foreignkey:ClassGroupID"`
}
