const mongoose = require('mongoose');
const bcrypt = require('bcrypt');

const userSchema = new mongoose.Schema({
  username: { 
    type: String, 
    required: true,
    validate: {
      validator: function(v) {
        // Only allow letters, numbers, and underscore
        return /^[a-zA-Z0-9_]+$/.test(v);
      },
      message: props => `${props.value} contains invalid characters. Only letters, numbers, and underscore are allowed.`
    }
  },
  email: { type: String, required: true, unique: true },
  password: { type: String, required: true },
  emailVerified: { type: Boolean, default: false },
  verificationToken: { type: String },  
  verificationTokenExpires: { type: Date },
  resetPasswordToken: String,
  resetPasswordExpires: Date
}, { timestamps: true });

// Add a case-insensitive unique index for username
userSchema.index({ username: 1 }, { 
  unique: true, 
  collation: { locale: 'en', strength: 2 } 
});

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
