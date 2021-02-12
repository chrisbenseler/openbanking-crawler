package crawlerservices

import (
	"encoding/json"
	"fmt"
	"openbankingcrawler/common"
	"openbankingcrawler/domain/businesscreditcard"
	"strconv"
)

//ForBusinessCreditCards crawl business credit cards from institution
func ForBusinessCreditCards(httpCrawlService func(string, string, int) ([]byte, common.CustomError), baseURL string, page int, accumulator []businesscreditcard.Entity) (*[]businesscreditcard.Entity, common.CustomError) {

	fmt.Println("Start crawl business credit cards for", baseURL, page)

	body, _ := httpCrawlService(baseURL, "products-services/v1/business-credit-cards", page)

	jsonData := &businessCreditCardJSON{}

	metaInfo := &MetaInfoJSON{}
	json.Unmarshal(body, &metaInfo)

	jsonUnmarshallErr := json.Unmarshal(body, &jsonData)

	if jsonUnmarshallErr != nil {
		fmt.Printf("Error crawl business credit cards for %s %s %s", baseURL, strconv.Itoa(page), jsonUnmarshallErr)
		return nil, common.NewInternalServerError("Unable to unmarshall data", jsonUnmarshallErr)
	}

	businesscreditcards := accumulator

	for i := range jsonData.Data.Brand.Companies {
		company := jsonData.Data.Brand.Companies[i]
		result := company.BusinessCreditCards

		businesscreditcards = append(businesscreditcards, result...)
	}

	if metaInfo.Meta.TotalPages > page {
		return ForBusinessCreditCards(httpCrawlService, baseURL, page+1, businesscreditcards)
	}

	fmt.Println("End crawl business credit cards for", baseURL, page)

	return &businesscreditcards, nil

}

type businessCreditCardJSON struct {
	Data struct {
		Brand struct {
			Companies []struct {
				BusinessCreditCards []businesscreditcard.Entity `json:"businessCreditCards"`
			} `json:"companies"`
		} `json:"brand"`
	} `json:"data"`
}
