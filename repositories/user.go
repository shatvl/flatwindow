package repositories

import (
	"flatwindow/config"
	"flatwindow/models"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// NewUserRepository returns UserRepository preference to "users" repository
func NewUserRepository(session *mgo.Session) *UserRepository {
	return &UserRepository{
		coll: session.DB("").C("users"),
	}
}

// UserRepository with "users" collection
type UserRepository struct {
	coll *mgo.Collection
}

// Create user by json body
func (r *UserRepository) Create(user models.User) (models.UserResource, error) {
	passsword, err := models.GeneratePassword(user.Password)

	if err != nil {
		return models.UserResource{}, err
	}

	user.ID = bson.NewObjectId()
	user.Password = string(passsword)
	user.Role = config.ROLE_USER
	user.Active = true //ACTIVATED USER BY DEFAULT

	err = r.coll.Insert(&user)

	if err != nil {
		return models.UserResource{}, err
	}

	result := models.UserResource{Data: user}

	return result, nil
}

// FindByEmail finds user by email
func (r *UserRepository) FindByEmail(email string) (models.User, error) {
	user := models.User{}
	err := r.coll.Find(bson.M{"email": email}).One(&user)

	return user, err
}
