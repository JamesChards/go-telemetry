# go-telemetry

# Logging System in Go

This repository contains a telemetry and logging system written in Go. The system is highly extensible and supports various output formats and drivers for logging (CLI, JSON, and text files). The system also includes a transaction system for tagging, tracing, and monitoring transactions across the application.

# Features

Multiple logging drivers: Supports CLI, JSON, and text file logging.
Log levels: Customizable log levels (Info, Warning, Error, Debug).
Tagging system: Allows adding, removing, and resetting tags dynamically to log entries and transactions.
Transaction tracking: Manage and track transactions with support for sub-transactions.
Driver configuration: Flexible configuration options for log storage, file rotation, and formatting.

# Table of Contents

Installation
Usage
Basic Logging
Transaction Management
Configuration
Logging Drivers
Testing
Contributing
License
Installation
Clone the repository into your $GOPATH:

bash
Copy code
git clone https://github.com/yourusername/logging-system.git
cd logging-system
Run the following command to install dependencies:

bash
Copy code
go mod tidy
You can now start using the logging system in your Go projects.

Usage
Basic Logging
To get started, initialize the LogManager with your desired driver (CLI, JSON, or text). Here is an example of how to log messages:

go
Copy code
package main

import (
"example/telemetry"
)

func main() {
logger := telemetry.NewLogger("cli")

    logger.SetDriverWithName("cli") // Set CLI as the output driver

    // Add tags
    logger.AddTag("environment", "production")

    // Log an info message
    logger.Log(telemetry.Info, "Application started successfully", "", "session123", nil)

}
Transaction Management
The logging system also supports transaction tracking, which can be helpful for tracing execution or workflows. Here’s an example:

go
Copy code
package main

import (
"example/telemetry"
)

func main() {
logger := telemetry.NewLogger("json")
transaction := telemetry.NewTransaction("trans123", logger)

    // Start the transaction
    transaction.Start()

    // Add tags specific to the transaction
    transaction.AddTag("user_id", "user123")

    // End the transaction
    transaction.End()

}
Advanced Usage (Sub-transactions)
You can create sub-transactions to track smaller units of work within a larger transaction:

go
Copy code
subTransaction := transaction.SubTransaction("subTrans456")
subTransaction.Start()

// Perform operations under sub-transaction

subTransaction.End()
Configuration
Logging drivers in this system are highly configurable. Each driver can be configured with different settings such as file paths, maximum log file size, and log rotation policies. Configuration is passed using the DriverConfig struct.

Example configuration for JSON driver:
go
Copy code
driverConfig := telemetry.DriverConfig{
LogFilePath: "app-logs.json",
MaxSize: 10, // 10 MB
MaxBackups: 5, // Store up to 5 rotated logs
MaxAge: 30, // Keep logs for 30 days
TimestampFormat: time.RFC3339, // Log timestamp format
}

logger.SetDriverWithConfig("json", driverConfig)
Logging Drivers
This system supports the following logging drivers:

CLI Driver: Prints log messages directly to the console.
JSON Driver: Logs data in JSON format, suitable for structured logging systems.
Text File Driver: Stores logs in a traditional text format, with optional file rotation and size limits.
Adding a New Driver
To add a new driver, create a new struct that implements the Driver interface. Here’s a minimal example:

go
Copy code
type MyCustomDriver struct{}

func (d \*MyCustomDriver) Log(entry LogEntry) {
// Custom logging logic
}
Testing
The code is fully unit-tested using Go’s testing package. Unit tests ensure that all core functionalities, including drivers, tagging, and transaction management, behave as expected.

Running Tests
You can run the tests using the following command:

bash
Copy code
go test ./... -v
This command runs all tests across the project. Tests are located in \*\_test.go files within each package.

Example of Unit Testing
Here's an example of a simple unit test for the LogManager class:

go
Copy code
func TestLogManager_Log(t \*testing.T) {
lm := telemetry.NewLogger("cli")

    tags := map[string]string{"user": "admin"}
    lm.AddTag("env", "test")

    lm.Log(telemetry.Info, "Test log message", "", "1234", tags)

    // This ensures that the log method works without error.

}
