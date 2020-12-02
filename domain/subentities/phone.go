package subentities

//Phone phone struct
type Phone struct {
	Type        string `json:"type"`
	CountryCode string `json:"countryCode"`
	AreCode     string `json:"areaCode"`
	Number      string `json:"number"`
}
