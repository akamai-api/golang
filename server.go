package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

// AkamaiPayload is a Golang representation of the Cloudmonitor JSON datastructure
type AkamaiPayload struct {
	// Content Provider ID
	CP string `json:"cp"`
	// Defines the format of the payload (?)
	Format string            `json:"format"`
	Geo    map[string]string `json:"geo"`
	// city
	// country
	// lat
	// long
	// region
	ID      string            `json:"id"`
	Message map[string]string `json:"message"`
	// UA
	// bytes
	// cliIP
	// fwdHost
	// proto
	// protoVer
	// reqHost
	// reqMethod
	// reqPath
	// reqPort
	// respCT
	// respLen
	// status
	NetPerf map[string]string `json:"netPerf"`
	// asnum
	// cacheStatus
	// downloadTime
	// edgeIP
	// firstByte
	// lastByte
	// lastMileRTT
	Network map[string]string `json:"network"`
	// asnum
	// edgeIP
	// network
	// networkType
	ReqHdr map[string]string `json:"reqHdr"`
	// cookie
	RespHdr map[string]string `json:"respHdr"`
	// server
	// contEnc
	Start   string `json:"start"`
	Type    string `json:"type"`
	Version string `json:"version"`
}

// CreateObject creates a list of AkamaiPayloads from a raw byte slice
func CreateObject(jsonFile []byte) (_ []AkamaiPayload, err error) {
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
	obj, err := CreateObject(body)
	if err != nil {
		log.Fatal(err)
	}

	for key, o := range obj {
		//payload[0].Geo["city"]
		fmt.Println("City is:", o.Geo["city"], "KEY:", key)
	}
}

func main() {
	http.HandleFunc("/", Handle)
	http.ListenAndServe(":9143", nil)
}
