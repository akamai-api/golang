package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type AkamaiPayload struct {
	CP     string            `json:"cp"`
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

func CreateObject(jsonFile []byte) (_ []AkamaiPayload, err error) {

	var arrayObject []AkamaiPayload
	err = json.Unmarshal(jsonFile, &arrayObject)
	return arrayObject, err
}

func Handle(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "OK! I got you Bro! -> ")
	body, _ := ioutil.ReadAll(r.Body)
	obj, err := CreateObject(body)
	if err != nil {
		log.Fatal(err)
	}

	for _, o := range obj {
		//payload[0].Geo["city"]
		fmt.Println(o.Geo["city"])
	}
}

func main() {
	http.HandleFunc("/", Handle)
	http.ListenAndServe(":9143", nil)
}
