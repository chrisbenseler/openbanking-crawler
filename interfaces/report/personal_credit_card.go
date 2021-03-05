package report

import (
	"openbankingcrawler/domain/institution"
	"openbankingcrawler/domain/personalcreditcard"
	"openbankingcrawler/dtos"
)

//Report Report interface for personal credit card
type Report interface {
	Fees() *[]OutputPrice
}

type report struct {
	institutionService        institution.Service
	personalCreditCardService personalcreditcard.Service
}

//NewPersonalCreditCard create a new report interface
func NewPersonalCreditCard(institutionService institution.Service, personalCreditCardService personalcreditcard.Service) Report {
	return &report{
		institutionService:        institutionService,
		personalCreditCardService: personalCreditCardService,
	}
}

//Fees list all fees from personal credirt cards all institutuons
func (r *report) Fees() *[]OutputPrice {
	institutions, err := r.institutionService.List()
	if err != nil {
		panic(err)
	}
	var entries []OutputPrice

	for _, institution := range institutions {
		iEntries := getFromInstitution(r.personalCreditCardService, &institution)
		entries = append(entries, iEntries...)
	}
	return &entries
}

func getFromInstitution(personalCreditCardService personalcreditcard.Service, institution *dtos.Institution) []OutputPrice {
	var entries []OutputPrice
	var accumulator []personalcreditcard.Entity

	result, pagination, _ := personalCreditCardService.FindByInstitution(institution.ID, 1)
	accumulator = append(accumulator, result...)
	for i := 2; i < pagination.Total; i++ {
		pageResult, _, _ := personalCreditCardService.FindByInstitution(institution.ID, i)
		accumulator = append(accumulator, pageResult...)
	}

	for _, creditCard := range accumulator {
		services := creditCard.Fees.Services
		for _, service := range services {
			for _, price := range service.Prices {
				entry := newOutputPrice(institution, creditCard.Name, &service, &price)
				entries = append(entries, *entry)
			}
		}
	}

	return entries
}

/*
//OutputPrice output price struct
type OutputPrice struct {
	InstitutionName string
	CreditCardName  string
	Name            string
	Code            string
	Interval        string
	Value           string
	Currency        string
}

func newOutputPrice(institution *dtos.Institution, creditCard *personalcreditcard.Entity, service *subentities.FeeService, price *subentities.ServicePrice) *OutputPrice {
	return &OutputPrice{
		InstitutionName: institution.Name,
		CreditCardName:  creditCard.Name,
		Name:            service.Name,
		Code:            service.Code,
		Interval:        price.Interval,
		Value:           price.Value,
		Currency:        price.Currency,
	}
}
*/
