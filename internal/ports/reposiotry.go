package ports

import "e-student/internal/domain"

type UserRepository interface {
	// GetAll() ([]*domain.User, error)
	GetOne(id int) (*domain.User, error)
	AddUser(user *domain.User) error
	// GetOneByEmail(email string) (*domain.User, error)
	// IsActive(user *domain.User) bool
}
