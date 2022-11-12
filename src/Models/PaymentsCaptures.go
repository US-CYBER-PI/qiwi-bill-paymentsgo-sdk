package Models

import "github.com/US-CYBER-PI/qiwi-bill-paymentsgo-sdk/src/Models/Ets"

type PaymentsCaptures struct {
	CallbackUrl string `json:"callbackUrl"`
	Comment     string `json:"comment"`
}

func NewPaymentsCaptures(callbackUrl string, comment string) *PaymentsCaptures {
	return &PaymentsCaptures{
		CallbackUrl: callbackUrl,
		Comment:     comment,
	}
}

type PaymentsCapturesResponses struct {
	CaptureId       string     `json:"captureId"`
	CreatedDatetime string     `json:"createdDatetime"`
	Amount          Ets.Amount `json:"amount"`
	Status          Ets.Status `json:"status"`
}
