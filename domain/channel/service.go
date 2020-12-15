package channel

import "openbankingcrawler/common"

//Service channel service
type Service interface {
	DeleteAllFromInstitution(string) common.CustomError
	InsertMany([]Entity, string) common.CustomError
	FindByInstitution(institututionID string) ([]Entity, common.CustomError)
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
		return common.NewInternalServerError("Could not delete channels from institution", deleteErr)
	}

	return nil
}

//InsertMany insert many channels in instition
func (s *service) InsertMany(channels []Entity, institututionID string) common.CustomError {

	for _, channel := range channels {
		channel.InstitutionID = institututionID
		s.repository.Save(channel)
	}

	return nil
}

//FindByInstitution insert many branches in instition
func (s *service) FindByInstitution(institututionID string) ([]Entity, common.CustomError) {
	return s.repository.FindByInstitution(institututionID)
}
