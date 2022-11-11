package ets

type CapturedAmount struct {
	Value    float64 `json:"value"`
	Currency string  `json:"currency"`
}

func NewCapturedAmount(value float64, currency string) *CapturedAmount {
	return &CapturedAmount{
		Value:    value,
		Currency: currency,
	}
}
