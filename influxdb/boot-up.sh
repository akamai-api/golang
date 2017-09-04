#!/bin/bash

influxd 2>/var/log/influxdb.log &
sleep 3
influx -execute "create database metrics"
tail -f /var/log/influxdb.log

