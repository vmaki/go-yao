package boot

import "go-yao/pkg/config"

func SetupConfig(env string) {
	config.LoadConfig(env)
}
