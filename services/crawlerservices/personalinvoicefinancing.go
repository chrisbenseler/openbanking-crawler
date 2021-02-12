package crawlerservices

import (
	"encoding/json"
	"fmt"
	"openbankingcrawler/common"
	"openbankingcrawler/domain/personalinvoicefinancing"
	"strconv"
)

//ForPersonalInvoiceFinancings crawl business credit cards from institution
func ForPersonalInvoiceFinancings(httpCrawlService func(string, string, int) ([]byte, common.CustomError), baseURL string, page int, accumulator []personalinvoicefinancing.Entity) (*[]personalinvoicefinancing.Entity, common.CustomError) {

	fmt.Println("Start crawl personal invoice financings for", baseURL, page)

	body, _ := httpCrawlService(baseURL, "products-services/v1/personal-invoice-financings", page)

	jsonData := &personalInvoiceFinancingJSON{}

	metaInfo := &MetaInfoJSON{}
	json.Unmarshal(body, &metaInfo)

	jsonUnmarshallErr := json.Unmarshal(body, &jsonData)

	if jsonUnmarshallErr != nil {
		fmt.Printf("Error crawl personal invoice financings %s %s %s", baseURL, strconv.Itoa(page), jsonUnmarshallErr)
		return nil, common.NewInternalServerError("Unable to unmarshall data", jsonUnmarshallErr)
	}

	personalinvoicefinancings := accumulator

	for i := range jsonData.Data.Brand.Companies {
		company := jsonData.Data.Brand.Companies[i]
		result := company.PersonalInvoiceFinancings

		personalinvoicefinancings = append(personalinvoicefinancings, result...)
	}

	if metaInfo.Meta.TotalPages > page {
		return ForPersonalInvoiceFinancings(httpCrawlService, baseURL, page+1, personalinvoicefinancings)
	}

	fmt.Println("End crawl personal invoice financings for", baseURL, page)

	return &personalinvoicefinancings, nil

}

type personalInvoiceFinancingJSON struct {
	Data struct {
		Brand struct {
			Companies []struct {
				PersonalInvoiceFinancings []personalinvoicefinancing.Entity `json:"personalInvoiceFinancings"`
			} `json:"companies"`
		} `json:"brand"`
	} `json:"data"`
}
