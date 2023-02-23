package dto

import (
	"github.com/thedevsaddam/govalidator"
	"go-yao/app/http/dto/validators"
	"go-yao/pkg/request"
)

type AuthLoginReq struct {
	Phone    string `json:"phone,omitempty" valid:"phone"`       // 手机号码
	Code     string `json:"code,omitempty" valid:"code"`         // 短信验证码
	Template string `json:"template,omitempty" valid:"template"` // 短信场景码
}

func (s *AuthLoginReq) Generate(data interface{}) string {
	rules := govalidator.MapData{
		"phone":    []string{"required", "digits:11"},
		"code":     []string{"required", "digits:6"},
		"template": []string{"required"},
	}

	messages := govalidator.MapData{
		"phone": []string{
			"required:手机号为必填项",
			"digits:手机号长度必须为 11 位的数字",
		},
		"code": []string{
			"required:验证码为必填项",
			"digits:验证码长度必须为 6 位的数字",
		},
		"template": []string{
			"required:场景为必填项",
		},
	}

	err := request.GoValidate(data, rules, messages)
	if err != "" {
		return err
	}

	_data := data.(*AuthLoginReq)

	return validators.ValidateVerifyCode(_data.Template, _data.Phone, _data.Code)
}

type AuthLoginResp struct {
	Token string `json:"token"` // jwt-token
}
