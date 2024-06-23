package weather

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	Key           string `json:"key"`
	Location      string `json:"location"`
	CacheDate     string `json:"cacheDate"`
	CacheLocation string `json:"cacheLocation"`
}

type File interface {
	Config | Weather
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

func Load[T File](fp string) (*T, error) {
	var config T
	file, err := os.ReadFile(fp)
	if err != nil {
		return &config, err
	}
	err = json.Unmarshal(file, &config)
	return &config, err
}

func Save[T File](config T, fp string) error {
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(fp), 0755); err != nil {
		return err
	}
	return os.WriteFile(fp, data, 0644)
}
