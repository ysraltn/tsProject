package services

import "tsProject/models"

type CycleService struct {
	cycleRepositories models.ICycleRepository
	logger            models.ILogger
}

func NewCycleService(cycleRepositories models.ICycleRepository, logger models.ILogger) models.ICycleService {
	return &CycleService{
		cycleRepositories: cycleRepositories,
		logger:            logger,
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
		s.logger.Error(models.ErrorLog{
			Message: "error adding cycle",
			Error:   err.Error(),
			Context: map[string]interface{}{
				"product_id": productID,
				"year":       year,
				"month":      month,
				"cycle_no":   cycleNo,
			},
		})
		return err
	}
	s.logger.CycleInfoLog(models.CycleInfoLog{
		Message:   "cycle added",
		Event:     "cycle_added",
		ProductID: productID,
		Year:      year,
		Month:     month,
		CycleNo:   cycleNo,
	})
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
