package handlers

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Cycle struct {
	ID        int `json:"id" db:"id"`
	ProductID int `json:"product_id" db:"product_id"`
	Year      int `json:"year" db:"year"`
	Month     int `json:"month" db:"month"`
	CycleNo   int `json:"cycle,omitempty" db:"cycle"`
}

type Institution struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	City string `json:"city" db:"city"`
}

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