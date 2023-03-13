package job

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"time"
)

func SendSMSTask(phone string, code int) (*asynq.Task, error) {
	payload, err := json.Marshal(SendSMSPayload{Phone: phone, Code: code})
	if err != nil {
		return nil, err
	}

	return asynq.NewTask(TypeSendSMS, payload), nil
}

func HandleSendSMSTask(ctx context.Context, t *asynq.Task) error {
	var p SendSMSPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("发短信失败, err: : %v: %w", err, asynq.SkipRetry)
	}

	fmt.Printf("开始发送短信. 手机号码: %s, 短信内容: %d\n", p.Phone, p.Code)
	return nil
}

func SendEmailTask(email, content string) (*asynq.Task, error) {
	payload, err := json.Marshal(SendEmailPayload{Email: email, Content: content})
	if err != nil {
		return nil, err
	}

	return asynq.NewTask(TypeSendEmail, payload, asynq.ProcessIn(2*time.Minute)), nil
}

func HandleSendEmailTask(ctx context.Context, t *asynq.Task) error {
	var p SendEmailPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("发邮件失败, err: : %v: %w", err, asynq.SkipRetry)
	}

	fmt.Printf("开始发送邮件. 邮箱: %s, 邮件内容: %s\n", p.Email, p.Content)
	return nil
}
