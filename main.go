package main

import (
	"fmt"
	"openbankingcrawler/domain/branch"
	"openbankingcrawler/domain/institution"
	"openbankingcrawler/interfaces"
	"openbankingcrawler/services"
	"os"

	"github.com/go-bongo/bongo"
)

func main() {
	fmt.Println("Start Open banking")

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	connectionString := os.Getenv("DBHOST")
	if connectionString == "" {
		connectionString = "localhost"
	}

	database := os.Getenv("DBNAME")
	if database == "" {
		database = "openbankingcrawler"
	}

	config := &bongo.Config{
		ConnectionString: connectionString,
		Database:         database,
	}

	fmt.Printf("Connect to database %s", database)
	connection, dbErr := bongo.Connect(config)

	fmt.Println(dbErr)

	if dbErr != nil {
		fmt.Println(dbErr)
	}

	institutionRepository := institution.NewRepository(connection.Collection("institution"))
	institutionService := services.NewInstitution(institutionRepository)

	branchRepository := branch.NewRepository(connection.Collection("branch"))
	branchService := services.NewBranch(branchRepository)

	err := branchService.UpdateAll("any")

	if err != nil {
		fmt.Println(err)
	}

	institutionInterface := interfaces.NewInstitution(institutionService, branchService)
	fmt.Println(institutionInterface)

	institution, err := institutionInterface.Get("5fbe441109114eb2c238017a")

	fmt.Println(institution, err)
}
