package weather

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/fatih/color"
)

var weather Weather

func fetchWeather(config Config, configPath string, cachePath string, date string) error {
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
	if err != nil {
		return err
	}
	if err := os.WriteFile(cachePath, body, 0644); err != nil {
		return err
	}
	config.CacheDate = date
	config.CacheLocation = config.Location
	err = json.Unmarshal(body, &weather)
	if err != nil {
		return err
	}
	config.SaveConfig(configPath)
	return nil
}

func loadWeather(cachePath string) error {
	body, err := os.ReadFile(cachePath)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &weather)
	if err != nil {
		return err
	}
	return nil
}

func (c Client) Process(date string, time string, hour int) []Output {
	if _, err := os.Stat(c.CachePath); errors.Is(err, os.ErrNotExist) || c.NoCache || c.Config.CacheDate != date {
		err := fetchWeather(c.Config, c.ConfigPath, c.CachePath, date)
		if err != nil {
			panic(err)
		}
	} else {
		err := loadWeather(c.CachePath)
		if err != nil {
			panic(err)
		}
	}

	color.Green(fmt.Sprintf(
		"%s\n%s\t%s\n",
		weather.ResolvedAddress, date, time,
	))

	output := []Output{}
	for i := hour; i < hour+5 && i < 24; i++ {
		output = append(output, Output{
			Time:       weather.Days[0].Hours[i].Time,
			Temp:       weather.Days[0].Hours[i].Temp,
			Precip:     weather.Days[0].Hours[i].Precipprob,
			Conditions: weather.Days[0].Hours[i].Conditions,
		})
	}
	return output
}

func (c Client) Display() {
	now := time.Now()
	date := now.Format("2006-01-02")
	time := now.Format("15:04:05")
	hour, _, _ := now.Clock()

	output := c.Process(date, time, hour)
	color.White(fmt.Sprintf(
		"%-10s%-8s%-8s%-20s\n",
		"Time", "Temp", "Precip", "Conditions",
	))
	for _, o := range output {
		message := fmt.Sprintf(
			"%-10s%-8.1f%-8.1f%-20s\n",
			o.Time,
			o.Temp,
			o.Precip,
			o.Conditions,
		)
		if o.Precip < 40 {
			fmt.Printf(message)
		} else {
			color.Blue(message)
		}
	}
}
