package services

import "tsProject/models"

type UserService struct {
	userRepositories models.IUserRepository
}

func NewUserService(userRepositories models.IUserRepository) models.IUserService {
	return &UserService{
		userRepositories: userRepositories,
	}
}

func (s *UserService) Add(username, password string) error {
	user := &models.User{
		Username: username,
		Password: password,
	}
	err := s.userRepositories.Add(user)
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

func (s *UserService) GetAll() ([]models.User, error) {
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
