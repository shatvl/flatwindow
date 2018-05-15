package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris"
	"github.com/shatvl/flatwindow/config"
	"github.com/shatvl/flatwindow/helpers"
)

func AgentRoleResolverMiddleware(ctx iris.Context) {
	println("Role resolver middleware")
	token, err := helpers.RetrieveTokenFromRequest(ctx.Request())

	if err != nil {
		ctx.StatusCode(iris.StatusUnauthorized)
		return
	}

	claims := token.Claims.(jwt.MapClaims)

	if claims["role"].(string) != config.ROLE_AGENT {
		ctx.StatusCode(iris.StatusUnauthorized)
		return
	}

	ctx.Next()
}
