package channel

import "github.com/go-bongo/bongo"

type channelIdentification struct {
	Type           string `json:"type"`
	AdditionalInfo string `json:"additionalInfo"`
	URL            string `json:"url"`
}

type channelService struct {
	Codes          []string `json:"codes"`
	AdditionalInfo string   `json:"additionalInfo"`
}

//Entity branch entity
type Entity struct {
	bongo.DocumentBase `bson:",inline"`
	InstitutionID      string                `json:"institutionid"`
	Identification     channelIdentification `json:"identification"`
	Service            channelService        `json:"service"`
}

//NewEntity create a new channel entity
func NewEntity(institutionID string) *Entity {

	return &Entity{
		InstitutionID: institutionID,
	}
}
