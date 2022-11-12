package Models

import "github.com/US-CYBER-PI/qiwi-bill-paymentsgo-sdk/src/Models/Ets"

type PaymentsStatusResponses struct {
	PaymentId       string             `json:"paymentId"`
	BillId          string             `json:"billId"`
	CreatedDatetime string             `json:"createdDatetime"`
	Amount          Ets.Amount         `json:"amount"`
	CapturedAmount  Ets.CapturedAmount `json:"capturedAmount"`
	RefundedAmount  Ets.RefundedAmount `json:"refundedAmount"`
	PaymentMethod   Ets.PaymentMethod  `json:"paymentMethod"`
	Status          Ets.Status         `json:"status"`
}
