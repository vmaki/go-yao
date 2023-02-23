package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"go-yao/common/response"
	"go-yao/pkg/limiter"
	"go-yao/pkg/logger"
)

// LimitIP 全局限流中间件，针对 IP 进行限流
func LimitIP(limit string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		key := limiter.GetKeyIP(ctx)
		if ok := limitHandler(ctx, key, limit); !ok {
			return
		}

		ctx.Next()
	}
}

// LimitPerRoute 限流中间件，用在单独的路由中
func LimitPerRoute(limit string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 针对单个路由，增加访问次数
		ctx.Set("limiter-once", false)

		// 针对 IP + 路由进行限流
		key := limiter.GetKeyRouteWithIP(ctx)
		if ok := limitHandler(ctx, key, limit); !ok {
			return
		}

		ctx.Next()
	}
}

// limit 为格式化字符串，如 "5-S" ，示例:
//
// * 5 reqs/second: "5-S"
// * 10 reqs/minute: "10-M"
// * 1000 reqs/hour: "1000-H"
// * 2000 reqs/day: "2000-D"
func limitHandler(ctx *gin.Context, key string, limit string) bool {
	// 获取超额的情况
	rate, err := limiter.CheckRate(ctx, key, limit)
	if err != nil {
		logger.LogIf(err)
		response.SysError(ctx)
		return false
	}

	// ---- 设置标头信息-----
	ctx.Header("X-RateLimit-Limit", cast.ToString(rate.Limit))         // 最大访问次数
	ctx.Header("X-RateLimit-Remaining", cast.ToString(rate.Remaining)) // 剩余的访问次数
	ctx.Header("X-RateLimit-Reset", cast.ToString(rate.Reset))         // 到该时间点，访问次数会重置为 X-RateLimit-Limit

	// 超额
	if rate.Reached {
		response.TooManyRequests(ctx)
		return false
	}

	return true
}
