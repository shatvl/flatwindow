package controllers

import (
	"github.com/shatvl/flatwindow/services"
	"github.com/shatvl/flatwindow/models/ads"

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
	
	ads, err := ac.AdService.Repo.GetAdsWithFilter(request)

	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	ctx.JSON(ads)
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