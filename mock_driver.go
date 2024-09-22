package telemetry

import "example/telemetry/drivers"

// MockDriver is a mock implementation of the Driver interface for testing.
type MockDriver struct {
	Logs []LogEntry
}

type LogEntry struct {
	Message string
	Level   drivers.LogLevel
	Tags    map[string]string
}

// NewMockDriver creates a new instance of the mock driver.
func NewMockDriver() *MockDriver {
	return &MockDriver{
		Logs: []LogEntry{},
	}
}

// Log implements the Driver interface for MockDriver.
func (d *MockDriver) Log(message string, level drivers.LogLevel, tags map[string]string) {
	d.Logs = append(d.Logs, LogEntry{
		Message: message,
		Level:   level,
		Tags:    tags,
	})
}
