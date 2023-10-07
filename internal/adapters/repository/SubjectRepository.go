package repository

import (
	"e-student/internal/app/domain"

	"gorm.io/gorm"
)

type GormSubjectRepository struct {
	db *gorm.DB
}

func NewGormSubjectRepository(db *gorm.DB) *GormSubjectRepository {
	db.AutoMigrate(&domain.Subject{})

	return &GormSubjectRepository{
		db: db,
	}
}

func (r *GormSubjectRepository) AddSubject(subject *domain.Subject) error {
	return r.db.Create(subject).Error
}

func (r *GormSubjectRepository) GetAll() ([]*domain.Subject, error) {
	var subjects []*domain.Subject

	res := r.db.Find(&subjects)
	if res.Error != nil {
		return subjects, res.Error
	}

	return subjects, nil
}
