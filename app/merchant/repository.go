package merchant

import (
	"fmt"

	"gorm.io/gorm"
)

type Repository interface {
	GetMerchant(clientKey string, token string) (Merchant, error)
	FindByID(merchantID int) (Merchant, error)
	FindByClientKey(clientKey string) (Merchant, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetMerchant(clientKey string, token string) (Merchant, error) {
	var merchant Merchant

	err := r.db.Where("api_key_sb = ? AND token = ?", clientKey, token).First(&merchant).Error

	if err != nil {
		return merchant, err
	}

	fmt.Printf("Merchant: %+v\n", merchant)

	return merchant, nil
}

func (r *repository) FindByID(merchantID int) (Merchant, error) {
	var merchant Merchant

	err := r.db.Where("merchant_id =?", merchantID).First(&merchant).Error

	if err != nil {
		return merchant, err
	}

	return merchant, nil
}

func (r *repository) FindByClientKey(clientKey string) (Merchant, error) {
	var merchant Merchant

	err := r.db.Where("api_key_sb = ?", clientKey).First(&merchant).Error

	if err != nil {
		return merchant, err
	}

	return merchant, nil
}
