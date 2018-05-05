package controllers

import (
	"github.com/shatvl/flatwindow/services"
)

type FeedController struct {
	AdService *services.AdService
}

func NewFeedController() *FeedController {
	return &FeedController{AdService: services.NewAdService()}
}