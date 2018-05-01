package repositories

import (
	"github.com/shatvl/flatwindow/config"
	"github.com/shatvl/flatwindow/mongo"
	"github.com/shatvl/flatwindow/models"
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

func (r *BidRepository) CreateBid(bid *models.Bid) (error){
	session := mongo.Session()
	defer session.Close()

	err := session.DB(config.Db).C(r.collName).Insert(bid)

	return err
}

func (r *BidRepository) GetPaginatedBids(filter *models.AdFilterRequest) ([]*models.Bid, int, error) {
	session := mongo.Session()
	defer session.Close()

	var bids []*models.Bid

	err := session.DB(config.Db).C(r.collName).Pipe([]bson.M{
		                                         bson.M {
												 	"$skip": filter.Paginate.Page * filter.Paginate.PerPage},
												 bson.M {
												 	"$limit": filter.Paginate.PerPage },
												 bson.M {
												 	"$lookup": bson.M {
														"from": "ads",
														"foreignField": "_id",
														"localField": "ad_id",
														"as": "ads",}},
												 bson.M { "$unwind": "$ads" },
											   }).
		                                       All(&bids)

	count, _ := session.DB(config.Db).C(r.collName).Count()

	return bids, count, err
}