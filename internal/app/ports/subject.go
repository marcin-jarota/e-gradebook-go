package ports

import "e-student/internal/app/domain"

type (
	SubjectRepository interface {
		AddSubject(subject *domain.Subject) error
		GetAll() ([]*domain.Subject, error)
		GetOneByName(name string) (*domain.Subject, error)
		GetOneByID(id int) (domain.Subject, error)
		AddTeacher(teacherID int, subjectID int) error
		Exists(name string) (bool, error)
		DeleteByID(id uint) error
	}

	SubjectTeacher struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		Surname string `json:"surname"`
	}

	SubjectOutput struct {
		ID       uint             `json:"id"`
		Name     string           `json:"name"`
		Teachers []SubjectTeacher `json:"teachers"`
	}

	SubjectBaseOutput struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	}

	SubjectAddPayload struct {
		Name string `json:"name"`
	}

	TeacherSubjectID struct {
		TeacherID int `json:"teacherID"`
		SubjectID int `json:"subjectID"`
	}

	SubjectService interface {
		AddSubject(name string) error
		AddTeacher(payload TeacherSubjectID) error
		GetOneByID(id int) (*SubjectBaseOutput, error)
		GetAll() ([]*SubjectOutput, error)
		Delete(id uint) error
	}
)
