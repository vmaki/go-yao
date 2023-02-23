package api

import (
	"github.com/gin-gonic/gin"
	"go-yao/app/http/controllers/api/v1"
	"go-yao/app/middlewares"
)

func RegisterV1CommonRoutes(g *gin.RouterGroup) {
	r := g.Group("/common")
	{
		cc := new(v1.CommonController)

		r.POST("/send-sms", middlewares.LimitPerRoute("1-M"), cc.SendSms)
	}
}
