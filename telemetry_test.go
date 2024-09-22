package telemetry

import (
	"example/telemetry/drivers"
	"testing"
	"time"
)

// Test basic logging functionality.
func TestTelemetry_Log(t *testing.T) {
	mockDriver := NewMockDriver()
	tel := NewLogger()
	tel.SetDriver(mockDriver)

	// Set global tags
	tel.SetTags(map[string]string{
		"origin":     "http",
		"customerId": "123",
	})

	// Log a message
	tel.Info("Test log entry")

	// Assert that log was recorded
	if len(mockDriver.Logs) != 1 {
		t.Errorf("expected 1 log entry, got %d", len(mockDriver.Logs))
	}

	// Validate log content
	log := mockDriver.Logs[0]
	if log.Message != "Test log entry" {
		t.Errorf("expected log message 'Test log entry', got '%s'", log.Message)
	}
	if log.Level != drivers.Info {
		t.Errorf("expected log level Info, got %v", log.Level)
	}

	// Validate tags
	if log.Tags["origin"] != "http" {
		t.Errorf("expected origin tag 'http', got '%s'", log.Tags["origin"])
	}
	if log.Tags["customerId"] != "123" {
		t.Errorf("expected customerId '123', got '%s'", log.Tags["customerId"])
	}
	if _, ok := log.Tags["timestamp"]; !ok {
		t.Errorf("expected log to have timestamp, but it doesn't")
	}
}

// Test that each log call gets a unique timestamp.
func TestTelemetry_LogUniqueTimestamps(t *testing.T) {
	mockDriver := NewMockDriver()
	tel := NewLogger()
	tel.SetDriver(mockDriver)

	tel.Info("First log")
	time.Sleep(1 * time.Second)
	tel.Info("Second log")

	if len(mockDriver.Logs) != 2 {
		t.Fatalf("expected 2 log entries, got %d", len(mockDriver.Logs))
	}

	if mockDriver.Logs[0].Tags["timestamp"] == mockDriver.Logs[1].Tags["timestamp"] {
		t.Error("expected different timestamps for logs, but they are the same")
	}
}
