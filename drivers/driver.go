package drivers

import "time"

// LogLevel defines the level of logging
type LogLevel int

const (
	Debug LogLevel = iota
	Info
	Warning
	Error
)

type DriverConfig struct {
	LogFilePath     string `json:"log_file_path"`
	MaxSize         int    `json:"max_size"`
	MaxBackups      int    `json:"max_backups"`
	MaxAge          int    `json:"max_age"`
	TimestampFormat string `json:"timestamp_format"`
}

func levelToString(level LogLevel) string {
	switch level {
	case Debug:
		return "DEBUG"
	case Info:
		return "INFO"
	case Warning:
		return "WARNING"
	case Error:
		return "ERROR"
	default:
		return "UNKNOWN"
	}
}

// LogEntry is now a standard structure across all drivers.
type LogEntry struct {
	Message             string            // The log message
	Level               LogLevel          // The severity level of the log
	Tags                map[string]string // Any additional tags (e.g., metadata) for the log
	Timestamp           time.Time         // The time when the log is created
	TransactionID       string            // ID for the transaction
	ParentTransactionID string            // ID for the parent transaction
}

type Driver interface {
	Log(entry LogEntry) // Unified log method for all drivers
}
