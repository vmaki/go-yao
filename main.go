package main

import (
	"flag"
	"fmt"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"go-yao/boot"
	"go-yao/common/global"
	"log"
	"os"
	"os/exec"
	"strconv"
	"sync"
)

func init() {
	flag.StringVar(&global.Env, "env", "", "加载 settings.yml，如 --env=dev 加载的是 settings.dev.yml")
	flag.Parse()

	boot.SetupConfig(global.Env)
	boot.SetupLogger()
	boot.SetupDB()
	boot.SetupRedis()
}

func main() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	boot.SetupRoute(r)

	w := sync.WaitGroup{}
	w.Add(1)

	go func() {
		err := endless.ListenAndServe(":"+cast.ToString(global.Conf.Application.Port), r)
		if err != nil {
			log.Println(err)
		}

		log.Println("Server on 5003 stopped")
		w.Done()
	}()

	pid := os.Getpid()
	fmt.Printf("进程 PID: %d \n", pid)

	prc := exec.Command("ps", "-p", strconv.Itoa(pid), "-v")
	out, err := prc.Output()
	if err != nil {
		panic("获取进场 id 失败, err:" + err.Error())
	}

	fmt.Println(string(out))

	w.Wait()
	log.Println("All servers stopped. Exiting.")
}
