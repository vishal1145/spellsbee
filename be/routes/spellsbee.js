const express = require('express');
const router = express.Router();
const axios = require('axios');

// Route to validate a word
router.get('/validate/:word', async (req, res) => {
    try {
        const word = req.params.word;
        const response = await axios.get(`https://api.datamuse.com/words?sp=${word}`);
        
        const isValid = response.data.some(item => item.word.toLowerCase() === word.toLowerCase());
        res.json({
            word: word,
            isValid: isValid
        });
    } catch (error) {
        res.status(500).json({
            error: 'Error validating word',
            message: error.message
        });
    }
});

module.exports = router; 