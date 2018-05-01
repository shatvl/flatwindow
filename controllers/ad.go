package controllers

import (
	"github.com/shatvl/flatwindow/services"
	"github.com/shatvl/flatwindow/models"

	"github.com/kataras/iris"
)

type AdController struct {
	AdService *services.AdService
}

func NewAdController() *AdController {
	
	return &AdController{AdService: services.NewAdService()}
}


func (ac *AdController) GetProductsHandler(ctx iris.Context) {
	request := &models.AdFilterRequest{}

	if err := ctx.ReadJSON(request); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}
	
	ads, count, err := ac.AdService.Repo.GetAdsWithFilter(request)

	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	ctx.JSON(iris.Map{"data": iris.Map{"ads": ads, "count": count}})
}

func (ac *AdController) GetProductHandler(ctx iris.Context) {
	id := ctx.Params().Get("_id")

	ad, err := ac.AdService.Repo.GetAdById(id)

	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	ctx.JSON(iris.Map{"data": ad})
}

type addAdToFeedRequest struct {
	AdId     string `json:"adId"`
	FeedType int    `json:"feedType"`
	Value    bool   `json:"value"`
}

func (ac *AdController) AddAdToFeed(ctx iris.Context) {
	var request addAdToFeedRequest

	if err := ctx.ReadJSON(&request); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	err := ac.AdService.AddAdToFeed(request.AdId, request.FeedType, request.Value)

	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	ctx.StatusCode(iris.StatusNoContent)
	ctx.JSON(iris.Map{"norm": "norm"})
}