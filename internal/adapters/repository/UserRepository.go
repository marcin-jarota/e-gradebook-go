package repository

import (
	"e-student/internal/domain"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func (u *UserRepository) GetOne(id int) (*domain.User, error) {
	var user domain.User

	res := u.db.First(&user)

	if res.Error != nil {
		return &user, res.Error
	}

	return &user, nil
}

func (u *UserRepository) AddUser(user *domain.User) error {
	hash, err := user.GeneratePassword(user.Password)

	user.Password = hash
	if err != nil {
		return err
	}

	res := u.db.Create(user)

	if res.Error != nil {
		return err
	}

	return nil
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}
