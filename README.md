# WinMetrics
---
It's annoying to get hardware sensors metrics from Windows as HTTP server, not anymore.

Now you can monitor your physical server sensors as part of Prometheus stack!

## Build

You can just do `go build cmd/winmetrics/main.go` and that's it

## Dependencies

- This program is Windows only.
- You must run [Libre Hardware Monitor](https://github.com/LibreHardwareMonitor/LibreHardwareMonitor), this would exposes new WMI entries as `root/LibreHardwareMonitor` namespace and `Sensor` and `Hardware` classes.

## Run

Run the executable and you can get the data through HTTP API at (default) `GET localhost:8123/api/v1/metrics` in JSON format:

```
{
    "sensors": [List of sensors],
    "hardwares": [List of hardwares]
}
```

Or `GET localhost:8123/metrics` in Prometheus format.
