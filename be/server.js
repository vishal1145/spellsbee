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

// Store the current daily letters
let currentDailyLetters = generateDailyString();

// Function to generate random 7-character string from A-Z
function generateDailyString() {
  const letters = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ';
  let result = '';
  const usedLetters = new Set();
  
  // Generate unique random letters
  while (usedLetters.size < 7) {
    const letter = letters.charAt(Math.floor(Math.random() * letters.length));
    usedLetters.add(letter);
  }
  
  return Array.from(usedLetters).join('');
}

// Update letters at midnight
cron.schedule('0 0 * * *', () => {
  currentDailyLetters = generateDailyString();
  console.log('New daily letters generated:', currentDailyLetters);
});

// Add new endpoint to get daily letters
app.get('/api/daily-letters', (req, res) => {
  // res.json({ 
  //   letters: currentDailyLetters,
  //   centerLetter: currentDailyLetters[Math.floor(Math.random() * 7)] // Randomly select center letter
  // });
  res.json({ 
    letters: currentDailyLetters,
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