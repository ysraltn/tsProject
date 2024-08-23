package models

type UserProductAssignment struct {
	ID        int `json:"id" db:"id"`
	UserID    int `json:"user_id" db:"user_id"`
	ProductID int `json:"product_id" db:"product_id"`
}
