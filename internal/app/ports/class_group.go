package ports

import "e-student/internal/app/domain"

type (
	ClassGroupRepository interface {
		GetAll() ([]domain.ClassGroup, error)
		GetOneByID(id int) (domain.ClassGroup, error)
		AddClassGroup(classGroup *domain.ClassGroup) error
		AddSubject(classGroupID int, subjectID int) error
		AddTeacher(classGroupID int, teacherID int) error
		GetTeachers(classGroupID int) ([]domain.Teacher, error)
		GetTeachersWithSubject(classGropID int) ([]struct {
			Teacher domain.Teacher
			Subject domain.Subject
		}, error)
		Delete(classGroupID int) error
		AddTeacherWithSubject(classGroupID int, teacherID int, subjectID int) error
	}

	ClassGroupOutput struct {
		ID            int               `json:"id"`
		Name          string            `json:"name"`
		StudentsCount int               `json:"studentsCount"`
		EducationYear int               `json:"educationYear,omitempty"`
		SchoolYears   []SchoolYearBasic `json:"schoolYears,omitempty"`
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

	TeacherClassGroup struct {
		TeacherID    int `json:"teacherID"`
		ClassGroupID int `json:"classGroupID"`
	}

	ClassGroupTeacher struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		Surname string `json:"surname"`
		Email   string `json:"email"`
	}

	TeacherSubjectClassgroupID struct {
		ClassGroupID int `json:"classGroupID"`
		TeacherID    int `json:"teacherID"`
		SubjectID    int `json:"subjectID"`
	}

	TeacherSubject struct {
		Teacher TeacherBaseOutput `json:"teacher"`
		Subject SubjectBaseOutput `json:"subject"`
	}
	ClassGroupService interface {
		GetAll() ([]ClassGroupOutput, error)
		GetOneByID(id int) (ClassGroupOutput, error)
		AddClassGroup(input AddClassGroupInput) error
		AddSubject(input AddSubjectToClassGroupPayload) error
		AddTeacher(input TeacherClassGroup) error
		GetTeachers(classGroupID int) ([]ClassGroupTeacher, error)
		AddTeacherWithSubject(input TeacherSubjectClassgroupID) error
		GetTeachersWithSubject(classGroupID int) ([]TeacherSubject, error)
		Delete(id int) error
	}
)
