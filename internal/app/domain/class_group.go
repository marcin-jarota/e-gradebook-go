package domain

import "gorm.io/gorm"

type ClassGroup struct {
	gorm.Model
	Name     string    `json:"name"`
	Students []Student `json:"students"`
}
