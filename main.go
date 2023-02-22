package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"go-yao/boot"
	"go-yao/pkg/global"
)

func init() {
	flag.StringVar(&global.Env, "env", "", "加载 settings.yml，如 --env=dev 加载的是 settings.dev.yml")
	flag.Parse()

	boot.SetupConfig(global.Env)
	boot.SetupLogger()
	boot.SetupDB()
}

func main() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	boot.SetupRoute(r)

	err := r.Run(":" + cast.ToString(global.Conf.Application.Port))
	if err != nil {
		panic("启动失败, err: " + err.Error())
	}
}
