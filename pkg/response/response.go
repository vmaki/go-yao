package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func JSON(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, data)
}

func Success(ctx *gin.Context) {
	JSON(ctx, gin.H{
		"code": 200,
		"msg":  "success",
	})
}

func Data(ctx *gin.Context, data interface{}) {
	JSON(ctx, gin.H{
		"code": 200,
		"msg":  "success",
		"data": data,
	})
}

func Error(ctx *gin.Context, msg ...string) {
	ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
		"code": 500,
		"msg":  defaultMessage("服务器内部错误，请稍后再试", msg...),
	})
}

// BadRequest 在解析用户请求，请求的格式或者方法不符合预期时调用
func BadRequest(ctx *gin.Context, err error, msg ...string) {
	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"code": 400,
		"msg":  defaultMessage("请求解析错误，请确认请求格式是否正确。上传文件请使用 multipart 标头，参数请使用 JSON 格式。", msg...),
	})
}

// ValidationError 处理表单验证不通过的错误
func ValidationError(ctx *gin.Context, errors map[string][]string) {
	ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
		"code": 422,
		"msg":  "请求验证不通过",
	})
}

// Unauthorized 未传参 msg 时使用默认消息
func Unauthorized(ctx *gin.Context, msg ...string) {
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"code": 401,
		"msg":  defaultMessage("请求解析错误，请确认请求格式是否正确。上传文件请使用 multipart 标头，参数请使用 JSON 格式。", msg...),
	})
}

func Abort404(ctx *gin.Context, msg ...string) {
	ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
		"code":    404,
		"message": defaultMessage("路由未定义，请确认 url 和请求方法是否正确", msg...),
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
