package log

import (
	"fmt"
)

var logger *Logger = StdLogger()

func Init(strLevel string, path string) error {
	if path != "" {
		handler, err := NewTimeRotatingFileHandler(path, WhenDay, 1)
		if err != nil {
			return err
		}
		logger = NewDefault(handler)
	}
	var level int
	switch strLevel {
	case "trace":
		level = LevelTrace
	case "debug":
		level = LevelDebug
	case "info":
		level = LevelInfo
	case "warn":
		level = LevelWarn
	case "error":
		level = LevelError
	case "fatal":
		level = LevelFatal
	default:
		return fmt.Errorf("invalid log level:%s", strLevel)
	}
	logger.setLevel(level)
	return nil
}

func SetLevel(level int) {
	logger.setLevel(level)
}

func Trace(format string, v ...interface{}) {
	logger.Output(2, LevelTrace, format, v...)
}

func Debug(format string, v ...interface{}) {
	logger.Output(2, LevelDebug, format, v...)
}

func Info(format string, v ...interface{}) {
	logger.Output(2, LevelInfo, format, v...)
}

func Warn(format string, v ...interface{}) {
	logger.Output(2, LevelWarn, format, v...)
}

func Error(format string, v ...interface{}) {
	logger.Output(2, LevelError, format, v...)
}

func Fatal(format string, v ...interface{}) {
	logger.Output(2, LevelFatal, format, v...)
}
