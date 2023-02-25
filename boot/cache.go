package boot

import (
	"fmt"
	"go-yao/common/global"
	"go-yao/pkg/cache"
)

// SetupCache 缓存
func SetupCache() {
	config := global.Conf.Redis

	// 初始化缓存专用的 redis client, 使用专属缓存 DB
	rds := cache.NewRedisStore(
		fmt.Sprintf("%v:%v", config.Host, config.Port),
		config.Username,
		config.Password,
		config.Database,
	)

	cache.InitWithCacheStore(rds)
}
