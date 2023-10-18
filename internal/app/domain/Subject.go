package domain

import (
	"gorm.io/gorm"
)

type Subject struct {
	gorm.Model
	Name  string
	Marks []Mark
}
