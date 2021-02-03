package interfaces

import (
	"openbankingcrawler/common"
	"openbankingcrawler/domain/electronicchannel"
	"openbankingcrawler/domain/subentities"
)

//ElectronicChannelInterface interface
type ElectronicChannelInterface interface {
	GetFromInstitution(string, int) ([]electronicchannel.Entity, *subentities.Pagination, common.CustomError)
}

type electronicChannelInterface struct {
	electronicChannelService electronicchannel.Service
}

//NewChannel create a new interface for electronicchannel
func NewChannel(electronicChannelService electronicchannel.Service) ElectronicChannelInterface {

	return &electronicChannelInterface{
		electronicChannelService: electronicChannelService,
	}
}

//GetFromInstitution get electronicChannels from institutution
func (c *electronicChannelInterface) GetFromInstitution(id string, page int) ([]electronicchannel.Entity, *subentities.Pagination, common.CustomError) {
	return c.electronicChannelService.FindByInstitution(id, page)
}
