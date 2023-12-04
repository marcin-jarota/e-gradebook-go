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

func (r *GormClassGroupRepository) GetOneByID(classGroupID int) (domain.ClassGroup, error) {
	var classGroup domain.ClassGroup

	if err := r.db.Find(&classGroup, "id = ?", classGroupID).Error; err != nil {
		return classGroup, err
	}

	return classGroup, nil
}

func (r *GormClassGroupRepository) AddSubject(classGroupID int, subjectID int) error {
	var subject domain.Subject
	var classGroup domain.ClassGroup

	if err := r.db.First(&subject, subjectID).Error; err != nil {
		return err
	}

	if err := r.db.Preload("Subjects").First(&classGroup, classGroupID).Error; err != nil {
		return err
	}

	for _, subject := range classGroup.Subjects {
		if int(subject.ID) == subjectID {
			return nil
		}
	}

	return r.db.Model(&classGroup).Association("Subjects").Append(&subject)
}

func (r *GormClassGroupRepository) AddTeacher(classGroupID int, teacherID int) error {
	var teacher domain.Teacher
	var classGroup domain.ClassGroup

	if err := r.db.First(&teacher, teacherID).Error; err != nil {
		return err
	}

	if err := r.db.Preload("Teachers").First(&classGroup, classGroupID).Error; err != nil {
		return err
	}

	for _, teacher := range classGroup.Teachers {
		if int(teacher.ID) == teacherID {
			return nil
		}
	}

	return r.db.Model(&classGroup).Association("Teachers").Append(&teacher)
}

func (r *GormClassGroupRepository) GetTeachers(classGroupID int) ([]domain.Teacher, error) {
	var classGroup domain.ClassGroup

	if err := r.db.Preload("Teachers.User").First(&classGroup, classGroupID).Error; err != nil {
		return nil, err
	}

	return classGroup.Teachers, nil
}
