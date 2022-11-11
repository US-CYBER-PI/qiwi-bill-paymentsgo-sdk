package Models

import "us-cyber-pi/qiwi-bill-paymentsgo-sdk/src/Models/ets"

type PaymentsComplete struct {
	ThreeDS ets.ThreeDS `json:"threeDS"`
}

func NewPaymentsComplete(threeDS ets.ThreeDS) *PaymentsComplete {
	return &PaymentsComplete{
		ThreeDS: threeDS,
	}
}

type PaymentsCompleteResponses struct {
	PaymentId       string             `json:"paymentId"`
	CreatedDatetime string             `json:"createdDatetime"`
	Amount          ets.Amount         `json:"amount"`
	CapturedAmount  ets.CapturedAmount `json:"capturedAmount"`
	RefundedAmount  ets.RefundedAmount `json:"refundedAmount"`
	PaymentMethod   ets.PaymentMethod  `json:"paymentMethod"`
	Status          ets.Status         `json:"status"`
}
