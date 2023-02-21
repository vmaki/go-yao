package v1

import (
	"github.com/gin-gonic/gin"
	"go-yao/app/http/controllers/api"
	"go-yao/app/http/dto"
	"go-yao/pkg/request"
	"go-yao/pkg/response"
)

type AuthController struct {
	api.BaseAPIController
}

func (c *AuthController) Login(ctx *gin.Context) {
	req := dto.AuthLoginReq{}
	if ok := request.Validate(ctx, &req); !ok {
		return
	}

	data := &dto.LoginResp{
		Token: req.Phone + "-" + req.Code,
	}
	response.Data(ctx, data)
}
