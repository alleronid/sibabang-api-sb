const { Merchant } = require('../models');
const jwt = require('jsonwebtoken'); 

class AuthService {
  async generateToken(req) {
    const key = req.headers['x-api-key'];
    const secret = req.headers['x-secret-key'];
    const merchant = await Merchant.findOne({
      where: {
        api_key_sb: key,
        token: secret
      }
    });
    if (!merchant) {
      return null;
    }
    const token = jwt.sign({
         merchant_id : merchant.merchant_id,
         merchant_name  : merchant.merchant_name,
         company_id : merchant.company_id,
         token : merchant.token,
         api_key_sb : merchant.api_key_sb,
         cb_key_sb: merchant.cb_key_sb,
    }, process.env.JWT_SECRET, { expiresIn: '1d' });

    return token
  }
}

module.exports = new AuthService();