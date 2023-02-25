package v1

import (
	"github.com/gin-gonic/gin"
	"go-yao/app/http/controllers/api"
	"go-yao/common/response"
	"go-yao/pkg/cache"
	"go-yao/pkg/redis"
)

type Test struct {
	api.BaseAPIController
}

func (ctr *Test) Hello(ctx *gin.Context) {
	response.Success(ctx, nil)
}

func (ctr *Test) Err(ctx *gin.Context) {
	panic("这是 panic 测试")
}

func (ctr *Test) Redis(ctx *gin.Context) {
	redis.Client.Set("msg", "hello redis", 64)

	response.Success(ctx, nil)
}

func (ctr *Test) Auth(ctx *gin.Context) {
	data := map[string]uint64{
		"uid": ctr.CurrentUID(ctx),
	}

	response.Success(ctx, data)
}

func (ctr *Test) Cache(ctx *gin.Context) {
	cache.Set("msg", "hello cache", 64)

	response.Success(ctx, nil)
}
