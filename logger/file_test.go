package logger

import (
	"testing"
)

func TestFileLogger(t *testing.T) {
	logger := NewFileLogger(LogLevelDebug, "/Users/xushengnan/go/src/test/logs", "test")
	logger.Debug("user id [%d] ", 123)
	logger.Warn("a warn %s", "msg")
	logger.Close()
}