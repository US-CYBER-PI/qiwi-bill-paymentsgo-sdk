package Models

import "github.com/US-CYBER-PI/qiwi-bill-paymentsgo-sdk/src/Models/Ets"

type SmsConfirm struct {
	RequestId string `json:"requestId"`
	SmsCode   string `json:"smsCode"`
}

func NewSmsConfirm(RequestId string, SmsCode string) *SmsConfirm {
	return &SmsConfirm{
		RequestId: RequestId,
		SmsCode:   SmsCode,
	}
}

type SmsConfirmResponses struct {
	TokenValue Ets.Token  `json:"token"`
	Status     Ets.Status `json:"status"`
}
