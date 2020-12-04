package services

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"openbankingcrawler/common"
	"openbankingcrawler/domain/branch"
	"openbankingcrawler/domain/channel"
)

//Crawler service
type Crawler interface {
	Branches(string) (*[]branch.Entity, common.CustomError)
	Channels(string) (*[]channel.Entity, common.CustomError)
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

	resp, err := s.httpClient.Get(baseURL + "/open-banking/channels/v1/branches")

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

//Channels crawl channels from institution
func (s *crawler) Channels(baseURL string) (*[]channel.Entity, common.CustomError) {

	resp, err := s.httpClient.Get(baseURL + "/open-banking/channels/v1/electronic-channels")

	if err != nil {
		return nil, common.NewInternalServerError("Unable to crawl channels from institution", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	jsonData := &channelJSON{}

	jsonUnmarshallErr := json.Unmarshal(body, &jsonData)

	if jsonUnmarshallErr != nil {

		return nil, common.NewInternalServerError("Unable to unmarshall data", jsonUnmarshallErr)
	}

	companies := jsonData.Data.Brand.Companies[0]
	return &companies.Channels, nil

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

type channelJSON struct {
	Data struct {
		Brand struct {
			Companies []struct {
				Channels []channel.Entity `json:"channels"`
			} `json:"companies"`
		} `json:"brand"`
	} `json:"data"`
}
