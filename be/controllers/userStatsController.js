const UserStats = require('../models/userStatsModel');
const User = require('../models/userModel'); // Make sure you have this model

const getUserStatsByUsername = async (req, res) => {
  try {
    const { username } = req.params;
    
    const user = await User.findOne({ username });
    if (!user) {
      return res.status(404).json({ message: 'User not found' });
    }

    const today = new Date();
    today.setHours(0, 0, 0, 0);
    const tomorrow = new Date(today);
    tomorrow.setDate(tomorrow.getDate() + 1);

    const userStats = await UserStats.findOne({
      user: user._id,
      createdAt: {
        $gte: today,
        $lt: tomorrow
      }
    }).populate('user', 'username');

    if (!userStats) {
      return res.status(404).json({ userStats: { points: 0, words:  [] } });
    }

    res.status(200).json(userStats);
  } catch (error) {
    res.status(500).json({ message: error.message });
  }
};

const updateUserStats = async (req, res) => {
    try {
      const { username } = req.params;
      const { points, word } = req.body;
      console.log(points);
      console.log(word);
      const user = await User.findOne({ username });
      if (!user) {
        return res.status(404).json({ message: 'User not found' });
      }
  
      const today = new Date();
      today.setHours(0, 0, 0, 0);
      const tomorrow = new Date(today);
      tomorrow.setDate(tomorrow.getDate() + 1);
      
      let userStats = await UserStats.findOne({
        user: user._id,
        createdAt: {
          $gte: today,
          $lt: tomorrow
        }
      });

      if (!userStats) {
        userStats = new UserStats({
          user: user._id,
          points: 0,
          words: []
        });
      }
  

      userStats.points += points;
      userStats.words.push(word);

      await userStats.save();
  
      res.status(200).json(userStats);
    } catch (error) {
      res.status(500).json({ message: error.message });
    }
};
  
  module.exports = {
    getUserStatsByUsername,
    updateUserStats
  }; 
 