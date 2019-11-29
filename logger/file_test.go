package logger

import (
	"testing"
	"math/rand"
)

func TestFileLogger(t *testing.T) {
	logger := NewFileLogger(LogLevelDebug, "/Users/xushengnan02/go/src/test/logs", "test")
	logger.Debug("a debug log [%d] ", 123)
	logger.Trace("a trace log %d", rand.Intn(1000))
	logger.Info("a Info log %d", rand.Intn(1000))
	logger.Warn("a warn log %s", "诶呦")
	logger.Error("a error log %s", "哇哦")
	logger.Fatal("a fatal log %s", "挂喽")
	logger.Close()
}