package Ets

type PaymentCardInfo struct {
	IssuingCountry string `json:"issuingCountry"`
	PaymentSystem  string `json:"paymentSystem"`
}

func NewPaymentCardInfo(issuingCountry string, paymentSystem string) *PaymentCardInfo {
	return &PaymentCardInfo{
		IssuingCountry: issuingCountry,
		PaymentSystem:  paymentSystem,
	}
}
