package global

import (
	"go-yao/config"
	"gorm.io/gorm"
)

var (
	Env  string
	Conf = new(config.AppConfig)
	DB   *gorm.DB
)
