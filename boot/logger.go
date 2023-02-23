package boot

import (
	"go-yao/common/global"
	"go-yao/pkg/logger"
)

func SetupLogger() {
	config := global.Conf.Log

	logger.InitLogger(config.Level, config.Type, config.Filename, config.MaxSize, config.MaxAge, config.MaxBackup, config.Compress)
}
