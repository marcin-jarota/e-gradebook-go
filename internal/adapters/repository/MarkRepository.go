package repository

import (
	"e-student/internal/app/domain"
	"log"

	"gorm.io/gorm"
)

type MarkRepository struct {
	db *gorm.DB
}

func NewMarkRepository(db *gorm.DB) *MarkRepository {
	err := db.AutoMigrate(&domain.Mark{})

	if err != nil {
		log.Panic(err)
	}

	return &MarkRepository{
		db: db,
	}
}

func (r *MarkRepository) AddMark(mark *domain.Mark) error {
	return r.db.Create(mark).Error
}
