package controllers

import (
	"github.com/kataras/iris"
	"github.com/shatvl/flatwindow/models"
	"github.com/shatvl/flatwindow/services"
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
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	err := bc.BidService.CreateBid(bid)

	if err != nil {
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	if bid.CopyEmail {
		//err = services.NewSmtpMailer().SendBidRequestMail()
	}

	if err != nil {
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	ctx.JSON(iris.Map{"data": bid})
}
