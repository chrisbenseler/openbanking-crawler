package services

import (
	"errors"
	"openbankingcrawler/common"
	"openbankingcrawler/domain/institution"
	"openbankingcrawler/dtos"
)

//InstitutionService service
type InstitutionService interface {
	Create(dtos.Institution) (*dtos.Institution, error)
	Delete(string) error
	Find(string) (*institution.Entity, common.CustomError)
}

type institutionService struct {
	repository institution.Repository
}

//NewInstitution create a new service for institutions
func NewInstitution(repository institution.Repository) InstitutionService {

	return &institutionService{
		repository: repository,
	}
}

//Create create a new institution
func (s *institutionService) Create(institutionDTO dtos.Institution) (*dtos.Institution, error) {

	quueriedInstitution, _ := s.repository.FindByName(institutionDTO.Name)

	if quueriedInstitution != nil {
		return nil, errors.New("There is already an institution saved with this name")
	}

	newInstitution := institution.NewEntity(institutionDTO.Name)
	savedInstitution, saveErr := s.repository.Save(*newInstitution)

	if saveErr != nil {
		return nil, saveErr
	}

	institutionDTO.ID = savedInstitution.RetrieveID()

	return &institutionDTO, nil
}

//Delete an institution
func (s *institutionService) Delete(institutionID string) error {

	newInstitution := institution.NewEntityWithID(institutionID)
	return s.repository.Delete(*newInstitution)
}

func (s *institutionService) Find(id string) (*institution.Entity, common.CustomError) {
	return s.repository.Find(id)
}
