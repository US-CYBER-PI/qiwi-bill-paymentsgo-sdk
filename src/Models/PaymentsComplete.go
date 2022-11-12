package Models

import "github.com/US-CYBER-PI/qiwi-bill-paymentsgo-sdk/src/Models/Ets"

type PaymentsComplete struct {
	ThreeDS Ets.ThreeDS `json:"threeDS"`
}

func NewPaymentsComplete(threeDS Ets.ThreeDS) *PaymentsComplete {
	return &PaymentsComplete{
		ThreeDS: threeDS,
	}
}

type PaymentsCompleteResponses struct {
	PaymentId       string             `json:"paymentId"`
	CreatedDatetime string             `json:"createdDatetime"`
	Amount          Ets.Amount         `json:"amount"`
	CapturedAmount  Ets.CapturedAmount `json:"capturedAmount"`
	RefundedAmount  Ets.RefundedAmount `json:"refundedAmount"`
	PaymentMethod   Ets.PaymentMethod  `json:"paymentMethod"`
	Status          Ets.Status         `json:"status"`
}
