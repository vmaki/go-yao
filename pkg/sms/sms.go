package sms

import (
	"go-yao/pkg/global"
	"sync"
)

type Message struct {
	Template string
	Data     map[string]string
	Content  string
}

type SMS struct {
	Driver IDriver
}

var once sync.Once
var sms *SMS

func NewSMS() *SMS {
	once.Do(func() {
		sms = &SMS{Driver: &Aliyun{}}
	})

	return sms
}

// Send 发送短信
// 调用示: sms.NewSMS().Send("手机号", sms.Message{ Template: "800820", Data: map[string]string{"code": "123456"}})
func (sms *SMS) Send(phone string, message Message) bool {
	config := map[string]string{
		"access_key_id":     global.Conf.Sms.Aliyun.AccessKeyId,
		"access_key_secret": global.Conf.Sms.Aliyun.AccessKeySecret,
		"sign_name":         global.Conf.Sms.Aliyun.SignName,
	}

	return sms.Driver.Send(phone, message, config)
}
