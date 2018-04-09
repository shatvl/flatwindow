package services

import (
	"errors"
	"time"

	"github.com/shatvl/flatwindow/config"
	"github.com/shatvl/flatwindow/models"
	"github.com/shatvl/flatwindow/repositories"

	jwt "github.com/dgrijalva/jwt-go"
)

// UserService for user management
type UserService struct {
	Repo *repositories.UserRepository
}

// NewUserService returns UserService preference
func NewUserService() *UserService {
	repo := repositories.NewUserRepository()

	return &UserService{
		Repo: repo,
	}
}

type jwtClaims struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

// GenerateToken generates token for user
func (s *UserService) GenerateToken(credentials *models.Credentials) (string, *models.User, error) {
	foundUser, err := s.Repo.FindByEmail(credentials.Email)

	if err != nil {
		return "", nil, errors.New(`Invalid email or password`)
	}

	valid, err := models.ValidatePassword(credentials.Password, []byte(foundUser.Password))

	if err != nil || !valid {
		return "", nil, errors.New(`Invalid email or password`)
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
		return "", nil, err
	}

	return t, foundUser, err
}


func (s *UserService) CreateUser(u *models.User) (*models.User, error) {
	return s.Repo.Create(u)
}