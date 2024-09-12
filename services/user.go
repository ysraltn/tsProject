package services

import (
	"database/sql"
	"errors"
	"tsProject/models"
	"tsProject/utils"
)

type UserService struct {
	userRepositories models.IUserRepository
	logger           models.ILogger
}

func NewUserService(userRepositories models.IUserRepository, logger models.ILogger) models.IUserService {
	return &UserService{
		userRepositories: userRepositories,
		logger:           logger,
	}
}

func (s *UserService) Add(username, password, role, name, surname string) error {
	_, err := s.userRepositories.GetByUsername(username)
	if errors.Is(err, sql.ErrNoRows) {
		hashedPassword, err := utils.HashPassword(password)
		if err != nil {
			return err
		}
		user := &models.User{
			Username: username,
			Password: hashedPassword,
			Role:     role,
			Name:     name,
			Surname:  surname,
		}
		err = s.userRepositories.Add(user)
		if err != nil {
			s.logger.Error(models.ErrorLog{
				Message: "error adding user",
				Error:   err.Error(),
				Context: map[string]interface{}{
					"username": username,
					"role":     role,
				},
			})
			return err
		}
		s.logger.UserInfoLog(models.UserInfoLog{
			Message:  "user added",
			Event:    "user_added",
			UserID:   user.ID,
			UserName: user.Username,
			Role:     user.Role,
		})
		return nil

	}
	return errors.New("username taken")
}

func (s *UserService) Delete(id int) error {
	err := s.userRepositories.Delete(id)
	if err != nil {
		s.logger.Error(models.ErrorLog{
			Message: "error deleting user",
			Error:   err.Error(),
			Context: map[string]interface{}{
				"user_id": id,
			},
		})
		return err
	}
	s.logger.Info(models.InfoLog{
		Message: "user deleted",
		Event:   "user_deleted",
		UserID:  id,
	})
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

func (s *UserService) GetByUsername(username string) (*models.User, error) {
	var user *models.User
	user, err := s.userRepositories.GetByUsername(username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) GetAllForUsers() ([]models.UserForUsers, error) {
	users, err := s.userRepositories.GetAllForUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}
