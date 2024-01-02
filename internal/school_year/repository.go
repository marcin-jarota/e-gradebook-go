package schoolyear

import (
	"e-student/internal/app/domain"

	"gorm.io/gorm"
)

type GormSchoolYearRepostiory struct {
	db *gorm.DB
}

func NewGormSchoolYearRepository(db *gorm.DB) *GormSchoolYearRepostiory {
	return &GormSchoolYearRepostiory{
		db: db,
	}
}

func (r *GormSchoolYearRepostiory) GetAll() ([]domain.SchoolYear, error) {
	var schoolYears []domain.SchoolYear
	if err := r.db.Preload("ClassGroups").Find(&schoolYears).Error; err != nil {
		return schoolYears, err
	}

	return schoolYears, nil
}
