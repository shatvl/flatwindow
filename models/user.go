package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

// User
type User struct {
	ID        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Email     string        `json:"email"`
	Password  string        `json:"password"`
	CreatedAt time.Time     `json:"created_at"`
	Role      string        `json:"role"`
	Active    bool          `json:"active"`
}

// UserResource
type UserResource struct {
	Data *User `json:"data"`
}

// GeneratePassword will generate a hashed password by users input password
func GeneratePassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func ValidatePassword(password string, hashed []byte) (bool, error) {
	if err := bcrypt.CompareHashAndPassword(hashed, []byte(password)); err != nil {
		return false, err
	}

	return true, nil
}
