package branch

import "openbankingcrawler/common"

//Service branch service
type Service interface {
	DeleteAllFromInstitution(string) common.CustomError
	InsertMany([]Entity, string) common.CustomError
	FindByInstitution(string) []Entity
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
func (s *service) DeleteAllFromInstitution(InstitutionID string) common.CustomError {

	deleteErr := s.repository.DeleteMany(InstitutionID)

	if deleteErr != nil {
		return common.NewInternalServerError("Could not delete branches from sintitution", deleteErr)
	}

	return nil
}

//InsertMany insert many branches in instition
func (s *service) InsertMany(branches []Entity, institututionID string) common.CustomError {

	for _, branch := range branches {
		branch.InstitutionID = institututionID
		s.repository.Save(branch)
	}

	return nil
}

//FindByInstitution insert many branches in instition
func (s *service) FindByInstitution(institututionID string) []Entity {
	return s.FindByInstitution(institututionID)
}
