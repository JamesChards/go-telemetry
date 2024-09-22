package telemetry

import (
	"encoding/json"
	"example/telemetry/drivers"
	"io"
	"os"
	"time"
)

// Driver is the interface that must be implemented by all log drivers.
type Driver interface {
	Log(message string, level drivers.LogLevel, tags map[string]string)
}

// Logger holds the logger configuration and the active driver.
type Logger struct {
	driver Driver
	tags   map[string]string
}

func CreateDefaultConfig() error {
	defaultConfig := DefaultConfig()
	bytes, _ := json.MarshalIndent(defaultConfig, "", "  ")

	if _, err := os.Stat(CONFIG_FILE); err != nil {
		return os.WriteFile(CONFIG_FILE, bytes, 0644)
	}
	_, err := LoadConfig(CONFIG_FILE)
	if err != nil {
		return os.WriteFile(CONFIG_FILE, bytes, 0644)
	}

	return nil
}

// NewLogger initializes the Logger instance with a specific driver.
func NewLogger() *Logger {
	config := DefaultConfig()

	// Check if the configuration file exists
	if _, err := os.Stat(CONFIG_FILE); err == nil {
		loadedConfig, cerr := LoadConfig(CONFIG_FILE)
		if cerr == nil {
			config = loadedConfig
		}
	}

	var driver Driver

	switch config.DefaultDriver {
	case "json":
		driver, _ = drivers.NewJSONDriver(config.LogFilePath)
	case "text":
		driver, _ = drivers.NewTextFileDriver(config.LogFilePath)
	default:
		driver = drivers.NewCLIDriver()
	}

	return &Logger{
		driver: driver,
		tags:   make(map[string]string),
	}
}

func (t *Logger) IsCloser() (io.Closer, bool) {
	closer, ok := t.driver.(io.Closer)
	return closer, ok
}

// SetDriverWithName assign a new driver with a name to the Logger instance.
func (t *Logger) SetDriverWithName(driver string) {
	config := DefaultConfig()

	// Check if the configuration file exists
	if _, err := os.Stat(CONFIG_FILE); err == nil {
		loadedConfig, cerr := LoadConfig(CONFIG_FILE)
		if cerr == nil {
			config = loadedConfig
		}
	}

	var newDriver Driver

	switch driver {
	case "json":
		newDriver, _ = drivers.NewJSONDriver(config.LogFilePath)
	case "text":
		newDriver, _ = drivers.NewTextFileDriver(config.LogFilePath)
	case "mock":
		newDriver = NewMockDriver()
	default:
		newDriver = drivers.NewCLIDriver()
	}
	t.driver = newDriver
}

// SetDriver assign a new driver to the Logger instance.
func (t *Logger) SetDriver(driver Driver) {
	t.driver = driver
}

// Log writes a log message with the given level and tags.
func (t *Logger) Log(level drivers.LogLevel, message string) {
	// Create a copy of the tags map
	tags := make(map[string]string)
	for key, value := range t.tags {
		tags[key] = value
	}

	// Add the timestamp to the copied map
	tags["timestamp"] = time.Now().Format(time.RFC3339)

	// Pass the copied map to the driver's Log function
	t.driver.Log(message, level, tags)
}

// Debug writes a log message with the Debug level and tags.
func (t *Logger) Debug(message string) {
	t.Log(drivers.Debug, message)
}

// Info writes a log message with the Info level and tags.
func (t *Logger) Info(message string) {
	t.Log(drivers.Info, message)
}

// Warning writes a log message with the Warning level and tags.
func (t *Logger) Warning(message string) {
	t.Log(drivers.Warning, message)
}

// Error writes a log message with the Error level and tags.
func (t *Logger) Error(message string) {
	t.Log(drivers.Error, message)
}

// SetTags allows adding global tags to the Logger instance.
func (t *Logger) SetTags(tags map[string]string) {
	t.tags = map[string]string{}

	for key, value := range tags {
		t.tags[key] = value
	}
}
