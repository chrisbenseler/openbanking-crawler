package main

import (
	"fmt"
	"openbankingcrawler/adapters"
	"openbankingcrawler/domain/branch"
	"openbankingcrawler/domain/institution"
	"openbankingcrawler/interfaces"
	"openbankingcrawler/services"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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

	if dbErr != nil {
		fmt.Println(dbErr)
	}

	institutionRepository := institution.NewRepository(connection.Collection("institution"))
	institutionService := institution.NewService(institutionRepository)

	branchRepository := branch.NewRepository(connection.Collection("branch"))
	branchService := branch.NewService(branchRepository)

	crawler := services.NewCrawler()

	institutionInterface := interfaces.NewInstitution(institutionService, branchService, crawler)

	router := gin.Default()
	ginConfig := cors.DefaultConfig()
	ginConfig.AllowAllOrigins = true
	ginConfig.AddAllowHeaders("Authorization", "Access-Control-Allow-Headers")
	ginConfig.AddExposeHeaders("Authorization")
	router.Use(cors.New(ginConfig))

	apiRoutes := router.Group("/api")

	controller := adapters.NewController(institutionInterface)

	apiRoutes.GET("/institutions/:id", controller.GetInstitution)
	apiRoutes.GET("/institutions/:id/branches/update", controller.UpdateInstitutionBranches)
	apiRoutes.POST("/institutions", controller.CreateInstitution)
	apiRoutes.PUT("/institutions/:id", controller.UpdateInstitution)

	router.Run(":3000")

}
