package ports

import "e-student/internal/app/domain"

type (
	UserRepository interface {
		GetAll() ([]*domain.User, error)
		GetOne(id int) (*domain.User, error)
		AddUser(user *domain.User) error
		GetOneByEmail(email string) (*domain.User, error)
		// IsActive(user *domain.User) bool
	}

	SessionUser struct {
		ID       uint            `json:"id"`
		Name     string          `json:"name"`
		Surname  string          `json:"surname"`
		FullName string          `json:"fullName"`
		Email    string          `json:"email"`
		Role     domain.UserRole `json:"role"`
	}
)
