package routes

import (
	"github.com/gin-gonic/gin"
	v1C "go-yao/app/http/controllers/api/v1"
	"go-yao/app/middlewares"
)

func RegisterAPIRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	v1.Use(middlewares.LimitIP("200-H"))
	{
		authGroup := v1.Group("/auth")
		{
			ac := new(v1C.AuthController)

			authGroup.POST("/login", ac.LoginByPhone)
			authGroup.POST("/register", ac.Register)
			authGroup.POST("/refresh-token", middlewares.AuthJWT(), ac.RefreshToken)
		}

		commonGroup := v1.Group("/common")
		{
			cc := new(v1C.CommonController)

			commonGroup.POST("/send-sms", middlewares.LimitPerRoute("20-H"), cc.SendSms)
		}

		testGroup := v1.Group("/test")
		{
			tc := new(v1C.TestController)

			testGroup.GET("/", tc.Hello)
			testGroup.GET("/500", tc.Err)
			testGroup.GET("/redis", tc.Redis)
			testGroup.GET("/is-auth", middlewares.AuthJWT(), tc.Auth)
		}
	}
}
