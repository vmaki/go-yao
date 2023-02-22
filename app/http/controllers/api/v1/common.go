package v1

import (
	"github.com/gin-gonic/gin"
	"go-yao/app/http/controllers/api"
	"go-yao/app/http/dto"
	"go-yao/common/verifycode"
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

	res := verifycode.NewVerifyCode().SendSMS(req.Template, req.Phone)
	if !res {
		response.BadRequest(ctx, "请求短信失败, 请稍后重试")
		return
	}

	response.Success(ctx)
}
