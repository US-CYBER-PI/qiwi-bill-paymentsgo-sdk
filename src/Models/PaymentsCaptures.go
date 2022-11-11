package Models

import "us-cyber-pi/qiwi-bill-paymentsgo-sdk/src/Models/ets"

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
	Amount          ets.Amount `json:"amount"`
	Status          ets.Status `json:"status"`
}
