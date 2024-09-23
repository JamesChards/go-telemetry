package telemetry

import (
	"encoding/json"
	"example/telemetry/drivers"
	"os"
	"time"
)

const CONFIG_FILE = "config.json"

type DriversConfig struct {
	Cli  drivers.DriverConfig `json:"cli"`
	Text drivers.DriverConfig `json:"text"`
	Json drivers.DriverConfig `json:"json"`
}

// Config struct defines the structure for the configuration file.
type Config struct {
	Drivers       DriversConfig `json:"drivers"`
	DefaultDriver string        `json:"default_driver"`
}

// Load default config
func defaultConfig() *Config {
	return &Config{
		Drivers: DriversConfig{
			Cli: drivers.DriverConfig{
				TimestampFormat: time.RFC822,
			},
			Text: drivers.DriverConfig{
				LogFilePath:     "app-text.log",
				MaxSize:         10,
				MaxBackups:      5,
				MaxAge:          30,
				TimestampFormat: time.RFC3339,
			},
			Json: drivers.DriverConfig{
				LogFilePath:     "app-json.log",
				MaxSize:         10,
				MaxBackups:      5,
				MaxAge:          30,
				TimestampFormat: time.RFC3339,
			},
		},
		DefaultDriver: "cli",
	}
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

func CreateDefaultConfig() error {
	defaultConfig := defaultConfig()
	bytes, _ := json.MarshalIndent(defaultConfig, "", "  ")

	_, err := LoadConfig(CONFIG_FILE)
	if err != nil {
		return os.WriteFile(CONFIG_FILE, bytes, 0644)
	}

	return nil
}
