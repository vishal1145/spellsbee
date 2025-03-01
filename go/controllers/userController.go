package controllers

import (
	"context"
	"encoding/hex"
	"log"
	"math/rand"
	"net/http"
	"spellsbee/database"
	"spellsbee/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GenerateRandomToken() string {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		log.Fatal(err)
	}
	return hex.EncodeToString(bytes)
}

func CreateUser(c *gin.Context) {
	var input models.RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	user := models.User{
		ID:            primitive.NewObjectID(),
		Username:      input.Username,
		Email:         input.Email,
		Password:      input.Password,
		EmailVerified: false,
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    user,
	})
}

func RegisterUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true})
}

func LoginUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true})
}

func VerifyEmail(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true})
}

func ForgotPassword(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true})
}

func ResetPassword(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true})
}

func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true})
}

func GetUserByID(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true})
}

func CheckUser(c *gin.Context) {
	userID := c.Param("id")

	collection := database.GetCollection("users")
	var user bson.M
	err := collection.FindOne(context.TODO(), bson.M{"_id": userID}).Decode(&user)
	if err != nil {
		log.Println("Error fetching user:", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}
