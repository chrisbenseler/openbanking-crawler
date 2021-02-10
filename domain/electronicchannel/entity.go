package electronicchannel

import "github.com/go-bongo/bongo"

type electronicChannelIdentification struct {
	Type           string   `json:"type" bson:"type"`
	AdditionalInfo string   `json:"additionalInfo" bson:"additionalInfo"`
	URLS           []string `json:"urls"`
}

type electronicChannelService struct {
	Code           string `json:"code"`
	Name           string `json:"name"`
	AdditionalInfo string `json:"additionalInfo" bson:"additionalInfo"`
}

//Entity branch entity
type Entity struct {
	bongo.DocumentBase `bson:",inline"`
	InstitutionID      string                          `json:"institutionid"`
	Identification     electronicChannelIdentification `json:"identification"`
	Services           []electronicChannelService      `json:"services"`
}

//NewEntity create a new electronicchannel entity
func NewEntity(institutionID string) *Entity {

	return &Entity{
		InstitutionID: institutionID,
	}
}
