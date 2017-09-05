#!/bin/bash

influxd 2>/var/log/influxdb.log &
sleep 3
influx -execute "create database myfirstdatabase"
tail -f /var/log/influxdb.log

