package repositories

import (
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
		INSERT INTO Users (username, password, role) 
        VALUES (:username, :password, :role)`
	_, err := r.db.NamedExec(query, user)
	return err
}

func (r *UserRepository) Delete(id int) error {
	query := `
		DELETE FROM Users WHERE id=$1`
	_, err := r.db.NamedExec(query, id)
	return err
}

func (r *UserRepository) GetAll() ([]models.User, error) {
	var users []models.User
	query := `
		SELECT * FROM Users`
	err := r.db.Get(&users, query)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) GetByID(id int) (*models.User, error) {
	var user *models.User
	query := `
		SELECT * FROM Users WHERE id=$1`
	err := r.db.Get(&user, query, id)
	if err != nil {
		return nil, err
	}
	return user, nil

}

func (r *UserRepository) GetByUsername(username string) (*models.User, error) {
	var user models.User
	query := `
		SELECT * FROM Users WHERE username=$1`
	err := r.db.Get(&user, query, username)
	if err != nil {
		return nil, err
	}
	return &user, nil

}
