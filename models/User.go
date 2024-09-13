package models

type User struct {
	ID       int    `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
	Role     string `json:"role" db:"role"`
	Name     string `json:"name" db:"name"`
	Surname  string `json:"surname" db:"surname"`
}

type UserResponse struct {
	ID       int    `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Role     string `json:"role" db:"role"`
	Name     string `json:"name" db:"name"`
	Surname  string `json:"surname" db:"surname"`
}

type UserForUsers struct {
	ID      int    `json:"id" db:"id"`
	Name    string `json:"name" db:"name"`
	Surname string `json:"surname" db:"surname"`
}

type IUserRepository interface {
	Add(user *User) error
	Delete(id int) error
	GetAll() ([]UserResponse, error)
	GetByID(id int) (*User, error)
	GetByUsername(username string) (*User, error)
	GetAllForUsers() ([]UserForUsers, error)
	Update(user *User) error
}

type IUserService interface {
	Add(username, password, role, name, surname string) error
	Delete(id int) error
	GetAll() ([]UserResponse, error)
	GetByID(id int) (*User, error)
	GetByUsername(username string) (*User, error)
	GetAllForUsers() ([]UserForUsers, error)
	Update(id int, username, password, role, name, surname string) error
}
