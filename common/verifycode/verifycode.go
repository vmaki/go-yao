package verifycode

import (
	"go-yao/common/helpers"
	"go-yao/common/sms"
	"go-yao/pkg/global"
	"go-yao/pkg/logger"
	"go-yao/pkg/redis"
	"sync"
)

type VerifyCode struct {
	Store IStore
}

var (
	once       sync.Once
	verifyCode *VerifyCode
)

// NewVerifyCode 单例模式获取
func NewVerifyCode() *VerifyCode {
	once.Do(func() {
		verifyCode = &VerifyCode{
			Store: &RedisStore{
				RedisClient: redis.Client,
				KeyPrefix:   "sms:",
			},
		}
	})

	return verifyCode
}

func (vc *VerifyCode) cacheKey(template, phone string) string {
	return template + ":" + phone
}

// SendSMS 发送短信
func (vc *VerifyCode) SendSMS(template, phone string) bool {
	code := vc.generateVerifyCode(vc.cacheKey(template, phone)) // 生成验证码

	if global.Conf.Application.Mode == "local" {
		logger.DebugString("VerifyCode", template, code)
		return true
	}

	// 发送短信
	return sms.NewSMS().Send(phone, sms.Message{
		Template: template,
		Data:     map[string]string{"code": code},
	})
}

// generateVerifyCode 生成验证码，并放置于 Redis 中
func (vc *VerifyCode) generateVerifyCode(key string) string {
	code := helpers.RandomNumber(6)
	vc.Store.Set(key, code, 60*10)

	return code
}

func (vc *VerifyCode) CheckAnswer(template, phone string, answer string) bool {
	return vc.Store.Verify(vc.cacheKey(template, phone), answer, false)
}
