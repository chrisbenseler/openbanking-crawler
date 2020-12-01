package institution

import (
	"testing"
)

func Test_NewEntity(t *testing.T) {

	entity := NewEntity("some random name")

	if entity.Name != "some random name" {
		t.Error("Wrong name")
	}

}
