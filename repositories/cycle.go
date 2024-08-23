package repositories

import (
	"tsProject/models"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type CycleRepository struct {
	db *sqlx.DB
}

func NewCycleRepository(db *sqlx.DB) *CycleRepository {
	return &CycleRepository{db: db}
}

func (r *CycleRepository) Add(cycle *models.Cycle) error {
	query := `
		INSERT INTO cycles (product_id, year, month, cycle)
		VALUES (:product_id, :year, :month, :cycle)
	`
	_, err := r.db.NamedExec(query, cycle)
	return err
}

func (r *CycleRepository) GetAll() ([]models.Cycle, error) {
	var cycles []models.Cycle
	query := `
		SELECT * FROM cycles`
	err := r.db.Select(&cycles, query)
	if err != nil {
		return nil, err
	}
	return cycles, nil
}

func (r *CycleRepository) GetProductsCyclesByYear(productID, year int) ([]models.Cycle, error) {
	var cycles []models.Cycle
	query := `
		SELECT * FROM cycles
		WHERE product_id = $1 AND year = $2
		ORDER BY month ASC
	`

	// Execute the query and map the result to the 'cycles' slice
	err := r.db.Select(&cycles, query, productID, year)

	if err != nil {
		return nil, err
	}

	return cycles, nil
}
