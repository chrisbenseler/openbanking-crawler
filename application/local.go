package application

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"openbankingcrawler/dtos"
	"openbankingcrawler/interfaces"
	"openbankingcrawler/services"
	"os"
	"time"

	"github.com/go-bongo/bongo"
)

//IF IF struct
type IF struct {
	Name    string `json:"name"`
	BaseURL string `json:"baseURL"`
}

//IFs struct - to read from file
type IFs struct {
	Institutions []IF `json:"institutions"`
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

	institutionService,
		branchService,
		electronicChannelService := CreateBasicServices(connection)

	personalAccountService,
		personalLoanService,
		personalFinancingService,
		personalCreditCardService,
		businessAccountService,
		businessLoanService := CreateProductsServicesServices(connection)

	httpClient := http.Client{}
	crawler := services.NewCrawler(&httpClient)

	institutionInterface := interfaces.NewInstitution(
		institutionService, branchService, electronicChannelService,
		personalAccountService, personalLoanService, personalFinancingService, personalCreditCardService,
		businessAccountService, businessLoanService,
		crawler)

	ifs := readFile()

	for _, _if := range *ifs {

		ifDTO := dtos.Institution{Name: _if.Name, BaseURL: _if.BaseURL}

		savedIF, _ := institutionService.FindByName(_if.Name)

		if savedIF == nil {
			savedIF, _ = institutionService.Create(ifDTO)
		}
		institutionService.Update(dtos.Institution{Name: savedIF.Name, BaseURL: _if.BaseURL, ID: savedIF.ID})

		fmt.Println("Start crawl for", _if.Name)

		//go institutionInterface.UpdatePersonalCreditCards(savedIF.ID)
		//time.NewTimer(1 * time.Second)
		//go institutionInterface.UpdatePersonalLoans(savedIF.ID)
		//time.NewTimer(1 * time.Second)
		// go institutionInterface.UpdatePersonalAccounts(savedIF.ID)
		// time.NewTimer(1 * time.Second)
		// go institutionInterface.UpdateBusinessAccounts(savedIF.ID)
		// time.NewTimer(1 * time.Second)
		// go institutionInterface.UpdatePersonalFinancings(savedIF.ID)
		// time.NewTimer(1 * time.Second)
		go institutionInterface.UpdateBusinessLoans(savedIF.ID)
		time.NewTimer(1 * time.Second)
	}

	fmt.Scanln()
	fmt.Println("done")

}

func readFile() *[]IF {
	jsonFile, err := os.Open("./application/localsrc.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully opened src file")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var ifs IFs
	unmarshallErr := json.Unmarshal(byteValue, &ifs)
	if unmarshallErr != nil {
		panic(unmarshallErr)
	}
	return &ifs.Institutions
}
