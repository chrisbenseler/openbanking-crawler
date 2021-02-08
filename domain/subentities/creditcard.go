package subentities

//RewardsProgram RewardsProgram structure
type RewardsProgram struct {
	HasRewardProgram  bool   `json:"hasRewardProgram"`
	RewardProgramInfo string `json:"rewardProgramInfo"`
}

//IdentificationProduct IdentificationProduct
type IdentificationProduct struct {
	Type           string `json:"type"`
	AdditionalInfo string `json:"additionalInfo"`
}

//CreditCardFlag CreditCardFlag
type CreditCardFlag struct {
	Network        string `json:"network"`
	AdditionalInfo string `json:"additionalInfo"`
}

//CreditCardIdentification CreditCardIdentification
type CreditCardIdentification struct {
	Product    IdentificationProduct `json:"product"`
	CreditCard CreditCardFlag        `json:"creditCard"`
}

//CreditCardInterests CreditCardInterests
type CreditCardInterests struct {
	Rates           []Rates `json:"rates"`
	InstalmentRates []Rates `json:"instalmentRates"`
}

//CreditCardOthers
type CreditCardOthers struct {
	Code string `json:"code"`
}
