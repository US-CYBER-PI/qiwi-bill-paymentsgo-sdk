package ets

type PaymentMethod struct {
	Type         string `json:"type"`
	PaymentToken string `json:"paymentToken"`
}

func NewPaymentMethod(type_ string, paymentToken string) *PaymentMethod {
	return &PaymentMethod{
		Type:         type_,
		PaymentToken: paymentToken,
	}
}
