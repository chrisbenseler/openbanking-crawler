package crawlerservices

import (
	"encoding/json"
	"fmt"
	"openbankingcrawler/common"
	"openbankingcrawler/domain/personalunarrangedaccountoverdraft"
	"strconv"
)

//ForPersonalUnarrangedAccountOverdrafts crawl personal loans from institution
func ForPersonalUnarrangedAccountOverdrafts(httpCrawlService func(string, string, int) ([]byte, common.CustomError), baseURL string, page int, accumulator []personalunarrangedaccountoverdraft.Entity) (*[]personalunarrangedaccountoverdraft.Entity, common.CustomError) {

	fmt.Println("Start crawl personal accounts for", baseURL, page)

	body, _ := httpCrawlService(baseURL, "products-services/v1/personal-unarranged-account-overdraft", page)

	jsonData := &personalUnarrangedAccountOverdraftJSON{}

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
		result := company.PersonalUnarrangedAccountOverdrafts

		items = append(items, result...)
	}

	if metaInfo.Meta.TotalPages > page {
		return ForPersonalUnarrangedAccountOverdrafts(httpCrawlService, baseURL, page+1, items)
	}

	fmt.Println("End crawl personal unarrenged account oversdrafts for", baseURL, page)

	return &items, nil

}

type personalUnarrangedAccountOverdraftJSON struct {
	Data struct {
		Brand struct {
			Companies []struct {
				PersonalUnarrangedAccountOverdrafts []personalunarrangedaccountoverdraft.Entity `json:"personalUnarrangedAccountOverdraft"`
			} `json:"companies"`
		} `json:"brand"`
	} `json:"data"`
}
