package businessunarrangedaccountoverdraft

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
	} `json:"fees"`
	InterestRates   []subentities.Rates `json:"interestRates" bson:"interestRates"`
	TermsConditions string              `json:"termsConditions" bson:"termsConditions"`
}

//NewEntity create a new business load entity
func NewEntity(institutionID string) *Entity {

	return &Entity{
		InstitutionID: institutionID,
	}
}
