package Models

import "github.com/US-CYBER-PI/qiwi-bill-paymentsgo-sdk/src/Models/Ets"

type PaymentsStatusCapturesResponses struct {
	CaptureId       string     `json:"captureId"`
	CreatedDatetime string     `json:"createdDatetime"`
	Amount          Ets.Amount `json:"amount"`
	Status          Ets.Status `json:"status"`
}
