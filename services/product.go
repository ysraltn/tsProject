package services

import "tsProject/models"

type ProductService struct {
	productRepositories models.IProductRepository
	logger              models.ILogger
}

func NewProductService(productRepositories models.IProductRepository, logger models.ILogger) models.IProductService {
	return &ProductService{
		productRepositories: productRepositories,
		logger:              logger,
	}
}

func (s *ProductService) Add(serial, brand, model, responsible, owner, status string, institutionID int) error {
	product := &models.Product{
		Serial:        serial,
		Brand:         brand,
		Model:         model,
		InstitutionID: institutionID,
		Responsible:   responsible,
		Owner:         owner,
		Status:        status,
	}
	err := s.productRepositories.Add(product)
	if err != nil {
		s.logger.Error(models.ErrorLog{
			Message: "error adding product",
			Error:   err.Error(),
			Context: map[string]interface{}{
				"product_serial": serial,
				"institution_id": institutionID,
				"responsible":    responsible,
			},
		})
		return err
	}
	s.logger.ProductInfoLog(models.ProductInfoLog{
		Message:       "product added",
		Event:         "product_added",
		ProductID:     product.ID,
		ProductName:   product.Serial,
		ProductSerial: product.Serial,
		InstitutionID: product.InstitutionID,
		Owner:         product.Owner,
		Status:        product.Status,
		Responsible:   product.Responsible,
	})
	return nil
}

func (s *ProductService) Delete(id int) error {
	err := s.productRepositories.Delete(id)
	if err != nil {
		s.logger.Error(models.ErrorLog{
			Message: "error deleting product",
			Error:   err.Error(),
			Context: map[string]interface{}{
				"product_id": id,
			},
		})
		return err
	}
	s.logger.Info(models.InfoLog{
		Message:   "product deleted",
		Event:     "product_deleted",
		ProductID: id,
	})
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
