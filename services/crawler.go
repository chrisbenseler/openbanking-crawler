package services

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"openbankingcrawler/common"
	"openbankingcrawler/domain/branch"
	"openbankingcrawler/domain/businessaccount"
	"openbankingcrawler/domain/businesscreditcard"
	"openbankingcrawler/domain/businessfinancing"
	"openbankingcrawler/domain/businessinvoicefinancing"
	"openbankingcrawler/domain/businessloan"
	"openbankingcrawler/domain/businessunarrangedaccountoverdraft"
	"openbankingcrawler/domain/electronicchannel"
	"openbankingcrawler/domain/personalaccount"
	"openbankingcrawler/domain/personalcreditcard"
	"openbankingcrawler/domain/personalfinancing"
	"openbankingcrawler/domain/personalinvoicefinancing"
	"openbankingcrawler/domain/personalloan"
	"openbankingcrawler/domain/personalunarrangedaccountoverdraft"
	"openbankingcrawler/services/crawlerservices"
	"strconv"
)

//Crawler service
type Crawler interface {
	Branches(string, int, []branch.Entity) (*[]branch.Entity, common.CustomError)
	ElectronicChannels(string, int, []electronicchannel.Entity) (*[]electronicchannel.Entity, common.CustomError)
	PersonalAccounts(string, int, []personalaccount.Entity) (*[]personalaccount.Entity, common.CustomError)
	PersonalLoans(string, int, []personalloan.Entity) (*[]personalloan.Entity, common.CustomError)
	PersonalFinancings(string, int, []personalfinancing.Entity) (*[]personalfinancing.Entity, common.CustomError)
	PersonalInvoiceFinancings(string, int, []personalinvoicefinancing.Entity) (*[]personalinvoicefinancing.Entity, common.CustomError)
	PersonalCreditCards(string, int, []personalcreditcard.Entity) (*[]personalcreditcard.Entity, common.CustomError)
	PersonalUnarrangedAccountOverdrafts(string, int, []personalunarrangedaccountoverdraft.Entity) (*[]personalunarrangedaccountoverdraft.Entity, common.CustomError)
	BusinessAccounts(string, int, []businessaccount.Entity) (*[]businessaccount.Entity, common.CustomError)
	BusinessLoans(string, int, []businessloan.Entity) (*[]businessloan.Entity, common.CustomError)
	BusinessFinancings(string, int, []businessfinancing.Entity) (*[]businessfinancing.Entity, common.CustomError)
	BusinessInvoiceFinancings(string, int, []businessinvoicefinancing.Entity) (*[]businessinvoicefinancing.Entity, common.CustomError)
	BusinessCreditCards(string, int, []businesscreditcard.Entity) (*[]businesscreditcard.Entity, common.CustomError)
	BusinessUnarrangedAccountOverdrafts(string, int, []businessunarrangedaccountoverdraft.Entity) (*[]businessunarrangedaccountoverdraft.Entity, common.CustomError)
	Do(string, string, int) ([]byte, common.CustomError)
}

type crawler struct {
	httpClient *http.Client
}

//NewCrawler create a new service for crawl
func NewCrawler(http1 *http.Client) Crawler {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	return &crawler{
		httpClient: client,
	}
}

//Branches crawl branches from institution
func (s *crawler) Branches(baseURL string, page int, accumulator []branch.Entity) (*[]branch.Entity, common.CustomError) {

	fmt.Println("Start crawl branches for", baseURL, page)

	body, _ := s.Do(baseURL, "channels/v1/branches", page)

	jsonData := &branchJSON{}

	metaInfo := &MetaInfoJSON{}
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

	fmt.Println("End crawl branches for", baseURL)

	return &branches, nil

}

//ElectronicChannels crawl electronicChannels from institution
func (s *crawler) ElectronicChannels(baseURL string, page int, accumulator []electronicchannel.Entity) (*[]electronicchannel.Entity, common.CustomError) {

	fmt.Println("Start crawl electronic channels for", baseURL, page)

	body, _ := s.Do(baseURL, "channels/v1/electronic-channels", page)

	jsonData := &electronicChannelJSON{}

	metaInfo := &MetaInfoJSON{}
	metaInfoErr := json.Unmarshal(body, &metaInfo)
	if metaInfoErr != nil {
		fmt.Printf("Error crawl electronic channel: %s %s", metaInfoErr.Error(), baseURL)
	}

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

	fmt.Println("End crawl electronic channels for", baseURL)

	return &channels, nil

}

//PersonalLoans crawl personal loans from institution
func (s *crawler) PersonalLoans(baseURL string, page int, accumulator []personalloan.Entity) (*[]personalloan.Entity, common.CustomError) {
	fmt.Println("Start crawl personal loans for", baseURL)
	result, err := crawlerservices.ForPersonalLoans(s.Do, baseURL, page, accumulator)
	if err != nil {
		fmt.Printf("Error crawl personal loans for %s %s %s", baseURL, strconv.Itoa(page), err.Message())
	}
	fmt.Println("End crawl personal loans for", baseURL)
	return result, err
}

//PersonalLoans crawl personal loans from institution
func (s *crawler) PersonalCreditCards(baseURL string, page int, accumulator []personalcreditcard.Entity) (*[]personalcreditcard.Entity, common.CustomError) {
	fmt.Println("Start crawl personal credit cards for", baseURL)
	result, err := crawlerservices.ForPersonalCreditCards(s.Do, baseURL, page, accumulator)
	if err != nil {
		fmt.Printf("Error crawl personal credit cards for %s %s %s", baseURL, strconv.Itoa(page), err.Message())
	}
	fmt.Println("End crawl personal credit cards for", baseURL)
	return result, err
}

func (s *crawler) PersonalAccounts(baseURL string, page int, accumulator []personalaccount.Entity) (*[]personalaccount.Entity, common.CustomError) {
	fmt.Println("Start crawl personal account cards for", baseURL)
	result, err := crawlerservices.ForPersonalAccounts(s.Do, baseURL, page, accumulator)
	if err != nil {
		fmt.Printf("Error crawl personal accounts for %s %s %s", baseURL, strconv.Itoa(page), err.Message())
	}
	fmt.Println("End crawl personal accounts for", baseURL)
	return result, err
}

//PersonalFinancings PersonalFinancings
func (s *crawler) PersonalFinancings(baseURL string, page int, accumulator []personalfinancing.Entity) (*[]personalfinancing.Entity, common.CustomError) {
	fmt.Println("Start crawl personal financings for", baseURL)
	result, err := crawlerservices.ForPersonalFinancings(s.Do, baseURL, page, accumulator)
	if err != nil {
		fmt.Printf("Error crawl personal financings for %s %s %s", baseURL, strconv.Itoa(page), err.Message())
	}
	fmt.Println("End crawl personal financings for", baseURL)
	return result, err
}

//PersonalInvoiceFinancings PersonalInvoiceFinancings
func (s *crawler) PersonalInvoiceFinancings(baseURL string, page int, accumulator []personalinvoicefinancing.Entity) (*[]personalinvoicefinancing.Entity, common.CustomError) {
	fmt.Println("Start crawl personal invoice financings for", baseURL)
	result, err := crawlerservices.ForPersonalInvoiceFinancings(s.Do, baseURL, page, accumulator)
	if err != nil {
		fmt.Printf("Error crawl personal invoice financings for %s %s %s", baseURL, strconv.Itoa(page), err.Message())
	}
	fmt.Println("End crawl personal invoice financings for", baseURL)
	return result, err
}

//PersonalUnarrangedAccountOverdrafts PersonalUnarrangedAccountOverdrafts
func (s *crawler) PersonalUnarrangedAccountOverdrafts(baseURL string, page int, accumulator []personalunarrangedaccountoverdraft.Entity) (*[]personalunarrangedaccountoverdraft.Entity, common.CustomError) {
	fmt.Println("Start crawl personal unarranged account overdrafts for", baseURL)
	result, err := crawlerservices.ForPersonalUnarrangedAccountOverdrafts(s.Do, baseURL, page, accumulator)
	if err != nil {
		fmt.Printf("Error crawl personal unarranged account overdrafts for %s %s %s", baseURL, strconv.Itoa(page), err.Message())
	}
	fmt.Println("End crawl personal unarranged account overdrafts for", baseURL)
	return result, err
}

//BusinessAccounts BusinessAccounts
func (s *crawler) BusinessAccounts(baseURL string, page int, accumulator []businessaccount.Entity) (*[]businessaccount.Entity, common.CustomError) {
	fmt.Println("Start crawl business account for", baseURL)
	result, err := crawlerservices.ForBusinessAccounts(s.Do, baseURL, page, accumulator)
	if err != nil {
		fmt.Printf("Error crawl business account for %s %s %s", baseURL, strconv.Itoa(page), err.Message())
	}
	fmt.Println("End crawl business accounts for", baseURL)
	return result, err
}

//BusinessLoans BusinessLoans
func (s *crawler) BusinessLoans(baseURL string, page int, accumulator []businessloan.Entity) (*[]businessloan.Entity, common.CustomError) {
	fmt.Println("Start crawl business loan for", baseURL)
	result, err := crawlerservices.ForBusinessLoans(s.Do, baseURL, page, accumulator)
	if err != nil {
		fmt.Printf("Error crawl business loan for %s %s %s", baseURL, strconv.Itoa(page), err.Message())
	}
	fmt.Println("End crawl business loan for", baseURL)
	return result, err
}

//BusinessFinancings BusinessFinancings
func (s *crawler) BusinessFinancings(baseURL string, page int, accumulator []businessfinancing.Entity) (*[]businessfinancing.Entity, common.CustomError) {
	fmt.Println("Start crawl business financings for", baseURL)
	result, err := crawlerservices.ForBusinessFinancings(s.Do, baseURL, page, accumulator)
	if err != nil {
		fmt.Printf("Error crawl business financing for %s %s %s", baseURL, strconv.Itoa(page), err.Message())
	}
	fmt.Println("End crawl business financings for", baseURL)
	return result, err
}

//BusinessInvoiceFinancings BusinessFinancings
func (s *crawler) BusinessInvoiceFinancings(baseURL string, page int, accumulator []businessinvoicefinancing.Entity) (*[]businessinvoicefinancing.Entity, common.CustomError) {
	fmt.Println("Start crawl business invoice financings for", baseURL)
	result, err := crawlerservices.ForBusinessInvoiceFinancings(s.Do, baseURL, page, accumulator)
	if err != nil {
		fmt.Printf("Error crawl business invoice financing for %s %s %s", baseURL, strconv.Itoa(page), err.Message())
	}
	fmt.Println("End crawl business invoice financings for", baseURL)
	return result, err
}

//BusinessCreditCards crawl business credit cards from institution
func (s *crawler) BusinessCreditCards(baseURL string, page int, accumulator []businesscreditcard.Entity) (*[]businesscreditcard.Entity, common.CustomError) {
	fmt.Println("Start crawl business business cards for", baseURL)
	result, err := crawlerservices.ForBusinessCreditCards(s.Do, baseURL, page, accumulator)
	if err != nil {
		fmt.Printf("Error crawl business cards for %s %s %s", baseURL, strconv.Itoa(page), err.Message())
	}
	fmt.Println("End crawl business business cards for", baseURL)
	return result, err
}

//BusinessUnarrangedAccountOverdrafts BusinessUnarrangedAccountOverdrafts
func (s *crawler) BusinessUnarrangedAccountOverdrafts(baseURL string, page int, accumulator []businessunarrangedaccountoverdraft.Entity) (*[]businessunarrangedaccountoverdraft.Entity, common.CustomError) {
	fmt.Println("Start crawl business unarranged account overdrafts for", baseURL)
	result, err := crawlerservices.ForBusinessUnarrangedAccountOverdrafts(s.Do, baseURL, page, accumulator)
	if err != nil {
		fmt.Printf("Error crawl business unarranged account overdrafts for %s %s %s", baseURL, strconv.Itoa(page), err.Message())
	}
	fmt.Println("End crawl business unarranged account overdrafts for", baseURL)
	return result, err
}

//Do do
func (s *crawler) Do(baseURL string, url string, page int) ([]byte, common.CustomError) {

	fullURL := baseURL + "/open-banking/" + url + "?page=" + strconv.Itoa(page)

	resp, err := s.httpClient.Get(fullURL)

	if err != nil {
		fmt.Printf("Error crawling: %s %s", fullURL, err.Error())
		return nil, common.NewInternalServerError("Unable to crawl", err)
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

//MetaInfoJSON MetaInfoJSON
type MetaInfoJSON struct {
	Meta struct {
		TotalRecords int `json:"totalRecords"`
		TotalPages   int `json:"totalPages"`
	} `json:"meta"`
}
