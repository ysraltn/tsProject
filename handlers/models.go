package handlers

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
}

type Cycle struct {
	ProductID  int `json:"product_id" db:"product_id"`
	Year       int `json:"year" db:"year"`
	Month      int `json:"month" db:"month"`
	CycleCount int `json:"cycle_count" db:"cycle_count"`
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
	ResponsibleID int    `json:"responsible_id,omitempty" db:"responsible_id"`
	OwnerID       int    `json:"owner_id,omitempty" db:"owner_id"`
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

//responsible ve owner id'leri string olarak tanımlanmış. Bu alanlar int olmalı.

type JWTClaimsUser struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Role     string `json:"role"`
}
