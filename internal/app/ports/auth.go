package ports

import (
	"e-student/internal/app/domain"
	"time"
)

type AuthService interface {
	Login(email string, password string) (string, error)
	Logout(userId int) error
	IsLoggedIn(token string) (bool, *domain.User)
	IsTokenValid(token string) (isValid bool, userID int)
	GeneratePasswordToken(email string, expiresIn time.Duration) (string, error)
	IsGenerateTokenValid(token string) (string, uint, bool)
}
