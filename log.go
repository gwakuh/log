package logger

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
	Logger        *logrus.Logger
}

func NewWithParams(filePath, fileExtension string, maxSize, maxBackups, maxAge int, compress bool) *Logger {
	return &Logger{
		FilePath:      filePath,
		FileExtension: fileExtension,
		MaxSize:       maxSize,
		MaxBackups:    maxBackups,
		MaxAge:        maxAge,
		Compress:      compress,
		Logger:        logrusLogger(lumberjackLogger(filePath+fileExtension, maxSize, maxBackups, maxAge, compress)),
	}
}

func New() *Logger {
	return &Logger{
		Logger: logrusLogger(os.Stdout),
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

func (logger *Logger) Log(level string, msg string, args ...interface{}) {
	l := logger.Logger

	switch level {
	case Info:
		l.Infof(msg+"\n", args...)
	case Warning:
		l.Warnf(msg+"\n", args...)
	case Error:
		l.Errorf(msg+"\n", args...)
	case Debug:
		l.Debugf(msg+"\n", args...)
	}
}
