package payment

import (
	"encoding/json"
	"fmt"
	"lanaya/api/app/ayolinx"
)

type PaymentService interface {
	SavePayment(input PaymentInput) (TrxPayment, error)
	GetTransaction(trxId string) (TrxPayment, error)
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

	qrisData := map[string]interface{}{
		"partnerReferenceNo": input.TrxId,
		"amount": map[string]interface{}{
			"currency": "IDR",
			"value":    input.Amount,
		},
		"additionalInfo": map[string]interface{}{
			"channel":       "BNC_QRIS",
			"subMerchantId": input.Merchant.MidQris,
		},
	}

	qrisResponse, err := s.ayolinxService.GenerateQris(qrisData)
	if err != nil {
		return transaction, err
	}

	var qrisResult map[string]interface{}
	if err := json.Unmarshal([]byte(qrisResponse), &qrisResult); err != nil {
		return transaction, err
	}

	if responseCode, ok := qrisResult["responseCode"].(string); ok {
		if responseCode != "2004700" {
			return transaction, fmt.Errorf("QRIS generation failed: %s", qrisResult["responseMessage"])
		}
	}

	if qrContent, ok := qrisResult["qrContent"].(string); ok {
		transaction.PaymentCode = qrContent
	}

	transaction.DataRaw = qrisResponse

	newTransaction, err := s.repository.CreatePayment(transaction)
	if err != nil {
		return newTransaction, err
	}

	return newTransaction, nil
}

func (s *servicePayment) GetTransaction(trxId string) (TrxPayment, error) {
	payment, err := s.repository.FindPayment(trxId)

	if err != nil {
		return payment, err
	}

	return payment, nil
}
