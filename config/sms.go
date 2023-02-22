package config

type SmsConfig struct {
	Aliyun AliyunConfig
}

type AliyunConfig struct {
	AccessKeyId     string
	AccessKeySecret string
	SignName        string
}
