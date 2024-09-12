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
	query1 := `DELETE FROM cycles WHERE product_id = $1 AND year = $2 AND month = $3`
	_, err := r.db.Exec(query1, cycle.ProductID, cycle.Year, cycle.Month)
	if err != nil {
		return err
	}
	query2 := `
		INSERT INTO cycles (product_id, year, month, cycle_count)
		VALUES (:product_id, :year, :month, :cycle_count)
	`
	_, err = r.db.NamedExec(query2, cycle)
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

func (r *CycleRepository) GetAllByYear(year int) ([]models.Cycle, error) {
	var cycles []models.Cycle
	query := `
		SELECT * FROM cycles
		WHERE year = $1
		ORDER BY product_id ASC
	`

	// Execute the query and map the result to the 'cycles' slice
	err := r.db.Select(&cycles, query, year)

	if err != nil {
		return nil, err
	}

	return cycles, nil
}
