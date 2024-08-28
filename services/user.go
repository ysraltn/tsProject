package services

import (
	"tsProject/models"
	"tsProject/utils"
)

type UserService struct {
	userRepositories models.IUserRepository
}

func NewUserService(userRepositories models.IUserRepository) models.IUserService {
	return &UserService{
		userRepositories: userRepositories,
	}
}

func (s *UserService) Add(username, password, role string) error {
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}
	user := &models.User{
		Username: username,
		Password: hashedPassword,
		Role:     role,
	}
	err = s.userRepositories.Add(user)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) Delete(id int) error {
	err := s.userRepositories.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) GetAll() ([]models.UserResponse, error) {
	users, err := s.userRepositories.GetAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UserService) GetByID(id int) (*models.User, error) {
	var user *models.User
	user, err := s.userRepositories.GetByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
