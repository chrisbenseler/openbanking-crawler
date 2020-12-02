package institution

import (
	"openbankingcrawler/common"

	"github.com/globalsign/mgo/bson"
	"github.com/go-bongo/bongo"
)

//Repository interface
type Repository interface {
	Save(Entity) (*Entity, common.CustomError)
	FindByName(string) (*Entity, common.CustomError)
	Delete(Entity) common.CustomError
	Find(string) (*Entity, common.CustomError)
}

type institutionRepository struct {
	dao *bongo.Collection
}

//NewRepository create a new repository for institution
func NewRepository(dao *bongo.Collection) Repository {

	return &institutionRepository{
		dao: dao,
	}
}

//Save save an entity
func (r *institutionRepository) Save(entity Entity) (*Entity, common.CustomError) {
	err := r.dao.Save(&entity)

	if err != nil {
		return nil, common.NewInternalServerError("Error on database", err)
	}

	return &entity, nil
}

//FindByName find an entity by name
func (r *institutionRepository) FindByName(name string) (*Entity, common.CustomError) {

	entity := NewEntity("")

	err := r.dao.FindOne(bson.M{"name": name}, entity)

	if err != nil {

		return nil, common.NewNotFoundError("No institution found with name " + name)
	}

	return entity, nil
}

//Delete delete an institution
func (r *institutionRepository) Delete(institution Entity) common.CustomError {
	err := r.dao.DeleteDocument(&institution)

	if err != nil {
		return common.NewInternalServerError("Error on database", err)
	}

	return nil
}

//Find find an entity
func (r *institutionRepository) Find(id string) (*Entity, common.CustomError) {

	entity := NewEntityWithID(id)

	if entity == nil {
		return nil, common.NewBadRequestError("The provided id is not valid: " + id)
	}

	err := r.dao.FindById(entity.Id, entity)

	if err != nil {

		if err.Error() == "Document not found" {
			return nil, common.NewNotFoundError("No institution found for id " + id)
		}
		return nil, common.NewInternalServerError("", err)
	}

	return entity, nil
}
