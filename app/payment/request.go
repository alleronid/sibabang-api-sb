package payment

import "lanaya/api/app/merchant"

type PaymentInput struct {
	TrxId          string `json:"no_transaction" binding:"required"`
	Amount         int    `json:"amount" binding:"required"`
	PaymentChannel string `json:"payment_channel" binding:"required"`
	Email          string `json:"email" binding:"required"`
	Fullname       string `json:"fullname" binding:"required"`
	PhoneNumber    string `json:"phone_number" binding:"required"`
	Merchant       merchant.Merchant
}
