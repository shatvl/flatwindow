package adminc

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris"
	"github.com/shatvl/flatwindow/helpers"
	"github.com/shatvl/flatwindow/models"
	"github.com/shatvl/flatwindow/services"
)

type AdController struct {
	AdService *services.AdService
}

func NewAdController() *AdController {

	return &AdController{AdService: services.NewAdService()}
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

func (ac *AdController) GetProductsHandler(ctx iris.Context) {
	token, err := helpers.RetrieveTokenFromRequest(ctx.Request())

	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	claims := token.Claims.(jwt.MapClaims)
	request := &models.AdFilterRequest{}
	request.Filter.AgentType = byte(claims["agent_type"].(float64))

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
