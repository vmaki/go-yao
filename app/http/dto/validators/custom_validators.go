package validators

import "go-yao/common/verifycode"

// ValidateVerifyCode 自定义规则，验证『手机/邮箱验证码』
func ValidateVerifyCode(template, phone, answer string) string {
	if ok := verifycode.NewVerifyCode().CheckAnswer(template, phone, answer); !ok {
		return "验证码错误"
	}

	return ""
}
