package crawlerservices

import (
	"encoding/json"
	"fmt"
	"openbankingcrawler/common"
	"openbankingcrawler/domain/businessunarrangedaccountoverdraft"
	"strconv"
)

//ForBusinessUnarrangedAccountOverdrafts crawl business loans from institution
func ForBusinessUnarrangedAccountOverdrafts(httpCrawlService func(string, string, int) ([]byte, common.CustomError), baseURL string, page int, accumulator []businessunarrangedaccountoverdraft.Entity) (*[]businessunarrangedaccountoverdraft.Entity, common.CustomError) {

	fmt.Println("Start crawl business accounts for", baseURL, page)

	body, _ := httpCrawlService(baseURL, "products-services/v1/business-unarranged-account-overdraft", page)

	jsonData := &businessUnarrangedAccountOverdraftJSON{}

	metaInfo := &MetaInfoJSON{}
	json.Unmarshal(body, &metaInfo)

	jsonUnmarshallErr := json.Unmarshal(body, &jsonData)

	if jsonUnmarshallErr != nil {
		fmt.Printf("Error crawl business accounts for %s %s %s", baseURL, strconv.Itoa(page), jsonUnmarshallErr)
		return nil, common.NewInternalServerError("Unable to unmarshall data", jsonUnmarshallErr)
	}

	items := accumulator

	for i := range jsonData.Data.Brand.Companies {
		company := jsonData.Data.Brand.Companies[i]
		result := company.BusinessUnarrangedAccountOverdrafts

		items = append(items, result...)
	}

	if metaInfo.Meta.TotalPages > page {
		return ForBusinessUnarrangedAccountOverdrafts(httpCrawlService, baseURL, page+1, items)
	}

	fmt.Println("End crawl business unarrenged account oversdrafts for", baseURL, page)

	return &items, nil

}

type businessUnarrangedAccountOverdraftJSON struct {
	Data struct {
		Brand struct {
			Companies []struct {
				BusinessUnarrangedAccountOverdrafts []businessunarrangedaccountoverdraft.Entity `json:"businessUnarrangedAccountOverdraft"`
			} `json:"companies"`
		} `json:"brand"`
	} `json:"data"`
}
