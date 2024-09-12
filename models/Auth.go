package models

type IAuthService interface {
	Login(username, password string) (int, string, string, error)
	Register(username, password, name, surname string) error
}
