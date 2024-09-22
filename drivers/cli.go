package drivers

import (
	"fmt"
)

// CLIDriver outputs logs to the CLI.
type CLIDriver struct{}

// NewCLIDriver initializes a new CLIDriver
func NewCLIDriver() *CLIDriver {
	return &CLIDriver{}
}

// Log implements the Driver interface for CLI logging.
func (d *CLIDriver) Log(message string, level LogLevel, tags map[string]string) {
	fmt.Printf("[%s] %s - %s\n", tags["timestamp"], levelToString(level), message)
	for k, v := range tags {
		if k != "timestamp" {
			fmt.Printf(" %s: %s\n", k, v)
		}
	}
}
