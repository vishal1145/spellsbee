package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID                     primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username               string             `bson:"username" json:"username"`
	Email                  string             `bson:"email" json:"email"`
	Password               string             `bson:"password" json:"-"`
	EmailVerified         bool               `bson:"email_verified" json:"email_verified"`
	VerificationToken     string             `bson:"verification_token,omitempty" json:"-"`
	VerificationTokenExpiry time.Time        `bson:"verification_token_expiry,omitempty" json:"-"`
	ResetPasswordToken    string             `bson:"reset_password_token,omitempty" json:"-"`
	ResetPasswordExpiry   time.Time          `bson:"reset_password_expiry,omitempty" json:"-"`
}

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
} 