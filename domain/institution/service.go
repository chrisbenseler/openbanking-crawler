package institution

import (
	"errors"
)

//Service service
type Service interface {
	Create(string) error
	Delete(string) error
}

type institutionService struct {
	repository Repository
}

//NewService create a new service for institutions
func NewService(repository Repository) Service {

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

	newInstitution := NewEntity(name)
	return s.repository.Save(*newInstitution)
}

//Delete
func (s *institutionService) Delete(institutionID string) error {

	newInstitution := NewEntityWithID(institutionID)
	return s.repository.Delete(*newInstitution)
}
