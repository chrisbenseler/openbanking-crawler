package crawlerservices

import (
	"encoding/json"
	"fmt"
	"openbankingcrawler/common"
	"openbankingcrawler/domain/businessaccount"
)

//ForBusinessAccounts crawl business loans from institution
func ForBusinessAccounts(httpCrawlService func(string, string, int) ([]byte, common.CustomError), baseURL string, page int, accumulator []businessaccount.Entity) (*[]businessaccount.Entity, common.CustomError) {

	fmt.Println("Start crawl business accounts for", baseURL, page)

	body, _ := httpCrawlService(baseURL, "products-services/v1/business-accounts", page)

	jsonData := &businessAccountJSON{}

	metaInfo := &MetaInfoJSON{}
	json.Unmarshal(body, &metaInfo)

	jsonUnmarshallErr := json.Unmarshal(body, &jsonData)

	if jsonUnmarshallErr != nil {
		fmt.Println(jsonUnmarshallErr, baseURL+"/products-services/v1/business-accounts", page)
		return nil, common.NewInternalServerError("Unable to unmarshall data", jsonUnmarshallErr)
	}

	businessaccounts := accumulator

	for i := range jsonData.Data.Brand.Companies {
		company := jsonData.Data.Brand.Companies[i]
		result := company.BusinessAccounts

		businessaccounts = append(businessaccounts, result...)
	}

	if metaInfo.Meta.TotalPages > page {
		return ForBusinessAccounts(httpCrawlService, baseURL, page+1, businessaccounts)
	}

	fmt.Println("End crawl business accounts for", baseURL)

	return &businessaccounts, nil

}

type businessAccountJSON struct {
	Data struct {
		Brand struct {
			Companies []struct {
				BusinessAccounts []businessaccount.Entity `json:"businessAccounts"`
			} `json:"companies"`
		} `json:"brand"`
	} `json:"data"`
}
