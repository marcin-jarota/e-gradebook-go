package domain

import (
	"time"

	"gorm.io/gorm"
)

// Model
type Student struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	UserID    uint           `json:"-"`
	User      User           `json:"user"`
	Marks     []Mark         `json:"marks"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

// Transport response
type StudentResponse struct {
	ID        uint           `json:"id"`
	Name      string         `json:"name"`
	Surname   string         `json:"surname"`
	FullName  string         `json:"fullName"`
	Email     string         `json:"email"`
	Active    bool           `json:"active"`
	Marks     []Mark         `json:"marks"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
