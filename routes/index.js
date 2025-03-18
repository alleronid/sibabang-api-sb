const express = require('express');
const authRoutes = require('./authRoutes');
const paymentRoutes = require('./paymentRoutes');

const router = express.Router();

router.use('/auth', authRoutes);
router.use('/payment', paymentRoutes);

module.exports = router;