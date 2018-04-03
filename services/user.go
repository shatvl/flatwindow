package services

import (
	"errors"
	"time"

	"flatwindow/config"
	"flatwindow/models"
	"flatwindow/repositories"

	jwt "github.com/dgrijalva/jwt-go"
)

type UserService struct {
	repo *repositories.UserRepository
}

// NewUserService returns UserService preference
func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

// Create user method
func (s *UserService) Create(user models.User) (models.UserResource, error) {
	if user.ID != "" || string(user.Password) == "" || user.Email == "" {
		return models.UserResource{}, errors.New("Unable to create this user")
	}

	return s.repo.Create(user)
}

type jwtClaims struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

// Generate token for user
func (s *UserService) GenerateToken(credentials models.Credentials) (string, error) {
	foundUser, err := s.repo.FindByEmail(credentials.Email)

	if err != nil {
		return "", errors.New(`Invalid email or password`)
	}

	valid, err := models.ValidatePassword(credentials.Password, []byte(foundUser.Password))

	if err != nil || !valid {
		return "", errors.New(`Invalid email or password`)
	}

	claims := &jwtClaims{
		foundUser.Email,
		foundUser.Role,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// Generate token with Claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token
	t, err := token.SignedString([]byte(config.Secret))

	if err != nil {
		return "", err
	}

	return t, err
}
