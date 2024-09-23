package drivers

import (
	"fmt"
)

// CLIDriver outputs logs to the CLI.
type CLIDriver struct {
	timestampFormat string
}

// NewCLIDriver initializes a new CLIDriver
func NewCLIDriver(driverConfig DriverConfig) *CLIDriver {
	return &CLIDriver{
		timestampFormat: driverConfig.TimestampFormat,
	}
}

// Log implements the Driver interface for CLI logging.
func (d *CLIDriver) Log(entry LogEntry) {
	var transString string
	if entry.ParentTransactionID != "" {
		transString = fmt.Sprintf("[Transaction %s] -> [SubTransaction %s]", entry.ParentTransactionID, entry.TransactionID)
	} else {
		transString = fmt.Sprintf("[Transaction %s]", entry.TransactionID)
	}
	fmt.Printf("[%s] %s - %s %s\n", entry.Timestamp.Format(d.timestampFormat), levelToString(entry.Level), transString, entry.Message)
	for k, v := range entry.Tags {
		if k != "timestamp" {
			fmt.Printf(" %s: %s\n", k, v)
		}
	}
}
