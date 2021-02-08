package interfaces

import (
	"openbankingcrawler/common"
	"openbankingcrawler/domain/branch"
	"openbankingcrawler/domain/electronicchannel"
	"openbankingcrawler/domain/subentities"
)

//ChannelsInterface interface
type ChannelsInterface interface {
	GetBranches(string, int) ([]branch.Entity, *subentities.Pagination, common.CustomError)
	GetElectronicChannels(string, int) ([]electronicchannel.Entity, *subentities.Pagination, common.CustomError)
}

type channelsInterface struct {
	branchService            branch.Service
	electronicChannelService electronicchannel.Service
}

//NewChannels create a new interface for channels
func NewChannels(branchService branch.Service, electronicChannelService electronicchannel.Service) ChannelsInterface {

	return &channelsInterface{
		branchService:            branchService,
		electronicChannelService: electronicChannelService,
	}
}

//GetBranches get branches from institutution
func (b *channelsInterface) GetBranches(id string, page int) ([]branch.Entity, *subentities.Pagination, common.CustomError) {
	return b.branchService.FindByInstitution(id, page)
}

//GetElectronicChannels get electronicChannels from institutution
func (b *channelsInterface) GetElectronicChannels(id string, page int) ([]electronicchannel.Entity, *subentities.Pagination, common.CustomError) {
	return b.electronicChannelService.FindByInstitution(id, page)
}
