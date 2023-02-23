package api

import (
	"github.com/gin-gonic/gin"
	"go-yao/app/http/controllers/api/v1"
	"go-yao/app/middlewares"
)

func RegisterV1AuthRoutes(g *gin.RouterGroup) {
	r := g.Group("/auth")
	{
		ac := new(v1.AuthController)

		r.POST("/login", ac.LoginByPhone)
		r.POST("/register", ac.Register)
		r.POST("/refresh-token", middlewares.AuthJWT(), ac.RefreshToken)
	}
}
