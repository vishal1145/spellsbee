package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"spellsbee/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

// ConnectDB initializes the MongoDB connection
func ConnectDB() {
	config.LoadEnv()
	mongoURI := config.GetMongoURI()

	if mongoURI == "" {
		log.Fatal("❌ MONGO_URI not found in environment variables")
	}

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal("❌ Error connecting to MongoDB:", err)
	}

	// Set timeout and verify connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("❌ Could not ping MongoDB:", err)
	}

	fmt.Println("✅ Connected to MongoDB!")

	// Assign global DB variable
	DB = client.Database("spellsbee_db")
}

// GetCollection returns a collection from the database
func GetCollection(collectionName string) *mongo.Collection {
	if DB == nil {
		log.Fatal("❌ Database connection is not initialized. Call ConnectDB() first.")
	}
	fmt.Println("📂 Getting collection:", collectionName)
	return DB.Collection(collectionName)
}
