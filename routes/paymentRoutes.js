const express = require('express');
const router = express.Router();
const paymentController = require('../controllers/paymentController');

router.post("/generate-qris", paymentController.generateQris);

module.exports = router;