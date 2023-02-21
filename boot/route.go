package boot

import (
	"github.com/gin-gonic/gin"
	"go-yao/app/http/middlewares"
	"go-yao/routes"
	"net/http"
	"strings"
)

func SetupRoute(r *gin.Engine) {
	registerGlobalMiddleware(r)

	routes.RegisterAPIRoutes(r)

	notFoundHandle(r)
}

// registerGlobalMiddleware 注册全局中间件
func registerGlobalMiddleware(r *gin.Engine) {
	r.Use(middlewares.Logger(), middlewares.Recovery())
}

// notFoundHandle 处理404请求
func notFoundHandle(r *gin.Engine) {
	r.NoRoute(func(ctx *gin.Context) {
		accept := ctx.GetHeader("Accept")
		if strings.Contains(accept, "text/html") {
			ctx.String(http.StatusNotFound, "页面返回 404")
		} else {
			ctx.JSON(http.StatusNotFound, gin.H{
				"code": 404,
				"msg":  "路由未定义，请确认 url 和请求方法是否正确",
			})
		}
	})
}
