package drivers

import (
	"fmt"
	"sync"

	"github.com/natefinch/lumberjack"
)

// TextFileDriver writes logs to a text file.
type TextFileDriver struct {
	logger          *lumberjack.Logger
	mu              sync.Mutex // Add a mutex to protect file writes
	timestampFormat string
}

// NewTextFileDriver creates a new instance of TextFileDriver.
func NewTextFileDriver(driverConfig DriverConfig) *TextFileDriver {
	return &TextFileDriver{
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

// Log implements the Driver interface for text file logging.
func (d *TextFileDriver) Log(entry LogEntry) {
	d.mu.Lock()         // Lock before writing to the file
	defer d.mu.Unlock() // Unlock after the write is done

	var transString string
	if entry.ParentTransactionID != "" {
		transString = fmt.Sprintf("[Transaction %s] -> [SubTransaction %s]", entry.ParentTransactionID, entry.TransactionID)
	} else {
		transString = fmt.Sprintf("[Transaction %s]", entry.TransactionID)
	}
	// Format the log entry with timestamp and level
	logEntry := fmt.Sprintf("[%s] %s: %s %s\n", entry.Timestamp.Format(d.timestampFormat), levelToString(entry.Level), transString, entry.Message)

	// Add the tags
	for k, v := range entry.Tags {
		logEntry += fmt.Sprintf("  %s: %s\n", k, v)
	}

	d.logger.Write([]byte(logEntry))
}

// Close closes the underlying file to prevent resource leaks.
func (d *TextFileDriver) Close() error {
	d.mu.Lock() // Lock during closing to prevent concurrent access
	defer d.mu.Unlock()

	return d.logger.Close()
}
