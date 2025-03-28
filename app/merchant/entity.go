package merchant

type Merchant struct {
	MerchantId   int
	MerchantName string
	Env          string
	ApiKeySb     string
	CbKeySb      string
	Token        string
	NmidQris     string
	MidQris      string
	MidDana      string
	CompanyId    int
}

func (Merchant) TableName() string {
	return "merchants"
}
