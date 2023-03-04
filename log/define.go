package log

import (
	"io"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var DefaultLogger = newDefaultLogger()

func newDefaultLogger() ILogger {
	encode := zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig())
	writer := zapcore.AddSync(os.Stderr)
	level := zapcore.DebugLevel
	zlog := zap.New(zapcore.NewCore(encode, writer, level)).WithOptions(zap.AddCaller(), zap.AddCallerSkip(2))
	return &Logger{log: zlog.Sugar(), ws: writer}
}

type ILogger interface {
	Debug(Fields, string)
	Error(Fields, string)
	Info(Fields, string)
	Warn(Fields, string)
	WithErr(error, string)
	Debugf(string, ...interface{})
	Errorf(string, ...interface{})
	Infof(string, ...interface{})
	Warnf(string, ...interface{})
	IoWriter() io.Writer
	Sync() error
}

type Logger struct {
	log *zap.SugaredLogger
	ws  io.Writer
}

func (f *Logger) keyAndValues(fields Fields) []any {
	m := make([]any, 0, len(fields))
	for k, v := range fields {
		m = append(m, k)
		m = append(m, v)
	}
	return m
}
func (f *Logger) Debug(fields Fields, message string) {
	f.log.Debugw(message, f.keyAndValues(fields)...)
}
func (f *Logger) Error(fields Fields, message string) {
	f.log.Errorw(message, f.keyAndValues(fields)...)
}
func (f *Logger) Info(fields Fields, message string) {
	f.log.Infow(message, f.keyAndValues(fields)...)
}
func (f *Logger) Warn(fields Fields, message string) {
	f.log.Warnw(message, f.keyAndValues(fields)...)
}
func (f *Logger) WithErr(err error, message string) {
	f.log.Errorw(message, err)
}

func (f *Logger) Debugf(format string, message ...interface{}) {
	f.log.Debugf(format, message)
}
func (f *Logger) Errorf(format string, message ...interface{}) {
	f.log.Errorf(format, message)

}
func (f *Logger) Infof(format string, message ...interface{}) {
	f.log.Infof(format, message)

}
func (f *Logger) Warnf(format string, message ...interface{}) {
	f.log.Warnf(format, message)

}

func (f *Logger) IoWriter() io.Writer {
	return f.ws
}
func (f *Logger) Sync() error {
	return f.log.Sync()
}

type Fields map[string]interface{}
