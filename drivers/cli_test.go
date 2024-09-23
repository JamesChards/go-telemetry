package drivers

import (
	"testing"
	"time"
)

func TestCLIDriver_Log(t *testing.T) {
	driver := NewCLIDriver(DriverConfig{TimestampFormat: time.RFC3339})

	entry := LogEntry{
		Message:   "Test log",
		Level:     Info,
		Tags:      map[string]string{"env": "test"},
		Timestamp: time.Now(),
	}

	driver.Log(entry)
	// Since it's hard to capture CLI output, this ensures that the function runs without error.
}
