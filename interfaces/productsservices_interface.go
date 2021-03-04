package interfaces

import (
	"openbankingcrawler/common"
	"openbankingcrawler/domain/personalaccount"
	"openbankingcrawler/domain/personalcreditcard"
	"openbankingcrawler/domain/personalloan"
	"openbankingcrawler/domain/subentities"
)

//ProductsServicesInterface interface
type ProductsServicesInterface interface {
	GetPersonalLoans(string, int) ([]personalloan.Entity, *subentities.Pagination, common.CustomError)
	GetPersonalAccounts(string, int) ([]personalaccount.Entity, *subentities.Pagination, common.CustomError)
	GetPersonalCreditCards(string, int) ([]personalcreditcard.Entity, *subentities.Pagination, common.CustomError)
}

type productsServicesInterface struct {
	personalLoanService       personalloan.Service
	personalAccountService    personalaccount.Service
	personalCreditCardService personalcreditcard.Service
}

//NewProductsServicesInterface create a new interface for NewPersonalLoan
func NewProductsServicesInterface(personalLoanService personalloan.Service, personalAccountService personalaccount.Service, personalCreditCardService personalcreditcard.Service) ProductsServicesInterface {

	return &productsServicesInterface{
		personalLoanService:       personalLoanService,
		personalAccountService:    personalAccountService,
		personalCreditCardService: personalCreditCardService,
	}
}

//GetFromInstitution get personal loans from institutution
func (c *productsServicesInterface) GetPersonalLoans(id string, page int) ([]personalloan.Entity, *subentities.Pagination, common.CustomError) {
	return c.personalLoanService.FindByInstitution(id, page)
}

//GetPersonalAccounts get personal accounts from institutution
func (c *productsServicesInterface) GetPersonalAccounts(id string, page int) ([]personalaccount.Entity, *subentities.Pagination, common.CustomError) {
	return c.personalAccountService.FindByInstitution(id, page)
}

//GetFromInstitution get personal credit cards from institutution
func (c *productsServicesInterface) GetPersonalCreditCards(id string, page int) ([]personalcreditcard.Entity, *subentities.Pagination, common.CustomError) {
	return c.personalCreditCardService.FindByInstitution(id, page)
}
