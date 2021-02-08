package interfaces

import (
	"openbankingcrawler/common"
	"openbankingcrawler/domain/personalcreditcard"
	"openbankingcrawler/domain/personalloan"
	"openbankingcrawler/domain/subentities"
)

//ProductsServicesInterface interface
type ProductsServicesInterface interface {
	GetPersonalLoans(string, int) ([]personalloan.Entity, *subentities.Pagination, common.CustomError)
	GetPersonalCreditCards(string, int) ([]personalcreditcard.Entity, *subentities.Pagination, common.CustomError)
}

type productsServicesInterface struct {
	personalLoanService       personalloan.Service
	personalCreditCardService personalcreditcard.Service
}

//NewProductsServicesInterface create a new interface for NewPersonalLoan
func NewProductsServicesInterface(personalLoanService personalloan.Service, personalCreditCardService personalcreditcard.Service) ProductsServicesInterface {

	return &productsServicesInterface{
		personalLoanService:       personalLoanService,
		personalCreditCardService: personalCreditCardService,
	}
}

//GetFromInstitution get personalLoans from institutution
func (c *productsServicesInterface) GetPersonalLoans(id string, page int) ([]personalloan.Entity, *subentities.Pagination, common.CustomError) {
	return c.personalLoanService.FindByInstitution(id, page)
}

//GetFromInstitution get personalLoans from institutution
func (c *productsServicesInterface) GetPersonalCreditCards(id string, page int) ([]personalcreditcard.Entity, *subentities.Pagination, common.CustomError) {
	return c.personalCreditCardService.FindByInstitution(id, page)
}
