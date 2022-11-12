package src

import (
	"encoding/json"
	"fmt"
	"github.com/US-CYBER-PI/qiwi-bill-paymentsgo-sdk/src/Models"
)

type Client struct {
	token  string
	host   string
	siteId string
}

func NewClient(token string, host string, siteId string) *Client {
	return &Client{
		token:  token,
		host:   host,
		siteId: siteId,
	}
}

func (s *Client) HttpRequest(js []byte, uri, method string) []byte {

	return HttpRequest(js,
		fmt.Sprintf("%s/%s", s.host, uri),
		method,
		map[string]string{
			"Authorization": fmt.Sprintf("Bearer %s", s.token),
			"Accept":        "application/json",
			"Content-type":  "application/json",
		},
	)

}

func (s *Client) SmsConfirm(bills Models.SmsConfirm) *Models.SmsConfirm {

	var smsConfirm Models.SmsConfirm

	b, err := json.Marshal(bills)
	if err == nil {
	}
	var body = s.HttpRequest(b, fmt.Sprintf("payin-tokenization-api/v1/sites/%s/token-requests/complete", s.siteId), "POST")

	json.Unmarshal(body, &smsConfirm)

	return &smsConfirm
}

func (s *Client) GeneratePayToken(bills Models.CreateBills) *Models.GeneratePayToken {

	var generatePayToken Models.GeneratePayToken

	b, err := json.Marshal(bills)
	if err == nil {
	}

	var body = s.HttpRequest(b, fmt.Sprintf("payin-tokenization-api/v1/sites/%s/token-requests", s.siteId), "POST")

	json.Unmarshal(body, &generatePayToken)

	return &generatePayToken
}

func (s *Client) CreateBills(bills Models.CreateBills, billId string) *Models.CreateBillsResponses {

	var createBillsResponses Models.CreateBillsResponses

	b, err := json.Marshal(bills)
	if err == nil {
	}
	var body = s.HttpRequest(b, fmt.Sprintf("payin/v1/sites/%s/bills/%s", s.siteId, billId), "PUT")

	json.Unmarshal(body, &createBillsResponses)

	return &createBillsResponses
}

func (s *Client) CreateBillsDetailsResponses(bills Models.CreateBills, billId string) *Models.CreateBillsDetailsResponses {

	var createBillsDetailsResponses Models.CreateBillsDetailsResponses

	_, err := json.Marshal(bills)
	if err == nil {
	}
	var body = s.HttpRequest(nil, fmt.Sprintf("payin/v1/sites/%s/bills/%s/details", s.siteId, billId), "GET")

	json.Unmarshal(body, &createBillsDetailsResponses)

	return &createBillsDetailsResponses
}

func (s *Client) BillsListResponses(bills Models.CreateBills, billId string) *Models.BillsListResponses {

	var billsListResponses Models.BillsListResponses
	_, err := json.Marshal(bills)
	if err == nil {
	}
	var body = s.HttpRequest(nil, fmt.Sprintf("/payin/v1/sites/%s/bills/%s/details", s.siteId, billId), "GET")
	json.Unmarshal(body, &billsListResponses)
	return &billsListResponses
}

func (s *Client) Payments(bills Models.CreateBills, billId string) *Models.Payments {

	var payments Models.Payments
	b, err := json.Marshal(bills)
	if err == nil {
	}
	var body = s.HttpRequest(b, fmt.Sprintf("/payin/v1/sites/%s/bills/%s/details", s.siteId, billId), "GET")
	json.Unmarshal(body, &payments)
	return &payments
}

func (s *Client) CreatePayment(bills Models.CreatePayment, paymentId string) *Models.PaymentsStatusResponses {

	var paymentsStatusResponses Models.PaymentsStatusResponses
	jsonString, err := json.Marshal(bills)
	if err != nil {
		panic(err)
	}

	var body = s.HttpRequest(jsonString, fmt.Sprintf("payin/v1/sites/%s/payments/%s", s.siteId, paymentId), "PUT")
	json.Unmarshal(body, &paymentsStatusResponses)
	return &paymentsStatusResponses
}

func (s *Client) PaymentsComplete(bills Models.CreateBills, paymentId string) *Models.PaymentsComplete {

	var paymentsComplete Models.PaymentsComplete
	b, err := json.Marshal(bills)
	if err == nil {
	}

	var body = s.HttpRequest(b, fmt.Sprintf("payin/v1/sites/%s/payments/%s/complete", s.siteId, paymentId), "POST")
	json.Unmarshal(body, &paymentsComplete)
	return &paymentsComplete
}
func (s *Client) PaymentsCaptures(bills Models.CreateBills, paymentId, captureId string) *Models.PaymentsCaptures {

	var payments Models.PaymentsCaptures
	b, err := json.Marshal(bills)
	if err == nil {
	}
	var body = s.HttpRequest(b, fmt.Sprintf("payin/v1/sites/%s/payments/%s/captures/%s", s.siteId, paymentId, captureId), "PUT")
	json.Unmarshal(body, &payments)
	return &payments
}

func (s *Client) PaymentsStatusCapturesResponses(bills Models.CreateBills, paymentId, captureId string) *Models.PaymentsStatusCapturesResponses {

	var paymentsStatusCapturesResponses Models.PaymentsStatusCapturesResponses
	_, err := json.Marshal(bills)
	if err == nil {
	}
	var body = s.HttpRequest(nil, fmt.Sprintf("payin/v1/sites/%s/payments/%s/captures/%s", s.siteId, paymentId, captureId), "GET")
	json.Unmarshal(body, &paymentsStatusCapturesResponses)
	return &paymentsStatusCapturesResponses
}
