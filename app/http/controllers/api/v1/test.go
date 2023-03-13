package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-yao/app/http/controllers/api"
	"go-yao/common/job"
	"go-yao/common/response"
	"go-yao/pkg/asynq"
	"go-yao/pkg/cache"
	"go-yao/pkg/redis"
)

type Test struct {
	api.BaseAPIController
}

func (ctr *Test) Hello(ctx *gin.Context) {
	response.Success(ctx, nil)
}

func (ctr *Test) Err(ctx *gin.Context) {
	panic("这是 panic 测试")
}

func (ctr *Test) Redis(ctx *gin.Context) {
	redis.Client.Set("msg", "hello redis", 64)

	response.Success(ctx, nil)
}

func (ctr *Test) Auth(ctx *gin.Context) {
	data := map[string]uint64{
		"uid": ctr.CurrentUID(ctx),
	}

	response.Success(ctx, data)
}

func (ctr *Test) Cache(ctx *gin.Context) {
	cache.Set("msg", "hello cache", 64)

	response.Success(ctx, nil)
}

func (ctr *Test) Job(ctx *gin.Context) {
	t1, err := job.SendSMSTask("15913395633", 123456)
	if err != nil {
		fmt.Println("创建任务失败, err: " + err.Error())
		return
	}

	err = asynq.EnqueueIn(t1, 0)
	if err != nil {
		fmt.Println("添加任务失败, err: " + err.Error())
		return
	}

	t2, err := job.SendEmailTask("15913395633@163.com", "hello world")
	if err != nil {
		fmt.Println("创建任务2失败, err: " + err.Error())
		return
	}

	err = asynq.EnqueueIn(t2, 3)
	if err != nil {
		fmt.Println("添加任务2失败, err: " + err.Error())
		return
	}

	response.Success(ctx, nil)
}
