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
		INSERT INTO Institution (name, city)
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
