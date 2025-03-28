package payment

type TrxPayment struct {
	ID          int
	TrxId       string
	Amount      int
	Email       string
	Fullname    string
	PhoneNumber string
	Method      string
	Status      string
	MerchantId  int
	CompanyId   int
	PaymentCode string
	DataRaw     string
	Ket         string
	CreatedAt   string
	UpdatedAt   string
}

func (TrxPayment) TableName() string {
	return "trx_payments"
}
