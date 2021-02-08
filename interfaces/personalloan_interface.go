package interfaces

import (
	"openbankingcrawler/common"
	"openbankingcrawler/domain/personalloan"
	"openbankingcrawler/domain/subentities"
)

//PersonalLoanInterface interface
type PersonalLoanInterface interface {
	GetFromInstitution(string, int) ([]personalloan.Entity, *subentities.Pagination, common.CustomError)
}

type personalLoanInterface struct {
	personalLoanService personalloan.Service
}

//NewPersonalLoan create a new interface for NewPersonalLoan
func NewPersonalLoan(personalLoanService personalloan.Service) PersonalLoanInterface {

	return &personalLoanInterface{
		personalLoanService: personalLoanService,
	}
}

//GetFromInstitution get personalLoans from institutution
func (c *personalLoanInterface) GetFromInstitution(id string, page int) ([]personalloan.Entity, *subentities.Pagination, common.CustomError) {
	return c.personalLoanService.FindByInstitution(id, page)
}
