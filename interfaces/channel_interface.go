package interfaces

import (
	"openbankingcrawler/common"
	"openbankingcrawler/domain/channel"
)

//ChannelInterface interface
type ChannelInterface interface {
	GetFromInstitution(string) ([]channel.Entity, common.CustomError)
}

type channelInterface struct {
	channelService channel.Service
}

//NewChannel create a new interface for channels
func NewChannel(channelService channel.Service) ChannelInterface {

	return &channelInterface{
		channelService: channelService,
	}
}

//GetFromInstitution get channels from institutution
func (c *channelInterface) GetFromInstitution(id string) ([]channel.Entity, common.CustomError) {
	return c.channelService.FindByInstitution(id)
}
