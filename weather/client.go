package weather

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
)

type Client struct {
	Config     Config
	ConfigPath string
	CachePath  string
	NoCache    bool
}

type Output struct {
	Time       string
	Temp       float64
	Precip     float64
	Conditions string
}

func (c Client) Process(date string, time string, hour int) []Output {
	weather := &Weather{}
	var err error
	if _, err = os.Stat(c.CachePath); errors.Is(err, os.ErrNotExist) || c.NoCache || c.Config.CacheDate != date {
		weather, err = fetchWeather(c.Config, c.ConfigPath, c.CachePath, date)
		if err != nil {
			panic(err)
		}
	} else {
		weather, err = Load[Weather](c.CachePath)
		if err != nil {
			panic(err)
		}
	}

	color.Green(fmt.Sprintf(
		"%s\n%s\t%s\n",
		(*weather).ResolvedAddress, date, time,
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
