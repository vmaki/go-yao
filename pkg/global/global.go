package global

import "go-yao/config"

var (
	Env  string
	Conf = new(config.AppConfig)
)
