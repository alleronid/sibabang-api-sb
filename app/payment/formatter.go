package payment

type PaymentFormatter struct {
	NoTransaction  string `json:"no_transaction"`
	CustomerName   string `json:"customer_name"`
	PaymentChannel string `json:"payment_channel"`
	Amount         int    `json:"amount"`
	QrContent      string `json:"qrContent"`
	QrLink         string `json:"qr_link"`
}

func FormatPayment(payment TrxPayment) PaymentFormatter {

	formatter := PaymentFormatter{
		NoTransaction:  payment.TrxId,
		CustomerName:   payment.Fullname,
		Amount:         payment.Amount,
		PaymentChannel: payment.Method,
		QrContent:      payment.PaymentCode,
		QrLink:         "www.facebook.com",
	}

	return formatter
}
