package main

import (
	"fmt"
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

	/*
		collection := connection.Collection("institution")
		institutionRepository := institution.NewRepository(collection)
		institutionService := institution.NewService(institutionRepository)

		err := institutionService.Delete("5fbbb65309114e01ce8d2cd6")

		if err != nil {
			fmt.Println(err)
		}
	*/

	/*

		collection := connection.Collection("branch")
		branchRepository := branch.NewRepository(collection)
		branchService := branch.NewService(branchRepository)

		err := branchService.UpdateAll("any")

		if err != nil {
			fmt.Println(err)
		}
	*/

}
