package middlewares

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"go-yao/pkg/helper"
	"go-yao/pkg/logger"
	"go.uber.org/zap"
	"io"
	"time"
)

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取 response 内容
		w := responseBodyWriter{
			ResponseWriter: ctx.Writer,
			body:           &bytes.Buffer{},
		}
		ctx.Writer = w

		// 获取请求数据
		var requestBody []byte
		if ctx.Request.Body != nil {
			// c.Request.Body 是一个 buffer 对象，只能读取一次. 因此,读取后需要重新赋值, 以供后续的其他操作
			requestBody, _ = io.ReadAll(ctx.Request.Body)
			ctx.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody))
		}

		startTime := time.Now()

		ctx.Next()

		// 开始记录日志
		endTime := time.Since(startTime)
		status := ctx.Writer.Status()

		logFields := []zap.Field{
			zap.Int("status", status),
			zap.String("request", ctx.Request.Method+" "+ctx.Request.URL.String()),
			zap.String("query", ctx.Request.URL.RawQuery),
			zap.String("ip", helper.GetClientIP(ctx)),
			zap.String("user-agent", ctx.Request.UserAgent()),
			zap.String("errors", ctx.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.String("time", helper.MicrosecondsStr(endTime)),
		}

		if ctx.Request.Method == "POST" || ctx.Request.Method == "PUT" || ctx.Request.Method == "DELETE" {
			logFields = append(logFields, zap.String("Request Body", string(requestBody)))
			logFields = append(logFields, zap.String("Response Body", w.body.String()))
		}

		if status > 400 && status <= 499 {
			logger.Warn("HTTP Warning "+cast.ToString(status), logFields...)
		} else if status >= 500 && status <= 599 {
			logger.Error("HTTP Error "+cast.ToString(status), logFields...)
		}
	}
}
