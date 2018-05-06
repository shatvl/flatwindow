package adminc

import (
	"github.com/kataras/iris"
	"github.com/shatvl/flatwindow/models"
	"github.com/shatvl/flatwindow/services"
)

// AuthController provides login, register api
type AuthController struct {
	UserService *services.UserService
}

// NewAuthController provides a reference to a AuthController
func NewAuthController() *AuthController {
	return &AuthController{UserService: services.NewUserService()}
}

// LoginHandler returns JWT
func (ac *AuthController) LoginHandler(ctx iris.Context) {
	credentials := models.Credentials{}
	if err := ctx.ReadJSON(&credentials); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	token, user, err := ac.UserService.GenerateAdminToken(&credentials)

	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map{"data": iris.Map{"token": token, "user": user.ToUserJSON()}})
}