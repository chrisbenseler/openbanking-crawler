package adapters

//InstitutionPayload institution payload
type InstitutionPayload struct {
	Name    string `json:"name"`
	BaseURL string `json:"baseurl"`
}

//AuthenticatePayload authenticate payload
type AuthenticatePayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
