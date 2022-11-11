package Models

import "us-cyber-pi/qiwi-bill-paymentsgo-sdk/src/Models/ets"

type BillsListResponses struct {
	NA ets.NA `json:"n/a"`
}

func NewBillsListResponses(NA ets.NA) *BillsListResponses {
	return &BillsListResponses{
		NA: NA,
	}
}
