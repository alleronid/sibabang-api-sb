class AyolinxEnums {
  //links
  static get URL_PROD() { return 'https://openapi.ayolinx.id'; }
  static get URL_DEV() { return 'https://sandbox.ayolinx.id'; }

  //channel
  static get QRIS() { return 'BNC_QRIS'; }
  static get EWALLET() { return 'EMONEY_DANA_SNAP'; }
  static get VABNI() { return 'VIRTUAL_ACCOUNT_BNI'; }
  static get VACIMB() { return 'VIRTUAL_ACCOUNT_CIMB'; }
  static get VAMANDIRI() { return 'VIRTUAL_ACCOUNT_MANDIRI'; }

  //partnerID
  static get BNI_SB() { return "98829172"; }
  static get BNI_PROD() { return "98828222"; }
  static get CIMB_SB() { return "2056"; }
  static get CIMB_PROD() { return "2056"; }
  static get MANDIRI_SB() { return "87319"; }
  static get MANDIRI_PROD() { return "87319"; }

  //status code
  static get SUCCESS_CODE() { return '00'; }
  static get INITIATED_CODE() { return '01'; }
  static get PAYING_CODE() { return '02'; }
  static get PENDING_CODE() { return '03'; }
  static get REFUNDED_CODE() { return '04'; }
  static get CANCEL_CODE() { return '05'; }
  static get FAILED_CODE() { return '06'; }
  static get NOT_FOUND() { return '07'; }

  //response code
  static get SUCCESS_DANA() { return '2005400'; }
  static get SUCCESS_QRIS() { return '2004700'; }
  static get SUCCESS_VA_BNI() { return '2002700'; }
  static get SUCCESS_VA_MANDIRI() { return '2002700'; }
  static get UNAUTHORIZED() { return '581000001'; }
  static get SUCCESS_GET_TOKENVA() { return '2007300'; }
  static get SUCCESS_CALLBACKVA() { return '2002500'; }
  static get SUCCESS_CALLBACK() { return '2005600'; }
  static get ERR_AYOLINK_PAYMENT_BAD_REQ() { return 4007300; }
  static get ERR_AYOLINK_TOKEN_NO_AUTH_ERROR() { return 4017300; }
  static get ERR_AYOLINK_PAYMENT_INVALID_SIGN() { return 4012501; }
}

module.exports = AyolinxEnums;