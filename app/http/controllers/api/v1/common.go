package v1

import (
	"github.com/gin-gonic/gin"
	"go-yao/app/http/controllers/api"
	"go-yao/app/http/dto"
	"go-yao/common/request"
	"go-yao/common/response"
	"go-yao/common/verifycode"
)

type Common struct {
	api.BaseAPIController
}

func (ctr *Common) SendSms(ctx *gin.Context) {
	req := dto.CommonSendSmsReq{}
	if err := request.Validate(ctx, &req); err != nil {
		response.Error(ctx, err)
		return
	}

	res := verifycode.NewVerifyCode().SendSMS(req.Template, req.Phone)
	if !res {
		response.SysError(ctx)
		return
	}

	response.Success(ctx, nil)
}
