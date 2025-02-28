const express = require('express');
const router = express.Router();
const {
  getUsers,
  getUserById,
  createUser,
  loginUser,
  registerUser,
  verifyEmail,
  forgotPassword,
  resetPassword,
  checkUser
} = require('../controllers/userController');

router.route('/').get(getUsers).post(createUser);
router.route('/login').put(loginUser);
router.route('/register').put(registerUser);    
router.route('/verify-email').put(verifyEmail);
router.route('/:id').get(getUserById);
router.post('/forgot-password', forgotPassword);
router.post('/reset-password', resetPassword);
router.route('/check-user/:username').get(checkUser);

module.exports = router; 