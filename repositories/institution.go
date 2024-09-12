package repositories

import (
	"tsProject/models"

	"github.com/jmoiron/sqlx"
)

type InstitutionRepository struct {
	db *sqlx.DB
}

func NewInstitutionRepository(db *sqlx.DB) *InstitutionRepository {
	return &InstitutionRepository{db: db}
}

func (r *InstitutionRepository) Add(institution *models.Institution) error {
	query := `
		INSERT INTO institutions (name, city)
		VALUES (:name, :city)
	`

	_, err := r.db.NamedExec(query, institution)
	return err
}

func (r *InstitutionRepository) GetByID(id int) (*models.Institution, error) {
	var institution models.Institution
	query := `SELECT * FROM institutions WHERE id = $1`
	err := r.db.Get(&institution, query, id)
	if err != nil {
		return nil, err
	}
	return &institution, nil
}

func (r *InstitutionRepository) GetAll() ([]models.Institution, error) {
	var institutions []models.Institution
	query := `SELECT * FROM institutions
			  ORDER BY institutions.name ASC`
	err := r.db.Select(&institutions, query)
	if err != nil {
		return nil, err
	}
	return institutions, nil
}

func (r *InstitutionRepository) Update(institution *models.Institution) error {
	query := `
		UPDATE institutions
		SET name = :name, city = :city
		WHERE id = :id
	`
	_, err := r.db.NamedExec(query, institution)
	return err
}

func (r *InstitutionRepository) Delete(id int) error {
	query := `DELETE FROM institutions WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *InstitutionRepository) GetInstitutionProducts(id int) ([]models.Product, error) {
	var products []models.Product
	query := `
		SELECT * FROM products
		WHERE institution_id = $1
	`
	err := r.db.Select(&products, query, id)
	if err != nil {
		return nil, err
	}
	return products, nil
}
