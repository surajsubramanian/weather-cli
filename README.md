## Weather CLI

Simple CLI application to output weather using [https://www.visualcrossing.com/weather-api](https://www.visualcrossing.com/weather-api)

### Build

Download the binary from the releases or build it using
```bash
go build -o wcli
```

### Initialization

* Get your free API key from [https://www.visualcrossing.com/weather-api](https://www.visualcrossing.com/weather-api)

* Set your key and location:

```bash
./wcli -set-key YOURSECRETKEY
./wcli -set-location Espoo
```

### Execution

* Execute the CLI using `./wcli`. Weather details are cached under `$HOME/.wcli` after initial execution.

* You can also execute `./wcli -refresh` to fetch details from the endpoint once again. Details are newly fetched if the city is changed by the user.

```bash
> ./wcli
Espoo, Etelä-Suomi, Suomi
2024-06-16      21:11:30
Time      Temp    Precip  Conditions
21:00:00  15.8    0.0     Overcast
22:00:00  15.3    0.0     Overcast
23:00:00  15.1    0.0     Overcast
```
