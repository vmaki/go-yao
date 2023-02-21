package boot

import (
	"go-yao/pkg/global"
	"go-yao/pkg/logger"
)

func SetupLogger() {
	log := global.Conf.Log
	logger.InitLogger(log.Level, log.Type, log.Filename, log.MaxSize, log.MaxAge, log.MaxBackup, log.Compress)
}
