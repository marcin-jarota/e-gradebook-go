package user

import (
	"e-student/internal/app/ports"
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
