package Ets

type RefundedAmount struct {
	Value    float64 `json:"value"`
	Currency string  `json:"currency"`
}

func NewRefundedAmount(value float64, currency string) *RefundedAmount {
	return &RefundedAmount{
		Value:    value,
		Currency: currency,
	}
}
