package request

type Verify2faCodeRequest struct {
	Code      string `json:"code"`
	AccountId string
}
