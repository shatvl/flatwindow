package services

import (
	"fmt"

	"github.com/shatvl/flatwindow/mongo"
	"github.com/shatvl/flatwindow/models/ads"
	"github.com/shatvl/flatwindow/repositories"
)

// AdService for ad management
type AdService struct {
	Repo *repositories.AdRepository
}

// NewUserService returns UserService preference
func NewAdService() *AdService {
	repo := repositories.NewAdRepository()

	return &AdService{
		Repo: repo,
	}
}

// CreateAd creates or updates ad for t-s ads
func (s *AdService) CreateAd(ad *models.Ad) (error) {
	session := mongo.Session()
	defer session.Close()
	
	foundAd, err := s.Repo.FindByTypeAndUID(repositories.TsType, ad.Unid)

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

func (s *AdService) GetAdsWithFilter(filter *models.AdFilterRequest) ([]*models.Ad, string, error){
	return s.Repo.GetAdsWithFilter(filter)
} 