package models

type User struct {
	ID       int    `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
	Role     string `json:"role" db:"role"`
}

type UserResponse struct {
	ID       int    `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Role     string `json:"role" db:"role"`
}

type IUserRepository interface {
	Add(user *User) error
	Delete(id int) error
	GetAll() ([]UserResponse, error)
	GetByID(id int) (*User, error)
	GetByUsername(username string) (*User, error)
}

type IUserService interface {
	Add(username, password, role string) error
	Delete(id int) error
	GetAll() ([]UserResponse, error)
	GetByID(id int) (*User, error)
	GetByUsername(username string) (*User, error)
}
