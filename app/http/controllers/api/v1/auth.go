package v1

import (
	"github.com/gin-gonic/gin"
	"go-yao/app/http/controllers/api"
	"go-yao/app/http/dto"
	"go-yao/app/services"
	"go-yao/common/request"
	"go-yao/common/response"
	"go-yao/pkg/jwt"
)

type AuthController struct {
	api.BaseAPIController
}

func (ctr *AuthController) Login(ctx *gin.Context) {
	req := dto.AuthLoginReq{}
	if ok := request.Validate(ctx, &req); !ok {
		return
	}

	// 使用手机号码登录
	us := new(services.UserService)
	user, err := us.LoginByPhone(req.Phone)
	if err != nil {
		response.Error(ctx, err)
		return
	}

	// 生成 Token
	token := jwt.NewJWT().IssueToken(jwt.UserInfo{
		UserID: user.ID,
	})

	data := &dto.AuthLoginResp{
		Token: token,
	}
	response.Success(ctx, data)
}

func (ctr *AuthController) Register(ctx *gin.Context) {
	req := dto.AuthRegisterReq{}
	if ok := request.Validate(ctx, &req); !ok {
		return
	}

	// 注册
	us := new(services.UserService)
	_, err := us.Register(req.Phone)
	if err != nil {
		response.Error(ctx, err)
		return
	}

	response.Success(ctx, nil)
}

func (ctr *AuthController) RefreshToken(ctx *gin.Context) {
	token, err := jwt.NewJWT().RefreshToken(ctx)
	if err != nil {
		response.Unauthorized(ctx, err.Error())
		return
	}

	data := &dto.AuthRefreshTokenResp{
		Token: token,
	}
	response.Success(ctx, data)
}
