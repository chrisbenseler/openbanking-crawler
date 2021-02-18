package crawlerservices

import (
	"encoding/json"
	"fmt"
	"openbankingcrawler/common"
	"openbankingcrawler/domain/businessfinancing"
	"strconv"
)

//ForBusinessFinancings crawl business loans from institution
func ForBusinessFinancings(httpCrawlService func(string, string, int) ([]byte, common.CustomError), baseURL string, page int, accumulator []businessfinancing.Entity) (*[]businessfinancing.Entity, common.CustomError) {

	fmt.Println("Start crawl business financings for", baseURL, page)

	body, crawlErr := httpCrawlService(baseURL, "products-services/v1/business-financings", page)

	if crawlErr != nil {
		fmt.Println(crawlErr)
	}

	jsonData := &businessFinancingJSON{}

	metaInfo := &MetaInfoJSON{}
	json.Unmarshal(body, &metaInfo)

	jsonUnmarshallErr := json.Unmarshal(body, &jsonData)

	if jsonUnmarshallErr != nil {
		fmt.Printf("Error crawl business financings for %s %s %s", baseURL, strconv.Itoa(page), jsonUnmarshallErr)
		return nil, common.NewInternalServerError("Unable to unmarshall data", jsonUnmarshallErr)
	}

	items := accumulator

	for i := range jsonData.Data.Brand.Companies {
		company := jsonData.Data.Brand.Companies[i]
		result := company.BusinessFinancing
		items = append(items, result...)
	}

	if metaInfo.Meta.TotalPages > page {
		return ForBusinessFinancings(httpCrawlService, baseURL, page+1, items)
	}

	fmt.Println("End crawl business financings for", baseURL, page)

	return &items, nil

}

type businessFinancingJSON struct {
	Data struct {
		Brand struct {
			Companies []struct {
				BusinessFinancing []businessfinancing.Entity `json:"businessFinancings"`
			} `json:"companies"`
		} `json:"brand"`
	} `json:"data"`
}
