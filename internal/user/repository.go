package user

import (
	"e-student/internal/app/domain"

	"gorm.io/gorm"
)

type GormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) *GormUserRepository {
	return &GormUserRepository{
		db: db,
	}
}

func (u *GormUserRepository) GetOneByEmail(email string) (*domain.User, error) {
	var user domain.User
	res := u.db.First(&user, "email = ?", email)

	if res.Error != nil {
		return &user, res.Error
	}

	return &user, nil
}

func (u *GormUserRepository) GetOne(id int) (*domain.User, error) {
	var user domain.User

	res := u.db.First(&user, id)

	if res.Error != nil {
		return &user, res.Error
	}

	return &user, nil
}

func (u *GormUserRepository) GetAll() ([]*domain.User, error) {
	var users []*domain.User

	res := u.db.Find(&users)

	if res.Error != nil {
		return users, res.Error
	}

	return users, nil
}

func (u *GormUserRepository) AddUser(user *domain.User) error {
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

func (u *GormUserRepository) Activate(userID uint) error {
	return u.db.Model(&domain.User{}).Where("id = ?", userID).Update("active", true).Error
}

func (u *GormUserRepository) Deactivate(userID uint) error {
	return u.db.Model(&domain.User{}).Where("id = ?", userID).Update("active", false).Error
}

func (r *GormUserRepository) ExistsByEmail(email string) (bool, error) {
	var user domain.User
	err := r.db.First(&user, "email = ?", email).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}

		return false, err
	}

	return true, nil
}
