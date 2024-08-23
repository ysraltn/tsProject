package services

import "tsProject/models"

type InstitutionService struct {
	institutionRepositories models.IInstitutionRepository
}

func NewInstitutionService(institutionRepositories models.IInstitutionRepository) models.IInstitutionService {
	return &InstitutionService{
		institutionRepositories: institutionRepositories,
	}
}

func (s *InstitutionService) Add(name, city string) error {
	institution := &models.Institution{
		Name: name,
		City: city,
	}
	err := s.institutionRepositories.Add(institution)
	if err != nil {
		return err
	}
	return nil
}

func (s *InstitutionService) GetByID(id int) (*models.Institution, error) {
	institution, err := s.institutionRepositories.GetByID(id)
	if err != nil {
		return nil, err
	}
	return institution, nil
}