package routes

import (
	"github.com/gin-gonic/gin"
	"go-yao/app/middlewares"
	"go-yao/routes/api"
)

func RegisterAPIRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	v1.Use(middlewares.LimitIP("200-H"))
	{
		api.RegisterV1AuthRoutes(v1)
		api.RegisterV1CommonRoutes(v1)
		api.RegisterV1UserRoutes(v1)

		api.RegisterV1TestRoutes(v1)
	}
}
