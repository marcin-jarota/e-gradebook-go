package repository

import (
	"e-student/internal/app/domain"
	"log"

	"gorm.io/gorm"
)

type GormMarkRepository struct {
	db *gorm.DB
}

func NewGormMarkRepository(db *gorm.DB) *GormMarkRepository {
	err := db.AutoMigrate(&domain.Mark{})

	if err != nil {
		log.Panic(err)
	}

	return &GormMarkRepository{
		db: db,
	}
}

func (r *GormMarkRepository) AddMark(mark *domain.Mark) error {
	return r.db.Create(mark).Error
}
