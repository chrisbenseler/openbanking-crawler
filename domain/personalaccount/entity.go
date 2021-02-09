package personalaccount

import (
	"openbankingcrawler/domain/subentities"

	"github.com/go-bongo/bongo"
)

//Entity branch entity
type Entity struct {
	bongo.DocumentBase `bson:",inline"`
	InstitutionID      string `json:"institutionid"`
	Type               string `json:"type"`
	Fees               struct {
		PriorityServices []subentities.FeeService `json:"priorityServices"`
		OtherServices    []subentities.FeeService `json:"otherServices"`
	} `json:"fees"`
	ServiceBundles         []subentities.ServiceBundle `json:"serviceBundles"`
	OpeningClosingChannels []string                    `json:"openingClosingChannels"`
	AdditionalInfo         string                      `json:"additionalInfo"`
	TransactionsMethods    []string                    `json:"transactionsMethods"`
	TermsConditions        subentities.TermsConditions `json:"termsConditions"`
	IncomeRate             subentities.IncomeRate      `json:"incomeRate"`
}

//NewEntity create a new personal load entity
func NewEntity(institutionID string) *Entity {

	return &Entity{
		InstitutionID: institutionID,
	}
}
