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

func (r *GormTeacherRepository) GetTeacherByUserID(id int) (domain.Teacher, error) {
	var teacher domain.Teacher

	if err := r.db.Joins("User").First(&teacher, "user_id = ?", id).Error; err != nil {
		return teacher, err
	}

	return teacher, nil
}
func (r *GormTeacherRepository) GetAll() ([]domain.Teacher, error) {
	var teachers []domain.Teacher

	if err := r.db.Joins("User").Find(&teachers).Error; err != nil {
		return nil, err
	}

	return teachers, nil
}

func (r *GormTeacherRepository) GetAllByClassGroup(classGroupID uint) ([]domain.Teacher, error) {
	var teachers []domain.Teacher

	if err := r.db.Preload("User").Preload("Marks").Where("class_group_id = ?", classGroupID).Find(&teachers).Error; err != nil {
		return nil, err
	}

	return teachers, nil
}
