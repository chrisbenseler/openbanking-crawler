package report

import (
	"openbankingcrawler/domain/institution"
	"openbankingcrawler/domain/personalcreditcard"
)

//Report Report interface for personal credit card
type Report interface {
	Fees() []Output
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

//Fees list all fees from all institutuons
func (r *report) Fees() []Output {
	institutions, err := r.institutionService.List()
	if err != nil {
		panic(err)
	}

	var entries []Output

	for _, institution := range institutions {
		// fmt.Println(institution)

		var accumulator []personalcreditcard.Entity

		result, pagination, _ := r.personalCreditCardService.FindByInstitution(institution.ID, 1)
		accumulator = append(accumulator, result...)
		for i := 2; i < pagination.Total; i++ {
			pageResult, _, _ := r.personalCreditCardService.FindByInstitution(institution.ID, i)
			accumulator = append(accumulator, pageResult...)
		}

		for _, creditCard := range accumulator {
			services := creditCard.Fees.Services

			for _, service := range services {
				for _, price := range service.Prices {

					entry := Output{
						InstitutionName: institution.Name,
						CreditCardName:  creditCard.Name,
						Name:            service.Name,
						Code:            service.Code,
						Interval:        price.Interval,
						Value:           price.Value,
						Currency:        price.Currency,
					}

					entries = append(entries, entry)
				}
			}
		}

	}

	return entries

}

//Output output struct
type Output struct {
	InstitutionName string
	CreditCardName  string
	Name            string
	Code            string
	Interval        string
	Value           string
	Currency        string
}
