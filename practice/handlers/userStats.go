package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"spellsbee/config"
	"spellsbee/models"
)

func GetUserStatsByUsername(w http.ResponseWriter, r *http.Request) {
	client, err := config.ConnectToMongoDB()
	if err != nil {
		http.Error(w, `{"error": "Database connection failed"}`, http.StatusInternalServerError)
		return
	}
	defer client.Disconnect(context.Background())

	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 5 { // Expected format: /api/stats/user/{username}
		http.Error(w, `{"error": "Invalid URL format"}`, http.StatusBadRequest)
		return
	}
	username := parts[4]

	userCollection := client.Database("spellsbee").Collection("users")
	var user models.User
	err = userCollection.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		http.Error(w, `{"error": "User not found"}`, http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, `{"error": "Database error"}`, http.StatusInternalServerError)
		return
	}

	now := time.Now()
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	endOfDay := startOfDay.Add(24 * time.Hour)

	userStatsCollection := client.Database("spellsbee").Collection("userstats")
	var userStats models.UserStats
	filter := bson.M{
		"user": user.ID,
		"createdAt": bson.M{
			"$gte": startOfDay,
			"$lt":  endOfDay,
		},
	}
	err = userStatsCollection.FindOne(context.Background(), filter).Decode(&userStats)

	if err == mongo.ErrNoDocuments {
		userStats = models.UserStats{
			User:   user.ID,
			Points: 0,
			Words:  []string{},
		}
	} else if err != nil {
		http.Error(w, `{"error": "Database error"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userStats)
}

func UpdateUserStats(w http.ResponseWriter, r *http.Request) {
	fmt.Print("djjdbd")

	client, err := config.ConnectToMongoDB()
	if err != nil {
		http.Error(w, `{"error": "Database connection failed"}`, http.StatusInternalServerError)
		return
	}
	defer client.Disconnect(context.Background())

	// Extract username from URL path
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 5 { // Expected: /api/stats/user/{username}
		http.Error(w, `{"error": "Invalid URL format"}`, http.StatusBadRequest)
		return
	}
	username := parts[4]

	// Decode request body
	var requestBody struct {
		Points int    `json:"points"`
		Word   string `json:"word"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, `{"error": "Invalid request body"}`, http.StatusBadRequest)
		return
	}

	// Get user details
	userCollection := client.Database("spellsbee").Collection("users")
	var user models.User
	err = userCollection.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		http.Error(w, `{"error": "User not found"}`, http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, `{"error": "Database error"}`, http.StatusInternalServerError)
		return
	}

	// Convert user ID to ObjectID
	userID, err := primitive.ObjectIDFromHex(user.ID.Hex())
	if err != nil {
		http.Error(w, `{"error": "Invalid user ID"}`, http.StatusInternalServerError)
		return
	}

	// Ensure UTC consistency in date filtering
	now := time.Now().UTC()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	tomorrow := today.Add(24 * time.Hour)

	// Log timestamps to debug
	fmt.Println("Today (UTC):", today)
	fmt.Println("Tomorrow (UTC):", tomorrow)

	// User stats collection
	collection := client.Database("spellsbee").Collection("userstats")

	// Find existing stats or create a new one
	filter := bson.M{
		"user": userID,
		"date": bson.M{"$gte": today, "$lt": tomorrow},
	}

	update := bson.M{
		"$inc":  bson.M{"points": requestBody.Points}, // Increment points
		"$push": bson.M{"words": requestBody.Word},    // Append word to array
		"$set":  bson.M{"updatedAt": now},             // Update timestamp
		"$setOnInsert": bson.M{
			"date":      today,  // Ensure date is set on first insert
			"createdAt": now,    // Set createdAt only when inserting a new document
			"user":      userID, // Ensure user ID is set for new document
		},
	}

	fmt.Print("djjdbd")

	opts := options.Update().SetUpsert(true)

	// Perform update
	result, err := collection.UpdateOne(context.Background(), filter, update, opts)
	if err != nil {
		http.Error(w, `{"error": "Failed to update user stats"}`, http.StatusInternalServerError)
		return
	}

	// Debugging logs
	fmt.Println("Matched Count:", result.MatchedCount)
	fmt.Println("Modified Count:", result.ModifiedCount)
	fmt.Println("Upserted Count:", result.UpsertedCount)

	// Fetch updated document
	var updatedStats models.UserStats
	err = collection.FindOne(context.Background(), filter).Decode(&updatedStats)
	if err == mongo.ErrNoDocuments {
		http.Error(w, `{"error": "Failed to fetch updated stats"}`, http.StatusInternalServerError)
		return
	} else if err != nil {
		http.Error(w, `{"error": "Database error"}`, http.StatusInternalServerError)
		return
	}

	// Ensure correct response structure
	response := map[string]interface{}{
		"_id":       updatedStats.ID.Hex(),
		"user":      bson.M{"_id": updatedStats.User.Hex(), "username": user.Username},
		"points":    updatedStats.Points,
		"words":     updatedStats.Words,
		"date":      updatedStats.Date,
		"createdAt": updatedStats.CreatedAt,
		"updatedAt": updatedStats.UpdatedAt,
		"__v":       0,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
