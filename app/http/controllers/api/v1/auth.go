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

type Auth struct {
	api.BaseAPIController
}

func (ctr *Auth) LoginByPhone(ctx *gin.Context) {
	req := dto.AuthLoginReq{}
	if err := request.Validate(ctx, &req); err != nil {
		response.Error(ctx, err)
		return
	}

	us := new(services.User)
	user, err := us.LoginByPhone(req.Phone)
	if err != nil {
		response.Error(ctx, err)
		return
	}

	token := jwt.NewJWT().IssueToken(jwt.UserInfo{
		UserID: user.ID,
	})

	data := &dto.AuthLoginResp{
		Token: token,
	}
	response.Success(ctx, data)
}

func (ctr *Auth) Register(ctx *gin.Context) {
	req := dto.AuthRegisterReq{}
	if err := request.Validate(ctx, &req); err != nil {
		response.Error(ctx, err)
		return
	}

	us := new(services.User)
	_, err := us.Register(req.Phone)
	if err != nil {
		response.Error(ctx, err)
		return
	}

	response.Success(ctx, nil)
}

func (ctr *Auth) RefreshToken(ctx *gin.Context) {
	token, err := jwt.NewJWT().RefreshToken(ctx)
	if err != nil {
		response.Unauthorized(ctx, err)
		return
	}

	data := &dto.AuthRefreshTokenResp{
		Token: token,
	}
	response.Success(ctx, data)
}
