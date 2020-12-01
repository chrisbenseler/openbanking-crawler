package services

import (
	"encoding/json"
	"io/ioutil"
	"openbankingcrawler/domain/branch"
	"os"
)

//Crawler service
type Crawler interface {
	Crawl(string) (*[]branch.Entity, error)
}

type crawler struct {
	// repository branch.Repository
}

//NewCrawler create a new service for crawl
func NewCrawler() Crawler {

	return &crawler{
		//repository: repository,
	}
}

type branchJSON struct {
	Data struct {
		Brand struct {
			Companies []struct {
				Branches []branch.Entity `json:"branches"`
			} `json:"companies"`
		} `json:"brand"`
	} `json:"data"`
}

type branchesList struct {
	Branches []branch.Entity
}

//Crawl crawl branches from institution
func (s *crawler) Crawl(InstitutionID string) (*[]branch.Entity, error) {

	jsonFile, err := os.Open("./domain/branch/branches.json")

	if err != nil {
		return nil, err
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	branchJSONData := &branchJSON{}

	jsonUnmarshallErr := json.Unmarshal(byteValue, &branchJSONData)

	if jsonUnmarshallErr != nil {
		return nil, jsonUnmarshallErr
	}

	companies := branchJSONData.Data.Brand.Companies[0]

	return &companies.Branches, nil

}
