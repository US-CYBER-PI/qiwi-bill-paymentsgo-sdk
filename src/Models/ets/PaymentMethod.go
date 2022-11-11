package ets

type PaymentMethod struct {
	Type      string `json:"type"`
	MaskedPan string `json:"maskedPan"`
}

func NewPaymentMethod(type_ string, maskedPan string) *PaymentMethod {
	return &PaymentMethod{
		Type:      type_,
		MaskedPan: maskedPan,
	}
}
