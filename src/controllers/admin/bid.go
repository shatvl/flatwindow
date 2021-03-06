package adminc

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris"
	"github.com/shatvl/flatwindow/helpers"
	"github.com/shatvl/flatwindow/models"
	"github.com/shatvl/flatwindow/services"
)

type BidController struct {
	BidService *services.BidService
}

func NewBidController() *BidController {

	return &BidController{BidService: services.NewBidService()}
}

func (bc *BidController) GetBidsHandler(ctx iris.Context) {
	token, err := helpers.RetrieveTokenFromRequest(ctx.Request())

	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	claims := token.Claims.(jwt.MapClaims)

	request := &models.BidFilterRequest{AgentType: byte(claims["agent_type"].(float64))}

	if err := ctx.ReadJSON(request); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	bids, count, err := bc.BidService.GetPaginatedBids(request)

	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	ctx.JSON(iris.Map{"data": iris.Map{"bids": bids, "count": count}})
}

func (bc *BidController) UpdateBidHandler(ctx iris.Context) {
	bid := &models.UpdatedBid{}

	if err := ctx.ReadJSON(bid); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	err := bc.BidService.UpdateBid(bid)

	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	ctx.StatusCode(iris.StatusNoContent)
}
