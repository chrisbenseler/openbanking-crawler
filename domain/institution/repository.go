package institution

import (
	"github.com/go-bongo/bongo"
	"go.mongodb.org/mongo-driver/bson"
)

//Repository interface
type Repository interface {
	Save(Entity) error
	FindByName(string) (*Entity, error)
	Delete(Entity) error
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
func (r *institutionRepository) Save(entity Entity) error {
	return r.dao.Save(&entity)
}

//FindByName find an entity by name
func (r *institutionRepository) FindByName(name string) (*Entity, error) {

	entity := NewEntity("")

	err := r.dao.FindOne(bson.M{"name": name}, entity)

	if err != nil {
		return nil, err
	}

	return entity, nil
}

//Delete delete an institution
func (r *institutionRepository) Delete(institution Entity) error {
	return r.dao.DeleteDocument(&institution)
}
