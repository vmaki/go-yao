package routes

import (
	"github.com/gin-gonic/gin"
	v1C "go-yao/app/http/controllers/api/v1"
)

func RegisterAPIRoutes(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		authGroup := v1.Group("/auth")
		{
			ac := new(v1C.AuthController)

			authGroup.POST("/login", ac.Login)
		}

		testGroup := v1.Group("/test")
		{
			tc := new(v1C.TestController)

			testGroup.GET("/", tc.Hello)
			testGroup.GET("/500", tc.Err)
		}
	}
}
