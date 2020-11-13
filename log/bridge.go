package log

import (
	"fmt"
	"io"
	"runtime"
)

func Debug(fields Fields, message interface{}) {
	fields["caller"] = printCallerName()
	DefaultLogger.Debug(fields, message)
}
func Error(fields Fields, message interface{}) {
	fields["caller"] = printCallerName()
	DefaultLogger.Error(fields, message)
}
func Info(fields Fields, message interface{}) {
	fields["caller"] = printCallerName()
	DefaultLogger.Info(fields, message)
}
func Warn(fields Fields, message interface{}) {
	fields["caller"] = printCallerName()
	DefaultLogger.Warn(fields, message)
}
func WithErr(err error, message interface{}) {
	DefaultLogger.Error(Fields{
		"err":    err,
		"caller": printCallerName(),
	}, message)
}

func Debugf(fields string, message ...interface{}) {
	DefaultLogger.Debugf(fields, message...)
}
func Errorf(fields string, message ...interface{}) {
	DefaultLogger.Errorf(fields, message...)
}
func Infof(fields string, message ...interface{}) {
	DefaultLogger.Infof(fields, message...)
}
func Warnf(fields string, message ...interface{}) {
	DefaultLogger.Warnf(fields, message...)
}
func IoWriter() io.Writer {
	return DefaultLogger.IoWriter()
}
func printCallerName() string {
	pc, _, _, _ := runtime.Caller(2)
	fn := runtime.FuncForPC(pc)
	filename, line := fn.FileLine(pc)

	return filename + ",line=" + fmt.Sprintf("%d", line)
}
