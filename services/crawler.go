package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"openbankingcrawler/common"
	"openbankingcrawler/domain/branch"
	"openbankingcrawler/domain/electronicchannel"
	"openbankingcrawler/domain/personalcreditcard"
	"openbankingcrawler/domain/personalloan"
	"strconv"
)

//Crawler service
type Crawler interface {
	Branches(string, int, []branch.Entity) (*[]branch.Entity, common.CustomError)
	ElectronicChannels(string, int, []electronicchannel.Entity) (*[]electronicchannel.Entity, common.CustomError)
	PersonalLoans(string, int, []personalloan.Entity) (*[]personalloan.Entity, common.CustomError)
	PersonalCreditCards(string, int, []personalcreditcard.Entity) (*[]personalcreditcard.Entity, common.CustomError)
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

	fmt.Println("Start crawl branches for", baseURL, page)

	body, _ := s.do(baseURL, "channels/v1/branches", page)

	jsonData := &branchJSON{}

	metaInfo := &metaInfoJSON{}
	json.Unmarshal(body, &metaInfo)

	jsonUnmarshallErr := json.Unmarshal(body, &jsonData)

	if jsonUnmarshallErr != nil {
		fmt.Println(jsonUnmarshallErr)
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

	fmt.Println("End crawl branches for", baseURL)

	return &branches, nil

}

//ElectronicChannels crawl electronicChannels from institution
func (s *crawler) ElectronicChannels(baseURL string, page int, accumulator []electronicchannel.Entity) (*[]electronicchannel.Entity, common.CustomError) {

	fmt.Println("Start crawl electronic channels for", baseURL, page)

	body, _ := s.do(baseURL, "channels/v1/electronic-channels", page)

	jsonData := &electronicChannelJSON{}

	metaInfo := &metaInfoJSON{}
	metaInfoErr := json.Unmarshal(body, &metaInfo)
	if metaInfoErr != nil {
		fmt.Println(metaInfoErr)
	}

	jsonUnmarshallErr := json.Unmarshal(body, &jsonData)

	if jsonUnmarshallErr != nil {
		fmt.Println(jsonUnmarshallErr)
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

	fmt.Println("End crawl electronic channels for", baseURL)

	return &channels, nil

}

//PersonalLoans crawl personal loans from institution
func (s *crawler) PersonalLoans(baseURL string, page int, accumulator []personalloan.Entity) (*[]personalloan.Entity, common.CustomError) {

	fmt.Println("Start crawl personal loans for", baseURL, page)

	body, crawlErr := s.do(baseURL, "products-services/v1/personal-loans", page)

	if crawlErr != nil {
		fmt.Println(crawlErr)
	}

	jsonData := &personalLoanJSON{}

	metaInfo := &metaInfoJSON{}
	json.Unmarshal(body, &metaInfo)

	jsonUnmarshallErr := json.Unmarshal(body, &jsonData)

	if jsonUnmarshallErr != nil {
		fmt.Println(jsonUnmarshallErr)
		return nil, common.NewInternalServerError("Unable to unmarshall data", jsonUnmarshallErr)
	}

	personalloans := accumulator

	for i := range jsonData.Data.Brand.Companies {
		company := jsonData.Data.Brand.Companies[i]
		result := company.PersonalLoans
		personalloans = append(personalloans, result...)
	}

	if metaInfo.Meta.TotalPages > page {
		return s.PersonalLoans(baseURL, page+1, personalloans)
	}

	fmt.Println("End crawl personal loans for", baseURL)

	return &personalloans, nil

}

//PersonalLoans crawl personal loans from institution
func (s *crawler) PersonalCreditCards(baseURL string, page int, accumulator []personalcreditcard.Entity) (*[]personalcreditcard.Entity, common.CustomError) {

	fmt.Println("Start crawl personal credit cards for", baseURL, page)

	body, _ := s.do(baseURL, "products-services/v1/personal-credit-cards", page)

	jsonData := &personalCreditCardJSON{}

	metaInfo := &metaInfoJSON{}
	json.Unmarshal(body, &metaInfo)

	jsonUnmarshallErr := json.Unmarshal(body, &jsonData)

	if jsonUnmarshallErr != nil {
		fmt.Println(jsonUnmarshallErr)
		return nil, common.NewInternalServerError("Unable to unmarshall data", jsonUnmarshallErr)
	}

	personalcreditcards := accumulator

	for i := range jsonData.Data.Brand.Companies {
		company := jsonData.Data.Brand.Companies[i]
		result := company.PersonalCreditCards

		personalcreditcards = append(personalcreditcards, result...)
	}

	if metaInfo.Meta.TotalPages > page {
		return s.PersonalCreditCards(baseURL, page+1, personalcreditcards)
	}

	fmt.Println("End crawl personal credit cards for", baseURL)

	return &personalcreditcards, nil

}

func (s *crawler) do(baseURL string, url string, page int) ([]byte, common.CustomError) {

	resp, err := s.httpClient.Get(baseURL + "/open-banking/" + url + "?&page=" + strconv.Itoa(page))

	fmt.Println(baseURL + "/open-banking/" + url + "?page=" + strconv.Itoa(page))

	if err != nil {
		fmt.Println(err)
		return nil, common.NewInternalServerError("Unable to crawl from institution", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	return body, nil

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

type personalLoanJSON struct {
	Data struct {
		Brand struct {
			Companies []struct {
				PersonalLoans []personalloan.Entity `json:"personalLoans"`
			} `json:"companies"`
		} `json:"brand"`
	} `json:"data"`
}

type personalCreditCardJSON struct {
	Data struct {
		Brand struct {
			Companies []struct {
				PersonalCreditCards []personalcreditcard.Entity `json:"personalCreditCards"`
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
