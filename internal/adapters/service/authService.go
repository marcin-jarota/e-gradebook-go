package service

import (
	"e-student/internal/app/domain"
	"e-student/internal/app/ports"
	"e-student/internal/common"
	"errors"
	"log"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
)

type AuthService struct {
	userRepo       ports.UserRepository
	sessionStorage ports.SessionStorage
	cfg            *common.Config
}

type claims struct {
	jwt.RegisteredClaims
	ID uint
}

func NewAuthService(userRepo ports.UserRepository, sessionStorage ports.SessionStorage, cfg *common.Config) *AuthService {
	return &AuthService{
		sessionStorage: sessionStorage,
		userRepo:       userRepo,
		cfg:            cfg,
	}
}

func (s *AuthService) IsLoggedIn(token string) (bool, *domain.User) {
	var userClaims claims

	parsed, err := jwt.ParseWithClaims(token, &userClaims, func(t *jwt.Token) (interface{}, error) {
		return []byte(s.cfg.Secret), nil
	})

	if err != nil {
		log.Println("Eerrror: ", err)
		return false, nil
	}

	if !parsed.Valid {
		log.Printf("Token %s invalid", token)

		return false, nil
	}

	exists, err := s.sessionStorage.Get(strconv.Itoa(int(userClaims.ID)))

	if err != nil {
		log.Println(err)
		return false, nil
	}

	if exists == nil {
		log.Printf("Token %s does not exists in storage", token)

		return false, nil
	}

	user, err := s.userRepo.GetOne(int(userClaims.ID))

	if err != nil {
		return false, nil
	}

	return true, user

}

func (s *AuthService) Login(email string, password string) (string, error) {
	user, err := s.userRepo.GetOneByEmail(email)

	if err != nil {
		return "", errors.New("password or email not found")
	}

	err = user.PaswordMatches(password)

	if err != nil {
		return "", err
	}

	if !user.Active {
		return "", errors.New("your account is not activated")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims{
		RegisteredClaims: jwt.RegisteredClaims{}, ID: user.ID,
	})

	signedToken, err := token.SignedString([]byte(s.cfg.Secret))

	if err != nil {
		return "", err
	}

	err = s.sessionStorage.Set(strconv.Itoa(int(user.ID)), signedToken)

	if err != nil {
		return "", errors.New("could not login")
	}

	return signedToken, nil
}
