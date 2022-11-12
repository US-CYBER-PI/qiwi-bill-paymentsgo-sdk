package Models

import "github.com/US-CYBER-PI/qiwi-bill-paymentsgo-sdk/src/Models/ets"

type PaymentsStatusResponses struct {
	PaymentId       string             `json:"paymentId"`
	BillId          string             `json:"billId"`
	CreatedDatetime string             `json:"createdDatetime"`
	Amount          ets.Amount         `json:"amount"`
	CapturedAmount  ets.CapturedAmount `json:"capturedAmount"`
	RefundedAmount  ets.RefundedAmount `json:"refundedAmount"`
	PaymentMethod   ets.PaymentMethod  `json:"paymentMethod"`
	Status          ets.Status         `json:"status"`
}
