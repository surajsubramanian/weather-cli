package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type hour struct {
	Time       string  `json:"datetime"`
	Temp       float64 `json:"temp"`
	Precipprob float64 `json:"precipprob"`
	Snow       float64 `json:"snow"`
	Conditions string  `json:"conditions"`
}

type day struct {
	Date       string  `json:"datetime"`
	Precipprob float64 `json:"precipprob"`
	Snow       float64 `json:"snow"`
	Sunrise    string  `json:"sunrise"`
	Sunset     string  `json:"sunset"`
	Hours      []hour  `json:"hours"`
}

type Weather struct {
	ResolvedAddress string `json:"resolvedAddress"`
	Days            []day  `json:"days"`
}

func fetchWeather(config Config, configPath string, cachePath string, date string) (*Weather, error) {
	var weather Weather
	url := fmt.Sprintf(
		"https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/%s/%s/%s?unitGroup=metric&include=hours&key=%s&contentType=json",
		config.Location, date, date, config.Key,
	)
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		fmt.Printf("Status Code: %dn", res.StatusCode)
		panic("Weather API not available")
	}

	body, err := io.ReadAll(res.Body)
	err = json.Unmarshal(body, &weather)
	if err != nil {
		return nil, err
	}
	if err := os.WriteFile(cachePath, body, 0644); err != nil {
		return nil, err
	}
	config.CacheDate = date
	config.CacheLocation = config.Location
	if err != nil {
		return nil, err
	}
	err = Save[Config](config, configPath)
	return &weather, err
}
