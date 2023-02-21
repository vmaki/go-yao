package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-yao/app/http/controllers/api"
	"go-yao/pkg/response"
)

type TestController struct {
	api.BaseAPIController
}

func (c *TestController) Hello(ctx *gin.Context) {
	response.Success(ctx)
}

func (c *TestController) Err(ctx *gin.Context) {
	panic("这是 panic 测试")
	fmt.Println("11111")
}
