package user

import (
	"e-student/internal/app/domain"
	"e-student/internal/app/ports"
	"errors"
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
			log.Println("[ERR]: Could not get session from storage", err)
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
		return errors.New("could not deactivate user")
	}

	return s.DestroySession(userID)
}

func (s *UserService) DestroySession(userID uint) error {
	return s.sessionStorage.Delete(strconv.Itoa(int(userID)))
}

func (s *UserService) AddAdmin(admin *ports.AdminCreatePayload) error {
	exists, err := s.repo.ExistsByEmail(admin.Email)

	if err != nil {
		return errors.New("could not verify if user exists")
	}

	if exists {
		return errors.New("user with this email exists")
	}

	return s.repo.AddUser(&domain.User{
		Name:     admin.Name,
		Surname:  admin.Surname,
		Email:    admin.Email,
		Password: admin.Password,
		Role:     domain.AdminRole,
		Active:   false,
	})
}
