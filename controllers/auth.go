package controllers

	import (
		"github.com/shatvl/flatwindow/models"
		"github.com/shatvl/flatwindow/services"

		"github.com/kataras/iris"
		"github.com/dgrijalva/jwt-go/request"
		"github.com/dgrijalva/jwt-go"
		"fmt"
		"github.com/shatvl/flatwindow/config"
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
	token, err := request.ParseFromRequest(ctx.Request(), request.AuthorizationHeaderExtractor, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(config.Secret), nil
	})

	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	ctx.JSON(token)
}
