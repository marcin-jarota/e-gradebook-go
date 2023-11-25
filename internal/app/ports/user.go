package ports

import "e-student/internal/app/domain"

type (
	UserRepository interface {
		GetAll() ([]*domain.User, error)
		GetOne(id int) (*domain.User, error)
		AddUser(user *domain.User) error
		GetOneByEmail(email string) (*domain.User, error)
		ExistsByEmail(email string) (bool, error)
		Activate(userID uint) error
		Deactivate(userID uint) error
		SetPassword(email string, password string) error
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

	UserCreatePayload struct {
		Name    string `json:"name"`
		Surname string `json:"surname"`
		Email   string `json:"email"`
	}

	UserOutput struct {
		ID            uint            `json:"id"`
		Name          string          `json:"name"`
		Surname       string          `json:"surname"`
		Email         string          `json:"email"`
		Role          domain.UserRole `json:"role,omitempty"`
		IsActive      bool            `json:"isActive,omitempty"`
		SessionActive bool            `json:"sessionActive,omitempty"`
	}

	SetupPasswordPayload struct {
		Password        string `json:"password"`
		PasswordConfirm string `json:"passwordConfirm"`
	}

	UserService interface {
		GetAll() ([]*UserOutput, error)
		AddAdmin(user UserCreatePayload) error
		Activate(userID uint) error
		Deactivate(userID uint) error
		DestroySession(userID uint) error
		SetupPassword(email string, password string, passwordConfirm string) error
	}
)
