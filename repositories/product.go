package repositories

import (
	"tsProject/models"

	"github.com/jmoiron/sqlx"
)

type ProductRepository struct {
	db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) Add(product *models.Product) error {
	query := `
		INSERT INTO Products (serial, brand, model, institution_id, responsible, owner, status)
		VALUES (:serial, :brand, :model, :institution_id, :responsible, :owner, :status)
	`

	_, err := r.db.NamedExec(query, product)
	return err
}

func (r *ProductRepository) Delete(id int) error {
	query := `
		DELETE FROM Products WHERE id=$1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *ProductRepository) GetAll() ([]models.Product, error) {
	var products []models.Product
	query := `
		SELECT * FROM Products`
	err := r.db.Select(&products, query)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductRepository) GetByID(id int) (*models.Product, error) {
	var product models.Product
	query := `
		SELECT * FROM Products WHERE id=$1`
	err := r.db.Get(&product, query, id)
	if err != nil {
		return nil, err
	}
	return &product, nil

}

func (r *ProductRepository) GetAllWithInstitutions() ([]models.ProductWithInstitution, error) {
	var products []models.ProductWithInstitution
	query := `
		SELECT 
			p.id, 
			p.serial, 
			p.brand, 
			p.model, 
			p.responsible, 
			p.owner, 
			p.status, 
			p.institution_id, 
			i.name AS institution_name
		FROM 
			products p
		JOIN 
			institutions i 
		ON 
			p.institution_id = i.id
	`
	err := r.db.Select(&products, query)
	if err != nil {
		return nil, err
	}
	return products, nil
}
