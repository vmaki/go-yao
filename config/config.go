package config

type AppConfig struct {
	Application ApplicationConfig
	Log         LogConfig
	DataBase    DataBaseConfig
	Redis       RedisConfig
	Sms         SmsConfig
}
