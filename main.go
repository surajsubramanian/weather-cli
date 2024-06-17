package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/surajsubramanian/weather-cli/weather"
)

var configPath string
var cachePath string

func init() {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	configPath = filepath.Join(home, ".weathercli/config.json")
	cachePath = filepath.Join(home, ".weathercli/weather.json")
}

func main() {

	setLocation := flag.String("set-location", "", "Set the default city")
	setKey := flag.String("set-key", "", "Set the default API key")
	noCache := flag.Bool("refresh", false, "Do not use cache")
	flag.Parse()

	config, err := weather.Load[weather.Config](configPath)
	client := weather.Client{
		Config:     *config,
		ConfigPath: configPath,
		CachePath:  cachePath,
		NoCache:    *noCache,
	}
	if *setKey == "" && *setLocation == "" {
		if err != nil {
			panic(err)
		} else if config.Key == "" {
			fmt.Printf("Please run the CLI with --set-key to initialize or edit/create %s", configPath)
			os.Exit(1)
		} else if config.Location == "" {
			fmt.Printf("Please run the CLI with --set-location to initialize or edit/create %s", configPath)
			os.Exit(1)
		} else {
			client.Display()
		}
	} else {
		if *setLocation != "" {
			config.Location = *setLocation
			if config.CacheLocation != config.Location {
				if _, err := os.Stat(cachePath); !errors.Is(err, os.ErrNotExist) {
					err := os.Remove(cachePath)
					if err != nil {
						panic(err)
					}
				}
			}
		} else if *setKey != "" {
			config.Key = *setKey
		}
		err := config.SaveConfig(configPath)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Config %s has been updated\n", configPath)
	}
}
