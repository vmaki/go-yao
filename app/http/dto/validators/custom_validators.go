package validators

import "go-yao/common/verifycode"

// ValidateVerifyCode 自定义规则，验证验证码是否正确
func ValidateVerifyCode(template, phone, answer string) string {
	if ok := verifycode.NewVerifyCode().CheckAnswer(template, phone, answer); !ok {
		return "验证码错误"
	}

	return ""
}
