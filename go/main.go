package main

import (
	"log"
	"math/rand"
	"time"

	"spellsbee/config"
	"spellsbee/database"
	"spellsbee/routes"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
)

var (
	currentDailyLetters []string
)

// Generates a unique 7-letter daily string
func generateDailyString() []string {
	letters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	usedLetters := make(map[string]bool)
	result := make([]string, 0, 7)

	for len(result) < 7 {
		randomIndex := rand.Intn(len(letters))
		letter := string(letters[randomIndex])
		if !usedLetters[letter] {
			usedLetters[letter] = true
			result = append(result, letter)
		}
	}
	return result
}

// Updates the global daily letters
func updateDailyLetters() {
	currentDailyLetters = generateDailyString()
	log.Printf("New daily letters generated: %v", currentDailyLetters)
}

// CORS setup to allow cross-origin requests
func setupCORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Accept, Origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Connect to MongoDB
	database.ConnectDB()

	// Generate first daily letters
	updateDailyLetters()

	// Setup cron job to update daily letters at midnight
	c := cron.New()
	if _, err := c.AddFunc("0 0 * * *", updateDailyLetters); err != nil {
		log.Fatal("Failed to setup cron job:", err)
	}
	c.Start()
	defer c.Stop()

	// Setup Gin router
	router := gin.Default()
	router.Use(setupCORS())

	// API route for fetching daily letters
	router.GET("/api/daily-letters", func(c *gin.Context) {
		if currentDailyLetters == nil || len(currentDailyLetters) != 7 {
			c.JSON(500, gin.H{"error": "Daily letters not properly initialized"})
			return
		}
		c.JSON(200, gin.H{"letters": currentDailyLetters})
	})

	// Setup other routes
	routes.SetupUserRouter(router)
	routes.SetupUserStatsRoutes(router)

	// Start server
	if err := router.Run(config.GetPort()); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
