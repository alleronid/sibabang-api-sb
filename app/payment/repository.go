package payment

import "gorm.io/gorm"

type Repository interface {
	CreatePayment(trxPayment TrxPayment) (TrxPayment, error)
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
