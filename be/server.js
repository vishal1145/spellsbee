const express = require('express');
const cors = require('cors');
const dotenv = require('dotenv');
const connectDB = require('./config/db');
const userRoutes = require('./routes/userRoutes');
const spellsbeeRoutes = require('./routes/spellsbee');
const userStatsRoutes = require('./routes/userStatsRoutes');
const cron = require('node-cron');
dotenv.config();

const app = express();

// Move CORS middleware to the top, before any routes
app.use(cors());
app.use(express.json());
app.use(express.urlencoded({ extended: true }));

// Store both the current daily letters and center letter
let currentDailyLetters = generateDailyString();
let currentCenterLetter = currentDailyLetters[1];

function generateDailyString() {
  const letters = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ';
  let result = '';
  const usedLetters = new Set();
  
  while (usedLetters.size < 7) {
    const letter = letters.charAt(Math.floor(Math.random() * letters.length));
    usedLetters.add(letter);
  }
  
  return Array.from(usedLetters).join('');
}

// Update letters at midnight
cron.schedule('0 0 * * *', () => {
  currentDailyLetters = generateDailyString();
  currentCenterLetter = currentDailyLetters[Math.floor(Math.random() * 7)];
  console.log('New daily letters generated:', currentDailyLetters);
  console.log('New center letter:', currentCenterLetter);
});

// Add new endpoint to get daily letters
app.get('/api/daily-letters', (req, res) => {
  res.json({ 
    letters: currentDailyLetters,
    centerLetter: currentCenterLetter  
  });
});

connectDB();

// Routes
app.use('/api/users', userRoutes);
app.use('/api/spellsbee', spellsbeeRoutes);
app.use('/api/stats', userStatsRoutes);
// Basic route
app.get('/', (req, res) => {
  res.send('API is running...');
});

const PORT = process.env.PORT || 5000;
app.listen(PORT, () => {
  console.log(`Server is running on port ${PORT}`);
}); 