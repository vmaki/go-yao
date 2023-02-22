package dto

import (
	"github.com/thedevsaddam/govalidator"
	"go-yao/pkg/request"
)

type AuthRegisterReq struct {
	Phone string `json:"phone,omitempty" valid:"phone"`
	Code  string `json:"code,omitempty" valid:"code"`
}

func (s *AuthRegisterReq) Generate(data interface{}) string {
	rules := govalidator.MapData{
		"phone": []string{"required", "digits:11"},
		"code":  []string{"required", "digits:6"},
	}

	messages := govalidator.MapData{
		"phone": []string{
			"required:手机号为必填项",
			"digits:手机号长度错误",
		},
		"code": []string{
			"required:验证码为必填项",
			"digits:验证码长度错误",
		},
	}

	return request.GoValidate(data, rules, messages)
}

type AuthRegisterResp struct {
	Token string `json:"token"`
}
