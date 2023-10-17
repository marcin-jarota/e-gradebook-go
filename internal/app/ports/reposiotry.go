package ports

import "e-student/internal/app/domain"

type UserRepository interface {
	GetAll() ([]*domain.User, error)
	GetOne(id int) (*domain.User, error)
	AddUser(user *domain.User) error
	GetOneByEmail(email string) (*domain.User, error)
	// IsActive(user *domain.User) bool
}

type StudentRepository interface {
	GetAll() ([]*domain.Student, error)
	AddStudent(student *domain.Student) error
}

type MarkRepository interface {
	AddMark(mark *domain.Mark) error
	GetMarksByStudent(studentId int) ([]*domain.Mark, error)
}

type SubjectRepository interface {
	AddSubject(subject *domain.Subject) error
	GetAll() ([]*domain.Subject, error)
}

// type ClassGroupRepository interface {
// 	AddClassGroup(classGroup *domain.ClassGroup) error
// }
