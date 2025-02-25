const mongoose = require('mongoose');

const userStatsSchema = new mongoose.Schema({
  user: {
    type: mongoose.Schema.Types.ObjectId,
    ref: 'User',
    required: true
  },
  points: {
    type: Number,
    default: 0
  },
  words: [{
    type: String
  }],
  date: {
    type: Date,
    default: Date.now
  }
}, {
  timestamps: true
});

module.exports = mongoose.model('UserStats', userStatsSchema); 