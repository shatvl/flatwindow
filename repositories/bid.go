package repositories

import (
	"github.com/shatvl/flatwindow/config"
	"github.com/shatvl/flatwindow/mongo"
	"github.com/shatvl/flatwindow/models"
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