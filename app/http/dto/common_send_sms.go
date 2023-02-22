package dto

import (
	"github.com/thedevsaddam/govalidator"
	"go-yao/common/helpers"
	"go-yao/pkg/global"
	"go-yao/pkg/request"
)

type CommonSendSmsReq struct {
	Phone string `json:"phone,omitempty" valid:"phone"`
	Scene string `json:"scene,omitempty" valid:"scene"`
}

func (s *CommonSendSmsReq) Generate(data interface{}) string {
	rules := govalidator.MapData{
		"phone": []string{"required", "digits:11"},
		"scene": []string{"required"},
	}

	messages := govalidator.MapData{
		"phone": []string{
			"required:手机号为必填项",
			"digits:手机号长度错误",
		},
		"scene": []string{
			"required:应用场景为必填项",
		},
	}

	err := request.GoValidate(data, rules, messages)

	if err != "" {
		return err
	}

	_data := data.(*CommonSendSmsReq)
	if i := helpers.InArray(global.SMSScene, _data.Scene); i == -1 {
		err = "无法申请验证码"
	}

	return err
}
