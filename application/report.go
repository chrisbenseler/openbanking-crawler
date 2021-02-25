package application

import (
	"bytes"
	"errors"
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

	fmt.Printf("Connect to database %s %s", database, "\n")
	connection, dbErr := bongo.Connect(config)

	if dbErr != nil {
		panic(dbErr)
	}

	institutionService,
		_,
		_ := CreateBasicServices(connection)

	_,
		_,
		_,
		_,
		personalCreditCardService,
		_,
		_,
		_,
		_,
		_,
		_,
		_ := CreateProductsServicesServices(connection)

	path := "./outputs/"
	filename := "report"

	reportType := os.Getenv("TYPE")

	writeErr := errors.New("No report type defined")

	if reportType == "PERSONAL_CC" {
		personalCreditCardReportInterface := report.NewPersonalCreditCard(institutionService, personalCreditCardService)
		filename = "report_personalcreditcard"
		output := *personalCreditCardReportInterface.Fees()

		writeErr = write(output, path, filename)
	}

	if writeErr != nil {
		panic(writeErr)
	}

	fmt.Println("Report done")
}

func write(st interface{}, path string, filename string) error {
	buff := &bytes.Buffer{}
	w := struct2csv.NewWriter(buff)
	err := w.WriteStructs(st)

	if err != nil {
		return err
	}

	writeErr := ioutil.WriteFile(path+filename+".csv", buff.Bytes(), 0644)
	if writeErr != nil {
		panic(writeErr)
	}

	return nil
}
