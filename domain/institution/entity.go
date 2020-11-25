package institution

import (
	"github.com/globalsign/mgo/bson"
	"github.com/go-bongo/bongo"
)

//Entity institution entity
type Entity struct {
	bongo.DocumentBase `bson:",inline"`
	Name               string `json:"string"`
}

//NewEntity create a new institution entity
func NewEntity(name string) *Entity {

	return &Entity{
		Name: name,
	}
}

//NewEntityWithID create an institution entity with ID
func NewEntityWithID(id string) *Entity {

	if !bson.IsObjectIdHex(id) {
		return nil
	}

	entity := &Entity{}
	bson.ObjectIdHex(id)
	entity.SetId(bson.ObjectIdHex(id))

	return entity
}
