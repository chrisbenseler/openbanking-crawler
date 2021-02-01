package interfaces

import (
	"openbankingcrawler/common"
	"openbankingcrawler/domain/electronicchannel"
)

//ElectronicChannelInterface interface
type ElectronicChannelInterface interface {
	GetFromInstitution(string) ([]electronicchannel.Entity, common.CustomError)
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
func (c *electronicChannelInterface) GetFromInstitution(id string) ([]electronicchannel.Entity, common.CustomError) {
	return c.electronicChannelService.FindByInstitution(id)
}
