package telemetry

import (
	"encoding/json"
	"os"
)

const CONFIG_FILE = "config.json"

// Config struct defines the structure for the configuration file.
type Config struct {
	DefaultDriver string `json:"default_driver"`
	LogFilePath   string `json:"log_file_path"`
}

// LoadConfig loads the telemetry configuration from a file.
func LoadConfig(filepath string) (*Config, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

// Load default config
func DefaultConfig() *Config {
	return &Config{
		"cli",
		"app.log",
	}
}
