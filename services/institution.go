package services

import "tsProject/models"

type InstitutionService struct {
	institutionRepositories models.IInstitutionRepository
	logger                  models.ILogger
}

func NewInstitutionService(institutionRepositories models.IInstitutionRepository, logger models.ILogger) models.IInstitutionService {
	return &InstitutionService{
		institutionRepositories: institutionRepositories,
		logger:                  logger,
	}
}

func (s *InstitutionService) GetAll() ([]models.Institution, error) {
	institutions, err := s.institutionRepositories.GetAll()
	if err != nil {
		return nil, err
	}
	return institutions, nil
}

func (s *InstitutionService) Add(name, city string) error {
	institution := &models.Institution{
		Name: name,
		City: city,
	}
	err := s.institutionRepositories.Add(institution)
	if err != nil {
		s.logger.Error(models.ErrorLog{
			Message: "error adding institution",
			Error:   err.Error(),
			Context: map[string]interface{}{
				"institution name": institution.Name,
			},
		})
		return err
	}
	s.logger.InstitutionInfoLog(models.InstitutionInfoLog{
		Message:         "institution added",
		Event:           "institution_added",
		InstitutionName: institution.Name,
		City:            institution.City,
	})
	return nil
}

func (s *InstitutionService) GetByID(id int) (*models.Institution, error) {
	institution, err := s.institutionRepositories.GetByID(id)
	if err != nil {
		return nil, err
	}
	return institution, nil
}
