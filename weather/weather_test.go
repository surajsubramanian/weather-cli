package weather

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestProcess(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "")
	if err != nil {
		t.Fatalf("Failed to create temporary directory")
	}
	defer os.RemoveAll(tempDir)

	configPath := filepath.Join(tempDir, "config.json")
	cachePath := filepath.Join(tempDir, "weather_cache.json")

	date := "2024-06-15"
	config := Config{
		Key:       "SampleKey",
		Location:  "SampleCity",
		CacheDate: date,
	}

	configData, _ := json.Marshal(config)
	err = os.WriteFile(configPath, configData, 0644)
	if err != nil {
		t.Fatalf("Failed to write config to temporary file: %v", err)
	}

	hours := []hour{}
	for i := 0; i < 19; i++ {
		time := fmt.Sprintf("%d:00:00", i)
		hours = append(hours, hour{time, 27.2, 20.0, 0.0, "Clear"})
	}
	for _, h := range []hour{
		{"19:00:00", 27.2, 20.0, 0.0, "Clear"},
		{"20:00:00", 27.1, 25.0, 0.0, "Clear"},
		{"21:00:00", 26.9, 35.0, 0.0, "Partly Cloudy"},
		{"22:00:00", 26.2, 40.0, 0.0, "Partly Cloudy"},
		{"23:00:00", 25.3, 50.0, 0.0, "Rain"},
	} {
		hours = append(hours, h)
	}

	sampleWeather := Weather{
		ResolvedAddress: "Sample City",
		Days: []day{
			{
				Date:       date,
				Precipprob: 30.0,
				Snow:       0.0,
				Sunrise:    "05:30",
				Sunset:     "19:30",
				Hours:      hours,
			},
		},
	}

	cacheData, _ := json.Marshal(sampleWeather)
	err = os.WriteFile(cachePath, cacheData, 0644)
	if err != nil {
		t.Fatalf("Failed to write cache to temporary file: %v", err)
	}

	c := Client{
		Config:     config,
		ConfigPath: configPath,
		CachePath:  cachePath,
		NoCache:    false,
	}
	get := c.Process(date, "20:23:10", 20)
	expected := []Output{
		{"20:00:00", 27.1, 25.0, "Clear"},
		{"21:00:00", 26.9, 35.0, "Partly Cloudy"},
		{"22:00:00", 26.2, 40.0, "Partly Cloudy"},
		{"23:00:00", 25.3, 50.0, "Rain"},
	}
	for i, e := range expected {
		if e != get[i] {
			t.Fatalf("Expected %v, got %v at index %d", e, get[i], i)
		}
	}
}
