package middlewares

import (
	"github.com/gin-gonic/gin"
	"go-yao/pkg/logger"
	"go-yao/pkg/response"
	"go.uber.org/zap"
	"net"
	"net/http/httputil"
	"os"
	"strings"
	"time"
)

func Recovery() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 获取用户的请求信息
				httpRequest, _ := httputil.DumpRequest(ctx.Request, true)

				// 链接中断，客户端中断连接为正常行为，不需要记录堆栈信息
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						errStr := strings.ToLower(se.Error())
						if strings.Contains(errStr, "broken pipe") || strings.Contains(errStr, "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				logFields := []zap.Field{
					zap.Time("time", time.Now()),
					zap.Any("error", err),
					zap.String("request", string(httpRequest)),
				}

				// 链接中断的情况
				if brokenPipe {
					logger.Error(ctx.Request.URL.Path, logFields...)
					_ = ctx.Error(err.(error))
					ctx.Abort()
					return
				}

				// 如果不是链接中断，就开始记录堆栈信息
				logFields = append(logFields, zap.Stack("stacktrace"))
				logger.Error("recovery from panic", logFields...)

				response.Error(ctx)
			}
		}()

		ctx.Next()
	}
}
