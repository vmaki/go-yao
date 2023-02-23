package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func JSON(ctx *gin.Context, data RespData) {
	ctx.JSON(http.StatusOK, data)
}

// Success 请求成功
func Success(ctx *gin.Context, data interface{}) {
	JSON(ctx, RespData{
		Code: CodeSuccess,
		Msg:  CodeSuccess.Msg(),
		Data: data,
	})
}

// Error 自定义错误
func Error(ctx *gin.Context, err error) {
	switch _err := err.(type) {
	case *RespData:
		JSON(ctx, RespData{
			Code: _err.Code,
			Msg:  _err.Msg,
		})
	default:
		SysError(ctx)
	}
}

// SysError 服务器内部错误
func SysError(ctx *gin.Context) {
	ctx.AbortWithStatusJSON(http.StatusInternalServerError, RespData{
		Code: CodeSysError,
		Msg:  CodeSysError.Msg(),
	})
}

// BadRequest 解析用户请求，请求的格式不符合预期时调用
func BadRequest(ctx *gin.Context) {
	ctx.AbortWithStatusJSON(http.StatusBadRequest, RespData{
		Code: CodeBadRequest,
		Msg:  CodeBadRequest.Msg(),
	})
}

// Unauthorized 解析 jwt 失败时调用
func Unauthorized(ctx *gin.Context, err error) {
	switch _err := err.(type) {
	case *RespData:
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, RespData{
			Code: _err.Code,
			Msg:  _err.Msg,
		})
	default:
		SysError(ctx)
	}
}

// Abort404 响应 404 请求
func Abort404(ctx *gin.Context) {
	ctx.AbortWithStatusJSON(http.StatusNotFound, RespData{
		Code: CodeNotFound,
		Msg:  CodeNotFound.Msg(),
	})
}

// TooManyRequests 请求过于频繁
func TooManyRequests(ctx *gin.Context) {
	ctx.AbortWithStatusJSON(http.StatusTooManyRequests, RespData{
		Code: CodeTooManyRequests,
		Msg:  CodeTooManyRequests.Msg(),
	})
}

// defaultMessage 内用的辅助函数，用以支持默认参数默认值
func defaultMessage(defaultMsg string, msg ...string) (message string) {
	if len(msg) > 0 {
		message = msg[0]
	} else {
		message = defaultMsg
	}

	return
}
