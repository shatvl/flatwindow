package repositories

import (
	"github.com/shatvl/flatwindow/config"
	"github.com/shatvl/flatwindow/mongo"
	"github.com/shatvl/flatwindow/models/ads"
	"gopkg.in/mgo.v2/bson"
)

const (
	TsType = 101
)

// AdRepository with "ads" collection
type AdRepository struct {
	collName string
}

// NewUserRepository returns UserRepository preference to "users" repository
func NewAdRepository() *AdRepository {

	return &AdRepository{collName: "ads"}
}

// Create user by json body
func (r *AdRepository) CreateTsAd(ad *models.TsAd) (error) {
	session := mongo.Session()
	defer session.Close()

	ad.AgentType = TsType

	err := session.DB(config.Db).C(r.collName).Insert(&ad)

	return err
}

// FindByTypeAndUID find ad by Type and Unique id
func (r *AdRepository) FindByTypeAndUID(t byte, uid string) (*models.TsAd, error) {
	session := mongo.Session()
	defer session.Close()

	switch t {
		case TsType:
			ad := models.TsAd{}
			err := session.DB(config.Db).C(r.collName).Find(bson.M{"agentType": TsType, "unid": uid}).One(&ad)
			
			if err != nil {
				return nil, err
			}
			
			return &ad, nil
	}

	return nil, nil
}