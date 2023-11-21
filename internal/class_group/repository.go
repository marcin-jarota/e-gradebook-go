package classgroup

import (
	"e-student/internal/app/domain"

	"gorm.io/gorm"
)

type GormClassGroupRepository struct {
	db *gorm.DB
}

func NewClassGroupRepository(db *gorm.DB) *GormClassGroupRepository {
	return &GormClassGroupRepository{
		db: db,
	}
}

func (r *GormClassGroupRepository) GetAll() ([]domain.ClassGroup, error) {
	var classGroups []domain.ClassGroup

	err := r.db.Preload("Students").Find(&classGroups).Error

	if err != nil {
		return nil, err
	}

	return classGroups, nil
}

func (r *GormClassGroupRepository) AddClassGroup(classGroup *domain.ClassGroup) error {
	return r.db.Create(classGroup).Error
}
