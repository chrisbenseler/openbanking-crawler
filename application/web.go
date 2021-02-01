package application

import (
	"fmt"
	"net/http"
	"openbankingcrawler/adapters"
	"openbankingcrawler/domain/branch"
	"openbankingcrawler/domain/electronicchannel"
	"openbankingcrawler/domain/institution"
	"openbankingcrawler/interfaces"
	"openbankingcrawler/services"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-bongo/bongo"
)

//NewWeb create a new web application
func NewWeb() {

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

	institutionRepository, branchRepository, electronicChannelRepository := createRepositories(connection)

	institutionService := institution.NewService(institutionRepository)
	branchService := branch.NewService(branchRepository)
	electronicChannelService := electronicchannel.NewService(electronicChannelRepository)
	authService := services.NewAuthService()

	httpClient := http.Client{}
	crawler := services.NewCrawler(&httpClient)

	institutionInterface := interfaces.NewInstitution(institutionService, branchService, electronicChannelService, crawler)
	branchInterface := interfaces.NewBranch(branchService)
	electronicChannelInterface := interfaces.NewChannel(electronicChannelService)

	router := gin.Default()
	ginConfig := cors.DefaultConfig()
	ginConfig.AllowAllOrigins = true
	ginConfig.AddAllowHeaders("Authorization", "Access-Control-Allow-Headers")
	ginConfig.AddExposeHeaders("Authorization")
	router.Use(cors.New(ginConfig))

	apiRoutes := router.Group("/api")

	controller := adapters.NewController(institutionInterface, branchInterface, electronicChannelInterface)

	authController := adapters.NewAuthenticateController(authService)

	authRequired := authMiddleware(authService)

	apiRoutes.GET("/institutions", controller.ListAllInstitutions)
	apiRoutes.GET("/institutions/:id", controller.GetInstitution)
	apiRoutes.GET("/institutions/:id/branches", controller.GetBranches)
	apiRoutes.GET("/institutions/:id/electronic-channels", controller.GetElectronicChannels)

	apiRoutes.PUT("/institutions/:id/branches/update", authRequired, controller.UpdateInstitutionBranches)
	apiRoutes.PUT("/institutions/:id/electronic-channels/update", authRequired, controller.UpdateInstitutionElectronicChannels)
	apiRoutes.POST("/institutions", authRequired, controller.CreateInstitution)
	apiRoutes.PUT("/institutions/:id", authRequired, controller.UpdateInstitution)

	apiRoutes.POST("/auth/signin", authController.SignIn)

	//router.Use(static.Serve("/open-banking", static.LocalFile("../mocks/open-banking", false)))

	router.Static("/open-banking", "./mocks/open-banking")

	router.Run(":" + port)
}

func createRepositories(connection *bongo.Connection) (institution.Repository, branch.Repository, electronicchannel.Repository) {
	institutionRepository := institution.NewRepository(connection.Collection("institution"))
	branchRepository := branch.NewRepository(connection.Collection("branch"))
	electronicChannelRepository := electronicchannel.NewRepository(connection.Collection("electronicChannel"))

	return institutionRepository, branchRepository, electronicChannelRepository
}

func authMiddleware(authService services.Auth) func(*gin.Context) {
	f := func(c *gin.Context) {
		_, validateErr := authService.ValidateAccessToken(c.Request)
		if validateErr != nil {
			c.AbortWithStatusJSON(validateErr.Status(), gin.H{"error": validateErr.Message()})
			return
		}
		c.Next()
	}
	return f
}
