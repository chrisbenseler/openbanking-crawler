package application

import (
	"fmt"
	"openbankingcrawler/domain/personalcreditcard"
	"os"
)

//NewReport create a new report application
func NewReport() {

	/*
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
			businessLoanService, businessFinancingService := CreateProductsServicesServices(connection)

		httpClient := http.Client{}
		crawler := services.NewCrawler(&httpClient)

		institutionInterface := interfaces.NewInstitution(
			institutionService, branchService, electronicChannelService,
			personalAccountService, personalLoanService, personalFinancingService, personalCreditCardService,
			businessAccountService, businessLoanService, businessFinancingService,
			crawler)

		institutions, _ := institutionInterface.ListAll()

		var output []toFile

		for _, institution := range institutions {
			// fmt.Println(institution)

			var accumulator []personalcreditcard.Entity

			result, pagination, _ := personalCreditCardService.FindByInstitution(institution.ID, 1)
			accumulator = append(accumulator, result...)
			for i := 2; i < pagination.Total; i++ {
				pageResult, _, _ := personalCreditCardService.FindByInstitution(institution.ID, i)
				accumulator = append(accumulator, pageResult...)
			}

			for _, creditCard := range accumulator {
				toFileEntry := toFile{Name: institution.Name, PersonalCreditCard: creditCard}
				output = append(output, toFileEntry)
			}
		}

		buff := &bytes.Buffer{}
		w := struct2csv.NewWriter(buff)
		err := w.WriteStructs(output)

		if err != nil {
			fmt.Println(err)
		}

		writeErr := ioutil.WriteFile("./report.csv", buff.Bytes(), 0644)
		if writeErr != nil {
			panic(writeErr)
		}

		fmt.Scanln()
		fmt.Println("done")
	*/

}

type toFile struct {
	Name               string
	PersonalCreditCard personalcreditcard.Entity
}

func printLines(filePath string, values []interface{}) error {
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()
	for _, value := range values {
		fmt.Fprintln(f, value) // print values to f, one per line
	}
	return nil
}

/*
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
*/
