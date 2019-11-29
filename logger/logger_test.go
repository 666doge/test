package logger

import (
	"testing"
)

func TestLogger (t *testing.T) {
	InitLogger("console", map[string]string {
		"log_level": "debug",
	})
}
