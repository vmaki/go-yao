package v1

import (
	"github.com/gin-gonic/gin"
	"go-yao/app/http/controllers/api"
	"go-yao/pkg/redis"
	"go-yao/pkg/response"
)

type TestController struct {
	api.BaseAPIController
}

func (c *TestController) Hello(ctx *gin.Context) {
	response.Success(ctx)
}

func (c *TestController) Err(ctx *gin.Context) {
	panic("这是 panic 测试")
}

func (c *TestController) Redis(ctx *gin.Context) {
	redis.Client.Set("msg", "hello world", 64)
	response.Success(ctx)
}

func (c *TestController) Auth(ctx *gin.Context) {
	data := map[string]uint64{
		"uid": c.CurrentUID(ctx),
	}
	response.Data(ctx, data)
}
