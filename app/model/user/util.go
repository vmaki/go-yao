package user

import "go-yao/pkg/global"

// IsPhoneExist 判断手机号已被注册
func IsPhoneExist(phone string) bool {
	var count int64
	global.DB.Model(User{}).Where("phone = ?", phone).Count(&count)
	return count > 0
}
