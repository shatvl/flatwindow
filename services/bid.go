package services

import (
	"github.com/shatvl/flatwindow/repositories"
	"github.com/shatvl/flatwindow/models"
)

// BidService for bid management
type BidService struct {
	Repo *repositories.BidRepository
}

// NewBidService returns UserService preference
func NewBidService() *BidService {
	repo := repositories.NewBidRepository()

	return &BidService{
		Repo: repo,
	}
}

// CreateAd creates or updates ad for t-s ads
func (bs *BidService) CreateBid(bid *models.Bid) (error) {
	return bs.Repo.CreateBid(bid)
}

func (bs *BidService) GetPaginatedBids(filter *models.AdFilterRequest) ([]*models.Bid, int, error) {
	return bs.Repo.GetPaginatedBids(filter)
}