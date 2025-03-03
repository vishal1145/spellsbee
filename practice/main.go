package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"spellsbee/handlers"
	"sync"
	"time"

	"github.com/robfig/cron/v3"
)

var (
	currentCode []string
	mu          sync.Mutex
)

// Generates a unique array of 7 letters
func generateUniqueCode() []string {
	const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterSet := make(map[byte]bool)
	var uniqueLetters []string

	for len(uniqueLetters) < 7 {
		letter := letters[rand.Intn(len(letters))]
		if !letterSet[letter] {
			letterSet[letter] = true
			uniqueLetters = append(uniqueLetters, string(letter))
		}
	}
	return uniqueLetters
}

func updateCode() {
	mu.Lock()
	defer mu.Unlock()
	currentCode = generateUniqueCode()
}

func getUniqueCode(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(currentCode)
}

// ✅ CORS Middleware to fix CORS issue
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins (Change to frontend URL in production)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle OPTIONS preflight request
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	rand.Seed(time.Now().UnixNano())
	updateCode()

	c := cron.New()
	_, err := c.AddFunc("0 0 * * *", updateCode) // Runs at 00:00 UTC
	if err != nil {
		log.Fatal("Failed to schedule cron job:", err)
	}
	c.Start()

	mux := http.NewServeMux()
	mux.HandleFunc("/api/daily-letters", getUniqueCode)
	mux.HandleFunc("/api/users", handlers.CreateUser)
	mux.HandleFunc("/api/users/register", handlers.RegisterUser)
	mux.HandleFunc("/api/users/login", handlers.LoginUser)
	mux.HandleFunc("/api/users/forgot-password", handlers.ForgotPassword)
	mux.HandleFunc("/api/users/reset-password", handlers.ResetPassword)
	mux.HandleFunc("/api/users/verify-email", handlers.VerifyEmail)
	mux.HandleFunc("/api/checkuser", handlers.CheckUser)
	mux.HandleFunc("/api/spellsbee/validate/", handlers.ValidateWord)
	mux.HandleFunc("/api/stats/user/", handlers.GetUserStatsByUsername)
	mux.HandleFunc("/api/stats/update-stats/", handlers.UpdateUserStats)

	// Wrap with CORS Middleware
	server := corsMiddleware(mux)

	fmt.Println("🚀 Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", server))
}
