package drivers

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sync"
)

// JSONDriver writes logs as JSON to a file.
type JSONDriver struct {
	file *os.File
	mu   sync.Mutex // Add a mutex to protect file writes
}

// NewJSONDriver creates a new instance of JSONDriver.
func NewJSONDriver(filepath string) (*JSONDriver, error) {
	file, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}

	return &JSONDriver{file: file}, nil
}

// Log implements the Driver interface for JSON logging.
func (d *JSONDriver) Log(message string, level LogLevel, tags map[string]string) {
	d.mu.Lock()         // Lock before writing to the file
	defer d.mu.Unlock() // Unlock after the write is done

	// Read the existing contents
	fileInfo, err := d.file.Stat()
	if err != nil {
		fmt.Println("Error getting file info:", err)
		return
	}
	fileSize := fileInfo.Size()

	// Move to the start of the file to read its contents
	_, err = d.file.Seek(0, io.SeekStart)
	if err != nil {
		fmt.Println("Error seeking file:", err)
		return
	}
	content, err := io.ReadAll(d.file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var jsonArray []map[string]interface{}
	if fileSize > 0 {
		// Try to decode existing JSON content
		if content[0] == '[' {
			json.Unmarshal(content, &jsonArray)
		}
	}

	// Create new log entry
	entry := map[string]interface{}{
		"message": message,
		"level":   levelToString(level),
		"tags":    tags,
	}
	jsonArray = append(jsonArray, entry)

	// Truncate the file and write the new JSON array
	err = d.file.Truncate(0)
	if err != nil {
		fmt.Println("Error truncating file:", err)
		return
	}
	_, err = d.file.Seek(0, io.SeekStart)
	if err != nil {
		fmt.Println("Error seeking file:", err)
		return
	}
	// Add indent to json array
	data, err := json.MarshalIndent(jsonArray, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}
	_, err = d.file.Write(data)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}

// Close closes the underlying file to prevent resource leaks.
func (d *JSONDriver) Close() error {
	d.mu.Lock() // Lock during closing to prevent concurrent access
	defer d.mu.Unlock()

	return d.file.Close()
}
