package log

import (
	"fmt"

	"github.com/gogap/logrus_mate"
	"github.com/sirupsen/logrus"
)

var logConfig = `release {
        level = "%s"
        formatter.name = "%s"
        formatter.options  {
                            force-colors      = false
                            disable-colors    = false
                            disable-timestamp = false
                            full-timestamp    = false
                            timestamp-format  = "2006-01-02 15:04:05"
                            disable-sorting   = false
        }
 		out.name = "%s"
        out.options {
                    path =  %s
                    link-name= %s
        }
}`

type Level string

type Formatter string

type Option func(c *Options)

const (
	PanicLevel Level = "panic"
	FatalLevel Level = "fatal"
	ErrorLevel Level = "error"
	WarnLevel  Level = "warn"
	InfoLevel  Level = "info"
	DebugLevel Level = "debug"
	TraceLevel Level = "trace"

	TextFormatter = "text"
	JsonFormatter = "json"
)

var DefaultOptions = Options{
	Level:     DebugLevel,
	Formatter: JsonFormatter,
	Out: OutputOption{
		Name:     "rotatelogs",
		Path:     "./logs/%Y%m%d.log",
		LinkName: "./logs/current.log",
	},
}

type Options struct {
	Level     Level
	Formatter Formatter
	Out       OutputOption
}
type OutputOption struct {
	Path     string
	LinkName string
	Name     string
}

func WithLevel(level Level) Option {
	return func(c *Options) {
		c.Level = level
	}
}

func WithFormatter(f Formatter) Option {
	return func(c *Options) {
		c.Formatter = f
	}
}
func WithOutputOption(f OutputOption) Option {
	return func(c *Options) {
		c.Out = f
	}
}

//设置高级日志
func UseAdvanceOptions(opts ...Option) {
	defaultOpts := &DefaultOptions
	if len(opts) > 0 {
		for _, v := range opts {
			v(defaultOpts)
		}
	}
	useAdvanceLogger(*defaultOpts)
}

func useAdvanceLogger(opts Options) {
	confString := fmt.Sprintf(logConfig, opts.Level, opts.Formatter, opts.Out.Name,opts.Out.Path, opts.Out.LinkName)
	mate, _ := logrus_mate.NewLogrusMate(
		logrus_mate.ConfigString(
			confString,
		),
	)
	lgrus := logrus.StandardLogger()
	mate.Hijack(
		lgrus,
		"release",
	)

	DefaultLogger = &Logger{log: lgrus}
}
