package Ets

type Amount struct {
	Value    float64 `json:"value"`
	Currency string  `json:"currency"` // 3 символа
}

func NewAmount(value float64, currency string) *Amount {
	return &Amount{
		Value:    value,
		Currency: currency,
	}
}
