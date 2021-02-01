package electronicchannel

import "github.com/go-bongo/bongo"

type electronicChannelIdentification struct {
	Type           string `json:"type"`
	AdditionalInfo string `json:"additionalInfo"`
	URL            string `json:"url"`
}

type electronicChannelService struct {
	Codes          []string `json:"codes"`
	AdditionalInfo string   `json:"additionalInfo"`
}

//Entity branch entity
type Entity struct {
	bongo.DocumentBase `bson:",inline"`
	InstitutionID      string                          `json:"institutionid"`
	Identification     electronicChannelIdentification `json:"identification"`
	Service            electronicChannelService        `json:"service"`
}

//NewEntity create a new electronicchannel entity
func NewEntity(institutionID string) *Entity {

	return &Entity{
		InstitutionID: institutionID,
	}
}
