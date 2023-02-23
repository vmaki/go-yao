package dto

import (
	"github.com/thedevsaddam/govalidator"
	"go-yao/pkg/request"
)

type CommonSendSmsReq struct {
	Phone    string `json:"phone,omitempty" valid:"phone"`       // 手机号码
	Template string `json:"template,omitempty" valid:"template"` // 短信场景码
}

func (s *CommonSendSmsReq) Generate(data interface{}) string {
	rules := govalidator.MapData{
		"phone":    []string{"required", "digits:11"},
		"template": []string{"required"},
	}

	messages := govalidator.MapData{
		"phone": []string{
			"required:手机号为必填项",
			"digits:手机号长度必须为 11 位的数字",
		},
		"template": []string{
			"required:场景为必填项",
		},
	}

	return request.GoValidate(data, rules, messages)
}
