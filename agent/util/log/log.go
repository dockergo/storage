package log

import (
	"fmt"
	"github.com/flyaways/storage/agent/util/log/log"
)

var logger *log.Logger = log.StdLogger()

func Init(strLevel string, path string) error {
	if path != "" {
		handler, err := log.NewTimeRotatingFileHandler(path, log.WhenDay, 1)
		if err != nil {
			return err
		}
		logger = log.NewDefault(handler)
	}
	var level int
	switch strLevel {
	case "trace":
		level = log.LevelTrace
	case "debug":
		level = log.LevelDebug
	case "info":
		level = log.LevelInfo
	case "warn":
		level = log.LevelWarn
	case "error":
		level = log.LevelError
	case "fatal":
		level = log.LevelFatal
	default:
		return fmt.Errorf("invalid log level:%s", strLevel)
	}
	logger.SetLevel(level)
	return nil
}

func SetLevel(level int) {
	logger.SetLevel(level)
}

func Trace(format string, v ...interface{}) {
	logger.Output(2, log.LevelTrace, format, v...)
}

func Debug(format string, v ...interface{}) {
	logger.Output(2, log.LevelDebug, format, v...)
}

func Info(format string, v ...interface{}) {
	logger.Output(2, log.LevelInfo, format, v...)
}

func Warn(format string, v ...interface{}) {
	logger.Output(2, log.LevelWarn, format, v...)
}

func Error(format string, v ...interface{}) {
	logger.Output(2, log.LevelError, format, v...)
}

func Fatal(format string, v ...interface{}) {
	logger.Output(2, log.LevelFatal, format, v...)
}
