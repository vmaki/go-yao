package dto

import (
	"github.com/thedevsaddam/govalidator"
	"go-yao/app/http/dto/validators"
	"go-yao/pkg/request"
)

type AuthRegisterReq struct {
	Phone    string `json:"phone,omitempty" valid:"phone"`
	Code     string `json:"code,omitempty" valid:"code"`
	Template string `json:"template,omitempty" valid:"template"`
}

func (s *AuthRegisterReq) Generate(data interface{}) string {
	rules := govalidator.MapData{
		"phone":    []string{"required", "digits:11", "not_exists:users,phone"},
		"code":     []string{"required", "digits:6"},
		"template": []string{"required"},
	}

	messages := govalidator.MapData{
		"phone": []string{
			"required:手机号为必填项",
			"digits:手机号长度错误",
		},
		"code": []string{
			"required:验证码为必填项",
			"digits:验证码长度必须为 6 位的数字",
		},
		"template": []string{
			"required:场景 code为必填项",
		},
	}

	err := request.GoValidate(data, rules, messages)
	if err != "" {
		return err
	}

	_data := data.(*AuthRegisterReq)
	return validators.ValidateVerifyCode(_data.Template, _data.Phone, _data.Code)
}

type AuthRegisterResp struct {
	Token string `json:"token"`
}
