package services

import "tsProject/models"

type CycleService struct {
	cycleRepositories models.ICycleRepository
}

func NewCycleService(cycleRepositories models.ICycleRepository) models.ICycleService {
	return &CycleService{
		cycleRepositories: cycleRepositories,
	}
}

func (s *CycleService) Add(productID, year, month, cycleNo int) error {
	cycle := &models.Cycle{
		ProductID: productID,
		Year:      year,
		Month:     month,
		CycleNo:   cycleNo,
	}
	err := s.cycleRepositories.Add(cycle)
	if err != nil {
		return err
	}
	return nil
}

func (s *CycleService) GetAll() ([]models.Cycle, error) {
	cycles, err := s.cycleRepositories.GetAll()
	if err != nil {
		return nil, err
	}
	return cycles, nil
}

func (s *CycleService) GetProductsCyclesByYear(productID, year int) ([]models.Cycle, error) {
	cycles, err := s.cycleRepositories.GetProductsCyclesByYear(productID, year)
	if err != nil {
		return nil, err
	}
	return cycles, nil
}
