package domain

import (
	"fmt"
	"openbankingcrawler/domain/institution"
	"os"
	"testing"

	"github.com/go-bongo/bongo"
	"github.com/mongo-go/testdb"
)

var testDb *testdb.TestDB

func Test_NewInstitutionRepository(t *testing.T) {

	repo := setup()

	entity := institution.NewEntity("my namessss")

	savedEntity, err := repo.Save(*entity)
	if err != nil {
		t.Error("Could not save institution")
	}
	if savedEntity.Name != entity.Name {
		t.Error("Wrong name")
	}

}

func setup() institution.Repository {
	coll := createCollection()

	repo := institution.NewRepository(coll)

	return repo
}

func createCollection() *bongo.Collection {

	connectionString := os.Getenv("DBHOST")
	if connectionString == "" {
		connectionString = "localhost"
	}

	database := os.Getenv("DBNAME")
	if database == "" {
		database = "openbankingcrawler_test_db"
	}

	config := &bongo.Config{
		ConnectionString: connectionString,
		Database:         database,
	}

	fmt.Printf("Connect to database %s", database)
	connection, _ := bongo.Connect(config)

	coll := connection.Collection("institution")

	coll.Delete(nil)

	return coll
}
