package job

import (
	"context"
	"github.com/hibiken/asynq"
)

type CronJob struct {
	ctx context.Context
}

func NewCronJob(ctx context.Context) *CronJob {
	return &CronJob{
		ctx: ctx,
	}
}

func (l *CronJob) Register() *asynq.ServeMux {
	mux := asynq.NewServeMux()

	mux.HandleFunc(TypeSendSMS, HandleSendSMSTask)
	mux.HandleFunc(TypeSendEmail, HandleSendEmailTask)

	return mux
}
