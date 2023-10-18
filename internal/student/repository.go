package student

import (
	"e-student/internal/app/domain"
	"errors"
	"log"

	"gorm.io/gorm"
)

var (
	ErrCouldNotCreateStudent = errors.New("could not create student")
)

type GormStudentRepository struct {
	db *gorm.DB
}

func NewGormStudentRepository(db *gorm.DB) *GormStudentRepository {
	if err := db.AutoMigrate(&domain.Student{}); err != nil {
		log.Panic(err)
	}

	return &GormStudentRepository{
		db,
	}
}

func (r *GormStudentRepository) GetAll() ([]*domain.Student, error) {
	var students []*domain.Student

	res := r.db.Preload("User").Preload("Marks").Find(&students)
	if res.Error != nil {
		return students, res.Error
	}

	return students, nil
}

func (r *GormStudentRepository) AddStudent(student *domain.Student) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		hash, err := student.User.GeneratePassword(student.User.Password)
		if err != nil {
			return errors.Join(err, ErrCouldNotCreateStudent)
		}

		student.User.Password = hash

		if err := tx.Create(&student.User).Error; err != nil {
			return err
		}

		student.UserID = student.User.ID

		if err := tx.Create(&student).Error; err != nil {
			return err
		}

		return nil

	})
}

func (r *GormStudentRepository) GetMarks(studentID int) ([]domain.Mark, error) {
	var student domain.Student
	err := r.db.Preload("Marks").Preload("Marks.Subject").First(&student, studentID).Error
	return student.Marks, err
}

func (r *GormStudentRepository) ExistsByEmail(email string) (bool, error) {
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
