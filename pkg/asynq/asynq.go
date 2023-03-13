package asynq

import (
	"github.com/hibiken/asynq"
	"sync"
	"time"
)

var (
	once   sync.Once
	Client *asynq.Client
	Srv    *asynq.Server
)

func ConnectAsynq(address string, username string, password string, db int) {
	r := asynq.RedisClientOpt{Addr: address, Username: username, Password: password, DB: db}

	once.Do(func() {
		Client = asynq.NewClient(r)
		Srv = asynq.NewServer(
			asynq.RedisClientOpt{Addr: address},
			asynq.Config{
				Concurrency: 10,
				Queues: map[string]int{
					"critical": 6,
					"default":  3,
					"low":      1,
				},
			},
		)
	})
}

func EnqueueIn(task *asynq.Task, timeout int) (err error) {
	if timeout == 0 {
		_, err = Client.Enqueue(task)
		return err
	}

	_, err = Client.Enqueue(task, asynq.ProcessIn(time.Duration(timeout)*time.Minute))
	return err
}
