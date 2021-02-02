package branch

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

type branchRepository struct {
	dao *bongo.Collection
}

//NewRepository create a new repository for branch
func NewRepository(dao *bongo.Collection) Repository {

	return &branchRepository{
		dao: dao,
	}
}

//Save save an entity
func (r *branchRepository) Save(entity Entity) error {
	return r.dao.Save(&entity)
}

//DeleteMany delete all branches from an institution
func (r *branchRepository) DeleteMany(institutionID string) error {
	_, err := r.dao.Delete(bson.M{"institutionid": institutionID})
	return err
}

//FindByInstitution find all branches from an institution
func (r *branchRepository) FindByInstitution(institutionID string, page int) ([]Entity, *subentities.Pagination, common.CustomError) {
	results := r.dao.Find(bson.M{"institutionid": institutionID})

	info, _ := results.Paginate(25, page)

	if results.Error != nil {
		return nil, nil, common.NewInternalServerError("Error on database", results.Error)
	}

	// branch := &Entity{}

	branches := make([]Entity, info.RecordsOnPage)

	for i := 0; i < info.RecordsOnPage; i++ {
		_ = results.Next(&branches[i])
	}

	pagination := subentities.Pagination{Total: info.TotalPages, Current: info.Current}

	return branches, &pagination, nil

	//	for info.RecordsOnPage.Next(branch) {
	//		branches = append(branches, *branch)
	//	}

	// return branches, nil
}
