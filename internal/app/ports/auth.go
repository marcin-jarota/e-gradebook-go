package ports

import "e-student/internal/app/domain"

type AuthService interface {
	Login(email string, password string) (string, error)
	Logout(userId int) error
	IsLoggedIn(token string) (bool, *domain.User)
}
