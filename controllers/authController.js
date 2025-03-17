const AuthService = require('../services/authService');

exports.generateToken = async (req, res) => {
  try{
    const result = await AuthService.generateToken(req);

    res.cookie('jwt', result.token, {
      httpOnly: true,
      secure: process.env.NODE_ENV === 'production',
      maxAge: 24 * 60 * 60 * 1000
    });

    return res.status(201).json({
      success: true,
      message: 'Token generated successfully',
      token : result
    });
  }catch(error){
    return res.status(error.message === 'User already exists' ? 400 : 500).json({
      success: false,
      message: error.message || 'Server Error'
    });
  }
}