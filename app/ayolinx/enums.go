package ayolinx

type AyolinxEnums struct {
	// Links
	URL_PROD string
	URL_DEV  string

	// Channel
	QRIS      string
	EWALLET   string
	VABNI     string
	VACIMB    string
	VAMANDIRI string

	// PartnerID
	BNI_SB       string
	BNI_PROD     string
	CIMB_SB      string
	CIMB_PROD    string
	MANDIRI_SB   string
	MANDIRI_PROD string

	// Status code
	SUCCESS_CODE   int
	INITIATED_CODE int
	PAYING_CODE    int
	PENDING_CODE   int
	REFUNDED_CODE  int
	CANCEL_CODE    int
	FAILED_CODE    int
	NOT_FOUND      int

	// Response code
	SUCCESS_DANA                     string
	SUCCESS_QRIS                     string
	SUCCESS_VA_BNI                   string
	SUCCESS_VA_MANDIRI               string
	UNAUTHORIZED                     string
	SUCCESS_GET_TOKENVA              string
	SUCCESS_CALLBACKVA               string
	SUCCESS_CALLBACK                 string
	ERR_AYOLINK_PAYMENT_BAD_REQ      int
	ERR_AYOLINK_TOKEN_NO_AUTH_ERROR  int
	ERR_AYOLINK_PAYMENT_INVALID_SIGN int
}

func NewAyolinxEnums() *AyolinxEnums {
	return &AyolinxEnums{
		// Links
		URL_PROD: "https://openapi.ayolinx.id",
		URL_DEV:  "https://sandbox.ayolinx.id",

		// Channel
		QRIS:      "BNC_QRIS",
		EWALLET:   "EMONEY_DANA_SNAP",
		VABNI:     "VIRTUAL_ACCOUNT_BNI",
		VACIMB:    "VIRTUAL_ACCOUNT_CIMB",
		VAMANDIRI: "VIRTUAL_ACCOUNT_MANDIRI",

		// PartnerID
		BNI_SB:       "98829172",
		BNI_PROD:     "98828222",
		CIMB_SB:      "2056",
		CIMB_PROD:    "2056",
		MANDIRI_SB:   "87319",
		MANDIRI_PROD: "87319",

		// Status code
		SUCCESS_CODE:   0,
		INITIATED_CODE: 1,
		PAYING_CODE:    2,
		PENDING_CODE:   3,
		REFUNDED_CODE:  4,
		CANCEL_CODE:    5,
		FAILED_CODE:    6,
		NOT_FOUND:      7,

		// Response code
		SUCCESS_DANA:                     "2005400",
		SUCCESS_QRIS:                     "2004700",
		SUCCESS_VA_BNI:                   "2002700",
		SUCCESS_VA_MANDIRI:               "2002700",
		UNAUTHORIZED:                     "581000001",
		SUCCESS_GET_TOKENVA:              "2007300",
		SUCCESS_CALLBACKVA:               "2002500",
		SUCCESS_CALLBACK:                 "2005600",
		ERR_AYOLINK_PAYMENT_BAD_REQ:      4007300,
		ERR_AYOLINK_TOKEN_NO_AUTH_ERROR:  4017300,
		ERR_AYOLINK_PAYMENT_INVALID_SIGN: 4012501,
	}
}
