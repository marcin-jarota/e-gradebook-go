package domain

import (
	"time"

	"gorm.io/gorm"
)

type Mark struct {
	ID        uint    `json:"id" gorm:"primaryKey"`
	Value     float32 `json:"value"`
	Subject   Subject `json:"subject" gorm:"foreignKey:SubjectID;references:ID"`
	Student   Student `json:"student" gorm:"foreignKey:StudentID;references:ID"`
	StudentID uint
	SubjectID uint
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
