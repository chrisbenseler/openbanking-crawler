package subentities

//FeeService fee service entity
type FeeService struct {
	Name                string         `json:"name"`
	Code                string         `json:"code"`
	ChargingTriggerInfo string         `json:"chargingTriggerInfo"`
	Prices              []ServicePrice `json:"prices"`
}

//ServicePrice service price
type ServicePrice struct {
	Interval  string    `json:"interval"`
	Value     string    `json:"value"`
	Currency  string    `json:"currency"`
	Customers Customers `json:"customers"`
}

//Customers customers
type Customers struct {
	Rate string `json:"rate"`
}

//Indexer indexer
type Indexer struct {
	Rate string `json:"rate"`
}

//Fees Fees struct
type Fees struct {
	Services []FeeService `json:"services"`
}

//Applications interests rates
type Applications struct {
	Interval  string    `json:"interval"`
	Customers Customers `json:"customers"`
	Indexer   Indexer   `json:"indexer"`
}

//InterestRates interests rates
type InterestRates struct {
	ReferentialRateIndexer string         `json:"referentialRateIndexer"`
	Rate                   string         `json:"rate"`
	Applications           []Applications `json:"applications"`
	MinimumRate            string         `json:"minimumRate"`
	MaximumRate            string         `json:"maximumRate"`
}
