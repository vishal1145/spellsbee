const mongoose = require('mongoose');
const bcrypt = require('bcrypt');

const userSchema = new mongoose.Schema({
  username: { type: String, required: true },
  email: { type: String, required: true, unique: true },
  password: { type: String, required: true },
  emailVerified: { type: Boolean, default: false },
  verificationToken: { type: String },  
  verificationTokenExpires: { type: Date } 
}, { timestamps: true });

userSchema.methods.comparePassword = async function (candidatePassword) {
  try {
    console.log('Entered Password:', candidatePassword);
    console.log('Stored Hashed Password:', this.password);

    const isMatch = await bcrypt.compare(candidatePassword, this.password);
    console.log('Password Match Result:', isMatch);

    return isMatch;
  } catch (error) {
    console.error('Error in password comparison:', error);
    return false;
  }
};

module.exports = mongoose.model('User', userSchema);
