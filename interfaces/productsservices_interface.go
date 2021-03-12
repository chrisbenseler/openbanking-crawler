package interfaces

import (
	"openbankingcrawler/common"
	"openbankingcrawler/domain/personalaccount"
	"openbankingcrawler/domain/personalcreditcard"
	"openbankingcrawler/domain/personalfinancing"
	"openbankingcrawler/domain/personalloan"
	"openbankingcrawler/domain/subentities"
)

//ProductsServicesInterface interface
type ProductsServicesInterface interface {
	GetPersonalAccounts(string, int) ([]personalaccount.Entity, *subentities.Pagination, common.CustomError)
	GetPersonalLoans(string, int) ([]personalloan.Entity, *subentities.Pagination, common.CustomError)
	GetPersonalFinancings(string, int) ([]personalfinancing.Entity, *subentities.Pagination, common.CustomError)
	GetPersonalCreditCards(string, int) ([]personalcreditcard.Entity, *subentities.Pagination, common.CustomError)
}

type productsServicesInterface struct {
	personalAccountService    personalaccount.Service
	personalLoanService       personalloan.Service
	personalFinancingService  personalfinancing.Service
	personalCreditCardService personalcreditcard.Service
}

//NewProductsServicesInterface create a new interface for products and services
func NewProductsServicesInterface(personalAccountService personalaccount.Service, personalLoanService personalloan.Service, personalFinancingService personalfinancing.Service, personalCreditCardService personalcreditcard.Service) ProductsServicesInterface {

	return &productsServicesInterface{
		personalAccountService:    personalAccountService,
		personalLoanService:       personalLoanService,
		personalFinancingService:  personalFinancingService,
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

//GetPersonalFinancings get personal financins from institutution
func (c *productsServicesInterface) GetPersonalFinancings(id string, page int) ([]personalfinancing.Entity, *subentities.Pagination, common.CustomError) {
	return c.personalFinancingService.FindByInstitution(id, page)
}

//GetFromInstitution get personal credit cards from institutution
func (c *productsServicesInterface) GetPersonalCreditCards(id string, page int) ([]personalcreditcard.Entity, *subentities.Pagination, common.CustomError) {
	return c.personalCreditCardService.FindByInstitution(id, page)
}
