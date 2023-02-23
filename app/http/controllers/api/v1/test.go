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

func (ctr *TestController) Hello(ctx *gin.Context) {
	response.Success(ctx)
}

func (ctr *TestController) Err(ctx *gin.Context) {
	panic("这是 panic 测试")
}

func (ctr *TestController) Redis(ctx *gin.Context) {
	redis.Client.Set("msg", "hello world", 64)

	response.Success(ctx)
}

func (ctr *TestController) Auth(ctx *gin.Context) {
	data := map[string]uint64{
		"uid": ctr.CurrentUID(ctx),
	}

	response.Data(ctx, data)
}
