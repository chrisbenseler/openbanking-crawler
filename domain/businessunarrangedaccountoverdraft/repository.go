package businessunarrangedaccountoverdraft

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

type businessUnarrangedAccountOverdraft struct {
	dao *bongo.Collection
}

//NewRepository create a new repository for businessUnarrangedAccountOverdraft
func NewRepository(dao *bongo.Collection) Repository {
	return &businessUnarrangedAccountOverdraft{
		dao: dao,
	}
}

//Save save an entity
func (r *businessUnarrangedAccountOverdraft) Save(entity Entity) error {
	return r.dao.Save(&entity)
}

//DeleteMany delete all branches from an institution
func (r *businessUnarrangedAccountOverdraft) DeleteMany(institutionID string) error {
	_, err := r.dao.Delete(bson.M{"institutionid": institutionID})
	return err
}

//FindByInstitution find all businessLoan from an institution
func (r *businessUnarrangedAccountOverdraft) FindByInstitution(institutionID string, page int) ([]Entity, *subentities.Pagination, common.CustomError) {
	results := r.dao.Find(bson.M{"institutionid": institutionID})

	info, _ := results.Paginate(25, page)

	if results.Error != nil {
		return nil, nil, common.NewInternalServerError("Error on database", results.Error)
	}

	businessUnarrangedAccountOverdrafts := make([]Entity, info.RecordsOnPage)

	for i := 0; i < info.RecordsOnPage; i++ {
		_ = results.Next(&businessUnarrangedAccountOverdrafts[i])
	}

	pagination := subentities.Pagination{Total: info.TotalPages, Current: info.Current}

	return businessUnarrangedAccountOverdrafts, &pagination, nil
}
