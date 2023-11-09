package auth

import (
	"e-student/internal/app"
	"e-student/internal/app/domain"
	"e-student/internal/app/ports"
	"errors"
	"log"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
)

type AuthService struct {
	userRepo       ports.UserRepository
	sessionStorage ports.SessionStorage
	cfg            *app.Config
}

type claims struct {
	SessionUser ports.SessionUser `json:"sessionUser"`
	jwt.RegisteredClaims
}

func NewAuthService(userRepo ports.UserRepository, sessionStorage ports.SessionStorage, cfg *app.Config) *AuthService {
	return &AuthService{
		sessionStorage: sessionStorage,
		userRepo:       userRepo,
		cfg:            cfg,
	}
}

func (s *AuthService) IsTokenValid(token string) (bool, int) {
	var userClaims claims

	parsed, err := jwt.ParseWithClaims(token, &userClaims, func(t *jwt.Token) (interface{}, error) {
		return []byte(s.cfg.Secret), nil
	})

	if err != nil {
		log.Println("Eerrror: ", err)
		return false, 0
	}

	if !parsed.Valid {
		log.Printf("Token %s invalid", token)

		return false, 0
	}

	exists, err := s.sessionStorage.Get(strconv.Itoa(int(userClaims.SessionUser.ID)))

	if err != nil {
		log.Println("Storage errpr:", err)
		return false, 0
	}

	return exists != nil, int(userClaims.SessionUser.ID)
}

func (s *AuthService) IsLoggedIn(token string) (bool, *domain.User) {
	exists, userId := s.IsTokenValid(token)

	if !exists {
		log.Printf("User does not exisrts i nstorage")
		return false, nil
	}

	user, err := s.userRepo.GetOne(userId)

	if err != nil {
		log.Println("Could not get user from storage", err)
		return false, nil
	}

	return true, user

}

func (s *AuthService) Logout(userId int) error {
	return s.sessionStorage.Delete(strconv.Itoa(userId))
}

func (s *AuthService) Login(email string, password string) (string, error) {
	user, err := s.userRepo.GetOneByEmail(email)

	if err != nil {
		return "", errors.New("password or email not found")
	}

	err = user.PaswordMatches(password)

	if err != nil {
		return "", errors.New("password or email mismatch")
	}

	if !user.Active {
		return "", errors.New("your account is not activated")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims{
		RegisteredClaims: jwt.RegisteredClaims{},
		SessionUser: ports.SessionUser{
			ID:      user.ID,
			Name:    user.Name,
			Surname: user.Surname,
			Email:   user.Email,
			Role:    user.Role,
		},
	})

	signedToken, err := token.SignedString([]byte(s.cfg.Secret))

	if err != nil {
		return "", err
	}

	err = s.sessionStorage.Set(strconv.Itoa(int(user.ID)), signedToken)

	if err != nil {
		return "", errors.Join(err, errors.New("could not login"))
	}

	return signedToken, nil
}
