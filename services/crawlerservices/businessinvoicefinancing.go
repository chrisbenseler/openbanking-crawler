package crawlerservices

import (
	"encoding/json"
	"fmt"
	"openbankingcrawler/common"
	"openbankingcrawler/domain/businessinvoicefinancing"
	"strconv"
)

//ForBusinessInvoiceFinancings crawl business credit cards from institution
func ForBusinessInvoiceFinancings(httpCrawlService func(string, string, int) ([]byte, common.CustomError), baseURL string, page int, accumulator []businessinvoicefinancing.Entity) (*[]businessinvoicefinancing.Entity, common.CustomError) {

	fmt.Println("Start crawl business invoice financings for", baseURL, page)

	body, _ := httpCrawlService(baseURL, "products-services/v1/business-invoice-financings", page)

	items := accumulator

	jsonData := &businessInvoiceFinancingJSON{}

	metaInfo := &MetaInfoJSON{}
	json.Unmarshal(body, &metaInfo)

	jsonUnmarshallErr := json.Unmarshal(body, &jsonData)

	if jsonUnmarshallErr != nil {
		fmt.Printf("Error crawl business invoice financings for %s %s %s", baseURL, strconv.Itoa(page), jsonUnmarshallErr)
		return nil, common.NewInternalServerError("Unable to unmarshall data", jsonUnmarshallErr)
	}

	for i := range jsonData.Data.Brand.Companies {
		company := jsonData.Data.Brand.Companies[i]
		result := company.BusinessInvoiceFinancings

		items = append(items, result...)
	}

	if metaInfo.Meta.TotalPages > page {
		return ForBusinessInvoiceFinancings(httpCrawlService, baseURL, page+1, items)
	}

	fmt.Println("End crawl business invoice financings for", baseURL, page)

	return &items, nil

}

type businessInvoiceFinancingJSON struct {
	Data struct {
		Brand struct {
			Companies []struct {
				BusinessInvoiceFinancings []businessinvoicefinancing.Entity `json:"businessInvoiceFinancings"`
			} `json:"companies"`
		} `json:"brand"`
	} `json:"data"`
}
