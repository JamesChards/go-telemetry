package telemetry

import (
	"os"
	"testing"
)

// Test configuration loading.
func TestLoadConfig(t *testing.T) {
	// Create a temporary config file
	configContent := `{"default_driver": "json", "log_file_path": "test.log"}`
	tmpFile, err := os.CreateTemp("", "config*.json")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	_, err = tmpFile.Write([]byte(configContent))
	if err != nil {
		t.Fatalf("failed to write to temp file: %v", err)
	}

	// Load the configuration
	config, err := LoadConfig(tmpFile.Name())
	if err != nil {
		t.Fatalf("failed to load config: %v", err)
	}

	// Assert the config values
	if config.DefaultDriver != "json" {
		t.Errorf("expected DefaultDriver to be 'json', got '%s'", config.DefaultDriver)
	}
	if config.LogFilePath != "test.log" {
		t.Errorf("expected LogFilePath to be 'test.log', got '%s'", config.LogFilePath)
	}
}

// Test loading an invalid config file.
func TestLoadConfig_InvalidFile(t *testing.T) {
	_, err := LoadConfig("nonexistent.json")
	if err == nil {
		t.Error("expected error when loading nonexistent config file, but got nil")
	}
}
