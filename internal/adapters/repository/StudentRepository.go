package repository

import (
	"e-student/internal/app/domain"
	"errors"
	"log"

	"gorm.io/gorm"
)

var (
	ErrCouldNotCreateStudent = errors.New("could not create student")
)

type StudentRepository struct {
	db *gorm.DB
}

func NewStudentRepository(db *gorm.DB) *StudentRepository {
	if err := db.AutoMigrate(&domain.Student{}); err != nil {
		log.Panic(err)
	}

	return &StudentRepository{
		db: db,
	}
}

func (r *StudentRepository) GetAll() ([]*domain.Student, error) {
	var students []*domain.Student

	res := r.db.Preload("User").Preload("Marks").Find(&students)
	if res.Error != nil {
		return students, res.Error
	}

	return students, nil
}

func (r *StudentRepository) AddStudent(student *domain.Student) error {
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
