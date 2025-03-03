package models

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID                       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username                 string             `bson:"username" json:"username"`
	Email                    string             `bson:"email" json:"email"`
	Password                 string             `bson:"password,omitempty" json:"-"`
	EmailVerified            bool               `bson:"emailVerified" json:"emailVerified"`
	VerificationToken        string             `bson:"verificationToken,omitempty" json:"-"`
	VerificationTokenExpires time.Time          `bson:"verificationTokenExpires,omitempty" json:"-"`
	CreatedAt                time.Time          `bson:"createdAt,omitempty" json:"createdAt"`
	UpdatedAt                time.Time          `bson:"updatedAt,omitempty" json:"updatedAt"`
}

func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func (u *User) Validate() error {
	if len(u.Username) < 3 || len(u.Username) > 30 {
		return errors.New("username must be between 3 and 30 characters")
	}
	if len(u.Email) == 0 {
		return errors.New("email is required")
	}
	if len(u.Password) < 6 {
		return errors.New("password must be at least 6 characters long")
	}
	return nil
}
