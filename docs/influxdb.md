# InfluxDB configuration
All configuration related InfluxDB can be found in this document.


## Install InfluxDB on macOS
```
brew update
brew install influxdb
```

## Make a link and luanch InfluxDB
```
ln -sfv /usr/local/opt/influxdb/*.plist ~/Library/LaunchAgents

Then luanch by this command:
launchctl load ~/Library/LaunchAgents/homebrew.mxcl.influxdb.plist

Or:
influxd -config /usr/local/etc/influxdb.conf
```

## Connect to the API on Port: 8086
```
influx: in the command-line
```

