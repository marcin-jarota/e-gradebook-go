package domain

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRole string

const (
	AdminRole   UserRole = "admin"
	StudentRole UserRole = "student"
)

type User struct {
	gorm.Model
	Name     string
	Surname  string
	Email    string
	Role     UserRole `gorm:"type:user_role"`
	Password string
	Active   bool
}

type SessionUser struct {
	ID       uint     `json:"id"`
	Name     string   `json:"name"`
	Surname  string   `json:"surname"`
	FullName string   `json:"fullName"`
	Email    string   `json:"email"`
	Role     UserRole `json:"role"`
}

func (u *User) GetFullName() string {
	return fmt.Sprintf("%s %s", u.Name, u.Surname)
}

func (u *User) GeneratePassword(plainText string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(plainText), bcrypt.DefaultCost)
	return string(bytes), err
}

func (u *User) PaswordMatches(plainText string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plainText))
	return err
}

func (u *User) IsAdmin() bool {
	return u.Role == AdminRole
}

func (u *User) IsStudent() bool {
	return u.Role == StudentRole
}
