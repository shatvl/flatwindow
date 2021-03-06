package repositories

import (
	"github.com/shatvl/flatwindow/config"
	"github.com/shatvl/flatwindow/models"
	"github.com/shatvl/flatwindow/mongo"
	"gopkg.in/mgo.v2/bson"
)

// BidRepository with "bids" collection
type BidRepository struct {
	collName string
}

// NewUserRepository returns UserRepository preference to "users" repository
func NewBidRepository() *BidRepository {

	return &BidRepository{collName: "bids"}
}

func (r *BidRepository) CreateBid(bid *models.Bid) error {
	session := mongo.Session()
	defer session.Close()

	var ad models.Ad

	err := session.DB(config.Db).C("ads").FindId(bid.AdId).One(&ad)

	if err != nil {
		return err
	}

	bid.AdId = ad.ID
	bid.AgentType = ad.AgentType

	err = session.DB(config.Db).C(r.collName).Insert(bid)

	return err
}

func (r *BidRepository) GetPaginatedBids(filter *models.BidFilterRequest) ([]*models.Bid, int, error) {
	session := mongo.Session()
	defer session.Close()

	bids := make([]*models.Bid, 0)

	err := session.DB(config.Db).C(r.collName).Pipe([]bson.M{
		bson.M{
			"$skip": filter.Paginate.Page * filter.Paginate.PerPage},
		bson.M{
			"$limit": filter.Paginate.PerPage},
		bson.M{
			"$lookup": bson.M{
				"from":         "ads",
				"foreignField": "_id",
				"localField":   "ad_id",
				"as":           "ads"}},
		bson.M{"$unwind": "$ads"},
	}).
		All(&bids)

	count, _ := session.DB(config.Db).C(r.collName).Count()

	return bids, count, err
}

func (r *BidRepository) UpdateBid(bid *models.UpdatedBid) error {
	session := mongo.Session()
	defer session.Close()

	err := session.DB(config.Db).C(r.collName).UpdateId(bid.ID, bson.M{"$set": bson.M{
		"fullname":  bid.Fullname,
		"email":     bid.Email,
		"phone":     bid.Phone,
		"city":      bid.City,
		"message":   bid.Message,
		"processed": bid.Processed}})

	return err
}
