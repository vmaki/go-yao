package helper

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net"
	"time"
)

func GetClientIP(ctx *gin.Context) string {
	clientIP := ctx.Request.RemoteAddr
	if ip := ctx.GetHeader("X-Real-IP"); ip != "" {
		clientIP = ip
	} else if ip = ctx.GetHeader("X-Forward-For"); ip != "" {
		clientIP = ip
	} else {
		clientIP, _, _ = net.SplitHostPort(clientIP)
	}

	if clientIP == "::1" {
		clientIP = "127.0.0.1"
	}

	return clientIP
}

// MicrosecondsStr 将 time.Duration 类型输出为小数点后 3 位的 ms
func MicrosecondsStr(elapsed time.Duration) string {
	return fmt.Sprintf("%.3fms", float64(elapsed.Nanoseconds())/1e6)
}
