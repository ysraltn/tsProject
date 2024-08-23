package models

type IAuthService interface {
	Login(username, password string) (string, string, error)
	Register(username, password string) error
}
