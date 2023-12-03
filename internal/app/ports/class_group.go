package ports

import "e-student/internal/app/domain"

type (
	ClassGroupRepository interface {
		GetAll() ([]domain.ClassGroup, error)
		GetOneByID(id int) (domain.ClassGroup, error)
		AddClassGroup(classGroup *domain.ClassGroup) error
		AddSubject(classGroupID int, subjectID int) error
		// AddStudent(studentID uint, classGroupID uint) error
		// RemoveStudent(studentID uint, classGroupID uint) error
		// DeleteByID(classGroupID uint) error
	}

	ClassGroupOutput struct {
		ID            int    `json:"id"`
		Name          string `json:"name"`
		StudentsCount int    `json:"studentsCount"`
	}

	AddClassGroupInput struct {
		Name string `json:"name"`
	}

	AddStudentToClassGroupPayload struct {
		StudentID int `json:"studentID"`
	}

	AddSubjectToClassGroupPayload struct {
		SubjectID    int `json:"subjectID"`
		ClassGroupID int `json:"classGroupID"`
	}

	ClassGroupService interface {
		GetAll() ([]ClassGroupOutput, error)
		GetOneByID(id int) (ClassGroupOutput, error)
		AddClassGroup(input AddClassGroupInput) error
		AddSubject(input AddSubjectToClassGroupPayload) error
	}
)
