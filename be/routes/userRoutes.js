const express = require('express');
const router = express.Router();
const {
  getUsers,
  getUserById,
  createUser,
  loginUser,
  registerUser,
  verifyEmail
} = require('../controllers/userController');

router.route('/').get(getUsers).post(createUser);
router.route('/login').put(loginUser);
router.route('/register').put(registerUser);    
router.route('/verify-email').put(verifyEmail);
router.route('/:id').get(getUserById);

module.exports = router; 