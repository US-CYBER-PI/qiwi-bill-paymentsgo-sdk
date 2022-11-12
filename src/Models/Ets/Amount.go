package Ets

type Amount struct {
	Value    string `json:"value"`
	Currency string `json:"currency"` // 3 символа
}

func NewAmount(value string, currency string) *Amount {
	return &Amount{
		Value:    value,
		Currency: currency,
	}
}
