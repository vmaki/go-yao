package boot

import (
	"fmt"
	"go-yao/common/global"
	"go-yao/pkg/redis"
)

func SetupRedis() {
	config := global.Conf.Redis

	redis.ConnectRedis(
		fmt.Sprintf("%s:%d", config.Host, config.Port),
		config.Username,
		config.Password,
		config.Database,
	)
}
