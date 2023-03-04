package log

import (
	"io"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Level string

type Option func(c *Options)

const (
	PanicLevel Level = "panic"
	FatalLevel Level = "fatal"
	ErrorLevel Level = "error"
	WarnLevel  Level = "warn"
	InfoLevel  Level = "info"
	DebugLevel Level = "debug"
)

type Options struct {
	Level         Level
	Out           *OutputOption
	AddCaller     bool
	AddCallerSkip int
}
type OutputOption struct {
	OutPath   io.Writer
	ErrorPath io.Writer
}

func WithLevel(level Level) Option {
	return func(c *Options) {
		c.Level = level
	}
}
func AddCaller() Option {
	return func(c *Options) {
		c.AddCaller = true
	}
}

func AddCallerSkip(skipStep int) Option {
	return func(c *Options) {
		c.AddCallerSkip = skipStep
	}
}

func WithOutputOption(f OutputOption) Option {
	return func(c *Options) {
		c.Out = &f
	}
}

// WithRotate 是否使用滚动日志
// size 最大size，默认1M
// day 最长保留天数
// maxBackup 最多保留日志文件数量
// file 文件路径
func WithRotate(size, day, maxBackup int, file string) Option {
	return func(c *Options) {
		c.Out.OutPath = &lumberjack.Logger{
			Filename:   file,
			MaxSize:    size, // megabytes
			MaxBackups: maxBackup,
			MaxAge:     day, //days
		}
	}
}

// WithErrorRotate 是否使用滚动日志
// size 最大size，默认1M
// day 最长保留天数
// maxBackup 最多保留日志文件数量
// file 文件路径
func WithErrorRotate(size, day, maxBackup int, file string) Option {
	return func(c *Options) {
		c.Out.ErrorPath = &lumberjack.Logger{
			Filename:   file,
			MaxSize:    size, // megabytes
			MaxBackups: maxBackup,
			MaxAge:     day, //days
		}
	}
}

// Production 设置高级日志
func Production(opts ...Option) {
	defaultOpts := &Options{
		Level: DebugLevel,
		Out: &OutputOption{
			OutPath:   os.Stderr,
			ErrorPath: os.Stderr,
		},
	}
	if len(opts) > 0 {
		for _, v := range opts {
			v(defaultOpts)
		}
	}
	hijack(*defaultOpts)
}

func hijack(opts Options) {
	l, _ := zapcore.ParseLevel(string(opts.Level))
	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= l && lvl < zapcore.ErrorLevel
	})
	enc := zap.NewProductionEncoderConfig()
	enc.EncodeTime = zapcore.RFC3339TimeEncoder

	outEncoder := zapcore.NewJSONEncoder(enc)

	c := zapcore.NewCore(outEncoder, zapcore.AddSync(opts.Out.OutPath), lowPriority)

	errEncoder := zapcore.NewJSONEncoder(enc)

	e := zapcore.NewCore(errEncoder, zapcore.AddSync(opts.Out.ErrorPath), highPriority)
	core := zapcore.NewTee(c, e)
	sugar := zap.New(core).Sugar()
	if opts.AddCaller {
		sugar = sugar.WithOptions(zap.AddCaller(), zap.AddCallerSkip(opts.AddCallerSkip))
	}
	DefaultLogger = &Logger{log: sugar, ws: opts.Out.OutPath}
}
