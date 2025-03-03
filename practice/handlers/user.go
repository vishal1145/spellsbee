package handlers

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"time"

	"spellsbee/config"
	"spellsbee/models"
	"spellsbee/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	client, err := config.ConnectToMongoDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer client.Disconnect(context.Background())

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user.ID = primitive.NewObjectID()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	collection := client.Database("spellsbee").Collection("users")
	_id, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Print("id", _id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func generateVerificationToken() (string, error) {
	token := make([]byte, 32)
	_, err := rand.Read(token)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(token), nil
}

func isValidUsername(username string) bool {
	validUsername := regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
	return validUsername.MatchString(username)
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	client, err := config.ConnectToMongoDB()
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}
	defer client.Disconnect(context.Background())

	// Parse request body
	var req struct {
		Username    string `json:"username"`
		NewUsername string `json:"newusername"`
		Email       string `json:"email"`
		Password    string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate required fields
	if req.Username == "" || req.NewUsername == "" || req.Email == "" || req.Password == "" {
		http.Error(w, "Please provide username, new username, email, and password", http.StatusBadRequest)
		return
	}

	// Validate new username format
	if !isValidUsername(req.NewUsername) {
		http.Error(w, "Username can only contain letters, numbers, and underscore", http.StatusBadRequest)
		return
	}

	collection := client.Database("spellsbee").Collection("users")

	// Check if user exists by username
	var user models.User
	err = collection.FindOne(context.Background(), bson.M{"username": req.Username}).Decode(&user)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Check if new username already exists (case insensitive)
	existingUsername := collection.FindOne(context.Background(), bson.M{
		"username": bson.M{"$regex": primitive.Regex{Pattern: "^" + req.NewUsername + "$", Options: "i"}},
	})
	if existingUsername.Err() == nil {
		http.Error(w, "Username already exists", http.StatusBadRequest)
		return
	}

	// Check if email already exists
	existingEmail := collection.FindOne(context.Background(), bson.M{"email": req.Email})
	if existingEmail.Err() == nil {
		http.Error(w, "Email already exists", http.StatusBadRequest)
		return
	}

	// Generate verification token
	verificationToken, err := generateVerificationToken()
	if err != nil {
		http.Error(w, "Error generating verification token", http.StatusInternalServerError)
		return
	}
	verificationTokenExpires := time.Now().Add(24 * time.Hour) // Token expires in 24 hours

	// Hash password before storing
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	// Update user with new details
	update := bson.M{
		"$set": bson.M{
			"username":                 req.NewUsername,
			"email":                    req.Email,
			"password":                 string(hashedPassword),
			"verificationToken":        verificationToken,
			"verificationTokenExpires": verificationTokenExpires,
			"updatedAt":                time.Now(),
		},
	}

	_, err = collection.UpdateOne(context.Background(), bson.M{"_id": user.ID}, update)
	if err != nil {
		http.Error(w, "Error updating user", http.StatusInternalServerError)
		return
	}

	verificationLink := fmt.Sprintf("%s/verify-email?token=%s", os.Getenv("BASE_URL"), verificationToken)
	subject := "Verify Your Email"
	message := fmt.Sprintf("Click here to verify your email: %s", verificationLink)

	err = utils.SendEmail(req.Email, subject, message)

	if err != nil {
		fmt.Println("Error sending email:", err)
	} else {
		fmt.Println("Email sent successfully!")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"_id":      user.ID,
		"email":    req.Email,
		"username": req.NewUsername,
	})
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	client, err := config.ConnectToMongoDB()
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}
	defer client.Disconnect(context.Background())

	var reqBody struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if reqBody.Email == "" || reqBody.Password == "" {
		http.Error(w, "Please provide email and password", http.StatusBadRequest)
		return
	}

	collection := client.Database("spellsbee").Collection("users")

	var user models.User
	err = collection.FindOne(context.Background(), bson.M{"email": reqBody.Email}).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(reqBody.Password))
	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	fmt.Print("user", user)

	if !user.EmailVerified {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message":         "Email not verified. Please verify your account before logging in.",
			"isEmailVerified": false,
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"username": user.Username,
		"email":    user.Email,
	})
}

func ForgotPassword(w http.ResponseWriter, r *http.Request) {
	client, err := config.ConnectToMongoDB()
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}
	defer client.Disconnect(context.Background())

	var reqBody struct {
		Email string `json:"email"`
	}

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if reqBody.Email == "" {
		http.Error(w, "Please provide an email", http.StatusBadRequest)
		return
	}

	collection := client.Database("spellsbee").Collection("users")

	var user models.User
	err = collection.FindOne(context.Background(), bson.M{"email": reqBody.Email}).Decode(&user)
	if err != nil {
		http.Error(w, "Email not found", http.StatusNotFound)
		return
	}

	resetToken, _ := generateToken()
	resetTokenExpires := time.Now().Add(1 * time.Hour) // 1-hour expiry

	update := bson.M{"$set": bson.M{
		"resetPasswordToken":   resetToken,
		"resetPasswordExpires": resetTokenExpires,
	}}

	_, err = collection.UpdateOne(context.Background(), bson.M{"email": reqBody.Email}, update)
	if err != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	// const resetLink = `${process.env.BASE_URL}/reset-password?token=${resetToken}`;
	// await sendEmail(
	//   email,
	//   'Password Reset Request',
	//   `Click here to reset your password: ${resetLink}\nThis link will expire in 1 hour.`
	// );

	resetLink := fmt.Sprintf("%s/reset-password?token=%s", os.Getenv("BASE_URL"), resetToken)
	err = utils.SendEmail(reqBody.Email, "Password Reset Request", "Click here to reset your password: "+resetLink)
	if err != nil {
		http.Error(w, "Failed to send email", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Password reset link sent to email",
	})
}

func ResetPassword(w http.ResponseWriter, r *http.Request) {
	client, err := config.ConnectToMongoDB()
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}
	defer client.Disconnect(context.Background())

	var reqBody struct {
		Token    string `json:"token"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if reqBody.Token == "" || reqBody.Password == "" {
		http.Error(w, "Please provide token and new password", http.StatusBadRequest)
		return
	}

	collection := client.Database("spellsbee").Collection("users")

	var user models.User
	err = collection.FindOne(context.Background(), bson.M{
		"resetPasswordToken": reqBody.Token,
		"resetPasswordExpires": bson.M{
			"$gt": time.Now(),
		},
	}).Decode(&user)

	if err != nil {
		http.Error(w, "Invalid or expired reset token", http.StatusBadRequest)
		return
	}

	// Hash new password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(reqBody.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	// Update user password & remove reset token
	update := bson.M{"$set": bson.M{
		"password":             string(hashedPassword),
		"resetPasswordToken":   "",
		"resetPasswordExpires": time.Time{},
	}}

	_, err = collection.UpdateOne(context.Background(), bson.M{"email": user.Email}, update)
	if err != nil {
		http.Error(w, "Failed to update password", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Password reset successful",
	})
}

func VerifyEmail(w http.ResponseWriter, r *http.Request) {
	client, err := config.ConnectToMongoDB()
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}
	defer client.Disconnect(context.Background())

	var reqBody struct {
		Token string `json:"token"`
	}

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if reqBody.Token == "" {
		http.Error(w, "Verification token is required", http.StatusBadRequest)
		return
	}

	collection := client.Database("spellsbee").Collection("users")

	var user models.User
	err = collection.FindOne(context.Background(), bson.M{
		"verificationToken": reqBody.Token,
		"verificationTokenExpires": bson.M{
			"$gt": time.Now(),
		},
	}).Decode(&user)

	if err != nil {
		http.Error(w, "Invalid or expired verification token", http.StatusBadRequest)
		return
	}

	update := bson.M{"$set": bson.M{
		"emailVerified":            true,
		"verificationToken":        "",
		"verificationTokenExpires": time.Time{},
	}}

	_, err = collection.UpdateOne(context.Background(), bson.M{"email": user.Email}, update)
	if err != nil {
		http.Error(w, "Failed to verify email", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":         "Email verified successfully",
		"isEmailVerified": true,
	})
}

func CheckUser(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	if username == "" {
		http.Error(w, "Username is required", http.StatusBadRequest)
		return
	}

	client, err := config.ConnectToMongoDB()
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}
	defer client.Disconnect(context.Background())

	collection := client.Database("spellsbee").Collection("users")

	var user models.User
	err = collection.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Error fetching user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func generateToken() (string, error) {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
