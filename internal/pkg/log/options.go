package log

type Options struct {
	// 日志级别, 默认为info
	Level string
	// 是否显示调用者信息(即打印日志的函数名称和行号), 默认为false
	DisableCaller bool
	// 是否打印调用栈信息, 默认为false
	DisableStacktrace bool
	// 日志格式, 默认为console, 可选值为console, json
	Format string
	// 日志输出路径, 默认为stdout, 可选值为stdout, stderr, 文件路径
	OutputPaths []string
}

// NewOptions 创建默认的日志选项
func NewOptions() *Options {
	return &Options{
		Level:             "info",
		DisableCaller:     false,
		DisableStacktrace: false,
		Format:            "console",
		OutputPaths:       []string{"stdout"},
	}
}
