package cache

import "fmt"

var (
	SmsCache = "sms" // 预约单
)

// GetSmsCacheKey  获取验证码缓存 key
func GetSmsCacheKey(scene, phone string) string {
	return fmt.Sprintf("%s:%s:%s", SmsCache, scene, phone)
}
