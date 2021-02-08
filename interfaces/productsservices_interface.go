package interfaces

import (
	"openbankingcrawler/common"
	"openbankingcrawler/domain/personalloan"
	"openbankingcrawler/domain/subentities"
)

//ProductsServicesInterface interface
type ProductsServicesInterface interface {
	GetPersonalLoans(string, int) ([]personalloan.Entity, *subentities.Pagination, common.CustomError)
}

type productsServicesInterface struct {
	personalLoanService personalloan.Service
}

//NewProductsServicesInterface create a new interface for NewPersonalLoan
func NewProductsServicesInterface(personalLoanService personalloan.Service) ProductsServicesInterface {

	return &productsServicesInterface{
		personalLoanService: personalLoanService,
	}
}

//GetFromInstitution get personalLoans from institutution
func (c *productsServicesInterface) GetPersonalLoans(id string, page int) ([]personalloan.Entity, *subentities.Pagination, common.CustomError) {
	return c.personalLoanService.FindByInstitution(id, page)
}
