package weather

type Config struct {
	Key           string `json:"key"`
	Location      string `json:"location"`
	CacheDate     string `json:"cacheDate"`
	CacheLocation string `json:"cacheLocation"`
}

type Weather struct {
	ResolvedAddress string `json:"resolvedAddress"`
	Days            []day  `json:"days"`
}

type day struct {
	Date       string  `json:"datetime"`
	Precipprob float64 `json:"precipprob"`
	Snow       float64 `json:"snow"`
	Sunrise    string  `json:"sunrise"`
	Sunset     string  `json:"sunset"`
	Hours      []hour  `json:"hours"`
}

type hour struct {
	Time       string  `json:"datetime"`
	Temp       float64 `json:"temp"`
	Precipprob float64 `json:"precipprob"`
	Snow       float64 `json:"snow"`
	Conditions string  `json:"conditions"`
}

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
