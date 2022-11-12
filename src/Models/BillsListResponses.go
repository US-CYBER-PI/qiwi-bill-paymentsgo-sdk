package Models

import "github.com/US-CYBER-PI/qiwi-bill-paymentsgo-sdk/src/Models/ets"

type BillsListResponses struct {
	NA ets.NA `json:"n/a"`
}

func NewBillsListResponses(NA ets.NA) *BillsListResponses {
	return &BillsListResponses{
		NA: NA,
	}
}
