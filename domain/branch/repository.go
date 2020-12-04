package branch

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
func (r *branchRepository) FindByInstitution(institutionID string) ([]Entity, common.CustomError) {
	results := r.dao.Find(bson.M{"institutionid": institutionID})

	if results.Error != nil {
		return nil, common.NewInternalServerError("Error on database", results.Error)
	}

	branch := &Entity{}

	branches := make([]Entity, 0)

	for results.Next(branch) {
		branches = append(branches, *branch)
	}

	return branches, nil
}
