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

// AkamaiPayload is a Golang representation of the Cloudmonitor JSON datastructure
type AkamaiPayload struct {
	// Content Provider ID
	CP string `json:"cp"`
	// Defines the format of the payload (?)
	Format  string `json:"format"`
	ID      string `json:"id"`
	Start   string `json:"start"`
	Type    string `json:"type"`
	Version string `json:"version"`

	Geo     GeoStruct     `json:"geo"`
	Message MessageStruct `json:"message"`
	NetPerf NetPerfStruct `json:"netPerf"`
	Network NetworkStruct `json:"network"`
	ReqHdr  ReqHdrStruct  `json:"reqHdr"`
	RespHdr RespHdrStruct `json:"respHdr"`
}

// GeoStruct is used for storing the JSON subfields
type GeoStruct struct {
	City    string `json:"city"`
	Country string `json:"country"`
	Lat     string `json:"lat"`
	Long    string `json:"long"`
	Region  string `json:"region"`
}

// MessageStruct is used for storing the JSON subfields
type MessageStruct struct {
	UA        string `json:"UA"`
	Bytes     string `json:"bytes"`
	CliIP     string `json:"cliIP"`
	FwdHost   string `json:"fwdHost"`
	Proto     string `json:"proto"`
	ProtoVer  string `json:"protoVer"`
	ReqHost   string `json:"reqHost"`
	ReqMethod string `json:"reqMethod"`
	ReqPath   string `json:"reqPath"`
	ReqPort   string `json:"reqPort"`
	RespCT    string `json:"respCT"`
	RespLen   string `json:"respLen"`
	Status    string `json:"status"`
}

// NetPerfStruct is used for storing the JSON subfields
type NetPerfStruct struct {
	Asnum        string `json:"asnum"`
	CacheStatus  string `json:"cacheStatus"`
	DownloadTime string `json:"downloadTime"`
	EdgeIP       string `json:"edgeIP"`
	FirstByte    string `json:"firstByte"`
	LastByte     string `json:"lastByte"`
	LastMileRTT  string `json:"lastMileRTT"`
}

// NetworkStruct is used for storing the JSON subfields
type NetworkStruct struct {
	Asnum       string `json:"asnum"`
	EdgeIP      string `json:"edgeIP"`
	Network     string `json:"network"`
	NetworkType string `json:"networkType"`
}

// ReqHdrStruct is used for storing the JSON subfields
type ReqHdrStruct struct {
	Cookie string `json:"cookie"`
}

// RespHdrStruct is used for storing the JSON subfields
type RespHdrStruct struct {
	Server  string `json:"server"`
	ContEnc string `json:"contEnc"`
}

// CreateObjects creates a list of AkamaiPayloads from a raw byte slice
func CreateObjects(jsonFile []byte) (_ []AkamaiPayload, err error) {
	var arrayObject []AkamaiPayload
	err = json.Unmarshal(jsonFile, &arrayObject)
	return arrayObject, err
}

// Handle parses the incoming request data
func Handle(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "OK! I got you Bro! -> ")
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
		//Precision: "s",
	})
	if err != nil {
		log.Fatal(err)
	}

	for key, o := range obj {
		//payload[0].Geo["city"]
		fmt.Println("Server is:", o.RespHdr.Server, "KEY:", key)

		if err != nil {
			log.Fatal(err)
		}
		status, err := strconv.ParseInt(o.Message.Status, 10, 32)

		// Create a point and add to batch
		tags := map[string]string{
			"cp":          o.CP,
			"format":      o.Format,
			"city":        o.Geo.City,
			"country":     o.Geo.Country,
			"status":      o.Message.Status,
			"cacheStatus": o.NetPerf.CacheStatus,
		}
		lat, err := strconv.ParseFloat(o.Geo.Lat, 64)
		long, err := strconv.ParseFloat(o.Geo.Long, 64)
		if err != nil {
			log.Fatal(err)
		}
		fields := map[string]interface{}{
			"lat":    lat,
			"long":   long,
			"status": status,
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
