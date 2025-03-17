const jwt = require('jsonwebtoken');
const {JWT_SECRET} = process.env;

exports.generateToken = (merchant) => {
  return jwt.sign({ merchant }, process.env.JWT_SECRET, {
    expiresIn: process.env.JWT_EXPIRES_IN
  });
};