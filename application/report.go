package application

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"openbankingcrawler/interfaces/report"
	"os"

	"github.com/go-bongo/bongo"
	"github.com/mohae/struct2csv"
)

//NewReport create a new report application
func NewReport() {

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
		_,
		_ := CreateBasicServices(connection)

	_,
		_,
		_,
		personalCreditCardService,
		_,
		_, _ := CreateProductsServicesServices(connection)

	path := "./"
	filename := "report"

	personalCreditCardReportInterface := report.NewPersonalCreditCard(institutionService, personalCreditCardService)
	filename = "report_personalcreditcard"

	output := personalCreditCardReportInterface.Fees()

	buff := &bytes.Buffer{}
	w := struct2csv.NewWriter(buff)
	err := w.WriteStructs(output)

	if err != nil {
		fmt.Println(err)
	}

	writeErr := ioutil.WriteFile(path+filename+".csv", buff.Bytes(), 0644)
	if writeErr != nil {
		panic(writeErr)
	}

	fmt.Scanln()
	fmt.Println("done")
}
