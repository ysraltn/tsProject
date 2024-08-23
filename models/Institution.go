package models

type Institution struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	City string `json:"city" db:"city"`
}

type IInstitutionRepository interface {
	Add(institution *Institution) error
	// Delete(id int) error
	// GetAll() ([]Institution, error)
	GetByID(id int) (*Institution, error)
}

type IInstitutionService interface {
	Add(name, city string) error
	// 	Delete(id int) error
	// 	GetAll() ([]Product, error)
	GetByID(id int) (*Institution, error)
}
