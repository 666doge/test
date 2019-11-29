package logger

import (
	"os"
	"fmt"
)

type ConsoleLogger struct {
	level int
}

func NewConsoleLogger (config map[string]string) (logger LogInterface, err error) {
	logLevel, ok := config["log_level"]
	if !ok {
		err = fmt.Errorf("config not found : log_level")
		return
	}
	level := GetLogLevel(logLevel)
	logger = &ConsoleLogger{
		level: level,
	}
	return
}

func (c *ConsoleLogger) Init() {}

func (c *ConsoleLogger) SetLevel(level int) {
	if level < LogLevelDebug || level > LogLevelFatal {
		c.level = LogLevelDebug
		return
	}
	c.level = level
}

func (c *ConsoleLogger) Debug(format string, args ...interface{}) {
	if c.level > LogLevelDebug {
		return
	}
	WriteLog(os.Stdout, LogLevelDebug, format, args...)
}

func (c *ConsoleLogger) Trace(format string, args ...interface{}) {
	if c.level > LogLevelTrace {
		return
	}
	WriteLog(os.Stdout, LogLevelTrace, format, args...)
}

func (c *ConsoleLogger) Info(format string, args ...interface{}) {
	if c.level > LogLevelInfo {
		return
	}
	WriteLog(os.Stdout, LogLevelInfo, format, args...)
}

func (c *ConsoleLogger) Warn(format string, args ...interface{}) {
	if c.level > LogLevelWarn {
		return
	}
	WriteLog(os.Stdout, LogLevelWarn, format, args...)
}

func (c *ConsoleLogger) Error(format string, args ...interface{}) {
	if c.level > LogLevelError {
		return
	}
	WriteLog(os.Stdout, LogLevelError, format, args...)
}

func (c *ConsoleLogger) Fatal(format string, args ...interface{}) {
	if c.level > LogLevelFatal {
		return
	}
	WriteLog(os.Stdout, LogLevelFatal, format, args...)
}

func (c *ConsoleLogger) Close() {
}