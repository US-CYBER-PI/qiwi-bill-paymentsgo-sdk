package Models

import "github.com/US-CYBER-PI/qiwi-bill-paymentsgo-sdk/src/Models/Ets"

type CreateBillsDetailsResponses struct {
	BillId  string     `json:"billId"`
	Amount  Ets.Amount `json:"amount"`
	Status  string     `json:"status"`
	Comment string     `json:"comment"`
	PayUrl  string     `json:"payUrl"`

	ExpirationDateTime string `json:"expirationDateTime"`
	payUrl             string `json:"payUrl"`
}
