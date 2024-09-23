package telemetry

import (
	"example/telemetry/drivers"
	"io"
	"time"
)

// LogManager holds the logger configuration and the active driver.
type LogManager struct {
	driver drivers.Driver
	tags   map[string]string
}

// NewLogger initializes the LogManager instance with a specific driver.
func NewLogger(driverName ...string) *LogManager {
	config := defaultConfig()

	// Check if the configuration file is set correctly.
	loadedConfig, err := LoadConfig(CONFIG_FILE)
	if err == nil {
		config = loadedConfig
	}

	defaultDriver := config.DefaultDriver
	if len(driverName) > 0 {
		defaultDriver = driverName[0]
	}

	var driver drivers.Driver

	switch defaultDriver {
	case "text":
		driver = drivers.NewTextFileDriver(config.Drivers.Text)
	case "json":
		driver = drivers.NewJSONDriver(config.Drivers.Json)
	default:
		driver = drivers.NewCLIDriver(config.Drivers.Cli)
	}

	return &LogManager{
		driver: driver,
		tags:   make(map[string]string),
	}
}

func (t *LogManager) ReloadConfig(filepath string) error {
	// Reload configuration from the file
	newConfig, err := LoadConfig(filepath)
	if err != nil {
		return err
	}

	var newDriver drivers.Driver
	// Set the driver based on the new configuration
	switch newConfig.DefaultDriver {
	case "text":
		newDriver = drivers.NewTextFileDriver(newConfig.Drivers.Text)
	case "json":
		newDriver = drivers.NewJSONDriver(newConfig.Drivers.Json)
	default:
		newDriver = drivers.NewCLIDriver(newConfig.Drivers.Cli)
	}
	t.setDriver(newDriver)

	return nil
}

func (t *LogManager) IsCloser() (io.Closer, bool) {
	closer, ok := t.driver.(io.Closer)
	return closer, ok
}

// SetDriverWithName assign a new driver with a name to the LogManager instance.
func (t *LogManager) SetDriverWithName(driver string) {
	config := defaultConfig()

	// Check if the configuration file is set correctly.
	loadedConfig, err := LoadConfig(CONFIG_FILE)
	if err == nil {
		config = loadedConfig
	}

	var newDriver drivers.Driver

	switch driver {
	case "text":
		newDriver = drivers.NewTextFileDriver(config.Drivers.Text)
	case "json":
		newDriver = drivers.NewJSONDriver(config.Drivers.Json)
	default:
		newDriver = drivers.NewCLIDriver(config.Drivers.Cli)
	}
	t.setDriver(newDriver)
}

// SetDriver assign a new driver to the LogManager instance.
func (t *LogManager) setDriver(driver drivers.Driver) {
	t.driver = driver
}

// Log writes a log message with the given level and tags.
func (t *LogManager) Log(level drivers.LogLevel, message string, parentID string, transID string, transactionTags map[string]string) {
	// Merge global tags and transaction-specific tags
	mergedTags := make(map[string]string)

	// Add global tags from LogManager
	for key, value := range t.tags {
		mergedTags[key] = value
	}

	// Add transaction-specific tags (override global tags if there's a conflict)
	for key, value := range transactionTags {
		mergedTags[key] = value
	}

	entry := drivers.LogEntry{
		Message:             message,
		Level:               level,
		Tags:                mergedTags,
		Timestamp:           time.Now(),
		TransactionID:       transID,
		ParentTransactionID: parentID,
	}

	// Pass the copied map to the driver's Log function
	t.driver.Log(entry)
}

// Debug writes a log message with the Debug level and related information.
func (t *LogManager) Debug(message string, parentID string, transID string, tags map[string]string) {
	t.Log(drivers.Debug, message, parentID, transID, tags)
}

// Info writes a log message with the Info level and related information.
func (t *LogManager) Info(message string, parentID string, transID string, tags map[string]string) {
	t.Log(drivers.Info, message, parentID, transID, tags)
}

// Warning writes a log message with the Warning level and related information.
func (t *LogManager) Warning(message string, parentID string, transID string, tags map[string]string) {
	t.Log(drivers.Warning, message, parentID, transID, tags)
}

// Error writes a log message with the Error level and related information.
func (t *LogManager) Error(message string, parentID string, transID string, tags map[string]string) {
	t.Log(drivers.Error, message, parentID, transID, tags)
}

// AddTag adds a new global tag to the LogManager instance.
func (t *LogManager) AddTag(key, value string) {
	t.tags[key] = value
}

// RemoveTag removes a tag from the LogManager instance.
func (t *LogManager) RemoveTag(key string) {
	delete(t.tags, key)
}

// SetTags allows adding global tags to the LogManager instance.
func (t *LogManager) SetTags(tags map[string]string) {
	t.tags = map[string]string{}

	for key, value := range tags {
		t.tags[key] = value
	}
}

// ResetTags remove all tags of the LogManager instance.
func (t *LogManager) ResetTags() {
	t.tags = map[string]string{}
}
