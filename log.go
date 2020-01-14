package log

import (
	"io"
	"os"

	gwformatter "github.com/gwakuh/logrus-gwakuh-formatter"
	"github.com/sirupsen/logrus"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

type Logger struct {
	FilePath      string `default:"log/app"`
	FileExtension string `default:".log"`
	MaxSize       int    `default:"100"` // megabytes
	MaxBackups    int    `default:"7"`
	MaxAge        int    `default:"30"` // days
	Compress      bool   `default:"false"`
	Logrus        *logrus.Logger
}

func NewWithParams(filePath, fileExtension string, maxSize, maxBackups, maxAge int, compress bool) *Logger {
	return &Logger{
		FilePath:      filePath,
		FileExtension: fileExtension,
		MaxSize:       maxSize,
		MaxBackups:    maxBackups,
		MaxAge:        maxAge,
		Compress:      compress,
		Logrus:        logrusLogger(lumberjackLogger(filePath+fileExtension, maxSize, maxBackups, maxAge, compress)),
	}
}

func New() *Logger {
	return &Logger{
		Logrus: logrusLogger(os.Stdout),
	}
}

func logrusLogger(out io.Writer) *logrus.Logger {
	return &logrus.Logger{
		Formatter: &gwformatter.Formatter{},
		Out:       out,
		Level:     logrus.TraceLevel, // Trace > Debug > Info > Warn > Error > Fatal > Panic
	}
}

func lumberjackLogger(filePath string, maxSize int, maxBackups int, maxAge int, compress bool) *lumberjack.Logger {
	return &lumberjack.Logger{
		Filename:   filePath,
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
		Compress:   compress,
	}
}

func (l *Logger) Info(msg string, args ...interface{}) {
	l.Logrus.Infof(msg+"\n", args...)
}

func (l *Logger) Warn(msg string, args ...interface{}) {
	l.Logrus.Warnf(msg+"\n", args...)
}

func (l *Logger) Error(msg string, args ...interface{}) {
	l.Logrus.Errorf(msg+"\n", args...)
}

func (l *Logger) Debug(msg string, args ...interface{}) {
	l.Logrus.Debugf(msg+"\n", args...)
}
