package repositories

import (
	"fmt"
	"tsProject/models"

	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Add(user *models.User) error {
	query := `
		INSERT INTO users (username, password, role, name, surname) 
        VALUES (:username, :password, :role, :name, :surname)`
	_, err := r.db.NamedExec(query, user)
	return err
}

func (r *UserRepository) Delete(id int) error {
	query := `
		DELETE FROM Users WHERE id=$1`
	_, err := r.db.NamedExec(query, id)
	return err
}

func (r *UserRepository) GetAll() ([]models.UserResponse, error) {
	var users []models.UserResponse
	query := `
		SELECT id, username, role, name, surname FROM users`
	err := r.db.Select(&users, query)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) GetByID(id int) (*models.User, error) {

	// burayı kontrol et
	// `user` değişkenini doğrudan başlatmak yerine, bir `models.User` nesnesi oluşturun.
	user := &models.User{}
	query := `
		SELECT id, username, role, name, surname FROM users WHERE id=$1`
	err := r.db.Get(user, query, id)
	if err != nil {
		return nil, err
	}
	return user, nil

}

func (r *UserRepository) GetByUsername(username string) (*models.User, error) {
	var user models.User
	fmt.Println(username)
	query := `
		SELECT id, username, password, role, name, surname FROM users WHERE username=$1`
	err := r.db.Get(&user, query, username)
	fmt.Println(user)
	if err != nil {
		return nil, err
	}
	return &user, nil

}

func (r *UserRepository) GetAllForUsers() ([]models.UserForUsers, error) {
	var users []models.UserForUsers
	query := `
		SELECT id, name, surname FROM users`
	err := r.db.Select(&users, query)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) Update(user *models.User) error {
	query := `
		UPDATE users SET username=:username, password=:password, role=:role, name=:name, surname=:surname WHERE id=:id`
	_, err := r.db.NamedExec(query, user)
	return err
}
