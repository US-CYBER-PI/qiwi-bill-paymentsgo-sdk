package Models

import "github.com/US-CYBER-PI/qiwi-bill-paymentsgo-sdk/src/Models/ets"

type PaymentsStatusCapturesResponses struct {
	CaptureId       string     `json:"captureId"`
	CreatedDatetime string     `json:"createdDatetime"`
	Amount          ets.Amount `json:"amount"`
	Status          ets.Status `json:"status"`
}
