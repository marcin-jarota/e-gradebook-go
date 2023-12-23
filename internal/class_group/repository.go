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

func (r *GormClassGroupRepository) AddTeacherWithSubject(classGroupID int, teacherID int, subjectID int) error {
	tx := r.db.Begin()
	var classGroup domain.ClassGroup
	var teacher domain.Teacher
	var subject domain.Subject

	if err := tx.Find(&classGroup, classGroupID).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Find(&teacher, teacherID).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Find(&subject, subjectID).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Create(&domain.SubjectTeacherClass{
		Teacher:    teacher,
		ClassGroup: classGroup,
		Subject:    subject,
	}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (r *GormClassGroupRepository) GetTeachersWithSubject(classGroupID int) ([]struct {
	Teacher domain.Teacher
	Subject domain.Subject
}, error) {
	var output []domain.SubjectTeacherClass

	var mapping []struct {
		Teacher domain.Teacher
		Subject domain.Subject
	}

	if err := r.db.Preload("Teacher.User").Preload("Subject").Find(&output, "class_group_id = ?", classGroupID).Error; err != nil {
		return nil, err
	}

	for _, entity := range output {
		mapping = append(mapping, struct {
			Teacher domain.Teacher
			Subject domain.Subject
		}{Teacher: entity.Teacher, Subject: entity.Subject})
	}

	return mapping, nil
}

func (r *GormClassGroupRepository) GetOneByID(classGroupID int) (domain.ClassGroup, error) {
	var classGroup domain.ClassGroup

	if err := r.db.Preload("SchoolYears").Find(&classGroup, "id = ?", classGroupID).Error; err != nil {
		return classGroup, err
	}

	return classGroup, nil
}

func (r *GormClassGroupRepository) Delete(id int) error {
	var classGroup domain.ClassGroup
	tx := r.db.Begin()

	if err := tx.Preload("Students").Find(&classGroup, "id = ?", id).Error; err != nil {
		tx.Rollback()
		return err
	}

	if len(classGroup.Students) != 0 {
		for _, s := range classGroup.Students {
			s.ClassGroupID = nil
			err := tx.Save(&s).Error

			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	if err := tx.Delete(&domain.SubjectTeacherClass{}, "class_group_id = ?", id).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Delete(&domain.Lesson{}, "class_group_id = ?", id).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Model(&classGroup).Association("SchoolYears").Clear(); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Delete(&classGroup).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r *GormClassGroupRepository) AddSubject(classGroupID int, subjectID int) error {
	// var subject domain.Subject
	// var classGroup domain.ClassGroup
	//
	// if err := r.db.First(&subject, subjectID).Error; err != nil {
	// 	return err
	// }

	// if err := r.db.Preload("Subjects").First(&classGroup, classGroupID).Error; err != nil {
	// 	return err
	// }

	return nil
	// for _, subject := range classGroup.Subjects {
	// 	if int(subject.ID) == subjectID {
	// 		return nil
	// 	}
	// }
	//
	// return r.db.Model(&classGroup).Association("Subjects").Append(&subject)
}

func (r *GormClassGroupRepository) AddTeacher(classGroupID int, teacherID int) error {
	return nil
	// var teacher domain.Teacher
	// var classGroup domain.ClassGroup
	//
	// if err := r.db.First(&teacher, teacherID).Error; err != nil {
	// 	return err
	// }
	//
	// if err := r.db.Preload("Teachers").First(&classGroup, classGroupID).Error; err != nil {
	// 	return err
	// }
	//
	// for _, teacher := range classGroup.Teachers {
	// 	if int(teacher.ID) == teacherID {
	// 		return nil
	// 	}
	// }
	//
	// return r.db.Model(&classGroup).Association("Teachers").Append(&teacher)
}

func (r *GormClassGroupRepository) GetTeachers(classGroupID int) ([]domain.Teacher, error) {
	// var classGroup domain.ClassGroup
	//
	// if err := r.db.Preload("Teachers.User").First(&classGroup, classGroupID).Error; err != nil {
	// 	return nil, err
	// }

	return nil, nil
}
