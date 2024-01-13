package lesson

import (
	"e-student/internal/app/domain"

	"gorm.io/gorm"
)

type GormLessonRepository struct {
	db *gorm.DB
}

func NewLessonRepository(db *gorm.DB) *GormLessonRepository {
	return &GormLessonRepository{
		db: db,
	}
}

func (r *GormLessonRepository) Create(lesson *domain.Lesson) error {
	return r.db.Create(lesson).Error
}

func (r *GormLessonRepository) GetByClassGroup(classGroupID int) ([]*domain.Lesson, error) {
	var lessons []*domain.Lesson

	if err := r.db.Preload("Teacher.User").Preload("Subject").Find(&lessons, "class_group_id = ?", classGroupID).Error; err != nil {
		return nil, err
	}

	return lessons, nil
}
