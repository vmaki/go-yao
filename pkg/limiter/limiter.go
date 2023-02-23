package limiter

import (
	"github.com/gin-gonic/gin"
	limiterLib "github.com/ulule/limiter/v3"
	limiterRedis "github.com/ulule/limiter/v3/drivers/store/redis"
	"go-yao/common/global"
	"go-yao/pkg/helper"
	"go-yao/pkg/logger"
	"go-yao/pkg/redis"
	"strings"
)

// GetKeyIP 限制流 key: IP，针对 ip 做限流
func GetKeyIP(ctx *gin.Context) string {
	return helper.GetClientIP(ctx)
}

// GetKeyRouteWithIP 限制流 key: 路由+IP，针对单个路由做限流
func GetKeyRouteWithIP(ctx *gin.Context) string {
	return routeToKeyString(ctx.FullPath()) + helper.GetClientIP(ctx)
}

// CheckRate 检测请求是否超额
func CheckRate(ctx *gin.Context, key string, formatted string) (limiterLib.Context, error) {
	var context limiterLib.Context
	rate, err := limiterLib.NewRateFromFormatted(formatted)
	if err != nil {
		logger.LogIf(err)
		return context, err
	}

	// 初始化存储，使用我们程序里共用的 redis.Redis 对象
	store, err := limiterRedis.NewStoreWithOptions(redis.Client.Client, limiterLib.StoreOptions{
		Prefix: global.Conf.Application.Name + ":limiter",
	})
	if err != nil {
		logger.LogIf(err)
		return context, err
	}

	// 使用上面的初始化的 limiter.Rate 对象和存储对象
	limiterObj := limiterLib.New(store, rate)

	// 获取限流的结果
	if ctx.GetBool("limiter-once") {
		// Peek() 取结果，不增加访问次数
		return limiterObj.Peek(ctx, key)
	}

	ctx.Set("limiter-once", true) // 确保多个路由组里调用 LimitIP 进行限流时，只增加一次访问次数。

	// Get() 取结果且增加访问次数
	return limiterObj.Get(ctx, key)
}

// routeToKeyString 辅助方法，将 URL 中的 / 格式为 -
func routeToKeyString(routeName string) string {
	routeName = strings.ReplaceAll(routeName, "/", "-")
	routeName = strings.ReplaceAll(routeName, ":", "_")

	return routeName
}
