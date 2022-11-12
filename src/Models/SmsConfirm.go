package Models

type SmsConfirm struct {
	RequestId string `json:"requestId"`
	SmsCode   string `json:"smsCode"`
}

func NewSmsConfirm(RequestId string, SmsCode string) *SmsConfirm {
	return &SmsConfirm{
		RequestId: RequestId,
		SmsCode:   SmsCode,
	}
}

type SmsConfirmResponses struct {
	TokenValue       string `json:"token.value"`
	TokenExpiredDate string `json:"token.expiredDate"`
	Status           string `json:"status.value"`
}
