package repositories

import (
	"encoding/json"
	"errors"
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
		INSERT INTO Products (serial, brand, model, institution_id, responsible_id, owner_id, status)
		VALUES (:serial, :brand, :model, :institution_id, :responsible_id, :owner_id, :status)
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

// responsible id ve owner id dönüyor, bunu isimler gelecek şekilde düzelt
func (r *ProductRepository) GetAllWithInstitutions() ([]models.ProductWithInstitution, error) {
	var products []models.ProductWithInstitution
	query := `
		SELECT 
			p.id, 
			p.serial, 
			p.brand, 
			p.model, 
			p.responsible_id, 
			p.owner_id, 
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

func (r *ProductRepository) GetAllProductsWithInstitutionAndCycle() ([]models.ProductWithInstitutionAndCycle, error) {
	var products []models.ProductWithInstitutionAndCycle

	query := `
        SELECT 
    p.id AS product_id, 
    p.serial AS product_serial, 
    p.brand AS product_brand, 
    p.model AS product_model, 
    i.id AS institution_id,
    i.name AS institution_name, 
    i.city AS institution_city, 
    COALESCE(json_agg(json_build_object(
        'year', c.year, 
        'month', c.month, 
        'cycle_count', c.cycle_count
    )) FILTER (WHERE c.cycle_count IS NOT NULL), '[]') AS cycles
FROM products p
JOIN institutions i ON p.institution_id = i.id
LEFT JOIN cycles c ON p.id = c.product_id
GROUP BY p.id, i.id;

    `

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product models.ProductWithInstitutionAndCycle
		var cyclesJSON []byte

		// İlk önce dönen JSON verisini []byte olarak okuyun
		err := rows.Scan(&product.ProductID, &product.ProductSerial, &product.ProductBrand, &product.ProductModel,
			&product.InstitutionID, &product.InstitutionName, &product.InstitutionCity, &cyclesJSON)
		if err != nil {
			return nil, err
		}

		// Daha sonra JSON'u []models.Cycle türüne parse edin
		var cycles []models.CycleForArray
		if err := json.Unmarshal(cyclesJSON, &cycles); err != nil {
			return nil, err
		}
		product.Cycles = cycles

		products = append(products, product)
	}

	return products, nil
}

func (r *ProductRepository) GetAssignedProducts(userID int) ([]models.ProductWithInstitutionAndCycle, error) {
	var products []models.ProductWithInstitutionAndCycle
	query := `
        SELECT 
    p.id AS product_id, 
    p.serial AS product_serial, 
    p.brand AS product_brand, 
    p.model AS product_model, 
    i.id AS institution_id,
    i.name AS institution_name, 
    i.city AS institution_city, 
    COALESCE(json_agg(json_build_object(
        'year', c.year, 
        'month', c.month, 
        'cycle_count', c.cycle_count
    )) FILTER (WHERE c.cycle_count IS NOT NULL), '[]') AS cycles
FROM products p
JOIN institutions i ON p.institution_id = i.id
LEFT JOIN cycles c ON p.id = c.product_id
JOIN user_product_assignments upa ON p.id = upa.product_id
WHERE upa.user_id = $1
GROUP BY p.id, i.id;


    `
	err := r.db.Select(&products, query, userID)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductRepository) FilterByInstitutionID(institutionID int) ([]models.ProductWithInstitutionAndCycle, error) {
	if institutionID == 0 {
		return nil, errors.New("invalid institution id")
	}
	var products []models.ProductWithInstitutionAndCycle
	query := `
        SELECT 
    p.id AS product_id, 
    p.serial AS product_serial, 
    p.brand AS product_brand, 
    p.model AS product_model, 
    i.id AS institution_id,
    i.name AS institution_name, 
    i.city AS institution_city, 
    COALESCE(json_agg(json_build_object(
        'year', c.year, 
        'month', c.month, 
        'cycle_count', c.cycle_count
    )) FILTER (WHERE c.cycle_count IS NOT NULL), '[]') AS cycles
FROM products p
JOIN institutions i ON p.institution_id = i.id
LEFT JOIN cycles c ON p.id = c.product_id
JOIN user_product_assignments upa ON p.id = upa.product_id
WHERE i.id = $1
GROUP BY p.id, i.id;
    `
	err := r.db.Select(&products, query, institutionID)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductRepository) Filter(institutionName, institutionCity, productModel string) ([]models.ProductWithInstitutionAndCycle, error) {
	var products []models.ProductWithInstitutionAndCycle
	query := `
        SELECT 
            p.id AS product_id, 
            p.serial AS product_serial, 
            p.brand AS product_brand, 
            p.model AS product_model, 
            i.id AS institution_id,
            i.name AS institution_name, 
            i.city AS institution_city, 
            COALESCE(json_agg(json_build_object(
                'year', c.year, 
                'month', c.month, 
                'cycle_count', c.cycle_count
            )) FILTER (WHERE c.cycle_count IS NOT NULL), '[]') AS cycles
        FROM products p
        JOIN institutions i ON p.institution_id = i.id
        LEFT JOIN cycles c ON p.id = c.product_id
        WHERE
            i.name LIKE ('%' || COALESCE(NULLIF($1, ''), i.name) || '%') AND
            i.city LIKE ('%' || COALESCE(NULLIF($2, ''), i.city) || '%') AND
            p.model LIKE ('%' || COALESCE(NULLIF($3, ''), p.model) || '%')
        GROUP BY p.id, i.id;
    `
	err := r.db.Select(&products, query, institutionName, institutionCity, productModel)
	if err != nil {
		return nil, err
	}
	return products, nil
}
