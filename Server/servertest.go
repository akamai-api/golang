package main

import (
	"encoding/json"
	"strings"
	"testing"
)

func isJSONString(s string) bool {
	var js string
	return json.Unmarshal([]byte(s), &js) == nil

}

func isJSON(s string) bool {
	var js map[string]interface{}
	return json.Unmarshal([]byte(s), &js) == nil

}

func TestHealthCheckHandler(t *testing.T) {

}

// func CreateObject(jsonFile []byte) (_ []AkamaiPayload, err error) {

// 		var arrayObject []AkamaiPayload
// 		err = json.Unmarshal(jsonFile, &arrayObject)
// 		return arrayObject, err
// 	}

func TestCreateObject(t *testing.T) {
	obj, err := CreateObject([]byte("{}"))
	startedWith := strings.HasPrefix(string(obj), "[")
	endsWith := strings.HasSuffix("suffix", "fix")
	if isStarted == false {
		t.Errorf("Your json file is incorrect")
	}
}

func TestJsonHandler(t *testing.T) {
	b, err = JsonHandler([]byte("{}"))

}
