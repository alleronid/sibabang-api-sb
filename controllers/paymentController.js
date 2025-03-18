const ayolinxService = require('../services/ayolinxService');

class AyolinxController{
  async generateQris(req, res) {
    try {
      const data = req.body;
      const result = await ayolinxService.generateQris(data);
      
      return res.status(200).json(JSON.parse(result));
    } catch (error) {
      console.error('Error generating QRIS:', error);
      return res.status(500).json({
        success: false,
        message: 'Error generating QRIS',
        error: error.message
      });
    }
  }
}

module.exports = new AyolinxController();