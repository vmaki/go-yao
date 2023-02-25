package api

import (
	"github.com/gin-gonic/gin"
	"go-yao/app/http/controllers/api/v1"
	"go-yao/app/middlewares"
)

func RegisterV1UserRoutes(g *gin.RouterGroup) {
	r := g.Group("/users", middlewares.AuthJWT())
	{
		ac := new(v1.Users)

		r.GET("/", ac.Index)
	}
}
