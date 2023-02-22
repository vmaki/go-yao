package api

import "github.com/gin-gonic/gin"

type BaseAPIController struct {
}

// CurrentUID 从 gin.context 中获取当前登录用户 ID
func (c BaseAPIController) CurrentUID(ctx *gin.Context) uint64 {
	return ctx.GetUint64("current_user_id")
}
