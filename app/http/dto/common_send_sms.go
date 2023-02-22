package dto

import (
	"github.com/thedevsaddam/govalidator"
	"go-yao/pkg/request"
)

type CommonSendSmsReq struct {
	Phone    string `json:"phone,omitempty" valid:"phone"`
	Template string `json:"template,omitempty" valid:"template"`
}

func (s *CommonSendSmsReq) Generate(data interface{}) string {
	rules := govalidator.MapData{
		"phone":    []string{"required", "digits:11"},
		"template": []string{"required"},
	}

	messages := govalidator.MapData{
		"phone": []string{
			"required:手机号为必填项",
			"digits:手机号长度错误",
		},
		"template": []string{
			"required:模板code为必填项",
		},
	}

	return request.GoValidate(data, rules, messages)
}
