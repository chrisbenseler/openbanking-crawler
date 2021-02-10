package personalaccount

import (
	"openbankingcrawler/domain/subentities"

	"github.com/go-bongo/bongo"
)

//Entity branch entity
type Entity struct {
	bongo.DocumentBase `bson:",inline"`
	InstitutionID      string `json:"institutionid"`
	Type               string `json:"type" bson:"type"`
	Fees               struct {
		PriorityServices []subentities.FeeService `json:"priorityServices"`
		OtherServices    []subentities.FeeService `json:"otherServices"`
	} `json:"fees"`
	ServiceBundles         []subentities.ServiceBundle `json:"serviceBundles" bson:"serviceBundles"`
	OpeningClosingChannels []string                    `json:"openingClosingChannels" bson:"openingClosingChannels"`
	AdditionalInfo         string                      `json:"additionalInfo" bson:"additionalInfo"`
	TransactionsMethods    []string                    `json:"transactionsMethods" bson:"transactionsMethods"`
	TermsConditions        subentities.TermsConditions `json:"termsConditions" bson:"termsConditions"`
	IncomeRate             subentities.IncomeRate      `json:"incomeRate" bson:"incomeRate"`
}

//NewEntity create a new personal load entity
func NewEntity(institutionID string) *Entity {

	return &Entity{
		InstitutionID: institutionID,
	}
}
