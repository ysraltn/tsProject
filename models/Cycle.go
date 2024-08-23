package models

type Cycle struct {
	ID        int `json:"id" db:"id"`
	ProductID int `json:"product_id" db:"product_id"`
	Year      int `json:"year" db:"year"`
	Month     int `json:"month" db:"month"`
	CycleNo   int `json:"cycle,omitempty" db:"cycle"`
}

type ICycleRepository interface {
	Add(cycle *Cycle) error
	GetAll() ([]Cycle, error)
	GetProductsCyclesByYear(productID, year int) ([]Cycle, error)
}

type ICycleService interface {
	Add(productID, year, month, cycleNo int) error
	GetAll() ([]Cycle, error)
	GetProductsCyclesByYear(productID, year int) ([]Cycle, error)
}
