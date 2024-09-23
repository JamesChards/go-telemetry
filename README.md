# Logging System in Go

This repository contains a **telemetry and logging system** written in Go. The system is highly extensible and supports various output formats and drivers for logging (CLI, JSON, and text files). The system also includes a transaction system for tagging, tracing, and monitoring transactions across the application.

## Features

- **Multiple logging drivers**: Supports CLI, JSON, and text file logging.
- **Log levels**: Customizable log levels (Info, Warning, Error, Debug).
- **Tagging system**: Allows adding, removing, and resetting tags dynamically to log entries and transactions.
- **Transaction tracking**: Manage and track transactions with support for sub-transactions.
- **Driver configuration**: Flexible configuration options for log storage, file rotation, and formatting.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
  - [Basic Logging](#basic-logging)
  - [Transaction Management](#transaction-management)
- [Configuration](#configuration)
- [Logging Drivers](#logging-drivers)
- [Testing](#testing)
- [Contributing](#contributing)
- [License](#license)

## Installation

Clone the repository into your workspace:

```bash
git clone https://github.com/JamesChards/go-telemetry.git
cd go-telemetry
```

Run the following command to install dependencies:

```bash
go mod tidy
```

You can now start using the logging system in your Go projects.

## Usage

### Basic Logging

To get started, initialize the `LogManager` with your desired driver (CLI, JSON, or text). Here is an example of how to log messages:

```go
package main

import (
    "example/telemetry"
)

func main() {
    logger := telemetry.NewLogger("cli")

    // in case when you need to change the driver
    logger.SetDriverWithName("json") // Set JSON as the output driver

    // Add tags
    logger.AddTag("environment", "production")

    // Log an info message
    logger.Log(telemetry.Info, "Application started successfully", "", "session123", nil)
}
```

### Transaction Management

The logging system also supports transaction tracking, which can be helpful for tracing execution or workflows. Here’s an example:

```go
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
```

### Advanced Usage (Sub-transactions)

You can create sub-transactions to track smaller units of work within a larger transaction:

```go
subTransaction := transaction.SubTransaction("subTrans456")
subTransaction.Start()

// Perform operations under sub-transaction

subTransaction.End()
```

## Configuration

Logging drivers in this system are highly configurable. Each driver can be configured with different settings such as file paths, maximum log file size, and log rotation policies. Configuration is passed using the `DriverConfig` struct.

### Example configuration for JSON driver:

```go
driverConfig := telemetry.DriverConfig{
    LogFilePath:     "app-logs.json",
    MaxSize:         10,   // 10 MB
    MaxBackups:      5,    // Store up to 5 rotated logs
    MaxAge:          30,   // Keep logs for 30 days
    TimestampFormat: time.RFC3339, // Log timestamp format
}

logger.SetDriverWithConfig("json", driverConfig)
```

## Logging Drivers

This system supports the following logging drivers:

- **CLI Driver**: Prints log messages directly to the console.
- **JSON Driver**: Logs data in JSON format, suitable for structured logging systems.
- **Text File Driver**: Stores logs in a traditional text format, with optional file rotation and size limits.

### Adding a New Driver

To add a new driver, create a new struct that implements the `Driver` interface. Here’s a minimal example:

```go
type MyCustomDriver struct{}

func (d *MyCustomDriver) Log(entry LogEntry) {
    // Custom logging logic
}
```

## Testing

The code is fully unit-tested using Go’s `testing` package. Unit tests ensure that all core functionalities, including drivers, tagging, and transaction management, behave as expected.

### Running Tests

You can run the tests using the following command:

```bash
go test ./... -v
```

This command runs all tests across the project. Tests are located in `*_test.go` files within each package.

### Example of Unit Testing

Here's an example of a simple unit test for the `LogManager` class:

```go
func TestLogManager_Log(t *testing.T) {
    lm := telemetry.NewLogger("cli")

    tags := map[string]string{"user": "admin"}
    lm.AddTag("env", "test")

    lm.Log(telemetry.Info, "Test log message", "", "1234", tags)

    // This ensures that the log method works without error.
}
```

## Areas for Improvement:

### Log Retrieval for a Specific Process:

Issue: While this code can tag logs with process-specific data (like transaction_id and parent_transaction_id), there is no mechanism in place to retrieve logs for a specific process or sub-process. The current code focuses on logging, but not on log retrieval.
Solution: To retrieve logs for a process, you need to introduce a mechanism to filter logs based on specific tags (e.g., transaction_id). This could be handled by extending the driver interface to support querying logs, or you could design a separate log processor that filters logs from the stored files based on tags.

### Asynchronous or Buffered Logging for Large Processes:

Issue: For long-running processes (e.g., an order import with many child processes), synchronous I/O operations may slow down the process, especially if logging is frequent or involves file I/O.
Solution: Consider adding asynchronous logging (using channels or a log queue) to decouple the log-writing process from the actual application logic. This will improve performance, especially in high-volume processes.

### Log Contextualization:

Issue: The logger writes logs based on the current tags, but there is no concept of a "logging context" that could be propagated across function calls and processes.
Solution: You could introduce a logging context (possibly using Go's context.Context) that automatically carries metadata (tags) across the entire process lifecycle, including sub-processes. This will ensure that tags are passed along naturally without needing explicit calls to SetTags.

## Contributing

Contributions are welcome! To contribute:

1. Fork the repository.
2. Create a new branch for your feature or bug fix (`git checkout -b feature/my-feature`).
3. Make the changes and ensure tests pass.
4. Commit your changes and push the branch (`git push origin feature/my-feature`).
5. Submit a pull request, and provide a detailed description of your changes.

Please ensure that your code follows Go best practices and includes tests for new features or fixes.
