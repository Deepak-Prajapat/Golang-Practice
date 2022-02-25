package salesforce

type GetAccountResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type ErrorResponse struct {
	ErrorCode  string `json:"errorCode"`
	ErrMessage string `json:"message"`
}
