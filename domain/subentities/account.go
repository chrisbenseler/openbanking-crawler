package subentities

//ServiceBundleService service bundle service
type ServiceBundleService struct {
	Code                string `json:"code"`
	ChargingTriggerInfo string `json:"chargingTriggerInfo"`
	EventLimitQuantity  string `json:"eventLimitQuantity"`
	FreeEventQuantity   string `json:"freeEventQuantity"`
}

//ServiceBundle ServiceBundle
type ServiceBundle struct {
	Name     string                 `json:"name"`
	Services []ServiceBundleService `json:"services"`
	Prices   []Price                `json:"prices"`
	Minimun  struct {
		Value    string `json:"value"`
		Currency string `json:"currency"`
	} `json:"minimum"`
	Maximun struct {
		Value    string `json:"value"`
		Currency string `json:"currency"`
	} `json:"maximum"`
}

//TermsConditions TermsConditions
type TermsConditions struct {
	MinimumBalance struct {
		Value    string `json:"value"`
		Currency string `json:"currency"`
	} `json:"minimumBalance"`
	ElegibilityCriteriaInfo string `json:"elegibilityCriteriaInfo"`
	ClosingProcessInfo      string `json:"closingProcessInfo"`
}

//IncomeRate IncomeRate
type IncomeRate struct {
	SavingAccount         string `json:"savingAccount"`
	PrepaidPaymentAccount string `json:"prepaidPaymentAccount"`
}
