package drivers

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/natefinch/lumberjack"
)

// JSONDriver writes logs as JSON to a file.
type JSONDriver struct {
	logger          *lumberjack.Logger
	mu              sync.Mutex // Add a mutex to protect file writes
	timestampFormat string
}

// NewJSONDriver creates a new instance of JSONDriver.
func NewJSONDriver(driverConfig DriverConfig) *JSONDriver {
	return &JSONDriver{
		logger: &lumberjack.Logger{
			Filename:   driverConfig.LogFilePath,
			MaxSize:    driverConfig.MaxSize,    // megabytes
			MaxBackups: driverConfig.MaxBackups, // number of backups
			MaxAge:     driverConfig.MaxAge,     // days
			Compress:   true,                    // compress old log files
		},
		timestampFormat: driverConfig.TimestampFormat,
	}
}

// Log implements the Driver interface for JSON logging.
func (d *JSONDriver) Log(entry LogEntry) {
	d.mu.Lock()         // Lock before writing to the file
	defer d.mu.Unlock() // Unlock after the write is done

	// Prepare the log entry in JSON format
	jsonEntry := map[string]interface{}{
		"message":   entry.Message,
		"level":     levelToString(entry.Level),
		"tags":      entry.Tags,
		"timestamp": entry.Timestamp.Format(d.timestampFormat),
	}
	if entry.ParentTransactionID != "" {
		jsonEntry["transaction_id"] = entry.ParentTransactionID
		jsonEntry["sub_transaction_id"] = entry.TransactionID
	} else {
		jsonEntry["transaction_id"] = entry.TransactionID
	}

	// Marshal the log entry to JSON
	data, err := json.Marshal(jsonEntry)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	// Write the JSON entry followed by a newline character
	if _, err := d.logger.Write(append(data, '\n')); err != nil {
		fmt.Println("Error writing to JSON log file:", err)
		return
	}
}

// Close closes the underlying file to prevent resource leaks.
func (d *JSONDriver) Close() error {
	d.mu.Lock() // Lock during closing to prevent concurrent access
	defer d.mu.Unlock()

	return d.logger.Close()
}
