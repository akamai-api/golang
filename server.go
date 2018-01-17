package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/influxdata/influxdb/client/v2"
)

// CreateObjects creates a list of AkamaiPayloads from a raw byte slice
func CreateObjects(jsonFile []byte) (_ []AkamaiPayload, err error) {
	var arrayObject []AkamaiPayload
	err = json.Unmarshal(jsonFile, &arrayObject)
	return arrayObject, err
}

// Handle parses the incoming request data
func Handle(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "OK! Received! -> ")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	obj, err := CreateObjects(body)
	if err != nil {
		log.Fatal(err)
	}
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: "http://influxDB:8086",
	})

	// Create a new point batch
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database: "metrics",
	})
	if err != nil {
		log.Fatal(err)
	}

	for key, o := range obj {
		fmt.Println("Server is:", o.RespHdr.Server, "KEY:", key)

		if err != nil {
			log.Fatal(err)
		}
		// Create a point and add to batch
		tags := map[string]string{
			"cp":            o.CP,
			"format":        o.Format,
			"city":          o.Geo.City,
			"country":       o.Geo.Country,
			"status":        o.Message.Status,
			"cacheStatus":   o.NetPerf.CacheStatus,
			"proto":         o.Message.Proto,
			"protoVer":      o.Message.ProtoVer,
			"reqPort":       o.Message.ReqPort,
			"reqPath":       o.Message.ReqPath,
			"reqMethod":     o.Message.ReqMethod,
			"region":        o.Geo.Region,
			"UA":            o.Message.UA,
			"reqHost":       o.Message.ReqHost,
			"respCT":        o.Message.RespCT,
			"netPerfasnum":  o.NetPerf.Asnum,
			"netPerfedgeIP": o.NetPerf.EdgeIP,
			"networkAsum":   o.Network.Asnum,
			"networkEdgeIP": o.Network.EdgeIP,
			"network":       o.Network.Network,
			"networkType":   o.Network.NetworkType,
			"cookie":        o.ReqHdr.Cookie,
			"server":        o.RespHdr.Server,
			"contEnc":       o.RespHdr.ContEnc,
			"type":          o.Type,
			"version":       o.Version,
		}

		lat, err := strconv.ParseFloat(o.Geo.Lat, 64)
		long, err := strconv.ParseFloat(o.Geo.Long, 64)
		downloadTime, err := strconv.ParseInt(o.NetPerf.DownloadTime, 10, 32)
		status, err := strconv.ParseInt(o.Message.Status, 10, 32)
		bytes, err := strconv.ParseInt(o.Message.Bytes, 10, 64)
		respLen, err := strconv.ParseInt(o.Message.RespLen, 10, 32)
		lastMileRTT, err := strconv.ParseFloat(o.NetPerf.LastMileRTT, 64)
		start, err := strconv.ParseFloat(o.Start, 64)

		if err != nil {
			log.Fatal(err)
		}
		fields := map[string]interface{}{
			"lat":          lat,
			"long":         long,
			"downloadtime": downloadTime,
			"status":       status,
			"bytes":        bytes,
			"respLen":      respLen,
			"lastMileRTT":  lastMileRTT,
			"start":        start,
			"respCT":       o.Message.RespCT,
		}

		pt, err := client.NewPoint("measurement", tags, fields, time.Now())
		if err != nil {
			log.Fatal(err)
		}
		bp.AddPoint(pt)
	}

	// Write the batch
	if err := c.Write(bp); err != nil {
		log.Fatal(err)
	}
}

func main() {

	http.HandleFunc("/", Handle)
	http.ListenAndServe(":9143", nil)
}
