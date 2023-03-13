package cmd

import (
	"fmt"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/hibiken/asynq"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"go-yao/boot"
	"go-yao/common/global"
	"go-yao/common/job"
	"go-yao/pkg/logger"
	"log"
	"sync"
)

var CmdServe = &cobra.Command{
	Use:   "serve",
	Short: "Start server",
	Run:   runWeb,
	Args:  cobra.NoArgs,
}

func runWeb(cmd *cobra.Command, args []string) {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	boot.SetupRoute(r)

	w := sync.WaitGroup{}
	w.Add(2)
	fmt.Println("启动端口: " + cast.ToString(global.Conf.Application.Port))
	go func() {
		err := endless.ListenAndServe(":"+cast.ToString(global.Conf.Application.Port), r)
		if err != nil {
			logger.ErrorString("CMD", "serve", err.Error())
		}

		w.Done()
	}()

	go func() {
		defer w.Done()

		srv := asynq.NewServer(
			asynq.RedisClientOpt{Addr: "127.0.0.1:6379"},
			asynq.Config{
				// Specify how many concurrent workers to use
				Concurrency: 10,
				// Optionally specify multiple queues with different priority.
				Queues: map[string]int{
					"critical": 6,
					"default":  3,
					"low":      1,
				},
				// See the godoc for other configuration options
			},
		)

		// mux maps a type to a handler
		mux := asynq.NewServeMux()
		mux.HandleFunc(job.TypeSendSMS, job.HandleSendSMSTask)
		mux.HandleFunc(job.TypeSendEmail, job.HandleSendEmailTask)
		// ...register other handlers...

		if err := srv.Run(mux); err != nil {
			log.Fatalf("could not run server: %v", err)
		}
	}()

	w.Wait()
}
