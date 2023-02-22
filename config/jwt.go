package config

type JWTConfig struct {
	Secret         string
	ExpireTime     int64
	MaxRefreshTime int64
}
