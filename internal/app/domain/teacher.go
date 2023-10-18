package domain

import (
	"gorm.io/gorm"
)

type Teacher struct {
	gorm.Model
	User  User   `json:"user"`
	Marks []Mark `json:"marks"`
}
