package services

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"openbankingcrawler/common"
	"openbankingcrawler/domain/branch"
	"openbankingcrawler/domain/electronicchannel"
)

//Crawler service
type Crawler interface {
	Branches(string) (*[]branch.Entity, common.CustomError)
	ElectronicChannels(string, int, []electronicchannel.Entity) (*[]electronicchannel.Entity, common.CustomError)
}

type crawler struct {
	httpClient *http.Client
}

//NewCrawler create a new service for crawl
func NewCrawler(http *http.Client) Crawler {

	return &crawler{
		httpClient: http,
	}
}

//Branches crawl branches from institution
func (s *crawler) Branches(baseURL string) (*[]branch.Entity, common.CustomError) {

	resp, err := s.httpClient.Get(baseURL + "/open-banking/electronicChannels/v1/branches")

	if err != nil {
		return nil, common.NewInternalServerError("Unable to crawl branches from institution", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, common.NewInternalServerError("Unable to crawl branches from institution", err)
	}

	branchJSONData := &branchJSON{}

	jsonUnmarshallErr := json.Unmarshal(body, &branchJSONData)

	if jsonUnmarshallErr != nil {

		return nil, common.NewInternalServerError("Unable to unmarshall data", jsonUnmarshallErr)
	}

	companies := branchJSONData.Data.Brand.Companies[0]

	return &companies.Branches, nil

}

//ElectronicChannels crawl electronicChannels from institution
func (s *crawler) ElectronicChannels(baseURL string, page int, accumulator []electronicchannel.Entity) (*[]electronicchannel.Entity, common.CustomError) {

	resp, err := s.httpClient.Get(baseURL + "/open-banking/channels/v1/electronic-channels")

	if err != nil {
		return nil, common.NewInternalServerError("Unable to crawl electronicchannel from institution", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	jsonData := &electronicChannelJSON{}

	metaInfo := &metaInfoJSON{}
	json.Unmarshal(body, &metaInfo)

	jsonUnmarshallErr := json.Unmarshal(body, &jsonData)

	if jsonUnmarshallErr != nil {
		return nil, common.NewInternalServerError("Unable to unmarshall data", jsonUnmarshallErr)
	}

	companies := accumulator

	for i := range jsonData.Data.Brand.Companies {
		company := jsonData.Data.Brand.Companies[i]
		result := company.ElectronicChannels
		companies = append(companies, result...)
	}

	if metaInfo.Meta.TotalPages > page {
		return s.ElectronicChannels(baseURL, page+1, companies)
	}

	return &companies, nil

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

type electronicChannelJSON struct {
	Data struct {
		Brand struct {
			Companies []struct {
				ElectronicChannels []electronicchannel.Entity `json:"electronicChannels"`
			} `json:"companies"`
		} `json:"brand"`
	} `json:"data"`
}

type metaInfoJSON struct {
	Meta struct {
		TotalRecords int `json:"totalRecords"`
		TotalPages   int `json:"totalPages"`
	} `json:"meta"`
}
