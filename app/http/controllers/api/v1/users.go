package v1

import (
	"github.com/gin-gonic/gin"
	"go-yao/app/http/controllers/api"
	"go-yao/app/http/dto"
	"go-yao/app/models/user"
	"go-yao/common/request"
	"go-yao/common/response"
)

type Users struct {
	api.BaseAPIController
}

func (ctr *Users) Index(ctx *gin.Context) {
	req := dto.PaginationReq{}
	if err := request.Validate(ctx, &req); err != nil {
		response.Error(ctx, err)
		return
	}

	data, pager := user.Paginate(ctx, 10)

	response.Success(ctx, gin.H{
		"data":  data,
		"pager": pager,
	})
}
