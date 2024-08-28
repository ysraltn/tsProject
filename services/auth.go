package services

import (
	"database/sql"
	"errors"
	"tsProject/models"
	"tsProject/utils"
)

type AuthService struct {
	userRepositories models.IUserRepository
	logger           models.ILogger
}

func NewAuthService(userRepositories models.IUserRepository, logger models.ILogger) models.IAuthService {
	return &AuthService{
		userRepositories: userRepositories,
		logger:           logger,
	}
}

func (s *AuthService) Login(username, password string) (string, string, error) {
	user, err := s.userRepositories.GetByUsername(username)
	if errors.Is(err, sql.ErrNoRows) {
		return "", "", errors.New("user not found")
	}
	if err != nil {
		return "", "", err
	}
	if user == nil {
		return "", "", errors.New("user not found")
	}
	if !utils.CheckPasswordHash(password, user.Password) {
		return "", "", errors.New("invalid credentials")
	}
	token, err := utils.GenerateJWT(username, user.Role)
	if err != nil {
		return "", "", errors.New("token generation error")
	}
	s.logger.UserInfoLog(models.UserInfoLog{
		Message:  "user logged in",
		Event:    "login",
		UserID:  user.ID,
		UserName: user.Username,
		Role:     user.Role,
	})
	return token, user.Role, nil
}

func (s *AuthService) Register(username, password string) error {
	_, err := s.userRepositories.GetByUsername(username)
	if errors.Is(err, sql.ErrNoRows) {
		hashedPassword, err := utils.HashPassword(password)
		if err != nil {
			return err
		}
		user := models.User{
			Username: username,
			Password: hashedPassword,
			Role:     "user",
		}
		s.userRepositories.Add(&user)
		s.logger.UserInfoLog(models.UserInfoLog{
			Message:  "user registered",
			Event:    "register",
			UserName: user.Username,
			Role:     user.Role,
		})
		return nil
	}
	return errors.New("username taken")
}
