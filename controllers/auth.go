package controllers

	import (
		"github.com/shatvl/flatwindow/models"
		"github.com/shatvl/flatwindow/services"

		"github.com/kataras/iris"
		"github.com/shatvl/flatwindow/helpers"
	)

// AuthController provides login, register api
type AuthController struct {
	UserService *services.UserService
}

// NewAuthController provides a reference to a AuthController
func NewAuthController() *AuthController {
	return &AuthController{UserService: services.NewUserService()}
}

// RegisterHandler creates a new user
func (ac *AuthController) RegisterHandler(ctx iris.Context) {
	//create user from json body
	user := models.User{}
	ctx.ReadJSON(&user)
	u, err := ac.UserService.Repo.Create(&user)

	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	ctx.JSON(iris.Map{"data": u.ToUserJSON()})
}

// LoginHandler returns JWT
func (ac *AuthController) LoginHandler(ctx iris.Context) {
	credentials := models.Credentials{}
	if err := ctx.ReadJSON(&credentials); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	token, user, err := ac.UserService.GenerateToken(&credentials)

	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map{"data": iris.Map{"token": token, "user": user.ToUserJSON()}})
}

func (ac *AuthController) MeHandler(ctx iris.Context) {
	token, err := helpers.RetrieveTokenFromRequest(ctx.Request())

	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	ctx.JSON(token)
}