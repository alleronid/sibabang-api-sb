package payment

import (
	"lanaya/api/app/ayolinx"
)

type PaymentService interface {
	SavePayment(input PaymentInput) (TrxPayment, error)
}

type servicePayment struct {
	repository     Repository
	ayolinxService *ayolinx.AyolinxService
}

func NewService(repository Repository, ayolinxService *ayolinx.AyolinxService) PaymentService {
	return &servicePayment{repository, ayolinxService}
}

func (s *servicePayment) SavePayment(input PaymentInput) (TrxPayment, error) {
	transaction := TrxPayment{}

	transaction.TrxId = input.TrxId
	transaction.Amount = input.Amount
	transaction.Email = input.Email
	transaction.Fullname = input.Fullname
	transaction.PhoneNumber = input.PhoneNumber
	transaction.Method = input.PaymentChannel
	transaction.Status = "UNPAID"
	transaction.MerchantId = input.Merchant.MerchantId
	transaction.CompanyId = input.Merchant.CompanyId

	// qrisData := map[string]interface{}{
	// 	"trx_id":       input.TrxId,
	// 	"amount":       input.Amount,
	// 	"fullname":     input.Fullname,
	// 	"email":        input.Email,
	// 	"phone_number": input.PhoneNumber,
	// }

	// qrisResponse, err := s.ayolinxService.GenerateQris(qrisData)
	// if err != nil {
	// 	return transaction, err
	// }

	// // Parse QRIS response
	// var qrisResult map[string]interface{}
	// if err := json.Unmarshal([]byte(qrisResponse), &qrisResult); err != nil {
	// 	return transaction, err
	// }

	// if qrContent, ok := qrisResult["qrContent"].(string); ok {
	// 	transaction.PaymentCode = qrContent
	// }

	// transaction.DataRaw = qrisResponse

	newTransaction, err := s.repository.CreatePayment(transaction)
	if err != nil {
		return newTransaction, err
	}

	return newTransaction, nil
}
