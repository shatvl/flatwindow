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

// CreateTSAd creates or updates ad for t-s ads
func (s *AdService) CreateTsAd(ad *models.TsAd) (error) {
	session := mongo.Session()
	defer session.Close()
	
	foundAd, err := s.Repo.FindByTypeAndUID(repositories.TsType, ad.Unid)

	if err == nil {
		fmt.Println("Ad is found: " + foundAd.Unid)
		return nil
	}		

	err = s.Repo.CreateTsAd(ad)

	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("Ad has been added successfully: " + ad.Unid)

	return nil
}