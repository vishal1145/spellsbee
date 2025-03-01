package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"	
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Define global collections (pass from main)
var UserCollection *mongo.Collection
var UserStatsCollection *mongo.Collection

// Get user stats by username
func GetUserStatsByUsername(c *gin.Context) {
	username := c.Param("username")

	// Find user
	var user struct {
		ID       primitive.ObjectID `bson:"_id,omitempty"`
		Username string             `bson:"username"`
	}
	err := UserCollection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	// Set today's date range
	today := time.Now().Truncate(24 * time.Hour)
	tomorrow := today.Add(24 * time.Hour)

	// Find user stats for today
	var userStats struct {
		ID        primitive.ObjectID `bson:"_id,omitempty"`
		UserID    primitive.ObjectID `bson:"user"`
		Points    int                `bson:"points"`
		Words     []string           `bson:"words"`
		CreatedAt time.Time          `bson:"createdAt"`
	}

	filter := bson.M{
		"user":      user.ID,
		"createdAt": bson.M{"$gte": today, "$lt": tomorrow},
	}

	err = UserStatsCollection.FindOne(context.TODO(), filter).Decode(&userStats)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"userStats": gin.H{"points": 0, "words": []string{}}})
		return
	}

	c.JSON(http.StatusOK, userStats)
}

// Update user stats
func UpdateUserStats(c *gin.Context) {
	username := c.Param("username")
	var input struct {
		Points int    `json:"points"`
		Word   string `json:"word"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}

	// Find user
	var user struct {
		ID       primitive.ObjectID `bson:"_id,omitempty"`
		Username string             `bson:"username"`
	}
	err := UserCollection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	// Set today's date range
	today := time.Now().Truncate(24 * time.Hour)
	tomorrow := today.Add(24 * time.Hour)

	// Find user stats for today
	var userStats struct {
		ID        primitive.ObjectID `bson:"_id,omitempty"`
		UserID    primitive.ObjectID `bson:"user"`
		Points    int                `bson:"points"`
		Words     []string           `bson:"words"`
		CreatedAt time.Time          `bson:"createdAt"`
	}

	filter := bson.M{
		"user":      user.ID,
		"createdAt": bson.M{"$gte": today, "$lt": tomorrow},
	}

	err = UserStatsCollection.FindOne(context.TODO(), filter).Decode(&userStats)
	if err != nil { // Create new if not found
		userStats = struct {
			ID        primitive.ObjectID `bson:"_id,omitempty"`
			UserID    primitive.ObjectID `bson:"user"`
			Points    int                `bson:"points"`
			Words     []string           `bson:"words"`
			CreatedAt time.Time          `bson:"createdAt"`
		}{
			UserID:    user.ID,
			Points:    0,
			Words:     []string{},
			CreatedAt: time.Now(),
		}
	}

	userStats.Points += input.Points
	userStats.Words = append(userStats.Words, input.Word)

	opts := options.Replace().SetUpsert(true)
	_, err = UserStatsCollection.ReplaceOne(context.TODO(), filter, userStats, opts)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update stats"})
		return
	}

	c.JSON(http.StatusOK, userStats)
}
