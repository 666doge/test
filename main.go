package main
import (
	"test/logger"
	"time"
	"test/utils"
	// "net/http"
)

func main() {
	logger.InitLogger("file", map[string]string {
		"log_level": "file",
		"log_path": utils.GetGopath() + "/src/test/logs",
		"log_name": "main",
		"log_split_type": "size",
		"log_split_size": "1024",
	})
	for {
		logger.Debug("this is a debug log called %s", "AChai")
		logger.Info("this is a info log called %s", "AChai")
		logger.Trace("this is a trace log called %s", "AChai")
		logger.Warn("this is a warn log called %s", "AChai")
		logger.Error("this is a error log called %s", "AChai")
		logger.Fatal("this is a fatal log called %s", "AChai")
		time.Sleep(1 * time.Second)
	}
}
