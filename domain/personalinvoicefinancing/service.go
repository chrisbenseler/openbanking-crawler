package personalinvoicefinancing

import (
	"openbankingcrawler/common"
	"openbankingcrawler/domain/subentities"
)

//Service service
type Service interface {
	DeleteAllFromInstitution(string) common.CustomError
	InsertMany([]Entity, string) common.CustomError
	FindByInstitution(string, int) ([]Entity, *subentities.Pagination, common.CustomError)
}

type service struct {
	repository Repository
}

//NewService create a new service
func NewService(repository Repository) Service {

	return &service{
		repository: repository,
	}
}

//DeleteAllFromInstitution update branches from institution
func (s *service) DeleteAllFromInstitution(InstitutionID string) common.CustomError {

	deleteErr := s.repository.DeleteMany(InstitutionID)

	if deleteErr != nil {
		return common.NewInternalServerError("Could not delete personalInvoiceFinancings from institution", deleteErr)
	}

	return nil
}

//InsertMany insert many personalInvoiceFinancings in instition
func (s *service) InsertMany(personalInvoiceFinancings []Entity, institututionID string) common.CustomError {

	for _, personalInvoiceFinancing := range personalInvoiceFinancings {
		personalInvoiceFinancing.InstitutionID = institututionID
		s.repository.Save(personalInvoiceFinancing)
	}

	return nil
}

//FindByInstitution insert many branches in instition
func (s *service) FindByInstitution(institututionID string, page int) ([]Entity, *subentities.Pagination, common.CustomError) {
	return s.repository.FindByInstitution(institututionID, page)
}
