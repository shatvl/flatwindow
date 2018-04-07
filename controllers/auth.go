package controllers

import (
	"github.com/shatvl/flatwindow/models"
	"github.com/shatvl/flatwindow/services"

	"github.com/kataras/iris"
)

// AuthController provides login, register api
type AuthController struct {
	userService *services.UserService
}

// NewAuthController provides a reference to a AuthController
func NewAuthController() *AuthController {
	userService := services.NewUserService()

	return &AuthController{userService}
}

// RegisterHandler creates a new user
func (ac *AuthController) RegisterHandler(ctx iris.Context) {
	//create user from json body
	user := models.User{}
	ctx.ReadJSON(&user)
	res, err := ac.userService.Repo.Create(&user)

	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": err.Error()})
	}

	ctx.JSON(res)
}

// LoginHandler returns JWT
func (ac *AuthController) LoginHandler(ctx iris.Context) {
	credentials := models.Credentials{}
	if err := ctx.ReadJSON(&credentials); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	token, user, err := ac.userService.GenerateToken(&credentials)

	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map{"token": token, "user": user})
}
