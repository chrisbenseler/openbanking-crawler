package electronicchannel

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

type electronicChannelRepository struct {
	dao *bongo.Collection
}

//NewRepository create a new repository for electronicchannel
func NewRepository(dao *bongo.Collection) Repository {
	return &electronicChannelRepository{
		dao: dao,
	}
}

//Save save an entity
func (r *electronicChannelRepository) Save(entity Entity) error {
	return r.dao.Save(&entity)
}

//DeleteMany delete all branches from an institution
func (r *electronicChannelRepository) DeleteMany(institutionID string) error {
	_, err := r.dao.Delete(bson.M{"institutionid": institutionID})
	return err
}

//FindByInstitution find all electronicChannels from an institution
func (r *electronicChannelRepository) FindByInstitution(institutionID string) ([]Entity, common.CustomError) {
	results := r.dao.Find(bson.M{"institutionid": institutionID})

	if results.Error != nil {
		return nil, common.NewInternalServerError("Error on database", results.Error)
	}

	electronicChannel := &Entity{}

	electronicChannels := make([]Entity, 0)

	for results.Next(electronicChannel) {
		electronicChannels = append(electronicChannels, *electronicChannel)
	}

	return electronicChannels, nil
}
