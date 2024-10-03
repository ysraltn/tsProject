package services

import (
	"bytes"
	"errors"
	"strconv"
	"time"
	"tsProject/models"

	"github.com/xuri/excelize/v2"
)

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

func (s *ProductService) Add(serial, brand, model, status string, institutionID, responsibleID, ownerID int) error {
	product := &models.Product{
		Serial:        serial,
		Brand:         brand,
		Model:         model,
		InstitutionID: institutionID,
		ResponsibleID: responsibleID,
		OwnerID:       ownerID,
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
				"responsible_id": responsibleID,
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
		OwnerID:       product.OwnerID,
		Status:        product.Status,
		ResponsibleID: product.ResponsibleID,
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

func (s *ProductService) GetAllProductsWithInstitutionAndCycle() ([]models.ProductWithInstitutionAndCycle, error) {
	products, err := s.productRepositories.GetAllProductsWithInstitutionAndCycle()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *ProductService) GetAssignedProducts(userID int) ([]models.ProductWithInstitutionAndCycle, error) {
	products, err := s.productRepositories.GetAssignedProducts(userID)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *ProductService) FilterByInstitutionID(institutionID int) ([]models.ProductWithInstitutionAndCycle, error) {
	products, err := s.productRepositories.FilterByInstitutionID(institutionID)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *ProductService) Filter(filter models.ProductFilter) ([]models.ProductWithInstitutionAndCycle, error) {
	products, err := s.productRepositories.Filter(filter.InstitutionName, filter.InstitutionCity, filter.ProductModel)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *ProductService) DownloadCycles(format string) (excelBuffer *bytes.Buffer, err error) {
	if format == "xlsx" {
		cycles, err := s.productRepositories.GetAllProductsWithInstitutionAndCycle()
		if err != nil {
			return nil, err
		}
		f := excelize.NewFile()
		sheetName := "Sheet1"
		headers := []string{"ID", "seri no", "marka", "model", "kurum", "sehir", "ocak", "subat", "mart", "nisan", "mayis", "haziran", "temmuz", "agustos", "eylul", "ekim", "kasim", "aralik"}
		for i, header := range headers {
			cell := string(rune('A'+i)) + "1"
			f.SetCellValue(sheetName, cell, header)
		}
		currentYear := time.Now().Year()

		// Verileri yaz
		row := 2
		for _, product := range cycles {
			f.SetCellValue(sheetName, "A"+strconv.Itoa(row), product.ProductID)
			f.SetCellValue(sheetName, "B"+strconv.Itoa(row), product.ProductSerial)
			f.SetCellValue(sheetName, "C"+strconv.Itoa(row), product.ProductBrand)
			f.SetCellValue(sheetName, "D"+strconv.Itoa(row), product.ProductModel)
			f.SetCellValue(sheetName, "E"+strconv.Itoa(row), product.InstitutionName)
			f.SetCellValue(sheetName, "F"+strconv.Itoa(row), product.InstitutionCity)

			// Ay isimlerine göre verileri yerleştir
			for _, cycle := range product.Cycles {
				if cycle.Year == currentYear {
					column := ""
					switch cycle.Month {
					case 1:
						column = "G" // Ocak
					case 2:
						column = "H" // Şubat
					case 3:
						column = "I" // Mart
					case 4:
						column = "J" // Nisan
					case 5:
						column = "K" // Mayıs
					case 6:
						column = "L" // Haziran
					case 7:
						column = "M" // Temmuz
					case 8:
						column = "N" // Ağustos
					case 9:
						column = "O" // Eylül
					case 10:
						column = "P" // Ekim
					case 11:
						column = "Q" // Kasım
					case 12:
						column = "R" // Aralık
					}

					// İlgili hücreye cycle sayısını yaz
					f.SetCellValue(sheetName, column+strconv.Itoa(row), cycle.Cycle_count)
				}
			}
			row++
		}

		// Excel dosyasını buffer'a kaydet
		excelBuffer, err := f.WriteToBuffer()
		if err != nil {
			return nil, err
		}
		return excelBuffer, nil
	}
	return nil, errors.New("invalid format")
}

func (s *ProductService) GetAllProductsWithInstitutionAndCycleByResponsibility(userID int) ([]models.ProductWithInstitutionAndCycle, error) {
	products, err := s.productRepositories.GetAllProductsWithInstitutionAndCycleByResponsibility(userID)
	if err != nil {
		return nil, err
	}
	return products, nil
}