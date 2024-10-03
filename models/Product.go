package models

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

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
	ResponsibleID int    `json:"responsible_id" db:"responsible_id"`
	OwnerID       int    `json:"owner_id" db:"owner_id"`
	Status        string `json:"status" db:"status"`
	InstitutionID int    `json:"institution_id" db:"institution_id"`
	Institution   string `json:"institution_name" db:"institution_name"`
}

type ProductWithInstitutionAndCycle struct {
	ProductID       int        `json:"product_id" db:"product_id"`
	ProductSerial   string     `json:"product_serial" db:"product_serial"`
	ProductBrand    string     `json:"product_brand" db:"product_brand"`
	ProductModel    string     `json:"product_model" db:"product_model"`
	InstitutionID   int        `json:"institution_id" db:"institution_id"`
	InstitutionName string     `json:"institution_name" db:"institution_name"`
	InstitutionCity string     `json:"institution_city" db:"institution_city"`
	Cycles          CyclesJSON `json:"cycles" db:"cycles"`
}

type CycleForArray struct {
	Year        int `json:"year" db:"year"`
	Month       int `json:"month" db:"month"`
	Cycle_count int `json:"cycle_count" db:"cycle_count"`
}

type ProductFilter struct {
	InstitutionName string `json:"institution_name" db:"institution_name"`
	InstitutionCity string `json:"institution_city" db:"institution_city"`
	ProductModel    string `json:"product_model" db:"product_model"`
}

// CyclesJSON wraps a slice of CycleForArray to provide custom scanning from JSON data
type CyclesJSON []CycleForArray

// Scan implements the Scanner interface for CyclesJSON
// This allows direct scanning of JSON data from the database into the struct
func (c *CyclesJSON) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan value, expected []byte, got %T", value)
	}
	return json.Unmarshal(bytes, c)
}

// Value implements the Valuer interface for CyclesJSON
// This allows the struct to be saved as JSON in the database
func (c CyclesJSON) Value() (driver.Value, error) {
	return json.Marshal(c)
}

type IProductRepository interface {
	Add(product *Product) error
	Delete(id int) error
	GetAll() ([]Product, error)
	GetByID(id int) (*Product, error)
	GetAllWithInstitutions() ([]ProductWithInstitution, error)
	GetAllProductsWithInstitutionAndCycle() ([]ProductWithInstitutionAndCycle, error)
	GetAssignedProducts(userID int) ([]ProductWithInstitutionAndCycle, error)
	FilterByInstitutionID(institutionID int) ([]ProductWithInstitutionAndCycle, error)
	Filter(institutionName, institutionCity, productModel string) ([]ProductWithInstitutionAndCycle, error)
	GetAllProductsWithInstitutionAndCycleByResponsibility(userID int) ([]ProductWithInstitutionAndCycle, error)
}

type IProductService interface {
	Add(serial, brand, model, status string, institutionID, responsibleID, ownerID int) error
	Delete(id int) error
	GetAll() ([]Product, error)
	GetByID(id int) (*Product, error)
	GetAllWithInstitutions() ([]ProductWithInstitution, error)
	GetAllProductsWithInstitutionAndCycle() ([]ProductWithInstitutionAndCycle, error)
	GetAssignedProducts(userID int) ([]ProductWithInstitutionAndCycle, error)
	FilterByInstitutionID(institutionID int) ([]ProductWithInstitutionAndCycle, error)
	Filter(filter ProductFilter) ([]ProductWithInstitutionAndCycle, error)
	DownloadCycles(format string) (excelBuffer *bytes.Buffer, err error)
	GetAllProductsWithInstitutionAndCycleByResponsibility(userID int) ([]ProductWithInstitutionAndCycle, error)
}
