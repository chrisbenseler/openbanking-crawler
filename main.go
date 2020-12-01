package main

import (
	"fmt"
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
	institutionService := services.NewInstitution(institutionRepository)

	branchRepository := branch.NewRepository(connection.Collection("branch"))
	branchService := services.NewBranch(branchRepository)

	err := branchService.UpdateAll("any")

	if err != nil {
		fmt.Println(err)
	}

	institutionInterface := interfaces.NewInstitution(institutionService, branchService)

	router := gin.Default()
	ginConfig := cors.DefaultConfig()
	ginConfig.AllowAllOrigins = true
	ginConfig.AddAllowHeaders("Authorization", "Access-Control-Allow-Headers")
	ginConfig.AddExposeHeaders("Authorization")
	router.Use(cors.New(ginConfig))

	apiRoutes := router.Group("/api")
	apiRoutes.GET("/institutions/:id", func(c *gin.Context) {
		id := c.Param("id")

		institution, err := institutionInterface.Get(id)

		if err != nil {
			c.JSON(err.Status(), gin.H{"error": err.Message()})
			return
		}

		c.JSON(200, institution)
	})

	apiRoutes.POST("/institutions", func(c *gin.Context) {

		type institutionPayload struct {
			Name string `json:"name"`
		}

		var payload institutionPayload

		c.BindJSON(&payload)

		institution, err := institutionInterface.Create(payload.Name)

		if err != nil {
			c.JSON(err.Status(), gin.H{"error": err.Message()})
			return
		}

		c.JSON(201, institution)
	})

	router.Run(":3000")

}
