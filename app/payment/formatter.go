package payment

import "encoding/json"

type PaymentFormatter struct {
	NoTransaction  string `json:"no_transaction"`
	CustomerName   string `json:"customer_name"`
	PaymentChannel string `json:"payment_channel"`
	Amount         int    `json:"amount"`
	QrContent      string `json:"qrContent"`
	QrLink         string `json:"qr_link"`
}

func FormatPayment(payment TrxPayment) PaymentFormatter {
	var dataRaw map[string]interface{}
	if err := json.Unmarshal([]byte(payment.DataRaw), &dataRaw); err != nil {
		dataRaw = make(map[string]interface{})
	}

	qrLink := ""
	if url, ok := dataRaw["qrUrl"].(string); ok {
		qrLink = url
	}

	formatter := PaymentFormatter{
		NoTransaction:  payment.TrxId,
		CustomerName:   payment.Fullname,
		Amount:         payment.Amount,
		PaymentChannel: payment.Method,
		QrContent:      payment.PaymentCode,
		QrLink:         qrLink,
	}

	return formatter
}