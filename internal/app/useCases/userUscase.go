package usecases

type UserUscase interface {
	Login(email string, password string) (string, error)
	IsLoggedIn(token string) bool
}
