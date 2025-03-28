package merchant

type AuthFormatter struct {
	MerchantName string `json:"merchant_name"`
	Token        string `json:"token"`
}

func FormatAuth(merchant Merchant, token string) AuthFormatter {

	formatter := AuthFormatter{
		MerchantName: merchant.MerchantName,
		Token:        token,
	}

	return formatter
}
