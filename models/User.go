package models

type User struct {
	ID       int    `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
	Role     string `json:"role" db:"role"`
}

type IUserRepository interface {
	Add(user *User) error
	Delete(id int) error
	GetAll() ([]User, error)
	GetByID(id int) (*User, error)
	GetByUsername(username string) (*User, error)
}

type IUserService interface {
	Add(username, password string) error
	Delete(id int) error
	GetAll() ([]User, error)
	GetByID(id int) (*User, error)
}
