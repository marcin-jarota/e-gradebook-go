package subject

import (
	"e-student/internal/app/domain"

	"gorm.io/gorm"
)

type GormSubjectRepository struct {
	db *gorm.DB
}

func NewGormSubjectRepository(db *gorm.DB) *GormSubjectRepository {
	return &GormSubjectRepository{
		db: db,
	}
}

func (r *GormSubjectRepository) AddSubject(subject *domain.Subject) error {
	return r.db.Create(subject).Error
}

func (r *GormSubjectRepository) GetAll() ([]*domain.Subject, error) {
	var subjects []*domain.Subject

	res := r.db.Find(&subjects)
	if res.Error != nil {
		return subjects, res.Error
	}

	return subjects, nil
}

func (r *GormSubjectRepository) DeleteByID(id uint) error {
	if err := r.db.Where("subject_id = ?", id).Delete(&domain.Mark{}).Error; err != nil {
		return err
	}

	if err := r.db.Delete(&domain.Subject{}, id).Error; err != nil {
		return err
	}

	return nil
}

func (r *GormSubjectRepository) GetOneByName(name string) (*domain.Subject, error) {
	var subject domain.Subject
	err := r.db.First(&subject, "name = ?", name).Error
	return &subject, err
}

func (r *GormSubjectRepository) Exists(name string) (bool, error) {
	_, err := r.GetOneByName(name)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}

		return false, err
	}

	return true, nil
}
