package logger
import (
	"fmt"
)

var logger LogInterface

func InitLogger(name string, config map[string]string) (err error) {
	switch name {
	case "file":
		logger, err = NewFileLogger(config)
	case "console":
		logger, err = NewConsoleLogger(config)
	default:
		err = fmt.Errorf("unsupport log name: %s", name)
	}
	return
}

func Debug(format string, args ...interface{}) {
	logger.Debug(format, args...)
}