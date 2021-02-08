package personalcreditcard

import (
	"openbankingcrawler/domain/subentities"

	"github.com/go-bongo/bongo"
)

//Entity branch entity
type Entity struct {
	bongo.DocumentBase `bson:",inline"`
	InstitutionID      string                               `json:"institutionid"`
	Name               string                               `json:"name"`
	Fees               subentities.Fees                     `json:"fees"`
	RewardsProgram     subentities.RewardsProgram           `json:"rewardsProgram"`
	Identification     subentities.CreditCardIdentification `json:"identification"`
	Interests          subentities.CreditCardInterests      `json:"interest"`
	OtherCredits       subentities.CreditCardOthers         `json:"otherCredits"`
}

//NewEntity create a new personal load entity
func NewEntity(institutionID string) *Entity {

	return &Entity{
		InstitutionID: institutionID,
	}
}
