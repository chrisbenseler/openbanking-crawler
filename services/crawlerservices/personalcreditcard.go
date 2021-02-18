package crawlerservices

import (
	"encoding/json"
	"fmt"
	"openbankingcrawler/common"
	"openbankingcrawler/domain/personalcreditcard"
	"strconv"
)

//ForPersonalCreditCards crawl personal credit cards from institution
func ForPersonalCreditCards(httpCrawlService func(string, string, int) ([]byte, common.CustomError), baseURL string, page int, accumulator []personalcreditcard.Entity) (*[]personalcreditcard.Entity, common.CustomError) {

	fmt.Println("Start crawl personal credit cards for", baseURL, page)

	body, _ := httpCrawlService(baseURL, "products-services/v1/personal-credit-cards", page)

	jsonData := &personalCreditCardJSON{}

	metaInfo := &MetaInfoJSON{}
	json.Unmarshal(body, &metaInfo)

	jsonUnmarshallErr := json.Unmarshal(body, &jsonData)

	if jsonUnmarshallErr != nil {
		fmt.Printf("Error crawl personal credit cards for %s %s %s", baseURL, strconv.Itoa(page), jsonUnmarshallErr)
		return nil, common.NewInternalServerError("Unable to unmarshall data", jsonUnmarshallErr)
	}

	items := accumulator

	for i := range jsonData.Data.Brand.Companies {
		company := jsonData.Data.Brand.Companies[i]
		result := company.PersonalCreditCards

		items = append(items, result...)
	}

	if metaInfo.Meta.TotalPages > page {
		return ForPersonalCreditCards(httpCrawlService, baseURL, page+1, items)
	}

	fmt.Println("End crawl personal credit cards for", baseURL, page)

	return &items, nil

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
