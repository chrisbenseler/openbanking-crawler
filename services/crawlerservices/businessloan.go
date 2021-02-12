package crawlerservices

import (
	"encoding/json"
	"fmt"
	"openbankingcrawler/common"
	"openbankingcrawler/domain/businessloan"
	"strconv"
)

//ForBusinessLoans crawl business loans from institution
func ForBusinessLoans(httpCrawlService func(string, string, int) ([]byte, common.CustomError), baseURL string, page int, accumulator []businessloan.Entity) (*[]businessloan.Entity, common.CustomError) {

	fmt.Println("Start crawl business loans for", baseURL, page)

	body, _ := httpCrawlService(baseURL, "products-services/v1/business-loans", page)

	jsonData := &businessLoanJSON{}

	metaInfo := &MetaInfoJSON{}
	json.Unmarshal(body, &metaInfo)

	jsonUnmarshallErr := json.Unmarshal(body, &jsonData)

	if jsonUnmarshallErr != nil {
		fmt.Printf("Error crawl business loans for %s %s %s", baseURL, strconv.Itoa(page), jsonUnmarshallErr)
		return nil, common.NewInternalServerError("Unable to unmarshall data", jsonUnmarshallErr)
	}

	businessloans := accumulator

	for i := range jsonData.Data.Brand.Companies {
		company := jsonData.Data.Brand.Companies[i]
		result := company.BusinessLoans

		businessloans = append(businessloans, result...)
	}

	if metaInfo.Meta.TotalPages > page {
		return ForBusinessLoans(httpCrawlService, baseURL, page+1, businessloans)
	}

	fmt.Println("End crawl business loans for", baseURL, page)

	return &businessloans, nil

}

type businessLoanJSON struct {
	Data struct {
		Brand struct {
			Companies []struct {
				BusinessLoans []businessloan.Entity `json:"businessLoans"`
			} `json:"companies"`
		} `json:"brand"`
	} `json:"data"`
}
