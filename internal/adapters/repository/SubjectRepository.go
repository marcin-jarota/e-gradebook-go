package repository

import (
	"e-student/internal/domain"

	"gorm.io/gorm"
)

type SubjectRepository struct {
	db *gorm.DB
}

func NewSubjectRepository(db *gorm.DB) *SubjectRepository {
	db.AutoMigrate(&domain.Subject{})

	return &SubjectRepository{
		db: db,
	}
}

func (r *SubjectRepository) AddSubject(subject *domain.Subject) error {
	return r.db.Create(subject).Error
}

func (r *SubjectRepository) GetAll() ([]*domain.Subject, error) {
	var subjects []*domain.Subject

	res := r.db.Find(&subjects)
	if res.Error != nil {
		return subjects, res.Error
	}

	return subjects, nil
}
