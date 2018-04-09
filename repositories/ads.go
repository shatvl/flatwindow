package repositories

import (
	"errors"

	"github.com/shatvl/flatwindow/config"
	"github.com/shatvl/flatwindow/mongo"
	"github.com/shatvl/flatwindow/models/ads"
	"gopkg.in/mgo.v2/bson"
)

const (
	TsType = 101
	AdsPerPage = 9
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
func (r *AdRepository) CreateAd(ad *models.Ad) (error) {
	session := mongo.Session()
	defer session.Close()

	ad.AgentType = TsType

	err := session.DB(config.Db).C(r.collName).Insert(&ad)

	return err
}

// FindByTypeAndUID find ad by Type and Unique id
func (r *AdRepository) FindByTypeAndUID(t byte, uid string) (*models.Ad, error) {
	session := mongo.Session()
	defer session.Close()

	switch t {
		case TsType:
			ad := models.Ad{}
			err := session.DB(config.Db).C(r.collName).Find(bson.M{"agentType": TsType, "unid": uid}).One(&ad)
			
			if err != nil {
				return nil, err
			}
			
			return &ad, nil
	}

	return nil, nil
}

func (r *AdRepository) GetAdsWithFilter(filter *models.AdFilterRequest) ([]*models.Ad, error) {
	session := mongo.Session()
	defer session.Close()

	//q := minquery.New(session.DB(config.Db), r.collName, bson.M{"rooms": rooms}).Sort("_id").Limit(5)
	var ads []*models.Ad

	query := getFilterQuery(&filter.Filter)

	session.DB(config.Db).C(r.collName).Find(&query).Limit(AdsPerPage).Skip(int(filter.Page * AdsPerPage)).All(&ads)

	return ads, nil
}

func (r *AdRepository) GetAdById(_id string) (*models.Ad, error) {
	if !bson.IsObjectIdHex(_id) {
		return nil, errors.New(`Invalid _id`)
	}
	
	session := mongo.Session()
	defer session.Close()

	var ad models.Ad

	err := session.DB(config.Db).C(r.collName).FindId(bson.ObjectIdHex(_id)).One(&ad)

	return &ad, err
}

func getFilterQuery(filter *models.AdFilter) (*bson.M){
	query := bson.M{}

	if filter.Rooms != 0 {
		query["rooms"] = bson.M{"$gt": filter.Rooms}	
	}
	if filter.MinPrice != 0 && filter.MaxPrice != 0 {
		query["price"] = bson.M{"$gte": filter.MinPrice, "$lte": filter.MaxPrice}
	} else if filter.MinPrice != 0 {
		query["price"] = bson.M{"$gte": filter.MinPrice}
	} else if filter.MaxPrice != 0 {
		query["price"] = bson.M{"$lte": filter.MaxPrice}
	}
	if filter.Text != "" {
		query["$text"] = bson.M{"$search": filter.Text}
	}

	return &query
}