package log

import (
	"io"
	"os"

	_ "github.com/gogap/logrus_mate/writers/rotatelogs"
	"github.com/sirupsen/logrus"
)

var DefaultLogger = newDefaultLogger()

func newDefaultLogger() ILogger {
	rus := logrus.New()
	rus.SetLevel(logrus.DebugLevel)
	rus.SetOutput(os.Stdout)
	rus.SetFormatter(&logrus.JSONFormatter{})
	return &Logger{
		log: rus,
	}
}

type ILogger interface {
	Debug(Fields, interface{})
	Error(Fields, interface{})
	Info(Fields, interface{})
	Warn(Fields, interface{})
	WithErr(error, interface{})
	Debugf(string, ...interface{})
	Errorf(string, ...interface{})
	Infof(string, ...interface{})
	Warnf(string, ...interface{})

	IoWriter() io.Writer
}

type Logger struct {
	log *logrus.Logger
}

func (f *Logger) Debug(fields Fields, message interface{}) {
	f.log.WithFields(logrus.Fields(fields)).Debugln(message)
}
func (f *Logger) Error(fields Fields, message interface{}) {
	f.log.WithFields(logrus.Fields(fields)).Errorln(message)
}
func (f *Logger) Info(fields Fields, message interface{}) {
	f.log.WithFields(logrus.Fields(fields)).Infoln(message)
}
func (f *Logger) Warn(fields Fields, message interface{}) {
	f.log.WithFields(logrus.Fields(fields)).Warnln(message)
}
func (f *Logger) WithErr(err error, message interface{}) {
	f.log.WithError(err).Errorln(message)
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
	f.log.Infof(format, message)

}

func (f *Logger) IoWriter() io.Writer {
	return f.log.Writer()
}

type Fields map[string]interface{}
