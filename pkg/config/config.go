package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go-yao/pkg/global"
	"go-yao/pkg/logger"
	"go.uber.org/zap"
	"os"
	"time"
)

var v *viper.Viper

func init() {
	v = viper.New()

	v.SetConfigType("yml")
	v.AddConfigPath(".")
	v.SetEnvPrefix("goYao")
	v.AutomaticEnv()
}

func LoadConfig(env string) {
	configName := "settings.yml"
	if env != "" {
		configName = fmt.Sprintf("settings.%s.yml", env)
	}

	// 检查配置文件是否存在
	if _, err := os.Stat(configName); err != nil {
		panic("启动失败，err: 配置文件 " + configName + " 不存在. " + err.Error())
	}

	// 读取配置文件
	v.SetConfigName(configName)
	if err := v.ReadInConfig(); err != nil {
		panic("启动失败，err: 读取配置文件 " + configName + " 失败. " + err.Error())
	}

	// 加载配置
	if err := v.Unmarshal(global.Conf); err != nil {
		panic("启动失败，err: 加载配置失败，" + err.Error())
	}

	// 监控配置文件，变更时重新加载，无需重启
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		logger.WarnString("配置文件", "重新加载", time.Now().Format("2006-01-02 15:04:05"))
		if err := v.Unmarshal(global.Conf); err != nil {
			logger.Error("配置文件", zap.Error(err))
		}
	})
}
