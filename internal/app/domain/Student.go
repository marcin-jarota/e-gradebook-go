package domain

import (
	"gorm.io/gorm"
)

// Model
type Student struct {
	gorm.Model
	UserID uint
	User   User
	Marks  []Mark
}
