package drivers

import (
	"os"
	"testing"
	"time"
)

func TestTextFileDriver_Log(t *testing.T) {
	filePath := "test-text-log.log"
	defer os.Remove(filePath) // Clean up

	driver := NewTextFileDriver(DriverConfig{
		LogFilePath:     filePath,
		MaxSize:         10,
		MaxBackups:      1,
		MaxAge:          1,
		TimestampFormat: time.RFC3339,
	})

	entry := LogEntry{
		Message:   "Test log",
		Level:     Info,
		Tags:      map[string]string{"env": "test"},
		Timestamp: time.Now(),
	}

	driver.Log(entry)

	// Check if the file is created
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		t.Errorf("Expected log file to be created, but it doesn't exist")
	}
}
