package config

type LogConfig struct {
	Level     string // 日志级别, 开发时推荐使用 "debug" 或者 "info" ，生产环境下使用 "error"
	Type      string // 日志的类型，可选： "single": 独立的文件; "daily": 按照日期每日一个
	Filename  string // 日志文件名跟路径
	MaxSize   int    // 每个日志文件保存的最大尺寸 单位：M
	MaxAge    int    // 最多保存多少天，0 表示不删
	MaxBackup int    // 最多保存日志文件数，0 为不限，MaxAge 到了还是会删
	Compress  bool   // 是否压缩
}
