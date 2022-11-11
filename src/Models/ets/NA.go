package ets

type NA struct {
	PaymentId       string          `json:"paymentId"`
	BillId          string          `json:"billId"`
	CreatedDatetime string          `json:"createdDatetime"`
	Amount          Amount          `json:"amount"`
	CapturedAmount  CapturedAmount  `json:"capturedAmount"`
	RefundedAmount  RefundedAmount  `json:"refundedAmount"`
	PaymentMethod   PaymentMethod   `json:"paymentMethod"`
	Status          Status          `json:"status"`
	PaymentCardInfo PaymentCardInfo `json:"paymentCardInfo"`
}

func NewNA(paymentId string, billId string, createdDatetime string, amount Amount, capturedAmount CapturedAmount, refundedAmount RefundedAmount, paymentMethod PaymentMethod, status Status, paymentCardInfo PaymentCardInfo) *NA {
	return &NA{
		PaymentId:       paymentId,
		BillId:          billId,
		CreatedDatetime: createdDatetime,
		Amount:          amount,
		CapturedAmount:  capturedAmount,
		RefundedAmount:  refundedAmount,
		PaymentMethod:   paymentMethod,
		Status:          status,
		PaymentCardInfo: paymentCardInfo,
	}
}
