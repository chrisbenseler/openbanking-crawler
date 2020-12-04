package channel

import (
	"testing"
)

func Test_NewEntity(t *testing.T) {

	entity := NewEntity("some random id")

	if entity.InstitutionID != "some random id" {
		t.Error("Wrong InstitutionID")
	}

}
