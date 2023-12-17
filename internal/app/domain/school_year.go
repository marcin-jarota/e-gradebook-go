package domain

import (
	"time"

	"gorm.io/gorm"
)

type SchoolYear struct {
	gorm.Model
	Name        string
	IsCurrent   bool
	Start       time.Time
	End         time.Time
	ClassGroups []ClassGroup `gorm:"many2many:school_year_class_group;"`
}
