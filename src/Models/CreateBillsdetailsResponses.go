package Models

import "us-cyber-pi/qiwi-bill-paymentsgo-sdk/src/Models/ets"

type CreateBillsDetailsResponses struct {
	BillId  string     `json:"billId"`
	Amount  ets.Amount `json:"amount"`
	Status  string     `json:"status"`
	Comment string     `json:"comment"`
	PayUrl  string     `json:"payUrl"`

	ExpirationDateTime string `json:"expirationDateTime"`
	payUrl             string `json:"payUrl"`
}
