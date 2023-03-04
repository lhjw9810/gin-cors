package log

import (
	"io"
)

func Debug(fields Fields, message string) {
	DefaultLogger.Debug(fields, message)
}
func Error(fields Fields, message string) {
	DefaultLogger.Error(fields, message)
}
func Info(fields Fields, message string) {
	DefaultLogger.Info(fields, message)
}
func Warn(fields Fields, message string) {
	DefaultLogger.Warn(fields, message)
}
func WithErr(err error, message string) {
	DefaultLogger.WithErr(err, message)
}

func Debugf(fields string, message ...string) {
	DefaultLogger.Debugf(fields, message...)
}
func Errorf(fields string, message ...string) {
	DefaultLogger.Errorf(fields, message...)
}
func Infof(fields string, message ...string) {
	DefaultLogger.Infof(fields, message...)
}
func Warnf(fields string, message ...string) {
	DefaultLogger.Warnf(fields, message...)
}
func IoWriter() io.Writer {
	return DefaultLogger.IoWriter()
}
func Sync() error {
	return DefaultLogger.Sync()
}
