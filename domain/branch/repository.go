package branch

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/go-bongo/bongo"
)

//Repository interface
type Repository interface {
	Save(Entity) error
	DeleteMany(string) error
	FindByInstitution(string) []Entity
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
func (r *branchRepository) FindByInstitution(institutionID string) []Entity {
	results := r.dao.Find(bson.M{"institutionid": institutionID})

	branch := &Entity{}

	var branches []Entity

	for results.Next(branch) {
		fmt.Print(branch)
		branches = append(branches, *branch)
	}

	return branches
}
