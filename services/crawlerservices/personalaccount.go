package crawlerservices

import (
	"encoding/json"
	"fmt"
	"openbankingcrawler/common"
	"openbankingcrawler/domain/personalaccount"
	"strconv"
)

//ForPersonalAccounts crawl personal loans from institution
func ForPersonalAccounts(httpCrawlService func(string, string, int) ([]byte, common.CustomError), baseURL string, page int, accumulator []personalaccount.Entity) (*[]personalaccount.Entity, common.CustomError) {

	fmt.Println("Start crawl personal accounts for", baseURL, page)

	body, _ := httpCrawlService(baseURL, "products-services/v1/personal-accounts", page)

	jsonData := &personalAccountJSON{}

	metaInfo := &MetaInfoJSON{}
	json.Unmarshal(body, &metaInfo)

	jsonUnmarshallErr := json.Unmarshal(body, &jsonData)

	if jsonUnmarshallErr != nil {
		fmt.Printf("Error crawl personal accounts for %s %s %s", baseURL, strconv.Itoa(page), jsonUnmarshallErr)
		return nil, common.NewInternalServerError("Unable to unmarshall data", jsonUnmarshallErr)
	}

	items := accumulator

	for i := range jsonData.Data.Brand.Companies {
		company := jsonData.Data.Brand.Companies[i]
		result := company.PersonalAccounts

		items = append(items, result...)
	}

	if metaInfo.Meta.TotalPages > page {
		return ForPersonalAccounts(httpCrawlService, baseURL, page+1, items)
	}

	fmt.Println("End crawl personal accounts for", baseURL, page)

	return &items, nil

}

type personalAccountJSON struct {
	Data struct {
		Brand struct {
			Companies []struct {
				PersonalAccounts []personalaccount.Entity `json:"personalAccounts"`
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
