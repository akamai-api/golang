package main

import (
	"testing"
)

// TestValidPayload tests if we can create a proper object from
// a valid JSON payload.
func TestValidPayload(t *testing.T) {
	json := `[{
		"message" : {
		   "reqHost" : "www.example.com",
		   "respLen" : "276248",
		   "cliIP" : "123.123.123.123",
		   "status" : "503",
		   "bytes" : "123440",
		   "protoVer" : "1.1",
		   "respCT" : "text/html",
		   "UA" : "Mozilla%2f5.0%20(Macintosh%3b%20Intel%20Mac%20OS%20X%2010.9%3b%20rv%3a28.0)%20Gecko%2f20100101%20Firefox%2f28.0%20(FlipboardProxy%2f1.1%3b%20+http%3a%2f%2fflipboard.com%2fbrowserproxy)",
		   "reqMethod" : "POST",
		   "fwdHost" : "www.example.com",
		   "proto" : "http",
		   "reqPort" : "80",
		   "reqPath" : "%2f"
		},
		"netPerf" : {
		   "asnum" : "8523",
		   "downloadTime" : "1",
		   "edgeIP" : "165.254.92.141",
		   "lastByte" : "0",
		   "lastMileRTT" : "102",
		   "firstByte" : "0",
		   "cacheStatus" : "0"
		},
		"network" : {
		   "asnum" : "8523",
		   "edgeIP" : "165.254.92.141",
		   "networkType" : "",
		   "network" : ""
		},
		"cp" : "123456",
		"id" : "915cfea5570f824cc27112-a",
		"version" : "1.0",
		"start" : "1460634188.565",
		"type" : "cloud_monitor",
		"format" : "default",
		"respHdr" : {
		   "server" : "Microsoft-IIS/8.5",
		   "contEnc" : "identity"
		},
		"geo" : {
		   "lat" : "59.33",
		   "region" : "AB",
		   "long" : "18.05",
		   "country" : "DE",
		   "city" : "dummy"
		},
		"reqHdr" : {
		   "cookie" : "drbanan%3d1"
		}
	 }]`

	payloads, err := CreateObjects([]byte(json))
	if err != nil {
		t.Errorf("Error while trying to decode valid JSON payload: %s", err)
	}

	if len(payloads) != 1 {
		t.Errorf("Unexpected number of payloads in JSON: Should be 1, is %d", len(payloads))
	}

	payload := payloads[0]
	if payload.CP != "123456" {
		t.Errorf("CP not correct in payload. Should be 123456, is %s", payload.CP)
	}
	if payload.ID != "915cfea5570f824cc27112-a" {
		t.Errorf("ID not correct in payload. Should be 915cfea5570f824cc27112-a, is %s", payload.ID)
	}
}
