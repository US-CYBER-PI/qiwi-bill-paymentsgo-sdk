package Models

import "github.com/US-CYBER-PI/qiwi-bill-paymentsgo-sdk/src/Models/Ets"

type Payments struct {
	PaymentMethod Ets.PaymentMethod `json:"paymentMethod"`
	Amount        Ets.Amount        `json:"amount"`
}

func NewPayments(paymentMethod Ets.PaymentMethod, amount Ets.Amount) *Payments {
	return &Payments{
		PaymentMethod: paymentMethod,
		Amount:        amount,
	}
}

type PaymentsResponses struct {
	PaymentId       string              `json:"paymentId"`
	CreatedDatetime string              `json:"createdDatetime"`
	Amount          Ets.Amount          `json:"amount"`
	CapturedAmount  Ets.CapturedAmount  `json:"capturedAmount"`
	RefundedAmount  Ets.RefundedAmount  `json:"refundedAmount"`
	PaymentMethod   Ets.PaymentMethod   `json:"paymentMethod"`
	Status          Ets.Status          `json:"status"`
	PaymentCardInfo Ets.PaymentCardInfo `json:"paymentCardInfo"`
}
