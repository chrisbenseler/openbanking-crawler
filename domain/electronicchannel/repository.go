package electronicchannel

import (
	"openbankingcrawler/common"
	"openbankingcrawler/domain/subentities"

	"github.com/globalsign/mgo/bson"
	"github.com/go-bongo/bongo"
)

//Repository interface
type Repository interface {
	Save(Entity) error
	DeleteMany(string) error
	FindByInstitution(string, int) ([]Entity, *subentities.Pagination, common.CustomError)
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
func (r *electronicChannelRepository) FindByInstitution(institutionID string, page int) ([]Entity, *subentities.Pagination, common.CustomError) {
	results := r.dao.Find(bson.M{"institutionid": institutionID})

	info, _ := results.Paginate(25, page)

	if results.Error != nil {
		return nil, nil, common.NewInternalServerError("Error on database", results.Error)
	}

	electronicChannels := make([]Entity, info.RecordsOnPage)

	for i := 0; i < info.RecordsOnPage; i++ {
		_ = results.Next(&electronicChannels[i])
	}

	pagination := subentities.Pagination{Total: info.TotalPages, Current: info.Current}

	return electronicChannels, &pagination, nil
}
