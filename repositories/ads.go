package repositories

import (
	"errors"

	"github.com/shatvl/flatwindow/config"
	"github.com/shatvl/flatwindow/mongo"
	"github.com/shatvl/flatwindow/models"
	"gopkg.in/mgo.v2/bson"
)

const (
	TsType = 101
	KnType = 102
	RtType = 103
	OnType = 104
	KfType = 105
)

var FeedTypeToName = map[int]string{
	101: "ts",
	102: "kn",
	103: "rt",
	104: "on",
	105: "kf",
}

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
			err := session.DB(config.Db).C(r.collName).Find(bson.M{"agentType": t, "unid": uid}).One(&ad)
			
			if err != nil {
				return nil, err
			}
			
			return &ad, nil
	}

	return nil, nil
}

func (r *AdRepository) GetAdsWithFilter(filter *models.AdFilterRequest) ([]*models.Ad, int, error) {
	session := mongo.Session()
	defer session.Close()

	//q := minquery.New(session.DB(config.Db), r.collName, bson.M{"rooms": rooms}).Sort("_id").Limit(5)
	ads := make([]*models.Ad,0)

	query := getFilterQuery(&filter.Filter)

	session.DB(config.Db).C(r.collName).Find(&query).Limit(filter.Paginate.PerPage).Skip(int(filter.Paginate.Page * filter.Paginate.PerPage)).All(&ads)
	count, _ := session.DB(config.Db).C(r.collName).Find(&query).Count()

	return ads, count, nil
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

// AddAdToFeed method add ad to feed for the any playform type (onliner, kufer and etc.)
func (r *AdRepository) AddAdToFeed(_id bson.ObjectId, feedType int, value bool) error {
	if FeedTypeToName[feedType] == "" {
		return errors.New("unexpected feed type")
	}

	session := mongo.Session()
	defer session.Close()

	err := session.DB(config.Db).C(r.collName).Update(bson.M{"_id": _id}, bson.M{"$set": bson.M{"rss." + FeedTypeToName[feedType]: value}})

	return err
}

func (r *AdRepository) GetAdsForFeedByAgentCode(code string) ([]*models.Ad, error) {
	session := mongo.Session()
	defer session.Close()

	ads := make([]*models.Ad,0)

	err := session.DB(config.Db).C(r.collName).Find(bson.M{"rss." + code: true}).All(&ads)

	return ads, err
}

func getFilterQuery(filter *models.AdFilter) (*bson.M){
	query := bson.M{}

	if filter.AgentType != 0 {
		query["agentType"] = bson.M{"$eq": filter.AgentType}
	}
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