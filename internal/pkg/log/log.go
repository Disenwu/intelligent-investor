package log

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var _ Logger = (*ZapLogger)(nil)

var logger *ZapLogger

func Init() {
	opts := &Options{
		Level:             viper.GetString("log.level"),
		Format:            viper.GetString("log.format"),
		OutputPaths:       viper.GetStringSlice("log.output-paths"),
		DisableCaller:     viper.GetBool("log.disable-caller"),
		DisableStacktrace: viper.GetBool("log.disable-stacktrace"),
	}
	s, _ := json.Marshal(opts)
	fmt.Printf("init zap logger: %s", string(s))
	logger = NewZapLogger(opts)
}

// Logger 日志接口, 定义了日志的基本操作
type Logger interface {
	// Debugw 打印调试日志
	Debugw(msg string, keysAndValues ...any)
	// Infow 打印信息日志, 用于记录普通的运行信息
	Infow(msg string, keysAndValues ...any)
	// Warnw 打印警告日志, 记录潜在问题
	Warnw(msg string, keysAndValues ...any)
	// Errorw 打印错误日志, 记录运行时错误，需要立即修复
	Errorw(msg string, keysAndValues ...any)
	// Panicw 打印严重错误的日志,表示系统无法继续运行， 打印后会触发程序Panic
	Panicw(msg string, keysAndValues ...any)
	// Fatalw 打印致命错误的日志,表示系统无法继续运行， 打印后会触发程序退出
	Fatalw(msg string, keysAndValues ...any)
	// Sync 刷新日志缓冲区, 确保日志完整记录
	Sync()
}

type ZapLogger struct {
	logger *zap.Logger
}

func NewZapLogger(opts *Options) *ZapLogger {
	if opts == nil {
		opts = NewOptions()
	}
	var zapLevel zapcore.Level
	if err := zapLevel.UnmarshalText([]byte(opts.Level)); err != nil {
		zapLevel = zapcore.InfoLevel
	}
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.MessageKey = "message"
	encoderConfig.TimeKey = "timestamp"
	//时间格式为2006-01-02 15:04:05.000
	encoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
	}
	//毫秒
	encoderConfig.EncodeDuration = func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendFloat64(float64(d) / float64(time.Second))
	}

	cfg := zap.Config{
		Level:             zap.NewAtomicLevelAt(zapLevel),
		DisableCaller:     opts.DisableCaller,
		DisableStacktrace: opts.DisableStacktrace,
		Encoding:          opts.Format,
		EncoderConfig:     encoderConfig,
		OutputPaths:       opts.OutputPaths,
		ErrorOutputPaths:  []string{"stderr"},
	}

	logger, err := cfg.Build(zap.AddStacktrace(zapcore.PanicLevel), zap.AddCallerSkip(2))
	if err != nil {
		panic(err)
	}
	zap.RedirectStdLog(logger)
	return &ZapLogger{logger: logger}
}

func Debugw(msg string, keysAndValues ...any) {
	logger.Debugw(msg, keysAndValues...)
}

func (l *ZapLogger) Debugw(msg string, keysAndValues ...any) {
	l.logger.Sugar().Debugw(msg, keysAndValues...)
}

func Infow(msg string, keysAndValues ...any) {
	logger.Infow(msg, keysAndValues...)
}

func (l *ZapLogger) Infow(msg string, keysAndValues ...any) {
	l.logger.Sugar().Infow(msg, keysAndValues...)
}

func Warnw(msg string, keysAndValues ...any) {
	logger.Warnw(msg, keysAndValues...)
}

func (l *ZapLogger) Warnw(msg string, keysAndValues ...any) {
	l.logger.Sugar().Warnw(msg, keysAndValues...)
}

func Errorw(msg string, keysAndValues ...any) {
	logger.Errorw(msg, keysAndValues...)
}

func (l *ZapLogger) Errorw(msg string, keysAndValues ...any) {
	l.logger.Sugar().Errorw(msg, keysAndValues...)
}

func Panicw(msg string, keysAndValues ...any) {
	logger.Panicw(msg, keysAndValues...)
}

func (l *ZapLogger) Panicw(msg string, keysAndValues ...any) {
	l.logger.Sugar().Panicw(msg, keysAndValues...)
}

func Fatalw(msg string, keysAndValues ...any) {
	logger.Fatalw(msg, keysAndValues...)
}

func (l *ZapLogger) Fatalw(msg string, keysAndValues ...any) {
	l.logger.Sugar().Fatalw(msg, keysAndValues...)
}

func Sync() {
	logger.Sync()
}

func (l *ZapLogger) Sync() {
	l.logger.Sync()
}
