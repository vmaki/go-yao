package api

import (
	"github.com/gin-gonic/gin"
	"go-yao/app/http/controllers/api/v1"
	"go-yao/app/middlewares"
)

func RegisterV1TestRoutes(g *gin.RouterGroup) {
	r := g.Group("/test")
	{
		tc := new(v1.Test)

		r.GET("/", tc.Hello)
		r.GET("/500", tc.Err)
		r.GET("/redis", tc.Redis)
		r.GET("/is-auth", middlewares.AuthJWT(), tc.Auth)
		r.GET("/cache", tc.Cache)
	}
}
