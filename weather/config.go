package weather

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func LoadConfig(configPath string) (Config, error) {
	var config Config
	file, err := os.ReadFile(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			return config, fmt.Errorf("configuration file not found. Please run the CLI with --set-key and --set-location to initialize or edit/create %s", configPath)
		}
		return config, err
	}
	err = json.Unmarshal(file, &config)
	return config, err
}

func (c Config) SaveConfig(configPath string) error {
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(configPath), 0755); err != nil {
		return err
	}
	return os.WriteFile(configPath, data, 0644)
}
