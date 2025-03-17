class AyolinxEnums {
  //links
  static URL_PROD = 'https://openapi.ayolinx.id';
  static URL_DEV = 'https://sandbox.ayolinx.id';

  //channel
  static QRIS = 'BNC_QRIS';
  static EWALLET = 'EMONEY_DANA_SNAP';
  static VABNI = 'VIRTUAL_ACCOUNT_BNI';
  static VACIMB = 'VIRTUAL_ACCOUNT_CIMB';
  static VAMANDIRI = 'VIRTUAL_ACCOUNT_MANDIRI';

  //partnerID
  static BNI_SB = "98829172";
  static BNI_PROD = "98828222";
  static CIMB_SB = "2056";
  static CIMB_PROD = "2056";
  static MANDIRI_SB = "87319";
  static MANDIRI_PROD = "87319";

  //status code
  static SUCCESS_CODE = '00';
  static INITIATED_CODE = '01';
  static PAYING_CODE = '02';
  static PENDING_CODE = '03';
  static REFUNDED_CODE = '04';
  static CANCEL_CODE = '05';
  static FAILED_CODE = '06';
  static NOT_FOUND = '07';

  //response code
  static SUCCESS_DANA = '2005400';
  static SUCCESS_QRIS = '2004700';
  static SUCCESS_VA_BNI = '2002700';
  static SUCCESS_VA_MANDIRI = '2002700';
  static UNAUTHORIZED = '581000001';
  static SUCCESS_GET_TOKENVA = '2007300';
  static SUCCESS_CALLBACKVA = '2002500';
  static SUCCESS_CALLBACK = '2005600';
  static ERR_AYOLINK_PAYMENT_BAD_REQ = 4007300;
  static ERR_AYOLINK_TOKEN_NO_AUTH_ERROR = 4017300;
  static ERR_AYOLINK_PAYMENT_INVALID_SIGN = 4012501;
}

module.exports = AyolinxEnums;