package branch

import "openbankingcrawler/common"

//Service branch service
type Service interface {
	DeleteAllFromInstitution(string) error
	InsertMany([]Entity, string) common.CustomError
}

type service struct {
	repository Repository
}

//NewService create a new service for branches
func NewService(repository Repository) Service {

	return &service{
		repository: repository,
	}
}

//DeleteAllFromInstitution update branches from institution
func (s *service) DeleteAllFromInstitution(InstitutionID string) error {

	deleteErr := s.repository.DeleteMany(InstitutionID)

	if deleteErr != nil {
		return deleteErr
	}

	return nil
}

//InsertMany insert many branches in instition
func (s *service) InsertMany(branches []Entity, institutuionID string) common.CustomError {

	for _, branch := range branches {
		branch.InstitutionID = institutuionID
		s.repository.Save(branch)
	}

	return nil
}
