package crawlerservices

import (
	"encoding/json"
	"fmt"
	"openbankingcrawler/common"
	"openbankingcrawler/domain/personalloan"
	"strconv"
)

//ForPersonalLoans crawl personal loans from institution
func ForPersonalLoans(httpCrawlService func(string, string, int) ([]byte, common.CustomError), baseURL string, page int, accumulator []personalloan.Entity) (*[]personalloan.Entity, common.CustomError) {

	fmt.Println("Start crawl personal loans for", baseURL, page)

	body, crawlErr := httpCrawlService(baseURL, "products-services/v1/personal-loans", page)

	if crawlErr != nil {
		fmt.Println(crawlErr)
	}

	jsonData := &personalLoanJSON{}

	metaInfo := &MetaInfoJSON{}
	json.Unmarshal(body, &metaInfo)

	jsonUnmarshallErr := json.Unmarshal(body, &jsonData)

	if jsonUnmarshallErr != nil {
		fmt.Printf("Error crawl personal loans for %s %s %s", baseURL, strconv.Itoa(page), jsonUnmarshallErr)
		return nil, common.NewInternalServerError("Unable to unmarshall data", jsonUnmarshallErr)
	}

	personalloans := accumulator

	for i := range jsonData.Data.Brand.Companies {
		company := jsonData.Data.Brand.Companies[i]
		result := company.PersonalLoans
		personalloans = append(personalloans, result...)
	}

	if metaInfo.Meta.TotalPages > page {
		return ForPersonalLoans(httpCrawlService, baseURL, page+1, personalloans)
	}

	fmt.Println("End crawl personal loans for", baseURL, page)

	return &personalloans, nil

}

type personalLoanJSON struct {
	Data struct {
		Brand struct {
			Companies []struct {
				PersonalLoans []personalloan.Entity `json:"personalLoans"`
			} `json:"companies"`
		} `json:"brand"`
	} `json:"data"`
}
