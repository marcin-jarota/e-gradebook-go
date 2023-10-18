package mark

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

func (r *GormMarkRepository) GetMarksByStudent(studentId int) ([]*domain.Mark, error) {
	var marks []*domain.Mark

	res := r.db.Preload("Subject").Where("student_id = ?", studentId).Find(&marks)

	if res.Error != nil {
		return nil, res.Error
	}

	return marks, nil
}