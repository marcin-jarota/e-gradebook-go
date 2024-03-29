package user

import (
	"e-student/internal/app/domain"
	"e-student/internal/app/ports"
	"errors"
	"fmt"
	"log"
	"strconv"
)

type UserService struct {
	repo           ports.UserRepository
	sessionStorage ports.SessionStorage
}

func NewUserService(repo ports.UserRepository, sessionStorage ports.SessionStorage) *UserService {
	return &UserService{
		repo:           repo,
		sessionStorage: sessionStorage,
	}
}

func (s *UserService) GetAll() ([]*ports.UserOutput, error) {
	var userList []*ports.UserOutput

	users, err := s.repo.GetAll()

	if err != nil {
		return nil, err
	}

	for _, dbUser := range users {
		session := false
		exists, err := s.sessionStorage.Get(strconv.Itoa(int(dbUser.ID)))

		if err != nil {
			// TODO log this error to system
			log.Println("[ERR]: Could not get session from storage", err.Error())
		}

		if exists != nil {
			session = true
		}

		userList = append(userList, &ports.UserOutput{
			ID:            dbUser.ID,
			Name:          dbUser.Name,
			Surname:       dbUser.Surname,
			Email:         dbUser.Email,
			IsActive:      dbUser.Active,
			Role:          dbUser.Role,
			SessionActive: session,
		})
	}

	return userList, nil
}

func (s *UserService) Activate(userID uint) error {
	return s.repo.Activate(userID)
}

func (s *UserService) Deactivate(userID uint) error {
	err := s.repo.Deactivate(userID)

	if err != nil {
		return errors.New("deactivate.errorFallback")
	}

	return s.DestroySession(userID)
}

func (s *UserService) DestroySession(userID uint) error {
	return s.sessionStorage.Delete(strconv.Itoa(int(userID)))
}

func (s *UserService) SetupPassword(email string, password string, passwordConfirm string) error {
	if password == "" {
		return errors.New("password.errorEmpty")
	}

	if password != passwordConfirm {
		return errors.New("password.errorNotEqual")
	}

	fmt.Println(password, email)
	return s.repo.SetPassword(email, password)
}

func (s *UserService) AddAdmin(admin ports.UserCreatePayload) error {
	exists, err := s.repo.ExistsByEmail(admin.Email)

	if err != nil {
		return errors.New("user.add.errorFallback")
	}

	if exists {
		return errors.New("user.add.errorExists")
	}

	return s.repo.AddUser(&domain.User{
		Name:    admin.Name,
		Surname: admin.Surname,
		Email:   admin.Email,
		Role:    domain.AdminRole,
		Active:  false,
	})
}
