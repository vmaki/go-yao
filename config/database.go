package config

type DataBaseConfig struct {
	Driver             string
	Host               string
	Port               int
	Database           string
	Username           string
	Password           string
	Charset            string
	MaxOpenConnections int // 最大连接数
	MaxIdleConnections int // 最大空闲连接数
	MaxLifeSeconds     int // 每个链接的过期时间
}
