package models

type Product struct {
	ID            int    `json:"id" db:"id"`
	Serial        string `json:"serial" db:"serial"`
	Brand         string `json:"brand" db:"brand"`
	Model         string `json:"model" db:"model"`
	InstitutionID int    `json:"institution_id" db:"institution_id"`
	Responsible   string `json:"responsible,omitempty" db:"responsible"`
	Owner         string `json:"owner,omitempty" db:"owner"`
	Status        string `json:"status,omitempty" db:"status"`
}

type ProductWithInstitution struct {
	ID            int    `json:"id" db:"id"`
	Serial        string `json:"serial" db:"serial"`
	Brand         string `json:"brand" db:"brand"`
	Model         string `json:"model" db:"model"`
	Responsible   string `json:"responsible" db:"responsible"`
	Owner         string `json:"owner" db:"owner"`
	Status        string `json:"status" db:"status"`
	InstitutionID int    `json:"institution_id" db:"institution_id"`
	Institution   string `json:"institution_name" db:"institution_name"`
}

type IProductRepository interface {
	Add(product *Product) error
	Delete(id int) error
	GetAll() ([]Product, error)
	GetByID(id int) (*Product, error)
	GetAllWithInstitutions() ([]ProductWithInstitution, error)
}

type IProductService interface {
	Add(serial, brand, model, responsible, owner, status string, institution int) error
	Delete(id int) error
	GetAll() ([]Product, error)
	GetByID(id int) (*Product, error)
	GetAllWithInstitutions() ([]ProductWithInstitution, error)
}
