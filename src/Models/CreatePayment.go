package Models

import "github.com/US-CYBER-PI/qiwi-bill-paymentsgo-sdk/src/Models/Ets"

type CreatePayment struct {
	Amount        Ets.Amount        `json:"amount"`
	PaymentMethod Ets.PaymentMethod `json:"paymentMethod"`
	Customer      Ets.Customer      `json:"customer"`
	Flags         []string          `json:"flags"`
}
