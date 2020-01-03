package session

import (
	"testing"
)

func TestSession(t *testing.T) {
	mS = MemorySession{
		data: "a memory session",
		id: "1"
	}
}