package repositories

import (
	"errors"
	
	"github.com/shatvl/flatwindow/config"
	"github.com/shatvl/flatwindow/models"
	"github.com/shatvl/flatwindow/mongo"

	"gopkg.in/mgo.v2/bson"
)

// UserRepository with "users" collection
type UserRepository struct {
	collName string
}

// NewUserRepository returns UserRepository preference to "users" repository
func NewUserRepository() *UserRepository {

	return &UserRepository{collName: "users"}
}

// Create user by json body
func (r *UserRepository) Create(user *models.User) (*models.User, error) {
	session := mongo.Session()
	defer session.Close()
	
	_, err := r.FindByEmail(user.Email)

	if err == nil || string(user.Password) == "" || user.Email == "" {
		return nil, errors.New("User already exists")
	}

	passsword, err := models.GeneratePassword(user.Password)

	if err != nil {
		return nil, err
	}

	user.ID = bson.NewObjectId()
	user.Password = string(passsword)
	user.Role = config.ROLE_USER
	user.Active = true //ACTIVATED USER BY DEFAULT

	err = session.DB(config.Db).C(r.collName).Insert(&user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

// FindByEmail finds user by email
func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	session := mongo.Session()
	defer session.Close()

	user := models.User{}
	err := session.DB(config.Db).C(r.collName).Find(bson.M{"email": email}).One(&user)

	return &user, err
}
