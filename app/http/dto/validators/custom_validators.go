package validators

import (
	"go-yao/common/response"
	"go-yao/common/verifycode"
)

// ValidateVerifyCode 自定义规则，验证验证码是否正确
func ValidateVerifyCode(template, phone, answer string) error {
	if ok := verifycode.NewVerifyCode().CheckAnswer(template, phone, answer); !ok {
		return response.New(response.CodeVerifyCodeErr)
	}

	return nil
}
