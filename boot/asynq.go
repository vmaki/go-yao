package boot

import (
	"fmt"
	"go-yao/common/global"
	"go-yao/pkg/asynq"
)

func SetupAsynq() {
	config := global.Conf.Redis
	asynq.ConnectAsynq(
		fmt.Sprintf("%s:%d", config.Host, config.Port),
		config.Username,
		config.Password,
		config.Database,
	)
}
