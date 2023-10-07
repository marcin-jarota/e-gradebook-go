package domain

import (
	"time"

	"gorm.io/gorm"
)

type Mark struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Value     float32        `json:"value"`
	SubjectID uint           `json:"-"`
	StudentID uint           `json:"-"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
