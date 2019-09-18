package logger

import (
	"os"

	gwformatter "github.com/gwakuh/logrus-gwakuh-formatter"
	"github.com/sirupsen/logrus"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

var (
	logger *logrus.Logger
)

func lumberjackLogger(filePath string, maxSize int, maxBackups int, maxAge int, compress bool) *lumberjack.Logger {
	return &lumberjack.Logger{
		Filename:   filePath,
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
		Compress:   compress,
	}
}

func initToStdout() {
	logger = logrus.New()
	logger.SetFormatter(&gwformatter.Formatter{})
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.TraceLevel) // Trace > Debug > Info > Warn > Error > Fatal > Panic
}

func initToFile() {
	logger = logrus.New()
	logger.SetFormatter(&gwformatter.Formatter{})
	out := lumberjackLogger(Config.Log.FilePath+Config.Log.FileExtension, Config.Log.MaxSize, Config.Log.MaxBackups, Config.Log.MaxAge, Config.Log.Compress)
	logger.SetOutput(out)
	logger.SetLevel(logrus.TraceLevel) // Trace > Debug > Info > Warn > Error > Fatal > Panic
}

func init() {
	initToFile()
}

func Log(level string, msg string, args ...interface{}) {
	switch level {
	case Info:
		logger.Infof(msg+"\n", args...)
	case Warning:
		logger.Warnf(msg+"\n", args...)
	case Error:
		logger.Errorf(msg+"\n", args...)
	case Debug:
		logger.Debugf(msg+"\n", args...)
	}
}
