package services

import (
	"errors"
	"openbankingcrawler/domain/institution"
)

//InstitutionService service
type InstitutionService interface {
	Create(string) error
	Delete(string) error
	Find(string) (*institution.Entity, error)
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
func (s *institutionService) Create(name string) error {

	savedInstitution, _ := s.repository.FindByName(name)

	if savedInstitution != nil {
		return errors.New("There is already an institution saved with this name")
	}

	newInstitution := institution.NewEntity(name)
	return s.repository.Save(*newInstitution)
}

//Delete an institution
func (s *institutionService) Delete(institutionID string) error {

	newInstitution := institution.NewEntityWithID(institutionID)
	return s.repository.Delete(*newInstitution)
}

func (s *institutionService) Find(id string) (*institution.Entity, error) {
	return s.repository.Find(id)
}
