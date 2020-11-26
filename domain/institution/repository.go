package institution

import (
	"fmt"
	"openbankingcrawler/common"

	"github.com/go-bongo/bongo"
	"go.mongodb.org/mongo-driver/bson"
)

//Repository interface
type Repository interface {
	Save(Entity) (*Entity, error)
	FindByName(string) (*Entity, error)
	Delete(Entity) error
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
func (r *institutionRepository) Save(entity Entity) (*Entity, error) {
	err := r.dao.Save(&entity)

	if err != nil {
		return nil, err
	}

	return &entity, nil
}

//FindByName find an entity by name
func (r *institutionRepository) FindByName(name string) (*Entity, error) {

	entity := NewEntity("")

	err := r.dao.FindOne(bson.M{"name": name}, entity)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return entity, nil
}

//Delete delete an institution
func (r *institutionRepository) Delete(institution Entity) error {
	return r.dao.DeleteDocument(&institution)
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
