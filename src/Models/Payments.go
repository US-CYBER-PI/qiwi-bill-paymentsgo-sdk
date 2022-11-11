package Models

import "us-cyber-pi/qiwi-bill-paymentsgo-sdk/src/Models/ets"

type Payments struct {
	PaymentMethod ets.PaymentMethod `json:"paymentMethod"`
	Amount        ets.Amount        `json:"amount"`
}

func NewPayments(paymentMethod ets.PaymentMethod, amount ets.Amount) *Payments {
	return &Payments{
		PaymentMethod: paymentMethod,
		Amount:        amount,
	}
}

type PaymentsResponses struct {
	PaymentId       string              `json:"paymentId"`
	CreatedDatetime string              `json:"createdDatetime"`
	Amount          ets.Amount          `json:"amount"`
	CapturedAmount  ets.CapturedAmount  `json:"capturedAmount"`
	RefundedAmount  ets.RefundedAmount  `json:"refundedAmount"`
	PaymentMethod   ets.PaymentMethod   `json:"paymentMethod"`
	Status          ets.Status          `json:"status"`
	PaymentCardInfo ets.PaymentCardInfo `json:"paymentCardInfo"`
}
