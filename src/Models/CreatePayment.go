package Models

import "github.com/US-CYBER-PI/qiwi-bill-paymentsgo-sdk/src/Models/ets"

type CreatePayment struct {
	Amount        ets.Amount        `json:"amount"`
	PaymentMethod ets.PaymentMethod `json:"paymentMethod"`
	Customer      ets.Customer      `json:"customer"`
	Flags         []string          `json:"flags"`
}
