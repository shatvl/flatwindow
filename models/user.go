package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
	"github.com/dgrijalva/jwt-go"
)

// User
type User struct {
	ID        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Email     string        `json:"email"`
	Password  string        `json:"password"`
	CreatedAt time.Time     `json:"created_at"`
	Role      string        `json:"role"`
	Active    bool          `json:"active"`
	AgentType byte 			`json:"agentType,omitempty" bson:"agent_type"`
	AgentCode string		`json:"agentCode,omitempty" bson:"agent_code"`
}

type UserJSON struct {
	ID 		   bson.ObjectId `json:"id"`
	Email 	   string 	     `json:"email"`
	CreatedAt  time.Time     `json:"created_at"`
	Role	   string 		 `json:"role"`
	Active	   bool			 `json:"active"`
}

type JwtClaims struct {
	Email     string `json:"email"`
	Role      string `json:"role"`
	AgentType byte   `json:"agent_type"`
	AgentCode string `json:"agent_code"`
	jwt.StandardClaims
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


func (u *User) ToUserJSON() (*UserJSON) {
	return &UserJSON{
		u.ID, u.Email, u.CreatedAt, u.Role, u.Active,
	}
}