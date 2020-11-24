package institution

import (
	"testing"
)

func Test_Institution(t *testing.T) {

	institution := NewEntity("name")

	if institution == nil {
		t.Error("Coulnd not create an institution entity")
	}

}
