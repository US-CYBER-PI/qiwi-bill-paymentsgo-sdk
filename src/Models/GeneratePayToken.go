package Models

import "github.com/US-CYBER-PI/qiwi-bill-paymentsgo-sdk/src/Models/Ets"

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

type GeneratePayTokenResponse struct {
	RequestId string     `json:"requestId"`
	Status    Ets.Status `json:"status"`
}
