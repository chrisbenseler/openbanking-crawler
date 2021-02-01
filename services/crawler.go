package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"openbankingcrawler/common"
	"openbankingcrawler/domain/branch"
	"openbankingcrawler/domain/electronicchannel"
	"strconv"
)

//Crawler service
type Crawler interface {
	Branches(string, int, []branch.Entity) (*[]branch.Entity, common.CustomError)
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
func (s *crawler) Branches(baseURL string, page int, accumulator []branch.Entity) (*[]branch.Entity, common.CustomError) {

	resp, err := s.httpClient.Get(baseURL + "/open-banking/channels/v1/branches?page-size=50&page=" + strconv.Itoa(page))

	if err != nil {
		return nil, common.NewInternalServerError("Unable to crawl branches from institution", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	jsonData := &branchJSON{}

	metaInfo := &metaInfoJSON{}
	json.Unmarshal(body, &metaInfo)

	jsonUnmarshallErr := json.Unmarshal(body, &jsonData)

	if jsonUnmarshallErr != nil {
		return nil, common.NewInternalServerError("Unable to unmarshall data", jsonUnmarshallErr)
	}

	branches := accumulator

	for i := range jsonData.Data.Brand.Companies {
		company := jsonData.Data.Brand.Companies[i]
		result := company.Branches
		branches = append(branches, result...)
	}

	if metaInfo.Meta.TotalPages > page {
		return s.Branches(baseURL, page+1, branches)
	}

	fmt.Println("end craw branches for", baseURL)

	return &branches, nil

}

//ElectronicChannels crawl electronicChannels from institution
func (s *crawler) ElectronicChannels(baseURL string, page int, accumulator []electronicchannel.Entity) (*[]electronicchannel.Entity, common.CustomError) {

	resp, err := s.httpClient.Get(baseURL + "/open-banking/channels/v1/electronic-channels?page-size=50&page=" + strconv.Itoa(page))

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

	channels := accumulator

	for i := range jsonData.Data.Brand.Companies {
		company := jsonData.Data.Brand.Companies[i]
		result := company.ElectronicChannels
		channels = append(channels, result...)
	}

	if metaInfo.Meta.TotalPages > page {
		return s.ElectronicChannels(baseURL, page+1, channels)
	}

	fmt.Println("end craw channels for", baseURL)

	return &channels, nil

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
