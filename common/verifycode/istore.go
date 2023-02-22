package verifycode

type IStore interface {
	Set(key string, value string, expiration int64) bool // Set 保存验证码
	Get(key string, clear bool) string                   // Get 获取验证码
	Verify(key, answer string, clear bool) bool          // Verify 检查验证码
}
