const express = require('express');
const router = express.Router();
const { getUserStatsByUsername, updateUserStats } = require('../controllers/userStatsController');

router.get('/user/:username', getUserStatsByUsername)
      .post('/user/:username', updateUserStats);
module.exports = router; 