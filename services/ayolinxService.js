const ayolinxEnums = require('../enums/AyolinxEnums');
const AppSetting = require('../models/appSetting');
const crypto = require('crypto');
const axios = require('axios');
const fs = require('fs');

class AyolinxService{
  constructor() {
    this.timestamp = new Date().toISOString();
    this.initialize();
  }

  async initialize() {
    try {
      const keySB = await AppSetting.findOne({ where: { key: 'ayolinx_key_sb' } });
      const secretSB = await AppSetting.findOne({ where: { key: 'ayolinx_secret_sb' } });
      const secretApp = await AppSetting.findOne({ where: { key: 'sibabang_secret' } });

      this.keySB = keySB ? keySB.value : null;
      this.secretSB = secretSB ? secretSB.value : null;
      this.secretApp = secretApp ? secretApp.value : null;
    } catch (error) {
      console.error('Error initializing AyolinxService:', error);
    }
  }

  signature() {
    try {
      const clientKey = this.secretSB;
      const requestTimestamp = this.timestamp;
      const stringToSign = `${clientKey}|${requestTimestamp}`;
      
      const privateKey = fs.readFileSync('/path/to/private_key.pem');
      
      const sign = crypto.createSign('SHA256');
      sign.update(stringToSign);
      const signature = sign.sign(privateKey);

      return signature.toString('base64');
    } catch (error) {
      console.error('Error generating signature:', error);
      throw error;
    }
  }

  async getToken() {
    try {
      const clientKey = this.keySB;
      const signature = this.signature();
      
      const headers = {
        'X-CLIENT-KEY': clientKey,
        'X-SIGNATURE': signature
      };
      
      const response = await this.api('/v1.0/access-token/b2b', headers);
      const result = response ? JSON.parse(response) : null;
      
      return result?.accessToken || null;
    } catch (error) {
      console.error('Error getting token:', error);
      return JSON.stringify({ error: error.message });
    }
  }

  async api(url, headers = [], post = {}) {
    try {
      const timestamp = this.timestamp;
      const defaultHeaders = {
        'Content-Type': 'application/json',
        'X-TIMESTAMP': timestamp
      };

      const mergedHeaders = { ...defaultHeaders };
      headers.forEach(header => {
        if (typeof header === 'string' && header.includes(': ')) {
          const [key, value] = header.split(': ');
          mergedHeaders[key] = value;
        } else if (typeof header === 'object') {
          Object.assign(mergedHeaders, header);
        }
      });
      
      const baseUrl = AyolinxEnums.URL_DEV + url;
      
      // Use axios for HTTP requests
      const response = await axios({
        method: 'POST',
        url: baseUrl,
        headers: mergedHeaders,
        data: post,
        timeout: 0,
        validateStatus: () => true
      });
      
      return JSON.stringify(response.data);
    } catch (error) {
      console.error('API Error:', error);
      throw error;
    }
  }

  async baseInterface(signature, timestamp, token, url, post) {
    try {
      const headers = {
        'X-TIMESTAMP': timestamp,
        'X-SIGNATURE': signature,
        'X-PARTNER-ID': this.keySB,
        'X-EXTERNAL-ID': this.randomNumber(),
        'Authorization': `Bearer ${token}`,
        'Content-Type': 'application/json'
      };
      
      const response = await axios({
        method: 'POST',
        url: AyolinxEnums.URL_DEV + url,
        headers: headers,
        data: post,
        timeout: 0,
        validateStatus: () => true
      });
            
      return JSON.stringify(response.data);
    } catch (error) {
      console.error('Base Interface Error:', error);
      throw error;
    }
  }

  async generateQris(data = {}) {
    const timestamp = this.timestamp;
    const method = 'POST';
    const urlSignature = "/v1.0/qr/qr-mpm-generate";
    const token = await this.getToken();
    const clientSecret = this.secretSB;
    const body = data;
    
    const hash = crypto.createHash('sha256');
    hash.update(JSON.stringify(body));
    const hexEncodedHash = hash.digest('hex').toLowerCase();
    
    const dataForSignature = `${method}:${urlSignature}:${token}:${hexEncodedHash}:${timestamp}`;
    
    const hmac = crypto.createHmac('sha512', clientSecret);
    hmac.update(dataForSignature);
    const signature = hmac.digest('base64');
    
    const response = await this.baseInterface(signature, timestamp, token, urlSignature, body);
    return response;
  }

  randomNumber() {
    return Date.now().toString() + Math.floor(Math.random() * 1000).toString();
  }

}

module.exports = new AyolinxService();