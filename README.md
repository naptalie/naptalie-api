# naptalie-api

this is an http api + different types of consumer clients approach that does stuff! I wanted to take a crack at a go web server from scratch to get more comfortable with the language. Enjoy!

## HTTP server 

please look at the various routes that are implemented in api/routes. you can run the server via the included makefile.

``` bash
make run_server
```

``` bash
curl http://localhost:8090/weather
```

``` json
{"message":"Command executed successfully","data":{"latitude":39.766758,"longitude":-86.14396,"generationtime_ms":0.25594234466552734,"utc_offset_seconds":0,"timezone":"GMT","timezone_abbreviation":"GMT","elevation":214,"daily_units":{"time":"iso8601","temperature_2m_max":"°F","temperature_2m_min":"°F","precipitation_sum":"inch","weathercode":"wmo code"},"daily":{"time":["2025-07-21","2025-07-22","2025-07-23","2025-07-24","2025-07-25","2025-07-26","2025-07-27"],"temperature_2m_max":[83.9,89.8,94,92.8,90.5,87.3,91],"temperature_2m_min":[72,65.4,67,72,73.4,73,73],"precipitation_sum":[0.008,0,0,0.004,0.386,0.055,0.035],"weathercode":[51,3,3,51,95,51,95]}},"success":true}

```

## Discord Client

the discord client implements calls to the http server and returns the data as structured responses back using the discord Go sdk! You 

