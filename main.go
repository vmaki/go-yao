package main

import (
	"github.com/gin-gonic/gin"
	"go-yao/boot"
)

func main() {
	r := gin.Default()

	boot.SetupRoute(r)

	err := r.Run(":7001")
	if err != nil {
		panic("启动失败, err: " + err.Error())
	}
}
