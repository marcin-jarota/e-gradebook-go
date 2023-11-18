package ports

import "e-student/internal/app/domain"

type (
	ClassGroupRepository interface {
		GetAll() ([]domain.ClassGroup, error)
		AddClassGroup(classGroup *domain.ClassGroup) error
		GetStudents(classGroupID uint) ([]domain.Student, error)
		// AddStudent(studentID uint, classGroupID uint) error
		// RemoveStudent(studentID uint, classGroupID uint) error
		// DeleteByID(classGroupID uint) error
	}

	ListStudentsOutput struct {
		Name    string `json:"name"`
		Surname string `json:"surname"`
	}

	ListClassGroupsOutput struct {
		ID            int    `json:"id"`
		Name          string `json:"name"`
		StudentsCount int    `json:"studentsCount"`
	}

	AddClassGroupInput struct {
		Name string `json:"name"`
	}

	ClassGroupService interface {
		GetAll() ([]*ListClassGroupsOutput, error)
		ListStudents(classGroupID uint) ([]*ListStudentsOutput, error)
		AddClassGroup(input *AddClassGroupInput) error
	}
)
