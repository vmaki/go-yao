package routes

import (
	"github.com/gin-gonic/gin"
	"go-yao/pkg/global"
	"net/http"
)

func RegisterAPIRoutes(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		v1.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "Hello " + global.Conf.Application.Name,
			})
		})
	}
}