package domain

import (
	"openbankingcrawler/domain/institution"
	"testing"
)

func Test_NewInstitutionEntity(t *testing.T) {

	entity := institution.NewEntity("some random name")
	if entity.Name != "some random name" {
		t.Error("Wrong name")
	}
}
