package application

import (
	"fmt"
	"net/http"
	"openbankingcrawler/adapters"
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

	institutionService,
		branchService,
		electronicChannelService,
		personalLoanService,
		personalCreditCardService,
		personalAccountService,
		businessAccountService := CreateServices(connection)

	authService := services.NewAuthService()

	httpClient := http.Client{}
	crawler := services.NewCrawler(&httpClient)

	institutionInterface := interfaces.NewInstitution(
		institutionService, branchService, electronicChannelService, personalLoanService, personalCreditCardService,
		personalAccountService, businessAccountService,
		crawler)

	channelsInterface := interfaces.NewChannels(branchService, electronicChannelService)
	productsServicesInterface := interfaces.NewProductsServicesInterface(personalLoanService, personalCreditCardService)

	router := gin.Default()
	ginConfig := cors.DefaultConfig()
	ginConfig.AllowAllOrigins = true
	ginConfig.AddAllowHeaders("Authorization", "Access-Control-Allow-Headers")
	ginConfig.AddExposeHeaders("Authorization")
	router.Use(cors.New(ginConfig))

	apiRoutes := router.Group("/api")

	controller := adapters.NewController(institutionInterface, channelsInterface, productsServicesInterface)

	authController := adapters.NewAuthenticateController(authService)

	authRequired := AuthMiddleware(authService)

	apiRoutes.GET("/health_check", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "OK"})
	})

	apiRoutes.GET("/institutions", controller.ListAllInstitutions)
	apiRoutes.GET("/institutions/:id", controller.GetInstitution)
	apiRoutes.GET("/institutions/:id/branches", controller.GetBranches)
	apiRoutes.GET("/institutions/:id/electronic-channels", controller.GetElectronicChannels)
	apiRoutes.GET("/institutions/:id/personal-loans", controller.GetPersonalLoans)
	apiRoutes.GET("/institutions/:id/personal-credit-cards", controller.GetPersonalCreditCards)

	apiRoutes.PUT("/institutions/:id/branches/update", authRequired, controller.UpdateInstitutionBranches)
	apiRoutes.PUT("/institutions/:id/electronic-channels/update", authRequired, controller.UpdateInstitutionElectronicChannels)
	apiRoutes.PUT("/institutions/:id/personal-loans/update", authRequired, controller.UpdatePersonalLoans)
	apiRoutes.PUT("/institutions/:id/personal-credit-cards/update", authRequired, controller.UpdatePersonalCreditCards)

	apiRoutes.POST("/institutions", authRequired, controller.CreateInstitution)
	apiRoutes.PUT("/institutions/:id", authRequired, controller.UpdateInstitution)

	apiRoutes.POST("/auth/signin", authController.SignIn)

	//router.Use(static.Serve("/open-banking", static.LocalFile("../mocks/open-banking", false)))

	router.Static("/open-banking", "./mocks/open-banking")

	router.Run(":" + port)
}

//AuthMiddleware AuthMiddleware
func AuthMiddleware(authService services.Auth) func(*gin.Context) {
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
