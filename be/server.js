const express = require('express');
const cors = require('cors');
const dotenv = require('dotenv');
const connectDB = require('./config/db');
const userRoutes = require('./routes/userRoutes');
const spellsbeeRoutes = require('./routes/spellsbee');
const userStatsRoutes = require('./routes/userStatsRoutes');
dotenv.config();

const app = express();

// Middleware
app.use(cors());
app.use(express.json());
app.use(express.urlencoded({ extended: true }));

// Connect to database
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