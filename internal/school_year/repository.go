package schoolyear

import (
	"e-student/internal/app/domain"
	"time"

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

func (r *GormSchoolYearRepostiory) AddSchoolYear(name string, start time.Time, end time.Time) error {
	return r.db.Exec("SELECT open_new_school_year(?, ?, ?)", name, start, end).Error
}
