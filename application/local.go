package application

import (
	"fmt"
	"net/http"
	"openbankingcrawler/domain/branch"
	"openbankingcrawler/domain/electronicchannel"
	"openbankingcrawler/domain/institution"
	"openbankingcrawler/domain/personalcreditcard"
	"openbankingcrawler/domain/personalloan"
	"openbankingcrawler/dtos"
	"openbankingcrawler/interfaces"
	"openbankingcrawler/services"
	"os"

	"github.com/go-bongo/bongo"
)

//IF
type IF struct {
	name    string
	baseURL string
}

//NewLocal create a new web application
func NewLocal() {

	connectionString := os.Getenv("DBHOST")
	if connectionString == "" {
		connectionString = "localhost"
	}

	database := os.Getenv("DBNAME")
	if database == "" {
		database = "openbankingcrawlerlocal"
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

	institutionRepository, branchRepository, electronicChannelRepository, personalLoanRepository, personalCreditCardRepository := CreateRepositories(connection)

	institutionService := institution.NewService(institutionRepository)
	branchService := branch.NewService(branchRepository)
	electronicChannelService := electronicchannel.NewService(electronicChannelRepository)
	personalLoanService := personalloan.NewService(personalLoanRepository)
	personalCreditCardService := personalcreditcard.NewService(personalCreditCardRepository)

	httpClient := http.Client{}
	crawler := services.NewCrawler(&httpClient)

	institutionInterface := interfaces.NewInstitution(institutionService, branchService, electronicChannelService, personalLoanService, personalCreditCardService, crawler)

	/*
		authService := services.NewAuthService()

		httpClient := http.Client{}
		crawler := services.NewCrawler(&httpClient)


		channelsInterface := interfaces.NewChannels(branchService, electronicChannelService)
		productsServicesInterface := interfaces.NewProductsServicesInterface(personalLoanService, personalCreditCardService)

	*/

	ifs := []IF{
		IF{name: "banco bv", baseURL: "https://api-openbanking.bvopen.com.br"},
		IF{name: "itau", baseURL: "https://api.itau"},
		IF{name: "banco do brasil", baseURL: "https://opendata.api.bb.com.br"},
		IF{name: "banco safra", baseURL: "https://api.safra.com.br"},
	}

	for _, _if := range ifs {
		ifDTO := dtos.Institution{Name: _if.name, BaseURL: _if.baseURL}
		savedIF, err := institutionService.Create(ifDTO)

		if err != nil {
			fmt.Println(err)
		}

		institutionService.Update(dtos.Institution{Name: savedIF.Name, BaseURL: _if.baseURL, ID: savedIF.ID})

		go institutionInterface.UpdatePersonalCreditCards(savedIF.ID)
		go institutionInterface.UpdatePersonalLoans(savedIF.ID)

	}

	fmt.Scanln()
	fmt.Println("done")

}
