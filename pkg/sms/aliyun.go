package sms

import (
	"encoding/json"
	openApi "github.com/alibabacloud-go/darabonba-openapi/client"
	smsApi "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	"go-yao/pkg/logger"
)

type Aliyun struct {
}

// NewClient 初始化客户端
func (s *Aliyun) NewClient(accessKeyId, accessKeySecret string) (*smsApi.Client, error) {
	apiConfig := &openApi.Config{
		Endpoint:        tea.String("dysmsapi.aliyuncs.com"),
		AccessKeyId:     tea.String(accessKeyId),
		AccessKeySecret: tea.String(accessKeySecret),
	}

	return smsApi.NewClient(apiConfig)
}

func (s *Aliyun) Send(phone string, message Message, config map[string]string) bool {
	client, err := s.NewClient(config["access_key_id"], config["access_key_secret"])
	if err != nil {
		logger.ErrorString("短信[阿里云]", "初始化客户端失败", err.Error())
		return false
	}

	// 短信模板参数解析
	templateParam, err := json.Marshal(message.Data)
	if err != nil {
		logger.ErrorString("短信[阿里云]", "短信模板参数解析失败", err.Error())
		return false
	}

	// 发送短信
	sendSmsRequest := &smsApi.SendSmsRequest{
		PhoneNumbers:  tea.String(phone),
		SignName:      tea.String(config["sign_name"]),
		TemplateCode:  tea.String(message.Template),
		TemplateParam: tea.String(string(templateParam)),
	}
	res, err := client.SendSms(sendSmsRequest)
	if err != nil {
		logger.ErrorString("短信[阿里云]", "发送短信失败", err.Error())
		return false
	}

	if *res.Body.Code != "OK" {
		logger.ErrorJSON("短信[阿里云]", "服务商返回错误", *res.Body)
		return false
	}

	return true
}
