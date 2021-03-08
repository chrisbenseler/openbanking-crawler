package report

import (
	"openbankingcrawler/domain/subentities"
	"openbankingcrawler/dtos"
)

//OutputPrice output price struct
type OutputPrice struct {
	InstitutionName string
	Label           string
	Name            string
	Code            string
	Interval        string
	Value           string
	Currency        string
}

func newOutputPrice(institution *dtos.Institution, label string, service *subentities.FeeService, price *subentities.ServicePrice) *OutputPrice {
	return &OutputPrice{
		InstitutionName: institution.Name,
		Label:           label,
		Name:            service.Name,
		Code:            service.Code,
		Interval:        price.Interval,
		Value:           price.Value,
		Currency:        price.Currency,
	}
}
