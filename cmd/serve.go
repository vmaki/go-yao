package cmd

import (
	"context"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"go-yao/boot"
	"go-yao/common/global"
	"go-yao/common/job"
	asynq2 "go-yao/pkg/asynq"
	"go-yao/pkg/logger"
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

	go func() {
		err := endless.ListenAndServe(":"+cast.ToString(global.Conf.Application.Port), r)
		if err != nil {
			logger.ErrorString("CMD", "serve", err.Error())
		}

		w.Done()
	}()

	go func() {
		jobs := job.NewCronJob(context.Background())
		mux := jobs.Register()
		if err := asynq2.Srv.Run(mux); err != nil {
			logger.ErrorString("CMD", "serve", err.Error())
		}

		w.Done()
	}()

	w.Wait()
}
