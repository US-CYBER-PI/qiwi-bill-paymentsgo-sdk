package Models

import "github.com/US-CYBER-PI/qiwi-bill-paymentsgo-sdk/src/Models/Ets"

type BillsListResponses struct {
	NA Ets.NA `json:"n/a"`
}

func NewBillsListResponses(NA Ets.NA) *BillsListResponses {
	return &BillsListResponses{
		NA: NA,
	}
}
