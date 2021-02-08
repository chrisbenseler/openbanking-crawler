package personalloan

import (
	"openbankingcrawler/domain/subentities"

	"github.com/go-bongo/bongo"
)

//Entity branch entity
type Entity struct {
	bongo.DocumentBase `bson:",inline"`
	InstitutionID      string                      `json:"institutionid"`
	Type               string                      `json:"type"`
	Fees               subentities.Fees            `json:"fees"`
	InterestRates      []subentities.InterestRates `json:"interestRates"`
	TermsConditions    string                      `json:"termsConditions"`
}

//NewEntity create a new personal load entity
func NewEntity(institutionID string) *Entity {

	return &Entity{
		InstitutionID: institutionID,
	}
}
