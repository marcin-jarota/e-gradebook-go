package mark

import (
	"e-student/internal/app/domain"

	"gorm.io/gorm"
)

type gormMarkRepository struct {
	db *gorm.DB
}

func NewGormMarkRepository(db *gorm.DB) *gormMarkRepository {

	return &gormMarkRepository{
		db: db,
	}
}

func (r *gormMarkRepository) AddMark(mark domain.Mark) error {
	return r.db.Create(mark).Error
}

func (r *gormMarkRepository) GetByStudent(studentID int) ([]domain.Mark, error) {
	var marks []domain.Mark

	if err := r.db.Joins("Teacher.User").Joins("Subject").Where("student_id = ?", studentID).Find(&marks).Error; err != nil {
		return nil, err
	}

	return marks, nil
}

func (r *gormMarkRepository) GetByClassGroup(classGroupID int) ([]domain.Mark, error) {
	var marks []domain.Mark

	if err := r.db.
		Joins("JOIN students ON students.id = marks.student_id").
		Where("students.class_group_id = ?", classGroupID).
		Find(&marks).Error; err != nil {
		return nil, err
	}

	return marks, nil
}
