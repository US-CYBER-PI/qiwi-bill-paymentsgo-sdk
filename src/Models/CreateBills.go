package Models

import "github.com/US-CYBER-PI/qiwi-bill-paymentsgo-sdk/src/Models/Ets"

type CreateBills struct {
	Amount             Ets.Amount `json:"amount"`
	ExpirationDateTime string     `json:"expirationDateTime"`
	Currency           string     `json:"currency"`
}

func NewCreateBills(expirationDateTime string, currency string, amount Ets.Amount) *CreateBills {
	return &CreateBills{
		Amount:             amount,
		ExpirationDateTime: expirationDateTime,
		Currency:           currency,
	}
}

type CreateBillsResponses struct {
	SiteId             string     `json:"siteId"`
	BillId             string     `json:"billId"`
	InvoiceUid         string     `json:"invoiceUid"`
	Amount             Ets.Amount `json:"amount"`
	Status             Ets.Status `json:"status"`
	Comment            string     `json:"comment"`
	CreationDateTime   string     `json:"creationDateTime"`
	ExpirationDateTime string     `json:"expirationDateTime"`
	PayUrl             string     `json:"payUrl"`
}
