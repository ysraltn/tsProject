package services

import "tsProject/models"

type ProductService struct {
	productRepositories models.IProductRepository
}

func NewProductService(productRepositories models.IProductRepository) models.IProductService {
	return &ProductService{
		productRepositories: productRepositories,
	}
}

func (s *ProductService) Add(serial, brand, model, responsible, owner, status string, institutionID int) error {
	product := &models.Product{
		Serial:      serial,
		Brand:       brand,
		Model:       model,
		InstitutionID:    institutionID,
		Responsible: responsible,
		Owner:       owner,
		Status:      status,
	}
	err := s.productRepositories.Add(product)
	if err != nil {
		return err
	}
	return nil
}

func (s *ProductService) Delete(id int) error {
	err := s.productRepositories.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *ProductService) GetAll() ([]models.Product, error) {
	products, err := s.productRepositories.GetAll()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *ProductService) GetAllWithInstitutions() ([]models.ProductWithInstitution, error) {
	products, err := s.productRepositories.GetAllWithInstitutions()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *ProductService) GetByID(id int) (*models.Product, error) {
	product, err := s.productRepositories.GetByID(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}
