package services

import (
	"fmt"

	"errors"
	"github.com/shatvl/flatwindow/models"
	"github.com/shatvl/flatwindow/mongo"
	"github.com/shatvl/flatwindow/repositories"
	"gopkg.in/mgo.v2/bson"
)

// AdService for ad management
type AdService struct {
	Repo *repositories.AdRepository
}

// NewBidService returns UserService preference
func NewAdService() *AdService {
	repo := repositories.NewAdRepository()

	return &AdService{
		Repo: repo,
	}
}

// CreateAd creates or updates ad for t-s ads
func (s *AdService) CreateAd(ad *models.Ad, agentType byte) error {
	session := mongo.Session()
	defer session.Close()

	foundAd, err := s.Repo.FindByTypeAndUID(agentType, ad.Unid)

	if err == nil {
		fmt.Println("Ad is found: " + foundAd.Unid)
		return nil
	}

	err = s.Repo.CreateAd(ad)

	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("Ad has been added successfully: " + ad.Unid)

	return nil
}

// Add ad to feed for any platform
func (s *AdService) AddAdToFeed(_id string, feedType int, value bool) error {
	if !bson.IsObjectIdHex(_id) {
		return errors.New(`Invalid _id`)
	}

	return s.Repo.AddAdToFeed(bson.ObjectIdHex(_id), feedType, value)
}

func (s *AdService) GetAgentAdsForFeedByCode(code string) ([]*models.Ad, error) {
	ads, err := s.Repo.GetAdsForFeedByAgentCode(code)

	if err != nil {
		return nil, err
	}

	return ads, err
}
