package v1

import (
	"github.com/gin-gonic/gin"
	"go-yao/app/http/controllers/api"
	"go-yao/app/http/dto"
	"go-yao/common/cache"
	"go-yao/common/helpers"
	"go-yao/pkg/redis"
	"go-yao/pkg/request"
	"go-yao/pkg/response"
)

type CommonController struct {
	api.BaseAPIController
}

func (c *CommonController) SendSms(ctx *gin.Context) {
	req := dto.CommonSendSmsReq{}
	if ok := request.Validate(ctx, &req); !ok {
		return
	}

	code := helpers.RandomNumber(6)
	res := redis.Client.Set(cache.GetSmsCacheKey(req.Scene, req.Phone), code, 60*5)
	if !res {
		response.BadRequest(ctx, "请求短信失败, 请稍后重试")
		return
	}

	response.Success(ctx)
}
