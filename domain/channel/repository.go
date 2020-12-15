package channel

import (
	"openbankingcrawler/common"

	"github.com/globalsign/mgo/bson"
	"github.com/go-bongo/bongo"
)

//Repository interface
type Repository interface {
	Save(Entity) error
	DeleteMany(string) error
	FindByInstitution(string) ([]Entity, common.CustomError)
}

type channelRepository struct {
	dao *bongo.Collection
}

//NewRepository create a new repository for channel
func NewRepository(dao *bongo.Collection) Repository {

	return &channelRepository{
		dao: dao,
	}
}

//Save save an entity
func (r *channelRepository) Save(entity Entity) error {
	return r.dao.Save(&entity)
}

//DeleteMany delete all branches from an institution
func (r *channelRepository) DeleteMany(institutionID string) error {
	_, err := r.dao.Delete(bson.M{"institutionid": institutionID})
	return err
}

//FindByInstitution find all channels from an institution
func (r *channelRepository) FindByInstitution(institutionID string) ([]Entity, common.CustomError) {
	results := r.dao.Find(bson.M{"institutionid": institutionID})

	if results.Error != nil {
		return nil, common.NewInternalServerError("Error on database", results.Error)
	}

	channel := &Entity{}

	channels := make([]Entity, 0)

	for results.Next(channel) {
		channels = append(channels, *channel)
	}

	return channels, nil
}
