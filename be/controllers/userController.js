const User = require('../models/userModel');
const jwt = require('jsonwebtoken');
const crypto = require('crypto');
const sendEmail = require('../utils/sendEmail');
const bcrypt = require('bcrypt');
// Get all users
const getUsers = async (req, res) => {
  try {
    const users = await User.find({});
    res.json(users);
  } catch (error) {
    res.status(500).json({ message: error.message });
  }
};

// Get single user
const getUserById = async (req, res) => {
  try {
    const user = await User.findById(req.params.id);
    if (user) {
      res.json(user);
    } else {
      res.status(404).json({ message: 'User not found' });
    }
  } catch (error) {
    res.status(500).json({ message: error.message });
  }
};

// Create user
const createUser = async (req, res) => {
  try {
    console.log('Received request body:', req.body);
    const user = await User.create(req.body);
    res.status(201).json(user);
  } catch (error) {
    res.status(400).json({ message: error.message });
  }
};


const registerUser = async (req, res) => {
  try {
    const { username, newusername, email, password } = req.body;

    if (!username || !newusername || !email || !password) {
      return res.status(400).json({ message: 'Please provide username, new username, email, and password' });
    }

    const user = await User.findOne({ username });
    if (!user) {
      return res.status(404).json({ message: 'User not found' });
    }

    const existingUser = await User.findOne({
      $or: [{ email }, { username: newusername }],
      _id: { $ne: user._id }  
    });

    if (existingUser) {
      return res.status(400).json({ message: 'Email or Username already exists' });
    }

    const verificationToken = crypto.randomBytes(32).toString('hex');
    const verificationTokenExpires = new Date(Date.now() + 24 * 60 * 60 * 1000); // 24 hours expiration

    user.username = newusername;
    user.email = email;
    user.password = password; 
    user.password = await bcrypt.hash(password, 10);
    user.verificationToken =  verificationToken;
    user.verificationTokenExpires = verificationTokenExpires;
    await user.save();

    const verificationLink = `${process.env.BASE_URL}/verify-email?token=${verificationToken}`;
    await sendEmail(email, 'Verify Your Email', `Click here to verify your email: ${verificationLink}`);

    res.status(200).json({
      _id: user._id,
      email: user.email,
      username: user.username
    });

  } catch (error) {
    console.error('Error updating user:', error);
    res.status(500).json({ message: 'Internal server error' });
  }
};

const loginUser = async (req, res) => {
  try {
    console.log('Received request body:', req.body);
    const { email, password } = req.body;

    if (!email || !password) {
      return res.status(400).json({ message: 'Please provide email and password' });
    }
 

    const user = await User.findOne({ email });
    if (!user) {
      return res.status(401).json({ message: 'Invalid email or password' });
    }

    const isPasswordValid = await user.comparePassword(password);

    if (!isPasswordValid) {
      return res.status(401).json({ message: 'Invalid username or password' });
    }

    if (!user.emailVerified) {
      return res.status(403).json({ 
        message: 'Email not verified. Please verify your account before logging in.',
        isEmailVerified: false
      });
    }

    res.status(200).json({
      _id: user._id,
      username: user.username,
      email: user.email,
      isEmailVerified: true
    });

  } catch (error) {
    console.error('Error during login:', error);
    res.status(500).json({ message: 'Internal server error' });
  }
};

const verifyEmail = async (req, res) => {
  try {
    const { token } = req.body;

    console.log(req.body);

    if (!token) {
      return res.status(400).json({ message: 'Verification token is required' });
    }

    const user = await User.findOne({
      verificationToken: token,
      verificationTokenExpires: { $gt: Date.now() }
    });

    if (!user) {
      return res.status(400).json({ 
        message: 'Invalid or expired verification token'
      });
    }

    user.emailVerified = true;
    user.verificationToken = undefined;
    user.verificationTokenExpires = undefined;
    await user.save();

    res.status(200).json({ 
      message: 'Email verified successfully',
      isEmailVerified: true
    });

  } catch (error) {
    console.error('Error verifying email:', error);
    res.status(500).json({ message: 'Internal server error' });
  }
};

module.exports = {
  getUsers,
  getUserById,
  createUser,
  registerUser,
  loginUser,
  verifyEmail
}; 