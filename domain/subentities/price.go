package subentities

//Price price
type Price struct {
	Interval   string `json:"interval"`
	MonthlyFee string `json:"monthlyFee"`
	Currency   string `json:"currency"`
	Customers  struct {
		Rate string `json:"rate"`
	} `json:"customers"`
}
