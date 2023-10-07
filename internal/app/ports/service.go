package ports

import "e-student/internal/app/domain"

type AuthService interface {
	Login(email string, password string) (string, error)
	IsLoggedIn(token string) (bool, *domain.User)
}

type StudentService interface {
	GetAll() ([]*domain.StudentResponse, error)
}
