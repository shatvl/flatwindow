package controllers

import (
	"github.com/shatvl/flatwindow/services"
	"github.com/kataras/iris"
	"github.com/shatvl/flatwindow/models"
)

type BidController struct {
	BidService *services.BidService
}

func NewBidController() *BidController {
	
	return &BidController{BidService: services.NewBidService()}
}

func (bc *BidController) BidAdHandler(ctx iris.Context) {
	bid := &models.Bid{}

	if err := ctx.ReadJSON(bid); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	err := bc.BidService.CreateBid(bid)	
	
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	if bid.CopyEmail {
		err = services.NewSmtpMailer().SendBidRequestMail()
	}

	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	ctx.JSON(iris.Map{"data" : bid})
}

func (bc *BidController) GetBidsHandler(ctx iris.Context) {
	ctx.JSON("GetBidsAction");
}