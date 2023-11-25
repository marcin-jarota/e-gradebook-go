package domain

import (
	"gorm.io/gorm"
)

// Model
type Student struct {
	gorm.Model
	UserID       uint
	User         User `gorm:"foreignkey:UserID"`
	Marks        []Mark
	ClassGroupID *uint `gorm:"default:null"`
	ClassGroup   ClassGroup
}
