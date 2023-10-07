package domain

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userRole string

const (
	admin   userRole = "admin"
	student userRole = "student"
)

type User struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name"`
	Surname   string         `json:"surname"`
	Email     string         `json:"email"`
	Role      userRole       `json:"role" gorm:"type:user_role"`
	Password  string         `json:"password"`
	Active    bool           `json:"active"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type SessionUser struct {
	ID       uint     `json:"id"`
	Name     string   `json:"name"`
	Surname  string   `json:"surname"`
	FullName string   `json:"fullName"`
	Email    string   `json:"email"`
	Role     userRole `json:"role"`
	Active   bool     `json:"active"`
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
	return u.Role == admin
}
