package main
import (
	"test/logger"
	"time"
	"test/utils"
	// "net/http"
)

func main() {
	logger.InitLogger("file", map[string]string {
		"log_level": "debug",
		"log_path": utils.GetGopath() + "/src/test/logs",
		"log_name": "main",
	})
	logger.Debug("this is a debug log called %s", "AChai")
	logger.Info("this is a info log called %s", "AChai")
	logger.Warn("this is a warn log called %s", "AChai")
	logger.Error("this is a error log called %s", "AChai")
	time.Sleep(1 * time.Second)
}