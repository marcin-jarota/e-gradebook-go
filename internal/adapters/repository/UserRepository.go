package repository

import (
	"e-student/internal/app/domain"
	"log"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	if err := db.AutoMigrate(&domain.User{}); err != nil {
		log.Panic(err)
	}

	return &UserRepository{
		db: db,
	}
}

func (u *UserRepository) GetOneByEmail(email string) (*domain.User, error) {
	var user domain.User
	res := u.db.First(&user, "email = ?", email)

	if res.Error != nil {
		return &user, res.Error
	}

	return &user, nil
}
func (u *UserRepository) GetOne(id int) (*domain.User, error) {
	var user domain.User

	res := u.db.First(&user, id)

	if res.Error != nil {
		return &user, res.Error
	}

	return &user, nil
}

func (u *UserRepository) GetAll() ([]*domain.User, error) {
	var users []*domain.User

	res := u.db.Find(&users)

	if res.Error != nil {
		return users, res.Error
	}

	return users, nil
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
