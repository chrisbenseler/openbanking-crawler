package personalfinancing

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
		Services []subentities.FeeService `json:"services"`
	} `json:"fees"`
	InterestRates   []subentities.Rates `json:"interestRates"`
	TermsConditions string              `json:"termsConditions"`
}

//NewEntity create a new personal load entity
func NewEntity(institutionID string) *Entity {

	return &Entity{
		InstitutionID: institutionID,
	}
}
