package v1

import (
	"github.com/gin-gonic/gin"
	"go-yao/app/http/controllers/api"
	"go-yao/app/http/dto"
	"go-yao/app/services"
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

	data := &dto.AuthLoginResp{
		Token: req.Phone + "-" + req.Code,
	}

	response.Data(ctx, data)
}

func (c *AuthController) Register(ctx *gin.Context) {
	req := dto.AuthRegisterReq{}
	if ok := request.Validate(ctx, &req); !ok {
		return
	}

	if isExist := services.IsPhoneExist(req.Phone); isExist {
		response.BadRequest(ctx, "该手机号码已经注册")
		return
	}

	data := &dto.AuthRegisterResp{
		Token: req.Phone + "-" + req.Code,
	}

	response.Data(ctx, data)
}
