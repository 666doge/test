package logger

import (
	"testing"
	"math/rand"
)

func TestFileLogger(t *testing.T) {
	logger,_ := NewFileLogger(map[string]string{
		"log_level": "debug",
		"log_path": "/Users/xushengnan02/go/src/test/logs",
		"log_name": "test",
	})
	logger.Debug("a debug log [%d] ", 123)
	logger.Trace("a trace log %d", rand.Intn(1000))
	logger.Info("a Info log %d", rand.Intn(1000))
	logger.Warn("a warn log %s", "诶呦")
	logger.Error("a error log %s", "哇哦")
	logger.Fatal("a fatal log %s", "挂喽")
	logger.Close()
}

func TestConsoleLogger(t *testing.T) {
	logger, _ := NewConsoleLogger(map[string]string{
		"log_level": "debug",
	})
	logger.Debug("a debug log [%d] ", 123)
	logger.Trace("a trace log %d", rand.Intn(1000))
	logger.Info("a Info log %d", rand.Intn(1000))
	logger.Warn("a warn log %s", "诶呦")
	logger.Error("a error log %s", "哇哦")
	logger.Fatal("a fatal log %s", "挂喽")
	logger.Close()
}