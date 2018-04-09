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
	}
	
	ads, _, err := ac.AdService.GetAdsWithFilter(request)

	if err != nil {
		ctx.JSON(err)
	}


	ctx.JSON(ads)
}