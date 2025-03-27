package payment

import "gorm.io/gorm"

type Repository interface {
	CreatePayment(trxPayment TrxPayment) (TrxPayment, error)
	FindPayment(trxId string) (TrxPayment, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreatePayment(trxPayment TrxPayment) (TrxPayment, error) {
	err := r.db.Create(&trxPayment).Error

	if err != nil {
		return trxPayment, err
	}

	return trxPayment, nil
}

func (r *repository) FindPayment(trxId string) (TrxPayment, error) {
	var payment TrxPayment
	err := r.db.Where("trx_id = ?", trxId).First(&payment).Error

	if err != nil {
		return payment, err
	}

	return payment, nil
}
