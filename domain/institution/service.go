package institution

import (
	"openbankingcrawler/common"
	"openbankingcrawler/dtos"
)

//Service service
type Service interface {
	Create(dtos.Institution) (*dtos.Institution, common.CustomError)
	Delete(string) common.CustomError
	Find(string) (*dtos.Institution, common.CustomError)
}

type service struct {
	repository Repository
}

//NewService create a new service for institutions
func NewService(repository Repository) Service {

	return &service{
		repository: repository,
	}
}

//Create create a new institution
func (s *service) Create(institutionDTO dtos.Institution) (*dtos.Institution, common.CustomError) {

	validateError := institutionDTO.Validate()
	if validateError != nil {
		return nil, validateError
	}

	queriedInstitution, _ := s.repository.FindByName(institutionDTO.Name)

	if queriedInstitution != nil {
		return nil, common.NewBadRequestError("There is already an institution saved with this name")
	}

	newInstitution := NewEntity(institutionDTO.Name)
	savedInstitution, saveErr := s.repository.Save(*newInstitution)

	if saveErr != nil {
		return nil, saveErr
	}

	institutionDTO.ID = savedInstitution.RetrieveID()

	return &institutionDTO, nil
}

//Delete an institution
func (s *service) Delete(institutionID string) common.CustomError {

	newInstitution := NewEntityWithID(institutionID)
	return s.repository.Delete(*newInstitution)
}

func (s *service) Find(id string) (*dtos.Institution, common.CustomError) {
	queriedInstitution, err := s.repository.Find(id)

	if err != nil {
		return nil, err
	}

	return &dtos.Institution{Name: queriedInstitution.Name, ID: queriedInstitution.RetrieveID()}, nil
}
