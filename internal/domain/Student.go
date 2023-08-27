package domain

import (
	"time"

	"gorm.io/gorm"
)

type Student struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	UserID    uint           `json:"-"`
	User      User           `json:"user"`
	Marks     []Mark         `json:"marks"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
