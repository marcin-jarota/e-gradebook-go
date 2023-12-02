package student

import (
	"e-student/internal/app/domain"
	"errors"

	"gorm.io/gorm"
)

var (
	ErrCouldNotCreateStudent = errors.New("could not create student")
)

type GormStudentRepository struct {
	db *gorm.DB
}

func NewGormStudentRepository(db *gorm.DB) *GormStudentRepository {
	return &GormStudentRepository{
		db,
	}
}

func (r *GormStudentRepository) GetAll() ([]domain.Student, error) {
	var students []domain.Student

	res := r.db.Preload("User").Preload("Marks").Find(&students)
	if res.Error != nil {
		return students, res.Error
	}

	return students, nil
}

func (r *GormStudentRepository) GetAllByClassGroup(classGroupID uint) ([]domain.Student, error) {
	var students []domain.Student

	if err := r.db.Preload("User").Preload("Marks").Where("class_group_id = ?", classGroupID).Find(&students).Error; err != nil {
		return nil, err
	}

	return students, nil
}

func (r *GormStudentRepository) AddStudent(student *domain.Student) error {
	return r.db.Create(student).Error
}

func (r *GormStudentRepository) GetMarks(studentID int) ([]domain.Mark, error) {
	var student domain.Student
	err := r.db.Preload("Marks").Preload("Marks.Subject").First(&student, studentID).Error
	return student.Marks, err
}

func (r *GormStudentRepository) SetClassGroup(studentID uint, classGroupID uint) error {
	var student domain.Student

	// Find the existing student by ID
	if err := r.db.First(&student, studentID).Error; err != nil {
		return err
	}

	// Update the student's ClassID to associate it with the desired class group
	student.ClassGroupID = &classGroupID

	// Save the updated student information to the database
	if err := r.db.Save(&student).Error; err != nil {
		return err
	}

	return nil
}

func (r *GormStudentRepository) RemoveClassGroup(studentID uint) error {
	student := domain.Student{}

	// Find the existing student by ID
	if err := r.db.First(&student, studentID).Error; err != nil {
		return err
	}

	var classGroupID *uint
	// Disassociate the student from the class group by setting ClassID to zero or nil
	student.ClassGroupID = classGroupID // or nil (depends on the type of ClassID field)

	// Save the updated student information to the database
	if err := r.db.Save(&student).Error; err != nil {
		return err
	}

	return nil
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

func (r *GormStudentRepository) GetByUserID(userID int) (domain.Student, error) {
	var student domain.Student
	err := r.db.First(&student, "user_id = ?", userID).Error
	return student, err
}
