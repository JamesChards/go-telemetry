package drivers

import (
	"fmt"
	"os"
	"sync"
)

// TextFileDriver writes logs to a text file.
type TextFileDriver struct {
	file *os.File
	mu   sync.Mutex // Add a mutex to protect file writes
}

// NewTextFileDriver creates a new instance of TextFileDriver.
func NewTextFileDriver(filepath string) (*TextFileDriver, error) {
	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	return &TextFileDriver{file: file}, nil
}

// Log implements the Driver interface for text file logging.
func (d *TextFileDriver) Log(message string, level LogLevel, tags map[string]string) {
	d.mu.Lock()         // Lock before writing to the file
	defer d.mu.Unlock() // Unlock after the write is done

	logEntry := fmt.Sprintf("[%s] %s: %s     <--- TAGS: {", tags["timestamp"], levelToString(level), message)
	for k, v := range tags {
		if k != "timestamp" {
			logEntry += fmt.Sprintf(" %s: %s", k, v)
		}
	}
	logEntry += "} --->\n"
	d.file.WriteString(logEntry)
}

// Close closes the underlying file to prevent resource leaks.
func (d *TextFileDriver) Close() error {
	d.mu.Lock() // Lock during closing to prevent concurrent access
	defer d.mu.Unlock()

	return d.file.Close()
}
