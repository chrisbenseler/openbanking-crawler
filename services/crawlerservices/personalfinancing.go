package crawlerservices

import (
	"encoding/json"
	"fmt"
	"openbankingcrawler/common"
	"openbankingcrawler/domain/personalfinancing"
	"strconv"
)

//ForPersonalFinancings crawl personal loans from institution
func ForPersonalFinancings(httpCrawlService func(string, string, int) ([]byte, common.CustomError), baseURL string, page int, accumulator []personalfinancing.Entity) (*[]personalfinancing.Entity, common.CustomError) {

	fmt.Println("Start crawl personal financings for", baseURL, page)

	body, crawlErr := httpCrawlService(baseURL, "products-services/v1/personal-financings", page)

	if crawlErr != nil {
		fmt.Println(crawlErr)
	}

	jsonData := &personalFinancingJSON{}

	metaInfo := &MetaInfoJSON{}
	json.Unmarshal(body, &metaInfo)

	jsonUnmarshallErr := json.Unmarshal(body, &jsonData)

	if jsonUnmarshallErr != nil {
		fmt.Printf("Error crawl personal financings for %s %s %s", baseURL, strconv.Itoa(page), jsonUnmarshallErr)
		return nil, common.NewInternalServerError("Unable to unmarshall data", jsonUnmarshallErr)
	}

	items := accumulator

	for i := range jsonData.Data.Brand.Companies {
		company := jsonData.Data.Brand.Companies[i]
		result := company.PersonalFinancing
		items = append(items, result...)
	}

	if metaInfo.Meta.TotalPages > page {
		return ForPersonalFinancings(httpCrawlService, baseURL, page+1, items)
	}

	fmt.Println("End crawl personal financings for", baseURL, page)

	return &items, nil

}

type personalFinancingJSON struct {
	Data struct {
		Brand struct {
			Companies []struct {
				PersonalFinancing []personalfinancing.Entity `json:"personalFinancings"`
			} `json:"companies"`
		} `json:"brand"`
	} `json:"data"`
}
