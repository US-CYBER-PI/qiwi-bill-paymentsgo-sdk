package Models

type GeneratePayToken struct {
	RequestId string `json:"requestId"`
	Phone     string `json:"phone"`
	AccountId string `json:"accountId"`
}

func NewGeneratePayToken(RequestId string, Phone string, AccountId string) *GeneratePayToken {
	return &GeneratePayToken{
		RequestId: RequestId,
		Phone:     Phone,
		AccountId: AccountId,
	}
}
