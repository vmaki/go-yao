package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"go-yao/boot"
	"go-yao/pkg/global"
)

func main() {
	boot.SetupConfig("")

	r := gin.New()
	boot.SetupRoute(r)

	err := r.Run(":" + cast.ToString(global.Conf.Application.Port))
	if err != nil {
		panic("启动失败, err: " + err.Error())
	}
}
