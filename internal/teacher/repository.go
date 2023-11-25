package teacher

import (
	"e-student/internal/app/domain"

	"gorm.io/gorm"
)

type GormTeacherRepository struct {
	db *gorm.DB
}

func NewTeacherRepository(db *gorm.DB) *GormTeacherRepository {
	return &GormTeacherRepository{
		db: db,
	}
}

func (r *GormTeacherRepository) AddTeacher(teacher domain.Teacher) error {
	return r.db.Create(&teacher).Error
}

func (r *GormTeacherRepository) ExistsByEmail(email string) (bool, error) {
	var user domain.User
	err := r.db.First(&user, "email = ?", email).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}

		return false, err
	}

	return true, nil
}
