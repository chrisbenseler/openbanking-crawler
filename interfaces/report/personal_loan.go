package report

import (
	"openbankingcrawler/domain/institution"
	"openbankingcrawler/domain/personalloan"
	"openbankingcrawler/dtos"
)

//PersonalLoanReport Report interface for personal loan
type PersonalLoanReport interface {
	PersonalLoanFees() *[]OutputPrice
}

type personalLoanReport struct {
	institutionService  institution.Service
	personalLoanService personalloan.Service
}

//NewPersonalLoan create a new report interface
func NewPersonalLoan(institutionService institution.Service, personalLoanService personalloan.Service) PersonalLoanReport {
	return &personalLoanReport{
		institutionService:  institutionService,
		personalLoanService: personalLoanService,
	}
}

//Fees list all fees from personal loans from all institutuons
func (r *personalLoanReport) PersonalLoanFees() *[]OutputPrice {
	institutions, err := r.institutionService.List()
	if err != nil {
		panic(err)
	}
	var entries []OutputPrice

	for _, institution := range institutions {
		iEntries := getAllPersonalLoansFromInstitution(r.personalLoanService, &institution)
		entries = append(entries, iEntries...)
	}
	return &entries
}

func getAllPersonalLoansFromInstitution(personalLoanService personalloan.Service, institution *dtos.Institution) []OutputPrice {
	var entries []OutputPrice
	var accumulator []personalloan.Entity

	result, pagination, _ := personalLoanService.FindByInstitution(institution.ID, 1)

	accumulator = append(accumulator, result...)
	for i := 2; i < pagination.Total; i++ {
		pageResult, _, _ := personalLoanService.FindByInstitution(institution.ID, i)
		accumulator = append(accumulator, pageResult...)
	}

	for _, personalLoan := range accumulator {
		services := personalLoan.Fees.Services

		for _, service := range services {
			for _, price := range service.Prices {

				entry := newOutputPrice(institution, personalLoan.Type, &service, &price)
				entries = append(entries, *entry)
			}
		}
	}

	return entries
}
